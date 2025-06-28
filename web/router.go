package web

import (
    "github.com/gin-gonic/gin"
)

//RunServer 启动 Gin Web 服务，并配置路由
func RunServer() {
	// 创建 Gin 引擎（带 Logger 和 Recovery 中间件）
    r := gin.Default()

    // API 路由分组
    api := r.Group("/api")
    {
		// 获取股票数据
        api.GET("/data", GetStockData)
		// 策略回测
        api.POST("/backtest", RunBacktest)
		// 图表接口
        api.GET("/chart", GetChart)
    }

    // 静态托管前端构建产物
    r.Static("/assets", "./web/views/assets")
	
    // SPA fallback: 任何非 /api /assets 的路径都返回 index.html
    r.NoRoute(func(c *gin.Context) {
        c.File("./web/views/index.html")
    })

	// 启动 Web 服务，监听 8080 端口
    r.Run(":8080")
}
