package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/jboursiquot/go-proverbs"
	"github.com/skratchdot/open-golang/open"
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

	// 2. Subcommands verarbeiten
	subcommand := flag.Args()[0]

	switch subcommand {
	case "print":
		// a) Flags parsen
		printCommand.Parse(flag.Args()[1:])

		// b) Subcommand auführen
		runPrintCommand(count, jsonFormat)
	case "watch":
		// a) Flags parsen
		watchCommand.Parse(flag.Args()[1:])

		// b) Subcommand auführen
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
	open.Run(proverbs.Random().Link)
}
