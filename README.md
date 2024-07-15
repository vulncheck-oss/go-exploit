# go-exploit: Go Exploit Framework

[![Go](https://github.com/vulncheck-oss/go-exploit/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/vulncheck-oss/go-exploit/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/vulncheck-oss/go-exploit)](https://goreportcard.com/report/github.com/vulncheck-oss/go-exploit)

`go-exploit` is an exploit development framework for [Go](https://go.dev/). The framework helps exploit developers create small, self-contained, portable, and consistent exploits. The framework was developed to simplify large scale scanning, exploitation, and integration with other tools. For API documentation, check out the package on [pkg.go.dev/github.com/vulncheck-oss/go-exploit](https://pkg.go.dev/github.com/vulncheck-oss/go-exploit).

## Go Exploit Phases

The Go Exploit Framework includes the following Phases which can be chained or executed independently:

* [Go Exploit Framework Phases](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/getting-started.md)
    * Step 1 - Target Verification
    * Step 2 - [Version Scanning](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/version-checking.md)
    * Step 3 - [Exploitation](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/exploit-types.md)
    * Step 4 - [Command & Control](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/c2.md)

## Go Exploit Features

The Go Exploit Framework includes these additional features:

* [Auto-detection](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/scanning.md#autodetect-ssl) of SSL/TLS on the remote target.
* Fully [proxy-aware](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/scanning.md#Proxy).
* Key-value or JSON [output](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/output.md) for easy integration into other automated systems.
* Builtin Java [gadgets](https://github.com/vulncheck-oss/go-exploit/blob/main/java/javagadget.go), [classes](https://github.com/vulncheck-oss/go-exploit/blob/main/java/javaclass.go), and [LDAP](https://github.com/vulncheck-oss/go-exploit/blob/main/java/ldapjndi/ldapjndi.go) infrastructure.
* Many [reverse shell](https://github.com/vulncheck-oss/go-exploit/blob/main/payload/reverse), [dropper](https://github.com/vulncheck-oss/go-exploit/tree/main/payload/dropper), and [bind shell](https://github.com/vulncheck-oss/go-exploit/blob/main/payload/bindshell) payloads.
* Functionality that integrates exploitation with other [tools](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/c2.md#using--o) or frameworks like [Metasploit](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/c2.md#using-httpservefile) and Sliver.
* Builtin ["c2"](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/c2.md) for catching encrypted/unencrypted shells or hosting implants.
* Supports multipe target [formats](https://github.com/vulncheck-oss/go-exploit/blob/main/docs/scanning.md#providing-targets) including lists, file-based, VulnCheck IP-Intel, and more.

## Examples

* [CVE-2023-22527](https://github.com/vulncheck-oss/cve-2023-22527): Three go-exploit implementations taking unique approaches to Atlassian Confluence CVE-2023-22527.
* [CVE-2023-25194](https://github.com/vulncheck-oss/cve-2023-25194): Demonstrates exploiting CVE-2023-25194 against Apache Druid (using Kafka).
* [CVE-2023-46604](https://github.com/vulncheck-oss/cve-2023-46604): Demonstrates exploiting CVE-2023-46604 and using the go-exploit HTTPServeFile c2.
* [CVE-2023-36845](https://github.com/vulncheck-oss/cve-2023-36845-scanner): Scans for Juniper firewalls to determine if they are vulnerable to CVE-2023-36845.
* [CVE-2023-51467](https://github.com/vulncheck-oss/cve-2023-51467): A go-exploit implementation of CVE-2023-51467 that lands a Nashorn reverse shell.

## Contributing

Community contributions in the form of issues and features are welcome. When submitting issues, please ensure they include sufficient information to reproduce the problem. For new features, provide a reasonable use case, appropriate unit tests, and ensure compliance with our `.golangci.yml` without generating any complaints.

Please also ensure that linting comes back clean, and all tests pass.

```sh
golangci-lint run --fix
go test ./...
```

## License

`go-exploit` is licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0). For more details, refer to the LICENSE file.

