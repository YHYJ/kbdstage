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

- 编译当前平台可执行文件：

```bash
go build main.go
```

- **交叉编译**指定平台可执行文件：

```bash
# 适用于Linux AArch64平台
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags='-s -w' -trimpath -o eniac main.go
```

```bash
# 适用于macOS amd64平台
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags='-s -w' -trimpath -o eniac main.go
```

```bash
# 适用于Windows amd64平台
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags='-s -w' -trimpath -o eniac main.go
```


## Screenshot

![Alt text](screenshots/1.png?raw=true "1.png")
