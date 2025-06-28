package web

import (
    "github.com/gin-gonic/gin"
    "golang-quant/data"
    "golang-quant/strategy"
    "golang-quant/backtest"
    "net/http"
)

/*
GetStockData 处理股票数据查询接口

- 通过 GET 请求参数 "symbol" 获取需要查询的股票代码
- 调用 data.FetchYahooData(symbol) 获取历史价格数据（当前为模拟数据）
- 返回格式为 JSON，包含股票代码和价格序列
- 用途：前端页面请求历史行情数据用于展示或分析
*/
func GetStockData(c *gin.Context) {
    symbol := c.Query("symbol")
	
    prices, err := data.FetchYahooData(symbol)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
            "symbol": symbol,
        })
        return
    }
    if len(prices) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No price data available",
            "symbol": symbol,
        })
        return
    }
	
    c.JSON(http.StatusOK, gin.H{
        "symbol": symbol,
        "prices": prices,
    })
}


/*
RunBacktest 处理量化策略回测接口

- 通过 GET 请求参数 "symbol" 获取回测的股票代码
- 获取指定股票的历史价格数据
- 调用 strategy.GenerateRSISignal(prices) 生成买卖信号（当前策略为 RSI 或其他简化策略）
- 使用 backtest.RunBacktest(prices, signals) 执行回测，计算账户最终总资产
- 返回格式为 JSON，包含股票代码和回测后的最终资产
- 用途：前端页面请求回测结果，用于策略性能评估
*/
func RunBacktest(c *gin.Context) {
    symbol := c.Query("symbol")
	
    prices, err := data.FetchYahooData(symbol)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
            "symbol": symbol,
        })
        return
    }
    if len(prices) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No price data available",
            "symbol": symbol,
        })
        return
    }

    signals := strategy.GenerateRSISignal(prices)
    finalValue := backtest.RunBacktest(prices, signals)

    c.JSON(http.StatusOK, gin.H{
        "symbol":      symbol,
        "final_value": finalValue,
        "signals":     signals,
        "prices":      prices,
    })
}


/*
GetChart 图表接口

- 通过 GET 请求参数 "symbol" 获取需要查询的股票代码
- 调用 data.FetchYahooData(symbol) 获取历史价格数据
- 返回格式为 JSON，包含股票代码和“chart”数组，chart为 [{ "index": 1, "close": xxx }, ...]
- 用途：前端请求图表数据用于可视化展示
*/
func GetChart(c *gin.Context) {
    symbol := c.Query("symbol")
	
    prices, err := data.FetchYahooData(symbol)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
            "symbol": symbol,
        })
        return
    }
    if len(prices) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No price data available",
            "symbol": symbol,
        })
        return
    }
	
    chart := make([]gin.H, len(prices))
    for i, v := range prices {
        chart[i] = gin.H{
            "index": i + 1,
            "close": v,
        }
    }
    
	c.JSON(http.StatusOK, gin.H{
        "symbol": symbol,
        "chart":  chart,
    })
}
