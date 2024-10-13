package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
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
	for i := 0; i < count; i++ {
		if jsonFormat {
			writer := bytes.NewBufferString("")
			json.NewEncoder(writer).Encode(proverbs.Random())
			fmt.Println(writer)
		} else {
			fmt.Println(proverbs.Random().Saying)
		}
	}

}
