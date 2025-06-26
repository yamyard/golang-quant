package strategy

// GenerateRSISignal 根据价格序列生成买卖信号
// 输入: prices []float64 —— 股票或标的的历史价格序列
// 输出: []string —— 与输入价格等长的信号序列，"SELL" 表示价格高于100, "HOLD" 表示价格不高于100
func GenerateRSISignal(prices []float64) []string {
    // 创建与价格序列长度相同的字符串切片，用于存储每个点的信号
    signals := make([]string, len(prices))
    // 遍历所有价格
    for i, price := range prices {
        // 如果价格大于100，发出 "SELL" 信号
        if price > 100 {
            signals[i] = "SELL"
        // 否则发出 "HOLD" 信号
	} else {
            signals[i] = "HOLD"
        }
    }
    // 返回信号序列
    return signals
}
