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

Modules will be covered in more detail in [chapter 10](../10/10_modules_packages_and_imports.md). For now, remember that a module isn't just source code, it's also a specification of the dependencies within the module. Running `go mod init` creates a *go.mod* file which declares the name of the module, the minimum version of Go for the module, and any other modules that your module depends on. Here's what the file might look like:
```
module hello_world

go 1.22.0
```

You won't edit the *go.mod* file directly. The `go get` and `go mod tidy` commands manage changes to this file. Again more on modules in [chapter 10](../10/10_modules_packages_and_imports.md).

### go build
Let's write a classic hello world program in go, we'll create a file called `hello.go` in our module:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Let's quickly go over the contents of this file. Within a Go module, code is organized into one or more packages. The `main` package in a Go module contains the code that starts a Go program.

The `import` statement lists the packages referenced in this file. We are using a function in the `fmt`(pronounced "fumpt") package from the standard library. Unlike other languages, Go imports only whole packages.

All Go programs start from the `main` function in the `main` package. You declare this function with `func main()`.

In the body of the function you are calling the `Println` function in the `fmt` package with the argument `"Hello, World!"`.

Once you've saved the file, go to your terminal or command prompt and enter the following command:
```
$ go build
```

This creates an executable `hello_world`(or hello_world.exe on Windows 🙁). Let's run it:
```
$ ./hello_world
Hello, world
```

The name of the binary matches the name in the module declaration. To change the name of the executable or store it in a different location use the `-o` flag. 
```
$ go build -o hello
```

### go fmt
One of the goals for Go was to create a language that made it easy to write code efficiently, this meant having a simple syntax and a fast compiler. To aid with that Go enforces a standard format which makes it easier to write tools that manipulate source code, simplifies the compiler, and also for better tools for generating code.(Additionally no more arguments over brace styling or indentation)

The `go fmt` command automatically fixes whitespace in your code to match the standard format and is run like so:
```
$ go fmt ./...
```

The `./...` tells a Go tool to apply the command to all the files in the current directory and all subdirectories. Most IDEs can be set up to run `go fmt` automatically on save using the `gopls` lsp(Language Server Protocol).

## Makefiles
Modern software development relies on repeatable, automatable builds that can be run by anyone, anywhere, at anytime. While there are many tools to do this make is a tried and tested one that has been used to build programs on Unix systems for decades.

Let's create a file called *Makefile* in the module we created earlier:
```make
.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean
```

The `.DEFAULT_GOAL` defines which target to run when no target is specified, in this case the default target is `build`. Next we have the target definitions. The word before the colon(`:`) is the name of the target. Any words that come after the target(like vet in build: vet) are the other targets that must be run before the specified target runs. That tasks that are performed by the target are on the indented line after the target. The `.PHONY` line keeps `make` from getting confused if a directory or file in your project has the same name as one of the listed targets.

We can run `make` like this:
```
$ make
go fmt ./...
go vet ./...
go build
```

Entering a single command formats the code correctly, checks it for errors with `go vet`, and compiles it. You can also vet the code with `make vet`, or just run the formatter with `make fmt`. This isn't a big improvement for in the context of our simple example but for larger projects with complicated build steps this can make life much easier.

If you want to learn more about writing Makefiles check out this [addendum](../addendums/makefiles.md)

## The Go Compatibility Promise
The [Go Compatibility Promise](https://go.dev/doc/go1compat) is a detailed description of how the Go team plans to avoid breaking Go code. It says that there won't be backward-breaking changes to the language or the standard library for any Go version that starts with 1, unless the change is required for a bug or security fix.

This guarantee doesn't apply to the `go` commands. There have been backward-incompatible changes to the flags and functionality of `go` commands and it is possible that it will happen again.

## Wrapping Up
This chapter covered how to install and configure your Go development environment as well as tools for building Go programs and ensuring code quality. In the [next chapter](../02/2_predeclared_types_and_declarations.md) we'll cover the built-in types in Go and how to declare variables.
