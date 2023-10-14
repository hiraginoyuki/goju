package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string

	greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %v!\n", name)
			if len(args) > 0 {
				fmt.Println(args[0])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(greetCmd)

	greetCmd.Flags().StringVar(&name, "name", "", "Source directory to read from")
	greetCmd.MarkFlagRequired("name")
}
