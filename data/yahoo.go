package data

func FetchYahooData(symbol string) ([]float64, error) {
    // 目前为演示用途，直接返回固定的价格序列
    return []float64{100, 101, 102, 103, 104}, nil
}
