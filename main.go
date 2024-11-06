package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	// 1. Flags definieren
	var count int
	flag.IntVar(&count, "count", 1, "Count of proverbs to print.")

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "verbose output")

	var jsonFormat bool
	flag.BoolVar(&jsonFormat, "json", false, "Count of proverbs to print.")

	// 2. Flags parsen
	flag.Parse()

	if verbose {
		fmt.Fprintf(os.Stdout, "Printing %v proverbs:\n", count)
	}

	// 3. Gewünschte Anzahl Proverbs ausgeben
	for range count {
		if jsonFormat {
			err := json.NewEncoder(os.Stdout).Encode(proverbs.Random())
			if err != nil {
				log.Fatalf("Could not encode json (%v)", err)
			}
		} else {
			fmt.Println(proverbs.Random().Saying)
		}
	}

}
