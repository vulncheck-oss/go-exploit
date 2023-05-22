# Version Checking

Version checking is a crucial step in the exploit process to ensure that the target is vulnerable and worth exploiting. The `CheckVersion` function in `go-exploit` determines whether the exploit will attempt to exploit the target. It has five possible return values defined in `framework.go`:

1. *NotVulnerable* - Indicates that the target is not vulnerable, and the exploit will not attempt to exploit it.
2. *Vulnerable* - Indicates that the target is vulnerable, and the exploit will attempt to exploit it.
3. *PossiblyVulnerable* - Indicates that the target might be vulnerable, and the exploit will attempt to exploit it.
4. *Unknown* - Indicates that an error occurred during version checking, and the exploit will not attempt to exploit the target.
5. *NotImplemented* - Indicates that no version check is implemented, and the exploit will attempt to exploit the target.

Here's an example of a version check function for CVE-2017-20149:

```go
func (sploit MTStackClash) CheckVersion(conf *config.Config) exploit.VersionCheckType {
	version, ok := getRouterOSVersion(conf)
	if !ok {
		return exploit.Unknown
	}
	output.PrintfStatus("The self-reported version is: %s", version)

	major, minor, point := versionToInt(version)
	if major == 0 {
		return exploit.Unknown
	}

	if major != 6 || minor >= 39 {
		return exploit.NotVulnerable
	}
	if minor == 38 && point >= 5 {
		return exploit.NotVulnerable
	}
	if minor == 37 && point >= 5 {
		return exploit.NotVulnerable
	}

	return exploit.Vulnerable
}
```

In this example, the function retrieves the target's version using the `getRouterOSVersion` function and performs a series of checks. If the version is unknown or if the target is not running RouterOS version 6 or has a minor version greater than or equal to 39, it returns *NotVulnerable*. If the minor version is 38 and the point version is greater than or equal to 5, or if the minor version is 37 and the point version is greater than or equal to 5, it also returns *NotVulnerable*. Otherwise, it concludes that the target is *Vulnerable*.

