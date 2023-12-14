/*
File: start.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-13 11:38:18

Description: 程序子命令'start'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yhyj/kbdstage/cli"
	"github.com/yhyj/kbdstage/general"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start keyboard input interception",
	Long:  `Start the kbdstage keyboard input interceptor.`,
	Run: func(cmd *cobra.Command, args []string) {
		if general.Platform == "linux" {
			if general.GetVariable("DISPLAY") != "" {
				cli.Start()
			} else {
				fmt.Println("Could not connect to display")
			}
		} else {
			fmt.Println("Current platform is not supported")
		}
	},
}

func init() {
	startCmd.Flags().BoolP("help", "h", false, "help for start command")
	rootCmd.AddCommand(startCmd)
}
