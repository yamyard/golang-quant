package broker

type Order struct {
    Type   string
    Price  float64
    Amount int
    Time   string
}

type Position struct {
    Cash     float64
    Holdings int
    History  []Order
}

func NewPaperBroker(initialCash float64) *Position {
    return &Position{
        Cash:     initialCash,
        Holdings: 0,
        History:  []Order{},
    }
}

func (p *Position) Execute(order Order) {
    if order.Type == "BUY" {
        cost := order.Price * float64(order.Amount)
        if p.Cash >= cost {
            p.Cash -= cost
            p.Holdings += order.Amount
            p.History = append(p.History, order)
        }
    } else if order.Type == "SELL" {
        if p.Holdings >= order.Amount {
            p.Cash += order.Price * float64(order.Amount)
            p.Holdings -= order.Amount
            p.History = append(p.History, order)
        }
    }
}

func (p *Position) TotalValue(currentPrice float64) float64 {
    return p.Cash + float64(p.Holdings)*currentPrice
}
