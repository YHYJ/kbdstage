/*
File: root.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-13 11:40:20

Description: 程序未带子命令或参数时执行
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kbdStage",
	Short: "keyboard input interceptor",
	Long:  `KbdStage intercepts keyboard input, then cleans the keyboard.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for KbdStage")
}
