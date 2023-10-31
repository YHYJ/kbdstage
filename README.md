# README

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-07-14 15:17:53 -->

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Usage](#usage)
* [Compile](#compile)
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)
    * [macOS](#macos)
    * [Windows](#windows)
* [Screenshot](#screenshot)

<!-- vim-markdown-toc -->

---

<!------------------------------------------------>
<!--  _    _         _     _                    -->
<!-- | | _| |__   __| |___| |_ __ _  __ _  ___  -->
<!-- | |/ / '_ \ / _` / __| __/ _` |/ _` |/ _ \ -->
<!-- |   <| |_) | (_| \__ \ || (_| | (_| |  __/ -->
<!-- |_|\_\_.__/ \__,_|___/\__\__,_|\__, |\___| -->
<!--                                |___/       -->
<!------------------------------------------------>

---

一个键盘输入拦截器，用来擦键盘的

## Usage

- `start`子命令

    开始拦截键盘输入

- `version`子命令

    查看程序版本信息

- `help`

    查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/kbdstage/function.buildTime=`date +%s` -X github.com/yhyj/kbdstage/function.buildBy=$USER" -o kbdstage main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/kbdstage/function.buildTime=`date +%s` -X github.com/yhyj/kbdstage/function.buildBy=$USER" -o kbdstage main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/kbdstage/function.buildTime=`date +%s` -X github.com/yhyj/kbdstage/function.buildBy=$USER" -o kbdstage main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

#### Windows

```powershell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -H windowsgui -X github.com/yhyj/kbdstage/function.buildTime=`date +%s` -X github.com/yhyj/kbdstage/function.buildBy=$USER" -o kbdstage main.go
```

> 使用`echo %PROCESSOR_ARCHITECTURE%`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

## Screenshot

![Alt text](screenshots/1.png?raw=true "1.png")
