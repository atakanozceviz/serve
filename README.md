
# serve

A simple file server with no dependencies.

### Installation

Download the appropriate version for your platform from [Releases](https://github.com/atakanozceviz/serve/releases). Once downloaded, the binary can be run from anywhere.

Ideally, you should install it somewhere in your PATH for easy use. /usr/local/bin is the most probable location.

### Installing from source

(requires [Go (golang)](https://golang.org) installed on your machine to build.)
```sh
$ go get -v github.com/atakanozceviz/serve
```

Once the `get` completes, you should find your `serve` (or `serve.exe`) executable inside `$GOPATH/bin/`.

### Usage

```sh
$ Usage of serve:
  -p string
        port on which the server will listen (default "4242")
  -s string
        filesystem path to read file(s) relative from (default ".")
```

You can use command line argument instead of `-s` flag.

#### Example:

```sh
$ serve /path/to/a/file/or/directory
```

Add `/u` to the end of the url to upload files to the server like so:
```
http://[server]:4242/u
```
