package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jboursiquot/go-proverbs"
	"github.com/spf13/cobra"
)

var jsonOutput bool

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints proverbs",
	Long:  `Nifty tool that prints the Proverbs in your terminal.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. count Flag abfragen
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			// 1a. Return mit Fehler
			return err
		}

		// 2. Gewünschte Anzahl Proverbs ausgeben
		for i := 0; i < count; i++ {
			if jsonOutput {
				writer := bytes.NewBufferString("")
				json.NewEncoder(writer).Encode(proverbs.Random())
				fmt.Println(writer)

			} else {
				fmt.Println(proverbs.Random().Saying)
			}
		}

		// 3. Return ohne Fehler
		return nil
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Flags
	printCmd.Flags().IntP("count", "c", 1, "Count of proverbs to print.")
	printCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Count of proverbs to print.")
}
