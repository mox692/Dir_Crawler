#!/bin/sh

# userの入力: cw jump test.txt
action=$1
kw=$2

# 実行
result=$(./dirWalk --$action=$kw)

# jumpの時だけ、cdする
if [ "$action" = "jump" ]; then
    echo $result
    if [ "$result" = "" ]; then
        return 0
    fi 
    # 抜き出し
    # arr=(${result// / })
    # echo ${arr}
    cd $result
else
    echo $result
fi

# jump以外の時は、そのまま出力