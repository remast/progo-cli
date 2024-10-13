package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	// 1. Flags definieren
	// a) print Subcommand
	printCommand := flag.NewFlagSet("print", flag.ExitOnError)

	var count int
	printCommand.IntVar(&count, "count", 1, "Count of proverbs to print.")

	var jsonFormat bool
	printCommand.BoolVar(&jsonFormat, "json", false, "Count of proverbs to print.")

	// a) watch Subcommand
	watchCommand := flag.NewFlagSet("watch", flag.ExitOnError)

	// 2. Flags parsen
	flag.Parse()

	subcommand := flag.Args()[0]

	switch subcommand {
	case "print":
		printCommand.Parse(flag.Args()[1:])
		runPrintCommand(count, jsonFormat)
	case "watch":
		watchCommand.Parse(flag.Args()[1:])
		runWatchCommand()
	default:
		log.Fatal("No matching subcommand found.")
	}

}

func runPrintCommand(count int, jsonFormat bool) {
	for i := 0; i < count; i++ {
		if jsonFormat {
			writer := bytes.NewBufferString("")
			err := json.NewEncoder(writer).Encode(proverbs.Random())
			if err != nil {
				log.Fatalf("Could not encode json (%v)", err)
			}
			fmt.Println(writer)
		} else {
			fmt.Println(proverbs.Random().Saying)
		}
	}
}

func runWatchCommand() {

}
