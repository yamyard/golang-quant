package broker

type Order struct {
    Type   string  // 订单类型，"BUY"或"SELL"
    Price  float64 // 下单价格
    Amount int     // 买/卖数量
}

type PaperBroker struct {
    Cash   float64 // 剩余现金
    Shares int     // 当前持有股票数量
}

func NewPaperBroker(initialCash float64) *PaperBroker {
    return &PaperBroker{
        Cash:   initialCash, // 用指定初始现金创建券商
        Shares: 0,           // 初始持仓为0
    }
}

func (b *PaperBroker) Execute(order Order) {
    switch order.Type { // 判断订单类型
    case "BUY":
        cost := order.Price * float64(order.Amount) // 买入所需资金
        if b.Cash >= cost {                         // 判断现金是否充足
            b.Cash -= cost        // 扣除买入资金
            b.Shares += order.Amount // 增加持仓数量
        }
    case "SELL":
        if b.Shares >= order.Amount { // 判断持仓是否足够
            b.Cash += order.Price * float64(order.Amount) // 卖出获得现金
            b.Shares -= order.Amount      // 减少持仓数量
        }
    }
}

// 按当前价格计算总资产（现金+股票市值）
func (b *PaperBroker) TotalValue(currentPrice float64) float64 {
    return b.Cash + float64(b.Shares)*currentPrice
}