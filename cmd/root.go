/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	myping_config "github.com/Juminiy/myping/config"
	"github.com/Juminiy/myping/ping_option"
	"github.com/Juminiy/myping/utils"

	"github.com/spf13/cobra"
)

const (
	MYPING_DEFAULT_SHORT = "💡 Software info, author info, dev info."
	MYPING_DEFAULT_LONG  = `😇 The myping is a CLI ( Command Line Interface ) software for ping,
	realizes all the ping's funciton regardless of OS,for supporting the Cross-Platform Operating System : MacOS,Linux various Distributions and Windows,
	the author of the software is Chisato, whose realname is Hua Lingnan , student ID number is 20197897, learned in in NEUQ CS major.`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myping",
	Short: MYPING_DEFAULT_SHORT,
	Long:  MYPING_DEFAULT_LONG,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😋 😍 🥰 You have installed myping successfully !")
		fmt.Println("😛 😜 😝 Please type 'myping help' for usage !")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myping.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(myping_config.AppConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(authorCmd)
	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(versionCmd)
	infoCmd.AddCommand(authorCmd)

	rootCmd.AddCommand(showCmd)

	rootCmd.AddCommand(pingCmd)
	pingCmd.Flags().IntVarP(&PingCount, "counts", "c", PING_DEFAULT_COUNT, "ping times count")
	pingCmd.Flags().StringVarP(&PingUrl, "url", "u", PING_DEFAULT_URL, "ping destination url")
	pingCmd.Flags().BoolVarP(&PingRecourd, "record", "r", PING_DEFAULT_RECORD, "ping whether record in redis")
	pingCmd.Flags().Int64VarP(&utils.DelayTime, "spinnerTime", "t", utils.DEFAULT_DELAY_TIME, "waiting ping spinner delay")
	pingCmd.Flags().StringVarP(&utils.TypeSpinner, "spinnerLogo", "l", utils.MazeSpinner, "waiting ping spinner logo")
	pingCmd.Flags().BoolVarP(&ping_option.ShowEachTime, "showEachTime", "s", ping_option.PING_SHOW_EACH_TIME, "show the every packet details")
	pingCmd.Flags().StringVarP(&ping_option.PingIpVersion, "IPVersion", "v", ping_option.PING_CMD_IP_VERSION_4, "choose the ip version")
}
