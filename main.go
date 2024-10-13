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
	// 1. count Flag definieren
	var count = flag.Int("count", 1, "Count of proverbs to print.")
	var verbose = flag.Bool("verbose", false, "verbose output")
	var jsonFormat bool
	flag.BoolVar(&jsonFormat, "json", false, "Count of proverbs to print.")

	// 2. Flags parsen
	flag.Parse()

	if *verbose {
		fmt.Fprintf(os.Stdout, "Printing %v proverbs:\n", count)
	}

	// 3. Gewünschte Anzahl Proverbs ausgeben
	for i := 0; i < *count; i++ {
		if jsonFormat {
			writer := bytes.NewBufferString("")
			json.NewEncoder(writer).Encode(proverbs.Random())
			fmt.Println(writer)
		} else {
			fmt.Println(proverbs.Random().Saying)
		}
	}

}
