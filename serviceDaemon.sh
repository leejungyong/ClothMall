#!/bin/bash
# 需要停止的程序名称
programname=ginbase
# 目录名称
folder=$programname
# 进入目录
cd /home/jh/GoProject/src/${folder}
# 程序名称对应的进程编号
pid=`ps -ef | grep $programname| grep -v grep | awk '{print $2}'`
echo $programname 的进程编号: $pid
if [ '' == "$pid" ]; then
    echo $programname 未启动
    ./serviceStart.sh
    echo $programname 启动完毕
else
    echo $programname 运行中
fi