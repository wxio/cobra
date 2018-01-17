package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:                "sub",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sub called with args %v", args)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
