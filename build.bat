:: 编译依赖--必需的，不然运行不了
rsrc -manifest test.exe.manifest -o rsrc.syso
:: 打包程序
go build -ldflags="-H windowsgui" -o builds\go-epub.exe
:: 运行
.\builds\go-epub.exe