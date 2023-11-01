/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-13 11:40:35

Description: 子命令`version`功能函数
*/

package general

import "fmt"

// 程序信息
const (
	name    = "Kbdstage"
	version = "v0.4.0"
	project = "github.com/yhyj/kbdstage"
)

// 编译信息
var (
	gitCommitHash string = "unknown"
	buildTime     string = "unknown"
	buildBy       string = "unknown"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", version)
	if !only {
		programInfo = fmt.Sprintf("%s version: %s\nGit commit hash: %s\nBuilt on: %s\nBuilt by: %s\n", name, version, gitCommitHash, buildTime, buildBy)
	}
	return programInfo
}
