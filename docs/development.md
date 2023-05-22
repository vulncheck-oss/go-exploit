# Developing `go-exploit`

The goal of `go-exploit` is to facilitate faster, more portable, and feature-rich exploit development. It is important to note that `go-exploit` is not a repository of exploits itself. The infosec community benefits from having a wide range of exploits for a single CVE, rather than relying solely on a single "approved" exploit within a massive framework.

Therefore, contributions to this repository should focus on making exploit development easier or improving the overall quality of developed exploits. This can be achieved through various means, such as:

* Adding more command and control (C2) options
* Supporting additional protocols
* Including more payloads
* Implementing obfuscation techniques
* And more.

If you have ideas that align with these goals, we welcome your pull requests. Development within `go-exploit` is relatively straightforward. The only slightly different aspect is that `go-exploit` is composed of multiple subpackages. This design choice aims to prevent unwanted or unused features from being included in individual exploits. For example, an exploit that utilizes the JNDI LDAP functionality must explicitly import `github.com/vulncheck-oss/java/ldapjndi`. This strict gating helps minimize dependencies and reduce the amount of imported code.

## Linting

It is important that the project passes linting without any warnings or errors. You can use our built-in `.golangci.yml` file by running the following command:

```sh
golangci-lint run --fix
```

## Testing

Go has a robust built-in unit testing framework. We strongly encourage the development of new tests for issue reproduction and to provide coverage for new features. You can execute all tests by running the following command:

```sh
go test ./...
```

