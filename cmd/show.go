/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "📈 show the ping records",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📈 We are going to show the view of ping records")
	},
}
