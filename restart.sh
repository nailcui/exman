#!/bin/bash

# 关闭当前服务
pid=`ps |grep "exman" |awk '{print $1}'`
kill -9 $pid && echo "服务下线"

export GIN_MODE=release
mkdir -p /logs/exman
touch /logs/exman/info.log
nohup ./exman >> /logs/exman/info.log 2>&1 &