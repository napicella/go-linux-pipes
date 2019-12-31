package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "uppercase",
	Short: "Transform the input to uppercase letters",
	Long: `Simple demo of the usage of linux pipes
Transform the input (pipe of file) to uppercase letters`,
	RunE: func(cmd *cobra.Command, args []string) error {
		print = logNoop
		if flags.verbose {
			print = logOut
		}
		return runCommand()
	},
}

var flags struct{
	filepath string
	verbose bool
}

var flagsName = struct{
	file, fileShort string
	verbose, verboseShort string
} {
	"file", "f",
	"verbose", "v",
}

var print func(s string)

func main() {
	rootCmd.Flags().StringVarP(
		&flags.filepath,
		flagsName.file,
		flagsName.fileShort,
		"", "path to the file")
	rootCmd.PersistentFlags().BoolVarP(
		&flags.verbose,
		flagsName.verbose,
		flagsName.verboseShort,
		false, "log verbose output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func logNoop(s string) {}

func logOut(s string) {
	log.Println(s)
}