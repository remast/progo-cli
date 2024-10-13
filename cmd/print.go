package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jboursiquot/go-proverbs"
	"github.com/spf13/cobra"
)

var jsonOutput bool

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints some proverbs",
	Long:  `Nifty tool that prints the Proverbs in your terminal.`,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(2), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. count Flag abfragen
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			// 1a. Return mit Fehler
			return err
		}

		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			// 1a. Return mit Fehler
			return err
		}
		if verbose {
			fmt.Fprintf(os.Stdout, "Printing %v proverbs:\n", count)
		}

		// 2. Gewünschte Anzahl Proverbs ausgeben
		for i := 0; i < count; i++ {
			if jsonOutput {
				writer := bytes.NewBufferString("")
				err := json.NewEncoder(writer).Encode(proverbs.Random())
				if err != nil {
					return err
				}
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
