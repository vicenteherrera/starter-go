package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	flag.Usage = Usage

	var svar string
	flag.StringVar(&svar, "svar", "bar", "dos")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

func Usage() {
	w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
	fmt.Fprintf(w, "Usage of %s: ...custom preamble... \n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(w, "...custom postamble ... \n")
}
