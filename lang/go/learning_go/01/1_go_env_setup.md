# Go Environment Setup

<!--toc:start-->
- [Go Environment Setup](#go-environment-setup)
  - [Installing the Go Tools](#installing-the-go-tools)
    - [Go Tooling](#go-tooling)
  - [Your First Go Program](#your-first-go-program)
    - [Making a Go Module](#making-a-go-module)
    - [go build](#go-build)
    - [go fmt](#go-fmt)
  - [Makefiles](#makefiles)
  - [The Go Compatibility Promise](#the-go-compatibility-promise)
  - [Wrapping Up](#wrapping-up)
<!--toc:end-->

## Installing the Go Tools
The latest version of the Go Tools can be found on the Go website. Windows(🙁) and Mac OS users can simply download the .msi or .pkg installers respectively.

For Linux systems the installers are gzipped TAR files that expand to a directory named go. Start by removing any previous Go installs and extract the archive into /usr/local/go.
```
$ sudo rm -rf /usr/local/go
$ sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```

Next add */usr/local/go/bin* to the `PATH` environment variable. Add the following to either  *$HOME/.profile* or */etc/profile*(for a system-wide install) for whichever shell you may be using(bash, zsh, fish, etc.):
```
export PATH=${PATH}:/usr/local/go/bin
```

Verify your install with the following command:
```
$ go version
```

### Go Tooling
All of the Go development tools are accessed via the `go` command. There's a compiler(`go build`), code formatter(`go fmt`), dependency manager(`go mod`), test runner(`go test`), a tool that scans for common coding errors(`go vet`), and more.

## Your First Go Program
Next we'll cover the basics of writing a Go program.

### Making a Go Module
To create a module you'll first need a directory, once inside run the following command to mark it as a Go module:
```
$ go mod init hello_world
go: creating new go.mod: module hello_world
```


### go build


### go fmt

## Makefiles

## The Go Compatibility Promise

## Wrapping Up
