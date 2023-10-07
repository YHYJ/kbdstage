/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-13 11:40:35

Description: 子命令`version`功能函数
*/

package function

import "fmt"

// 程序信息
var (
	name    string = "KbdStage"
	version string = "v0.2.5"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", version)
	if !only {
		programInfo = fmt.Sprintf("%s version %s\n", name, version)
	}
	return programInfo
}
