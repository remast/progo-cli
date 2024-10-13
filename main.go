package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	// 1. Flags definieren
	var count int
	flag.IntVarP(&count, "count", "c", 1, "Count of proverbs to print.")

	var verbose bool
	flag.BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	var jsonFormat bool
	flag.BoolVarP(&jsonFormat, "json", "j", false, "Count of proverbs to print.")

	// 2. Flags parsen
	flag.Parse()

	if verbose {
		fmt.Fprintf(os.Stdout, "Printing %v proverbs:\n", count)
	}

	// 3. Gewünschte Anzahl Proverbs ausgeben
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
