# Scanning Multiple Hosts

`go-exploit` exploits can be used for scanning multiple hosts and ports. This can be done three ways:

1. Providing a comma delimited list of IP addresses to `rhosts`.
2. Same as 1 but also IP addresses using CIDR notation.
3. Providing a file of hosts (ip,port,ssl) via `rhosts-file`.

Examples follow:

## Comma Delimited `rhosts` options

```sh
albinolobster@mournland:~/go-exploit/examples/cve-2019-15107$ ./cve-2019-15107 -rhosts 10.9.49.174,10.9.49.205 -rports 80,10000 -a -v
[*] Starting target 0: 10.9.49.174:80
[*] Validating the remote target is a Webmin installation
[-] HTTP request error: Get "http://10.9.49.174:80/": dial tcp 10.9.49.174:80: connect: connection refused
[-] The target isn't recognized as Webmin, quitting
[*] Starting target 1: 10.9.49.174:10000
[*] Validating the remote target is a Webmin installation
[+] Target validation succeeded!
[*] Starting target 2: 10.9.49.205:80
[*] Validating the remote target is a Webmin installation
[-] The HTTP header doesn't appear to be Webmin
[-] The target isn't recognized as Webmin, quitting
[*] Starting target 3: 10.9.49.205:10000
[*] Validating the remote target is a Webmin installation
[-] HTTP request error: Get "http://10.9.49.205:10000/": dial tcp 10.9.49.205:10000: connect: connection refused
[-] The target isn't recognized as Webmin, quitting
```

Note that `-a` is useful when using `rhosts` or `rports` because it will autodiscover if the target supports SSL.

## CIDR Notation in `rhosts`

```sh
albinolobster@mournland:~/initial-access/feed/cve-2023-38646$ ./build/cve-2023-38646_linux-arm64 -v -rhosts 192.168.1.0/24 -rport 80
[*] Starting target 0: 192.168.1.0:80
[*] Validating the remote target is a Metabase installation
[-] HTTP request error: Get "http://192.168.1.0:80/": dial tcp 192.168.1.0:80: connect: no route to host
[-] The target isn't recognized as Metabase, quitting
[*] Starting target 1: 192.168.1.1:80
[*] Validating the remote target is a Metabase installation
[-] Missing the Set-Cookie header
[-] The target isn't recognized as Metabase, quitting
[*] Starting target 2: 192.168.1.2:80
[*] Validating the remote target is a Metabase installation
[-] HTTP request error: Get "http://192.168.1.2:80/": dial tcp 192.168.1.2:80: connect: no route to host
[-] The target isn't recognized as Metabase, quitting
[*] Starting target 3: 192.168.1.3:80
[*] Validating the remote target is a Metabase installation
[-] HTTP request error: Get "http://192.168.1.3:80/": dial tcp 192.168.1.3:80: connect: no route to host
[-] The target isn't recognized as Metabase, quitting
[*] Starting target 4: 192.168.1.4:80
[*] Validating the remote target is a Metabase installation
[-] HTTP request error: Get "http://192.168.1.4:80/": dial tcp 192.168.1.4:80: connect: no route to host
[-] The target isn't recognized as Metabase, quitting
etc.
```

## Hosts Provided via `rhosts-file` (Shodan)

The file format is supposed to be:

```csv
ip,port,<empty if no ssl | any text if ssl>
ip,port,<empty if no ssl | any text if ssl>
ip,port,<empty if no ssl | any text if ssl>
```

That might seem odd, but the goal here is two-fold:

1. Support getting targets from Shodan.
2. Support individual targets having SSL enabled or not.

That works like the following (note that `-v` is just validation, e.g. checks that the target is openfire, no target was harmed):

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
[*] Starting target 0: 202.39.62.186:9090
[*] Validating the remote target is a Openfire installation
[+] Target validation succeeded!
[*] Starting target 1: 124.72.94.36:9090
[*] Validating the remote target is a Openfire installation
[+] Target validation succeeded!
[*] Starting target 2: 103.82.30.158:9091
[*] Validating the remote target is a Openfire installation
[+] Target validation succeeded!
[*] Starting target 3: 88.203.168.164:9091
[*] Validating the remote target is a Openfire installation
[+] Target validation succeeded!
[*] Starting target 4: 41.77.78.157:9091
[*] Validating the remote target is a Openfire installation
[+] Target validation succeeded!
[*] Starting target 5: 193.93.123.117:9091
[*] Validating the remote target is a Openfire installation
^C
albinolobster@mournland:~$
```
