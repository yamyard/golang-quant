package strategy

func GenerateRSISignal(prices []float64) []string {
    signals := make([]string, len(prices))
    for i, price := range prices {
		if price > 100 {
			signals[i] = "SELL"
		} else {
			signals[i] = "HOLD"
		}
	}
    return signals
}
