package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doccano-to-automl-jsonl",
	Short: "Transfrorm doccano JSONL to the format expected by AutoML.",
	Long:  `Transfrorm doccano JSONL to the format expected by AutoML.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// Do Stuff Here
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
