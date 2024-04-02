/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints proverbs",
	Long: `Nifty tool that prints the Proverbs in your terminal.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. count Flag abfragen
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			// 1a. Return mit Fehler
			return err
		}

		// 2. Gewünschte Anzahl Proverbs ausgeben
		for i := 0; i < count; i++ {
			fmt.Println(proverbs.Random().Saying)
		}

		// 3. Return ohne Fehler
		return nil
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// count Flag definieren
	printCmd.Flags().IntP("count", "c", 1, "Count of proverbs to print.")
}

// Here you will define your flags and configuration settings.

// Cobra supports Persistent Flags which will work for this command
// and all subcommands, e.g.:
// printCmd.PersistentFlags().String("foo", "", "A help for foo")

// Cobra supports local flags which will only run when this command
// is called directly, e.g.:
