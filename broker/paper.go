package broker

type Order struct {
    // 订单类型，"BUY"或"SELL"
    Type   string
    // 下单价格
    Price  float64
    // 买/卖数量
    Amount int
}

type PaperBroker struct {
    // 剩余现金
    Cash   float64
    // 当前持有股票数量
    Shares int
}

func NewPaperBroker(initialCash float64) *PaperBroker {
    return &PaperBroker{
        // 用指定初始现金创建券商
        Cash:   initialCash,
        // 初始持仓为0
        Shares: 0,
    }
}

func (b *PaperBroker) Execute(order Order) {
    // 判断订单类型
    switch order.Type {
    case "BUY":
        // 买入所需资金
        cost := order.Price * float64(order.Amount)
        // 判断现金是否充足
        if b.Cash >= cost {
            // 扣除买入资金
            b.Cash -= cost
            // 增加持仓数量
            b.Shares += order.Amount
        }
    case "SELL":
        // 判断持仓是否足够
        if b.Shares >= order.Amount {
            // 卖出获得现金
            b.Cash += order.Price * float64(order.Amount)
            // 减少持仓数量
            b.Shares -= order.Amount
        }
    }
}

// 按当前价格计算总资产（现金+股票市值）
func (b *PaperBroker) TotalValue(currentPrice float64) float64 {
    return b.Cash + float64(b.Shares)*currentPrice
}
