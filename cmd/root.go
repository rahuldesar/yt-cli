package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yt-cli",
	Short: "ytcli short desc",
	Long:  `ytcli long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this should run bubbletea?")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
