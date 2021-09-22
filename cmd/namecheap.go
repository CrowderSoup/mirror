package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// namecheapCmd represents the namecheap command
var namecheapCmd = &cobra.Command{
	Use:   "namecheap",
	Short: "Mirror your Dynamic IP in NameCheap DNS",
	Long: `Ensure that your DNS in NameCheap stays in sync with your dynamic IP
address.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(namecheapCmd)

	// domain flag
	namecheapCmd.Flags().StringP(
		"domain",
		"d",
		"",
		"The domain you want to keep in sync",
	)

	// password flag
	namecheapCmd.Flags().StringP(
		"password",
		"p",
		"",
		"The unique password for updating your DNS with NameCheap",
	)
}
