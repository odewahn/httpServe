A go program that:

* serves the current working directory as a static site
* proxies websockets from the specified backend


The goal is to be able to put a single, simple server on an alpine docker container so that it can serve a static site. Nothing fancy.

To use it:

```
$ httpServe -h
Usage of ./httpServe:
  -backend string
    	Backend URL for proxying (default "ws://localhost:8888")
  -dir string
    	directory to serve (default "/Users/odewahn/golang/src/github.com/odewahn/httpServe")
  -port string
    	port (default "8000")
  -version
    	show version
```

To serve a sub-directory:

```
$ httpServe -dir $(pwd)/static
```


## Building for linux

```
env GOOS=linux GOARCH=amd64 go build -o httpServe-0.0.3-linux-amd64 httpServe.go
```

For more detail, see this great article on cross-compilation:

http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
