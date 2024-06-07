# Updating the go-exploit HTTP User Agent

This main.go fetches user agents from the Project Discovery [useragent](https://github.com/projectdiscovery/useragent) package (using the [MIT license](https://github.com/projectdiscovery/useragent/blob/main/LICENSE)), and filters them down to the most recent Windows Chrome User-Agent. The output is written to `./protocol/http-user-agent.txt`.

Usage example:

```console
albinolobster@mournland:~/go-exploit/_uaupdate$ go run .
albinolobster@mournland:~/go-exploit/_uaupdate$ cat ../protocol/http-user-agent.txt 
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36
```