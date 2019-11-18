#!/bin/bash
# 需要停止的程序名称
programname=ginbase
# 程序名称对应的进程编号
tmp=`ps -ef | grep $programname| grep -v grep | awk '{print $2}'`
echo $programname 的进程编号: ${tmp}
# 杀掉进程
kill -9 ${tmp}
echo $programname 已停止