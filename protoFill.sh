#!/bin/bash

#!/bin/bash



# 获取当前运行所在目录
current_dir="$( pwd )"

# 查找并输出所有的.proto文件名
proto_files=$(find . -type f -name "*.proto")

for file in $proto_files; do
  # 提取文件名部分
  filename=$(basename "$file")
  name="${filename%.proto}"
  # 获取变量值的首字母
  first_letter="${name:0:1}"

  # 将首字母转换为大写
  first_letter_upper="$(echo "$first_letter" | tr '[:lower:]' '[:upper:]')"

  # 获取变量值的剩余部分
  rest_of_string="${name:1}"

  # 将首字母大写和剩余部分组合起来
  name="$first_letter_upper$rest_of_string"
done

protoPath=$current_dir/$filename

echo $protoPath

go run /Users/backend001/go/src/goshell/main.go -path=$protoPath -service=$name

echo "goctl rpc protoc *.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style goZero --home ../../../common/GoctlTemplate"

goctl rpc protoc *.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style goZero --home ../../../common/GoctlTemplate -m





