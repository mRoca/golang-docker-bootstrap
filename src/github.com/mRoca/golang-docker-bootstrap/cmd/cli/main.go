package main

import (
	"flag"
	"fmt"
	"github.com/mRoca/golang-docker-bootstrap/internal/cli"
	"math/rand"
	"os"
	"time"
)

func printDefault(flagset *flag.FlagSet) {
	fmt.Printf("\nUsage: cli %s [OPTIONS]\n", flagset.Name())
	flagset.PrintDefaults()
}

func printDefaults(flagsets []*flag.FlagSet) {
	for _, flagset := range flagsets {
		printDefault(flagset)
	}
}

var commands []*flag.FlagSet

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Subcommands
	fooCommand := flag.NewFlagSet("say", flag.ExitOnError)
	commands = append(commands, fooCommand)

	// Import command Flags
	fooTextPtr := fooCommand.String("text", "", "(required) Text to display")
	fooSuffixPtr := fooCommand.String("suffix", "", "Text to append")

	if len(os.Args) < 2 {
		fmt.Println("A subcommand is required")
		printDefaults(commands)
		os.Exit(1)
	}

	// Parse the flags for appropriate FlagSet
	switch os.Args[1] {
	case "say":
		_ = fooCommand.Parse(os.Args[2:])
	default:
		printDefaults(commands)
		os.Exit(1)
	}

	if fooCommand.Parsed() {
		if *fooTextPtr == "" {
			printDefault(fooCommand)
			os.Exit(1)
		}

		displayedString := cli.BuildCliString(*fooTextPtr, *fooSuffixPtr)
		fmt.Println(displayedString)
	}
}
