 package main
    import (
        "github.com/kataras/iris"
        "github.com/kataras/iris/middleware/logger"
        "github.com/kataras/iris/middleware/recover"
    )
    func main() {
        app := iris.New()
        app.Use(recover.New())
        app.Use(logger.New())
        //输出html
        // 请求方式: GET
        // 访问地址: http://localhost:8080/welcome
        app.Handle("GET", "/welcome", func(ctx iris.Context) {
            ctx.HTML("<h1>Welcome</h1>")
        })
        //输出字符串
        // 类似于 app.Handle("GET", "/ping", [...])
        // 请求方式: GET
        // 请求地址: http://localhost:8080/ping
        app.Get("/", func(ctx iris.Context) {
            ctx.WriteString("hello")
        })
        //输出json
        // 请求方式: GET
        // 请求地址: http://localhost:8080/hello
        app.Get("/hello", func(ctx iris.Context) {
            ctx.JSON(iris.Map{"message": "Hello Iris!"})
        })
        app.Run(iris.Addr(":8090"))//8080 监听端口
    }
