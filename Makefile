#定义变量
CURRENT_TIME=`date "+%Y-%m-%d %H:%M:%S"`

#默认执行第一个target
default: build_smallest
	@echo 当前时间: ${CURRENT_TIME}

#编译前的工作
before_build:
	go generate

#编译
build: before_build
	go build

#扔掉可执行文件中的无用信息, 以达到瘦身目的
build_smaller: before_build
	ls -lh ./apiproject
#		-ldflags的参数说明: -s去掉符号信息, -w去掉DWARF调试信息, 所以最终的程序无法使用gdb调试
	go build -ldflags "-s -w"
	ls -lh ./apiproject

#生成体积最小的可执行程序
build_smallest: build_smaller compress

#压缩可执行程序体积
compress:
#	使用压缩工具upx压缩可执行程序
	upx --brute ./apiproject
	ls -lh ./apiproject