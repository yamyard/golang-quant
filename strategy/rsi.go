package strategy

// GenerateRSISignal 根据价格序列生成RSI买卖信号
// "BUY"：RSI < 30，"SELL"：RSI > 70，其他为 "HOLD"
func GenerateRSISignal(prices []float64) []string {
    period := 14
    signals := make([]string, len(prices))
    rsi := calcRSI(prices, period)
    for i := range prices {
        if i < period {
            signals[i] = "HOLD"
        } else if rsi[i] < 30 {
            signals[i] = "BUY"
        } else if rsi[i] > 70 {
            signals[i] = "SELL"
        } else {
            signals[i] = "HOLD"
        }
    }
    return signals
}

// 计算RSI指标
func calcRSI(prices []float64, period int) []float64 {
    n := len(prices)
    rsi := make([]float64, n)
    if n <= period {
        return rsi
    }
    var gain, loss float64
    // 计算初始平均涨跌幅
    for i := 1; i <= period; i++ {
        diff := prices[i] - prices[i-1]
        if diff > 0 {
            gain += diff
        } else {
            loss -= diff
        }
    }
    avgGain := gain / float64(period)
    avgLoss := loss / float64(period)
    if avgLoss == 0 {
        rsi[period] = 100
    } else {
        rs := avgGain / avgLoss
        rsi[period] = 100 - 100/(1+rs)
    }
    // 递推平滑
    for i := period + 1; i < n; i++ {
        diff := prices[i] - prices[i-1]
        var currGain, currLoss float64
        if diff > 0 {
            currGain = diff
            currLoss = 0
        } else {
            currGain = 0
            currLoss = -diff
        }
        avgGain = (avgGain*float64(period-1) + currGain) / float64(period)
        avgLoss = (avgLoss*float64(period-1) + currLoss) / float64(period)
        if avgLoss == 0 {
            rsi[i] = 100
        } else {
            rs := avgGain / avgLoss
            rsi[i] = 100 - 100/(1+rs)
        }
    }
    return rsi
}
