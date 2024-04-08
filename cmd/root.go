/*
File: root.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-13 11:40:20

Description: 执行程序
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kbdStage",
	Short: "keyboard input interceptor",
	Long:  `kbdstage intercepts keyboard input, then cleans the keyboard.`,
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
	rootCmd.Flags().BoolP("help", "h", false, "help for kbdstage")
}
