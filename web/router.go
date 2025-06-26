package web

import (
    "github.com/gin-gonic/gin"
)

func RunServer() {
    r := gin.Default()

    r.LoadHTMLGlob("web/views/*")
    r.Static("/static", "./web/static")

    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })

    api := r.Group("/api")
    {
        api.GET("/data", GetStockData)
        api.POST("/backtest", RunBacktest)
        api.GET("/chart", GetChart)
    }

    r.Run(":8080")
}
