#!/bin/sh

module_name=$1
mkdir src/$module_name
if [ $2 = "main" ];then
    touch src/$module_name/main.go
else
    touch src/$module_name/$module_name.go
fi