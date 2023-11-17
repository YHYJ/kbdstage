<h1 align="center">Kbdstage</h1>

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-07-14 15:17:53 -->

---

<p align="center">
  <a href="https://github.com/YHYJ/kbdstage/actions/workflows/release.yml"><img src="https://github.com/YHYJ/kbdstage/actions/workflows/release.yml/badge.svg" alt="Go build and release by GoReleaser"></a>
</p>

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Install](#install)
  * [一键安装](#一键安装)
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

## Install

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/kbdstage/main/install.sh | sudo bash -s
```

## Usage

- `start`子命令

  开始拦截键盘输入

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/kbdstage/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/kbdstage/general.BuildTime=`date +%s` -X github.com/yhyj/kbdstage/general.BuildBy=$USER" -o build/kbdstage main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/kbdstage/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/kbdstage/general.BuildTime=`date +%s` -X github.com/yhyj/kbdstage/general.BuildBy=$USER" -o build/kbdstage main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/kbdstage/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/kbdstage/general.BuildTime=`date +%s` -X github.com/yhyj/kbdstage/general.BuildBy=$USER" -o build/kbdstage main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### Windows

```powershell
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -H windowsgui -X github.com/yhyj/kbdstage/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/kbdstage/general.BuildTime=`date +%s` -X github.com/yhyj/kbdstage/general.BuildBy=$USER" -o build/kbdstage.exe main.go
```

> 使用`echo %PROCESSOR_ARCHITECTURE%`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

## Screenshot

![Screenshot](resources/screenshots/1.png)
