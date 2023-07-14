/*
File: start.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-13 11:38:18

Description: 程序子命令'start'时执行
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/kbdstage/function"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start keyboard input interception",
	Long:  `Start the KbdStage keyboard input interceptor.`,
	Run: func(cmd *cobra.Command, args []string) {
		function.Start()
	},
}

func init() {
	startCmd.Flags().BoolP("help", "h", false, "help for start")
	rootCmd.AddCommand(startCmd)
}
