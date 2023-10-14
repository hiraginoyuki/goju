package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jugo-go",
	Short: "jugo in go experimentation",
	Long:  `"jugo" asdf hogehoge foo bar baz ok ok ok`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
