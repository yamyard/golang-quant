package data

import (
    "errors"
    "github.com/yamyard/golang-finance/yahoo"
)

// FetchYahooData 通过调用 golang-finance 的 Yahoo 模块获取历史收盘价（例如近30个交易日）
func FetchYahooData(symbol string) ([]float64, error) {
    // 这里 interval="1d"（日线），rangeStr="30d"（近30天）
    resp, err := yahoo.GetHistory(symbol, "1d", "30d")
    if err != nil {
        return nil, err
    }
    // 解析收盘价
    if len(resp.Chart.Result) == 0 || len(resp.Chart.Result[0].Indicators.Quote) == 0 {
        return nil, errors.New("无有效历史数据")
    }
    closes := resp.Chart.Result[0].Indicators.Quote[0].Close
    // 移除为nil的收盘价（Yahoo偶尔数据有缺失，返回为nil）
    res := make([]float64, 0, len(closes))
    for _, v := range closes {
        if v != 0 {
            res = append(res, v)
        }
    }
    if len(res) == 0 {
        return nil, errors.New("历史数据全部为0")
    }
    return res, nil
}