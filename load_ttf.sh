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

package="general"                    # 生成的go文件的package
input="resources/font/LCD_Solid.ttf" # 字体文件
output="general/resource_ttf.go"     # 生成的go文件

go-bindata -pkg $package -o $output $input
