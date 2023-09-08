package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/vulncheck-oss/go-exploit/output"
	"github.com/vulncheck-oss/go-exploit/random"
)

func GenerateCertificate() (tls.Certificate, bool) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		output.PrintFrameworkError(err.Error())

		return tls.Certificate{}, false
	}

	template := x509.Certificate{
		SerialNumber:          big.NewInt(8),
		Subject:               pkix.Name{CommonName: random.RandLetters(12)},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	certificateBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		output.PrintFrameworkError(err.Error())

		return tls.Certificate{}, false
	}

	certificate := tls.Certificate{
		Certificate: [][]byte{certificateBytes},
		PrivateKey:  privateKey,
	}

	return certificate, true
}
