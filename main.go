package main

import (
	"flag"
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	// 1. Flags definieren
	var count int
	flag.IntVar(&count, "count", 1, "Count of proverbs to print.")

	// 2. Flags parsen
	flag.Parse()

	// 3. Gewünschte Anzahl Proverbs ausgeben
	for range count {
		fmt.Println(proverbs.Random().Saying)
	}

}
