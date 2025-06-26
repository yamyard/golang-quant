package strategy

func GenerateRSISignal(prices []float64) []string {
    signals := make([]string, len(prices))
    for i := range prices {
        signals[i] = "HOLD"
    }
    return signals
}
