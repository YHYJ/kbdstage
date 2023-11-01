/*
File: start.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-14 13:23:00

Description: 子命令`start`功能函数
*/

package cli

import (
	"math/rand"
	"time"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/yhyj/kbdstage/general"
)

var (
	// 界面文本字体
	ttf = "resource/LCD_Solid.ttf"
	// 界面文本字体大小
	fontSize = 0.0
	// 界面文本字体放大倍率
	fontScale = 100.0
	// 界面文本
	message = "Capturing keyboard input, type Control-Alt-ESC to exit."
)

func Start() {
	// 初始化随机数种子
	rand.New(rand.NewSource(time.Now().UnixNano()))

	X, err := xgbutil.NewConn()
	if err != nil {
		panic(err)
	}

	// 创建一个随机渐变色的窗口
	geometry, err := general.RawGeometry(X, xproto.Drawable(X.RootWin()))
	if err != nil {
		panic(err)
	}

	// 根据窗口大小计算字体大小
	fontSize = float64(geometry.Height()) / float64(geometry.Width()) * fontScale

	general.NewWindow(X, geometry.Width(), geometry.Height(), general.RandomColorRGBA(), general.RandomColorRGBA(), ttf, message, fontSize)

	// 监听键盘事件
	xevent.Main(X)
}
