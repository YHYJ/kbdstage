#!/usr/bin/env bash

: << !
Name: load_ttf.sh
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-07-14 09:52:31

Description: 加载字体资源为go文件

Attentions:
-

Depends:
-
!

package="function"
input="resource/LCD_Solid.ttf"
output="function/ttf.go"

go-bindata -pkg $package -o $output $input
