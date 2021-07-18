package main

import (
	"flag"
	"fmt"
	"os"
)

var name *string

// var question *string
// var cmdLine *flag.FlagSet

func init() {

	// cmdLine = flag.NewFlagSet("", flag.ExitOnError)
	// cmdLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// }
	// question = cmdLine.String("question", "good question", "Usage of Question")

	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "name")
		flag.PrintDefaults()
	}

	//flag.StringVar(&name, "name", "ybq", "print name of author by flag.StringVar.")
	name = flag.String("name", "ybq", "Usage of print name of author by flag.String.")
}

func main() {
	// cmdLine.Parse(os.Args[1:])
	flag.Parse()
	//fmt.Printf("Hello, %s.", *name)
	// fmt.Printf("This question is %s", *question)
}
