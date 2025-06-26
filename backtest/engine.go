package backtest

import "golang-quant/broker"

// RunBacktest 用于回测交易策略，输入价格序列和信号序列，返回最终总资产
// prices: 历史价格序列
// signals: 与价格序列等长的信号序列，"BUY" 或 "SELL"
func RunBacktest(prices []float64, signals []string) float64 {
    // 创建初始资金为 10000 的模拟券商账户
    b := broker.NewPaperBroker(10000)
    // 遍历价格和信号序列
    for i := range prices {
        if signals[i] == "BUY" {
            // 收到买入信号，按当前价格买入 1 单位
            b.Execute(broker.Order{Type: "BUY", Price: prices[i], Amount: 1})
        } else if signals[i] == "SELL" {
            // 收到卖出信号，按当前价格卖出 1 单位
            b.Execute(broker.Order{Type: "SELL", Price: prices[i], Amount: 1})
        }
        // 其他信号不做操作
    }
    // 返回当前总资产价值（现金+最后价格下的持仓市值）
    return b.TotalValue(prices[len(prices)-1])
}