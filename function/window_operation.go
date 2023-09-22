/*
File: window_operation.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-13 13:42:38

Description: 窗口操作

由于我对键盘事件不是很熟悉，某些代码及其注释并不是很理解
*/

package function

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xrect"
	"github.com/BurntSushi/xgbutil/xwindow"
)

// 返回一个随机的color.RGBA
func randomColorRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 0xff,
	}
}

// 渲染渐变
func renderGradient(X *xgbutil.XUtil, window xproto.Window, width, height int, start_color, end_color color.RGBA, ttf string, message string, size float64) {
	// xgraphics.New()创建一个新的xgraphics.Image
	// img.Rect()绘制一个矩形，前两个参数是Pt(x0, y0)，后两个参数是Pt(x1, y1)
	img := xgraphics.New(X, image.Rect(0, 0, width, height))

	// 计算起始颜色start_color和结束颜色end_color之间的渐变步进长度
	rinc := (0xff * (int(end_color.R) - int(start_color.R))) / width
	ginc := (0xff * (int(end_color.G) - int(start_color.G))) / width
	binc := (0xff * (int(end_color.B) - int(start_color.B))) / width

	// 将渐变应用到图像
	img.ForExp(func(x, y int) (uint8, uint8, uint8, uint8) {
		return uint8(int(start_color.R) + (rinc*x)/0xff),
			uint8(int(start_color.G) + (ginc*x)/0xff),
			uint8(int(start_color.B) + (binc*x)/0xff),
			0xff
	})

	// 将图像设置到窗口
	// XSurfaceSet()（包含CreatePixmap）需要在XDraw()之前调用
	img.XSurfaceSet(window)

	// 渲染消息文本
	renderText(img, ttf, message, size, rand.Intn(width/3), rand.Intn(height-100))

	// 将消息文本写入图像（写入的是缓冲区，需要调用XPaint()才会将所绘制的内容显示在屏幕上）
	img.XDraw()

	// 将XDraw()写入缓冲区的内容绘制在屏幕上
	img.XPaint(window)

	// 绘制完成，释放资源
	img.Destroy()
}

// 渲染文本
func renderText(img *xgraphics.Image, ttf string, text string, size float64, x, y int) {
	// 加载字体文件
	fontData, err := Asset(ttf)
	if err != nil {
		panic(err)
	}

	// 创建一个字体文件阅读器
	reader := bytes.NewReader(fontData)

	// 解析字体
	font, err := xgraphics.ParseFont(reader)
	if err != nil {
		panic(err)
	}

	// 绘制文本
	_, _, err = img.Text(x, y, randomColorRGBA(), size, font, text)
	if err != nil {
		panic(err)
	}

	// 计算文本行最适当的宽度和高度
	secWidth, secHeight := xgraphics.Extents(font, size, text)

	// 绘制文字区域
	bounds := image.Rect(x, y, x+secWidth, y+secHeight)
	img.SubImage(bounds).(*xgraphics.Image).XDraw()
}

// 获取窗口的几何信息
func rawGeometry(X *xgbutil.XUtil, window xproto.Drawable) (xrect.Rect, error) {
	geometry, err := xproto.GetGeometry(X.Conn(), window).Reply()
	if err != nil {
		return nil, err
	}

	return xrect.New(int(geometry.X), int(geometry.Y), int(geometry.Width), int(geometry.Height)), nil
}

// 创建一个窗口
func newWindow(X *xgbutil.XUtil, width, height int, start, end color.RGBA, ttf string, message string, size float64) {
	// 获取当前根窗口
	rootWindow := X.RootWin()
	// 对X连接调用一次keybind.Initialize
	keybind.Initialize(X)

	// 调用keybind.GrabKeyboard拦截指定窗口的键盘输入
	if err := keybind.GrabKeyboard(X, rootWindow); err != nil {
		panic(err)
	}

	// 生成一个新窗口ID
	newWindow, err := xwindow.Generate(X)
	if err != nil {
		panic(err)
	}
	// 使用该ID创建一个新窗口
	if err := newWindow.CreateChecked(rootWindow, 0, 0, width, height, 0); err != nil {
		panic(err)
	}

	// 监听键盘按下和释放事件
	newWindow.Listen(xproto.EventMaskKeyPress, xproto.EventMaskKeyRelease)

	go func() {
		for {
			// 绘制渐变窗口
			renderGradient(X, newWindow.Id, width, height, start, end, ttf, message, size)
			time.Sleep(1 * time.Second)
		}
	}()

	// 调用Map()绘制窗口
	newWindow.Map()

	// 因为在调用Map()之后，窗口会接收到Expose事件，所以需要再次调用Listen()来监听Expose事件
	newWindow.Listen(xproto.EventMaskKeyPress)

	xevent.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			// ModifierString()返回一个字符串，表示按下的修饰键
			modStr := keybind.ModifierString(e.State)
			// LookupString()返回一个字符串，表示按下的键（英文字符串）
			keyStr := keybind.LookupString(X, e.State, e.Detail)
			// 如果按下的是Ctrl-Alt-ESC组合键，则退出（9代表ESC键）
			if e.State&xproto.ModMaskControl > 0 && e.State&xproto.ModMask1 > 0 && e.Detail == 9 {
				caser := cases.Title(language.English)
				if caser.String(modStr) == "Control-Mod1" {
					modStr = "Control-Alt"
				}
				fmt.Printf("%s-%s pressed, exiting...\n", caser.String(modStr), caser.String(keyStr))
				xevent.Quit(X)
			}
		}).Connect(X, rootWindow)

	// 发送一个窗口管理器状态请求信息，请求窗口切换到全屏模式
	if err := ewmh.WmStateReq(X, newWindow.Id, ewmh.StateToggle, "_NET_WM_STATE_FULLSCREEN"); err != nil {
		panic(err)
	}
	// 发送一个窗口管理器状态请求信息，请求窗口切换到最上层
	if err := ewmh.WmStateReq(X, newWindow.Id, ewmh.StateToggle, "_NET_WM_STATE_ABOVE"); err != nil {
		panic(err)
	}
}
