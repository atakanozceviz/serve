
# serve

A simple golang file server with no dependencies.

### Installation

Download the appropriate version for your platform from [Releases](https://github.com/atakanozceviz/serve/releases). Once downloaded, the binary can be run from anywhere.

### Build and Install the Binaries from Source

(requires [Go (golang)](https://golang.org) installed on your machine to build.)
```sh
$ go get -v github.com/atakanozceviz/serve
```

Once the `get` completes, you should find your new `serve` (or `serve.exe`) executable inside `$GOPATH/bin/`.

### Usage

```sh
$ serve -h
Usage of serve:
  -d string
    	the directory or file to host (default ".")
  -p string
    	port to serve on (default "8080")
```

You can use command line argument instead of `-d` flag.

#### Example:

```sh
$ serve /path/to/a/file/or/directory
```

