package cmd

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch the video for the proverb",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		randomProverb := proverbs.Random()
		fmt.Println("Opening proverb: ", randomProverb.Saying)
		open.Run(randomProverb.Link)

	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
