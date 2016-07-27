A go program that serves the current working directory as a static site.  The goal is to be able to put a single, simple server on an alpine docker container so that it can serve a static site. Nothing fancy.

To use it:

```
$ httpServe -h
Usage of httpServe:
  -dir string
    	directory to serve (defaults to current working directory)
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
env GOOS=linux GOARCH=amd64 go build -o httpServe-linux-amd64 httpServe.go
```

For more detail, see this great article on cross-compilation:

http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
