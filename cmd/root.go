package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/brokeyourbike/doccano-to-automl-jsonl/models"
	"github.com/spf13/cobra"
)

var outputPath string

var rootCmd = &cobra.Command{
	Use:   "doccano-to-automl-jsonl",
	Short: "Transfrorm doccano JSONL to the format expected by AutoML.",
	Long:  `Transfrorm doccano JSONL to the format expected by AutoML.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := args[0]

		if outputPath == "" {
			outputPath = path.Base(inputPath)
		}

		inputFile, err := os.Open(inputPath)
		if err != nil {
			return err
		}
		defer inputFile.Close()

		outputFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		fileScanner := bufio.NewScanner(inputFile)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			var d models.Doccano
			if err := json.Unmarshal([]byte(fileScanner.Text()), &d); err != nil {
				log.Fatal(err)
			}

			a, err := json.Marshal(d.Convert())
			if err != nil {
				log.Fatal(err)
			}

			outputFile.WriteString(fmt.Sprintf("%s\n", string(a)))
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file")
}
