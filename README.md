# serve

A simple golang file server with no dependencies.

### Installation
(requires [Go (golang)](https://golang.org) to build.)

```sh
$ go get -u github.com/atakanozceviz/serve
```

### Usage
```sh
$ serve -h
Usage of serve:
  -d string
    	the directory or file to host (default ".")
  -p string
    	port to serve on (default "8080")
```
You can use command line argument instead of "-d" flag.

#### Example:
```sh
$ serve /path/to/a/file/or/directory
```

