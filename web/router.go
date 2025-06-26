package web

import (
    "github.com/gin-gonic/gin"
)

/*
RunServer 启动 Gin Web 服务，并配置路由

- 使用 gin.Default() 创建带日志和恢复中间件的 Gin 实例
- r.LoadHTMLGlob("web/views/*") 加载模板文件（如 index.html），用于渲染 HTML 页面
- r.Static("/static", "./web/static") 将本地 ./web/static 目录映射为网站的 /static 路径，供前端静态资源（JS/CSS/图片等）访问

- r.GET("/", ...) 设置网站根路径 GET 请求，渲染 index.html 首页
- 以 /api 为前缀分组 RESTful API 路径，包括：
    - api.GET("/data", GetStockData)：GET 请求获取股票数据
    - api.POST("/backtest", RunBacktest)：POST 请求进行策略回测
    - api.GET("/chart", GetChart)：GET 请求用于获取图表数据（或占位）

- r.Run(":8080") 启动 Web 服务，监听 8080 端口
*/
func RunServer() {
    r := gin.Default() // 创建 Gin 引擎（带 Logger 和 Recovery 中间件）

    r.LoadHTMLGlob("web/views/*") // 加载 HTML 模板文件
    r.Static("/static", "./web/static") // 设置静态文件服务

    // 网站首页，返回 index.html
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })

    // API 路由分组
    api := r.Group("/api")
    {
        api.GET("/data", GetStockData)         // 获取股票数据
        api.POST("/backtest", RunBacktest)     // 策略回测
        api.GET("/chart", GetChart)            // 图表接口
    }

    r.Run(":8080") // 启动 Web 服务，监听 8080 端口
}
