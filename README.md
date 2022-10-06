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

* On `go.mod` change `module github.com/vicenteherrera/starter-go` for your repo and folder location.
* Edit `cmd/starter-go/main.go`, change references from `github.com/vicenteherrera/starter-go` to your repo and folder.
* Rename `cmd/starter-go` to the name you want for the project.
* Edit `makefile` and change values for `TARGET_BIN`, `MAIN_DIR` and `CONTAINER_IMAGE`.
* Edit `build/Containerfile` and change `ENTRYPOINT` value with your binary name.
* Execute `go mod tidy`.
* Check you can execute tests, build, run from local and in a container (see `makefile`).

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

### Directory structure

For a binary project the recommended directory structure is this:

* `go.mod` (defines project base URL and requirements)
* `go.sum` (locks specific dependecies version)
* `build/`
  * `Containerfile` (yes, this is a Dockerfile)
  * `Containerfile.containerignore`
* `cmd/`
  * `starter-go.go` (use the name of your binary/project)
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

# Create suits and specs
# ...

# Execute tests
ginkgo -randomize-all -randomize-suites -fail-on-pending -trace -race -progress -cover -r
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
