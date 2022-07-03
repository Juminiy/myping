/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "The myping app information",
	Long:  MYPING_DEFAULT_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😉 😙 🤩 😎 😆 The myping app information lists : ")
		fmt.Println("😉 App's Name: ", viper.GetString("app.name"))
		fmt.Println("😙 App's Version: ", viper.GetString("app.version"))
		fmt.Println("🤩 App's Author: ", viper.GetString("app.author.name"))
		fmt.Println("😎 App's gitRepoUrl: ", viper.GetString("app.gitRepo.url"))
		fmt.Println("😆 App's created: ", viper.GetString("app.createdTime"))
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The myping version and publishment",
	Long:  MYPING_DEFAULT_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🐵 App' Version: ", viper.GetString("app.version"))
	},
}

var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "The myping author and creator",
	Long:  MYPING_DEFAULT_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🤠 App' Author Name : ", viper.GetString("app.author.name"))
		fmt.Println("🏔️  App' Author Own : ", viper.GetBool("app.author.original"))
		fmt.Println("🔥 App' Author Github: ", viper.GetString("app.author.github"))
	},
}

