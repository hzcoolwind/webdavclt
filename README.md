# webdavclt

这是一个远程删除webdav文件服务器上指定文件或者指定目录下的所有文件
的客户端软件，应用例子为在计划任务中定时删除webdav上某临时目录里的临时文件。

# 使用方法
./webdavclt -h http://192.168.1.10:5244/dav -u guest -p guestpwd -c DELETE -d mydir/temp -f "*"

-c 命令 目前支持DELETE 

-d 目录 需要处理的目录

-f 文件 需要处理的文件，如需删除目录下所有文件“*"

