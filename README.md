# starter-go

[![Go build](https://github.com/vicenteherrera/starter-go/actions/workflows/go-build.yaml/badge.svg?branch=main&event=push)](https://github.com/vicenteherrera/starter-go/actions/workflows/go-build.yaml)
[![Go test unit](https://github.com/vicenteherrera/starter-go/actions/workflows/go-test-unit.yaml/badge.svg?branch=main&event=push)](https://github.com/vicenteherrera/starter-go/actions/workflows/go-build.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/vicenteherrera/starter-go)](https://goreportcard.com/report/github.com/vicenteherrera/starter-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/vicenteherrera/starter-go.svg)](https://pkg.go.dev/github.com/vicenteherrera/starter-go)

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

### Set PATH for local user go directory

When globally installing go binaries with `go install` they will by default install to `$HOME/go/bin`.

Modify your configuration to set the `PATH` variable to also include it:

```bash
export PATH="$GOPATH"/bin:/usr/local/go/bin:$PATH
```

### Is it neccesary to set GOPATH, GOROOT or GO111MODULE?

On modern Go versions there is no need to set `GOPATH`, `GOROOT` or `GO111MODULE` variables. Update yours if you installed it a long time ago.

Some instructions may include specifics steps for Go version <1.16 that define `GO111MODULE` environmen variable temporarily, like:

```bash
GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0
```

But again, if you use a recent version of Go it is not neccesary.

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

### Using this repo as a starting point

To use the files in this repo as a starting point and skip everything else, you have to do to it some modifications.

* Rename `cmd/starter-go` directory to the name you want for the project.
* Edit file previous located at `cmd/starter-go/root.go`:
  * change references from `github.com/vicenteherrera/starter-go/...` to your repo and folder.
  * Change references from `starter-go` to the name of your program.
* Edit `main.go` and change `github.com/vicenteherrera/starter-go/cmd/starter-go` to your repo and folder.
* Edit `makefile` and change values for `GH_REPO`, `TARGET_BIN` and `CONTAINER_IMAGE`.
* Edit `build/Containerfile` and change `ENTRYPOINT` value with your binary name.
* Execute `go mod tidy`.
* Check you can execute tests, build, run from local and in a container (see `makefile`).
* Change installations scripts `install/install.sh` and `install/install.ps1` to point to your own repo.
* Change `README.md` header badges and installation instructions to point to your own repo.

### Creating a project from scratch

#### First step and Go modules

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

### Using Cobra

If you are building a Command Line Interface (CLI) binary as this example, you can use:

[Viper](https://github.com/spf13/cobra) library will help you handle configuration from environment variables, a config file, or command line parameters.
[Cobra](https://github.com/spf13/viper) library will help you setting a hierarchy of nested commands, with their own parameters that are locally, or can be inherit by any subcommands.

You can use Viper on its own, or Cobra and Viper in conjuntion, even if you just set a root command with any further subcommands (as this example does).

You can write your own Go code that uses the Cobra library, or install the [cobra-cli](https://github.com/spf13/cobra) library using `go install github.com/spf13/cobra-cli@latest`.

You can use it to help create code for commands:

```bash
# Initialize a project with Cobra library code
cobra-cli init
# Add a new serve command
cobra-cli add serve
```

See more infor in [cobra-cli documentation](https://github.com/spf13/cobra-cli/blob/main/README.md).

### Directory structure

For a binary project the recommended directory structure is this:

* `go.mod` (defines project base URL and requirements)
* `go.sum` (locks specific dependecies version)
* `main.go` (main file that bootstrap Cobra commands from cmd/starter-go)
* `build/`
  * `Containerfile` (yes, this is a Dockerfile)
  * `Containerfile.containerignore`
* `cmd/`
  * `starter-go/`
    * `root.go` (root command that loads parameters starts real execution)
* `docs/` (documentation in markdown format)
* `pkg/`
  * `samplepkg/` (short package names preferred)
    * `client.go` (starts with `package samplepkg`)
    * `client_test.go` (tests for `client.go`)
    * `model.go` (structs)
    * `mocks/`
      * `client.go` (code generated by MockGen)
* `release/`
  * `starter-go` (build binary)
  * `config.yaml` (additional config files needed to run)
* `test/`
  * `config.yaml` (misc files used during testing)

### Compiling from the command line

Just run:

```bash
go build -o ./release/starter-go ./main.go
# or using Makefile
make build
```

The `makefile` target will use `ldflags` to stablish default value for the `version` variable inside the binary using the latest repo tag.

And a new binary `starter-go` will be generated on the `./release/` directory.

The binary requires configuration to run as an example. To use the provided one do:

```bash
cd ./release && ./starter-go
# or using Makefile
make run
```

## Installation

_These are the installation instructions you should use in your project._

To install the latest release, go to [releases](https://github.com/vicenteherrera/starter-go/releases) and download the latest version for your platform, or use these script:

```bash
# Linux / MacOs (Bash): install binary to /usr/local/bin
curl -fsSL https://raw.githubusercontent.com/vicenteherrera/starter-go/main/install/install.sh | sudo bash -s

# Windows (Powershell)
iwr https://raw.githubusercontent.com/vicenteherrera/starter-go/main/install/install.ps1 -useb | iex
```

You could install with `go install`, but you may get an unstable version not yet tagged.
```bash
go install github.com/vicenteherrera/go-starter@latest
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

~~We will use [logrus](github.com/sirupsen/logrus).~~

Logrus is now on maintenance mode, I will update with an alternative in the future.

## Testing

### Ginkgo

Install [Ginkgo test suite](https://onsi.github.io/ginkgo/) with

```bash
go install github.com/onsi/ginkgo/v2/ginkgo
go install github.com/onsi/gomega/...
```
Taps into existing Go testing infrastructure (tests live in *_test.go), using BDD.

spec = individual test
suite = collection of specs

```bash
# Initialize the test suit for the package
cd pkg/analyzer
ginkgo bootstrap

# Create a spec for a file in the package
ginko generate analyzer

# Dowload required packages
go mod tidy

# Create suits and specs
# (see below)

# Execute tests
ginkgo -randomize-all -randomize-suites -fail-on-pending -trace -race -progress -cover -r -v
```

Read more about [writing specs](https://onsi.github.io/ginkgo/#writing-specs).

Example spec:
```go
var _ = Describe("Books", func() {
  var book *books.Book

  BeforeEach(func() {
    book = &books.Book{
      Title: "Les Miserables",
      Author: "Victor Hugo",
      Pages: 2783,
    }
    Expect(book.IsValid()).To(BeTrue())
  })

  Describe("Extracting the author's first and last name", func() {
    Context("When the author has both names", func() {
      It("can extract the author's last name", func() {
        fmt.Fprintf(GinkgoWriter, "Author Last Name:\n%s", book.AuthorLastName()) // output on test run when -v used
        Expect(book.AuthorLastName()).To(Equal("Hugo"))
      })

      It("can extract the author's first name", func() {
        Expect(book.AuthorFirstName()).To(Equal("Victor"))
      })      
    })

    Context("When the author only has one name", func() {
      BeforeEach(func() {
        book.Author = "Hugo"
      })  

      It("interprets the single author name as a last name", func() {
        Expect(book.AuthorLastName()).To(Equal("Hugo"))
      })

      It("returns empty for the first name", func() {
        Expect(book.AuthorFirstName()).To(BeZero())
      })
    })

  })
})
```

More information:
* [Ginkgo](https://onsi.github.io/ginkgo/)
* [Getting Started with BDD in Go Using Ginkgo](https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo)
* [Gomega](https://onsi.github.io/gomega/)

### Gomock

Automatic Mock generation:
* [Gomock](https://github.com/golang/mock/)

Install binary
```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

More information:
* [Mocking techniques for go](https://www.myhatchpad.com/insight/mocking-techniques-for-go/)

## GitHub Actions workflows

For automation we include some GitHub Actions workflows

* **release.yaml**: Using `goreleaser/goreleaser-action`, prepares a release on repo tag push
* **go-test-unit.yaml**: Executes `make test` to validate unit testing
* **go-build.yaml**: Executes `make build` to validate unit testing

Although you can have workflows that use specified published actions for building and/or testing, using the same `make` targets as you can execute locally makes it easier to validate their execution on your machine before reaching the workflow pipeline.

Build and test are executed in separated workflow so it's easier which one of the two is failing without getting into the workflow's logs.

Release will be prepare binaries for your project on selected architectures and operating systems when you push a tag to your repo. Using repo tags is also the required way of setting package versions in Go.

You can use the badges in the beginning of this readme as example to include in you own project.


## Code Style

Read [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).

## Debugging Go with Visual Studio Code

It's not difficult to achieve, [this guide is obsolete in some regards](https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code) but useful anyway.

## Architecture

Read [The Twelve-Factor App](https://12factor.net/).

## Read more

* [Go by example](https://gobyexample.com/)
* [Golang programs](https://www.golangprograms.com/)
* [Notes Go](https://notes.shichao.io/gopl/)
* [Go first steps](https://docs.microsoft.com/en-us/learn/paths/go-first-steps/)
* [GoLang Notes](https://sharbeargle.gitbooks.io/golang-notes/content/)
* [ZetCode Go Tutorial](https://zetcode.com/all/#go)
* [Iterate a YAML without knowing its structure](https://stackoverflow.com/questions/36765842/how-to-iterate-over-all-the-yaml-values-in-golang)
* [Flexible YAML shapes in go](https://abhinavg.net/posts/flexible-yaml/)
