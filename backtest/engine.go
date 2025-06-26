package backtest

import "golang-quant/broker"

func RunBacktest(prices []float64, signals []string) float64 {
    b := broker.NewPaperBroker(10000)
    for i := range prices {
        if signals[i] == "BUY" {
            b.Execute(broker.Order{Type: "BUY", Price: prices[i], Amount: 1})
        } else if signals[i] == "SELL" {
            b.Execute(broker.Order{Type: "SELL", Price: prices[i], Amount: 1})
        }
    }
    return b.TotalValue(prices[len(prices)-1])
}
