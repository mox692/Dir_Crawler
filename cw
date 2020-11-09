#!/bin/sh

# userの入力: cw jump test.txt
action=$1
kw=$2

# 実行
result=$(./dirWalk --$action=$kw)
status=$? ## 必須
# jumpの時だけ、cdする
if [ "$action" = "jump" ]; then
    if [ "$result" = "" ]; then
        return
    # 以上終了の時    
    elif [ $status != "0" ]; then  # !=だとうまくいかない...。
    echo "$result"
    # 正常時
    else
    cd $result
    fi
else
    echo  "$result"
fi