# 基于 GIN 设计的 API 框架

    基于 Gin 进行模块化设计的 API 框架，封装了常用功能，使用简单，致力于进行快速的业务研发。

### 热编译

如果你想在测试时候不每次都执行 <code>go run main.go</code> 。可以安装 <code>gowatch</code>, 然后项目根目录下运行 <code>gowatch</code> 命令就可以了。

安装 <code>gowatch</code>:
```shell
    go get github.com/silenceper/gowatch
```

命令行参数
> -o : 非必须，指定build的目标文件路径
> -p : 非必须，指定需要build的package（也可以是单个文件）
例子:

    gowatch -o ./bin/go-gin-api -p ./cmd/go-gin-api