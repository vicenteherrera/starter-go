# starter-go

A starter Go language project.

Thanks to [Nestor](https://twitter.com/nestorsalceda) for teaching me how to structure this while working together.

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

This will be set once for the current user. For different Go projects under the same use, we will use _Go modules_. This will also allow updating the Go version by just removing `/usr/local/go` directory and extracting a new version there without having to backup anything.

### Is it neccesary to set GO111MODULE?

Not any more, it's the default setting.

Some instructions may include specifics steps for Go version <1.16 that define that variable temporarily, like:

```bash
GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0
```

If you use a recent version of Go it is not neccesary.

### IDEs and code editors

Some popular options to edit Go code are:

* [Goland](https://www.jetbrains.com/go/buy/#commercial) (not free)
* [Visual Studio Code](https://code.visualstudio.com/)
  * [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) extension
* [vim](https://danielmiessler.com/study/vim/)
  * [vim-go](https://github.com/fatih/vim-go) plugin

The Go extension for VS Code will install additional Go binaries when it detects a Go program, and will use our local `GOPATH` if configured correctly. So you have to start VS Code from a terminal or environment where your configuration is set up as explained before.

Take into consideration that if your code editor is erasing an import you added when you save a file, it's because it is detecting that the import is not being used. Try first referencing the library inside the code, then when you save the import may even be automatically added for you.

## Creating a Go program

### First step and Go modules.

When starting a program you should use _Go modules_. It was an optional feature added some time ago, but now is the standard for handling and locating dependencies. [Learn more about Go modules here](https://go.dev/doc/tutorial/create-module).

If for example you are creating a program that will be commited to a git repository located at [github.com/vicenteherrera/starter-go](github.com/vicenteherrera/starter-go) you should initialize this directory using:

```bash
# (You don't really need to do this on this repo, it has already been executed)
go mod init github.com/vicenteherrera/starter-go
```

This will create a `go.mod` file defining that name and go version.

A _Go module_ may include different _packages_, not only the file that includes the `main()` function. Each of them will have for package name (what you use for `import`) the same prefix name as the project module name, adding the folder structure where the file for the package is. Also, the package should have the same name as the child directory where it lives.

You can also define modules in other library projects that you reference from the one that have a `main()` and can create a binary.

The alternative to this is installing all dependencies for a Go project globally or on a `GOPATH` variable defined directory that you change for each project specified directory, which takes time and prune to errors.

### Adding package dependencies

Whenever you add a package dependency using `import`, you can track them in `go.sum` file just running from the command line:

```bash
go mod tidy
```

Or you can explicitely add a dependency using something like:

```bash
go get github.com/spf13/viper
```

If you create a dependency on a different project that uses a naming scheme like `example.org/mydep`, but you want to import it locally (maybe you are also making non-pushed changes to it) you can specify the local directory to fetch it from using:

```bash
go mod edit -replace example.com/mydep=../path-to-mydep
```

### Directory structure

For a binary project the recommended directory structure is this:

* build
  * `Dockerfile`
* cmd
  * `starter-go` (use the name of your binary/project)
* docs
* pkg
  * samplepkg (short package names preferred)
    * `file_name.go`
    * `file_name_test.go` (unit tests over `file_name.go`)
    * mocks
      * `file_name.go` (code generated by MockGen)
* release
  * `starter-go` (build binary)
  * `config.yaml` (additional config files needed to run)
* test
  * `config.yaml` (misc files used during testing)

### Compiling from the command line

Just run:

```bash
go build -o ./release/starter-go cmd/starter-go/main.go
# or using Makefile
make build
```

And a new binary `starter-go` will be generated on the project directory.

The binary requires configuration to run as an example. To use the provided one do:

```bash
cd ./release && ./starter-go
# or using Makefile
make run
```

## Building a container image

Use `build/Containerfile` as an example starting point (it's what has been previously called a Dockerfile). It uses two build stages and a lean distroless go base image for execution.

Build it for example using:

```bash
docker build -f build/Containerfile -t vicenteherrera/starter-go .
# or using Makefile
make container-build
```

You can test its execution with:

```bash
docker run --rm -it vicenteherrera/starter-go --help
```

## Developing in Go

### Reading configuration

For reading configuration from a file, command line parameters, or environment variables, we will use the powerful libraries [Viper](https://github.com/spf13/viper) and [pflag](https://github.com/spf13/pflag).

See the example usage on this repo in `cmd/starter-go/main.go`.

### Logging

We will use [logrus](github.com/sirupsen/logrus).

## Testing

Test suite:
* [Ginkgo test suite](https://onsi.github.io/ginkgo/)

Automatic Mock generation:
* [Gomock](https://github.com/golang/mock/)


## Code Style

Read [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).

## Architecture

Read [The Twelve-Factor App](https://12factor.net/).
