#!/bin/bash
#!/bin/bash
# 需要停止的程序名称
programname=ginbase
# 启动程序
./${programname} >> logs/service.log 2>&1 &