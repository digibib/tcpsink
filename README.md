# tcpsink

Very simple tcp server for slurping TCP connections that need no reponse

Usage of tcpsink:
```
  -l string
    	local address (default "9999")
  -h string
      host to listen on (default "localhost")
  -p string
    	prefix for logging (default "tcpsink: ")
```

To build docker image:
```
CGO_ENABLED=0 go build .
docker build -t <tag> .
```