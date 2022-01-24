# starter-go

A starter Go language project

## Introduction

Go is not a difficult language to learn, compared to other options. The syntaxis was specifically designed to _take things out_ so you have a minimal set of tools that prevent developers to overengineer programs. And it doesn't offer a ton out-of-the-box libraries you are required to wrap your code around.

But many of the things that makes a great Go program are beyond its syntaxis and libraries, and more on _how it is supposed to be used_.

I created this repo for my own sake to take notes on this recommendations, common best practices, important third party libraries, and starting codebase.

## Setting up your development environment

### Install Go

The official _Go_ website [https://go.dev/doc/install](https://go.dev/doc/install) provides steps to install required Go binaries and files, for Linux, Mac or Windows.

It will also instruct you to add the Go binaries directory to your path, something like this on Linux

```bash
export PATH=$PATH:/usr/local/go/bin
```

Then you can verify the installation worked by running:

```bash
go version
```

If you are using go in WSL2 Linux inside Windows, it may be useful to install it also on the Windows part.

### Set GOPATH for local user

On Linux when you want to gobally install a Go binary using `go install` it will try to use `/usr/local/go/bin/` directory to create binaries and `/usr/local/go/pkg/` to cache non-main packages. But it will fail as it requires root to write there. Instead of messing up with `sudo`, let's create a directory local to the user for this:

```bash
mkdir -p $HOME/go/bin
```
Modify your configuration to set the `GOPATH` variable and path:

```bash
export GOPATH=/usr/local/go
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

This will be set once for the current user. For different Go projects under the same use, we will use _Go modules_.

### IDEs and code editors

Some popular options to edit Go code are:

* [Goland](https://www.jetbrains.com/go/buy/#commercial) (not free)
* [Visual Studio Code](https://code.visualstudio.com/)
  * [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) extension
* [vim](https://danielmiessler.com/study/vim/)
  * [vim-go](https://github.com/fatih/vim-go) plugin

The Go extension for VS Code will install additional Go binaries when it detects a Go program, and will use our local `GOPATH` if configured correctly. So you have to start VS Code from a terminal or environment where your configuration is set up as explained before.

