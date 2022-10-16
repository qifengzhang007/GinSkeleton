#说明：makefile 文件只能在linux系统运行，windows 系统无法执行本文件定义的相关命令
# 使用文档参考：https://www.yuque.com/xiaofensinixidaouxiang/bkfhct/zso6xo

# 定义 makefile 的命名列表, 只需要将外部调用的公布在这里即可
.PHONY:  build-api build-web    build-cli help

# 设置 cmd/api/main.go 入口文件编译后的可执行文件名
apiBinName="ginskeleton-api.linux64"

# 设置 cmd/web/main.go 入口文件编译后的可执行文件名
webBinName="ginskeleton-web.linux64"

# 设置 cmd/cli/main.go 入口文件编译后的可执行文件名
cliBinName="ginskeleton-cli.linux64"

# 统一设置编译的目标平台公共参数
all:
	go env -w GOARCH=amd64
	go env -w GOOS=linux
	go env -w CGO_ENABLED=0
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod  tidy

build-api:all clean-api build-api-bin
build-api-bin:
	go build -o ${apiBinName}    -ldflags "-w -s"  -trimpath  ./cmd/api/main.go

build-web:all clean-web build-web-bin
build-web-bin:
	go build -o ${webBinName}   -ldflags "-w -s"  -trimpath  ./cmd/web/main.go

build-cli:all clean-cli build-cli-bin
build-cli-bin:
	go build -o ${cliBinName}   -ldflags "-w -s"  -trimpath  ./cmd/cli/main.go

# 编译前清理可能已经存在的旧文件
clean-api:
	@if [ -f ${apiBinName} ] ; then rm -rf ${apiBinName} ; fi
clean-web:
	@if [ -f ${webBinName} ] ; then rm -rf ${webBinName} ; fi
clean-cli:
	@if [ -f ${cliBinName} ] ; then rm -rf ${cliBinName} ; fi

help:
	@echo "make hep 查看编译命令列表"
	@echo "make build-api  编译 cmd/api/main.go 入口文件 "
	@echo "make build-web  编译 cmd/web/main.go 入口文件 "
	@echo "make build-cli  编译 cmd/cli/main.go 入口文件 "