#!/bin/bash

apiName=$1.api

# 获取当前运行所在目录
current_dir="$( pwd )"

apiPath=$current_dir/doc/$apiName

go run /Users/backend001/go/src/goshell/tool/apigen/apigen.go -path=$apiPath

echo "done!"