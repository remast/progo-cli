package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	// 1. count Flag definieren
	var count = flag.IntP("count", "c", 1, "Count of proverbs to print.")
	var verbose = flag.BoolP("verbose", "v", false, "verbose output")

	// 2. Flags parsen
	flag.Parse()

	if *verbose {
		fmt.Fprintf(os.Stdout, "Printing %v proverbs:\n", *count)
	}

	// 3. Gewünschte Anzahl Proverbs ausgeben
	for i := 0; i < *count; i++ {
		fmt.Println(proverbs.Random().Saying)
	}

}
