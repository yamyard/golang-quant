package web

import (
    "github.com/gin-gonic/gin"
    "golang-quant/data"
    "golang-quant/strategy"
    "golang-quant/backtest"
    "net/http"
)

func GetStockData(c *gin.Context) {
    symbol := c.Query("symbol")
    prices, _ := data.FetchYahooData(symbol)
    c.JSON(http.StatusOK, gin.H{
        "symbol": symbol,
        "prices": prices,
    })
}

func RunBacktest(c *gin.Context) {
    symbol := c.DefaultQuery("symbol", "AAPL")
    prices, _ := data.FetchYahooData(symbol)
    signals := strategy.GenerateRSISignal(prices)
    finalValue := backtest.RunBacktest(prices, signals)
    c.JSON(http.StatusOK, gin.H{
        "symbol":      symbol,
        "final_value": finalValue,
    })
}

func GetChart(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "chart": "（此处预留图表接口）",
    })
}
