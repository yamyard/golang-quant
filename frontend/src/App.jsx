import { useState } from 'react';
import ReactECharts from 'echarts-for-react';

function App() {
  const [symbol, setSymbol] = useState('');
  const [priceResult, setPriceResult] = useState('');
  const [backtestResult, setBacktestResult] = useState('');
  const [chartData, setChartData] = useState([]);
  const [chartLoading, setChartLoading] = useState(false);
  const [showChart, setShowChart] = useState(false);

  const getPrice = async () => {
    if (!symbol.trim()) {
      setPriceResult('请输入股票代码');
      return;
    }
    setPriceResult('正在获取价格...');
    try {
      const resp = await fetch('/api/data?symbol=' + encodeURIComponent(symbol.trim().toUpperCase()));
      let data;
      try {
        data = await resp.json();
      } catch (err) {
        throw new Error('服务器返回异常格式');
      }
      if (!resp.ok) {
        setPriceResult('请求失败: ' + (data.error || ('网络错误：' + resp.status)));
        return;
      }
      setPriceResult(JSON.stringify(data, null, 2));
    } catch (e) {
      setPriceResult('请求失败: ' + e.message);
    }
  };

  const runBacktest = async () => {
    if (!symbol.trim()) {
      setBacktestResult('请输入股票代码');
      return;
    }
    setBacktestResult('正在运行回测...');
    try {
      const resp = await fetch('/api/backtest?symbol=' + encodeURIComponent(symbol.trim().toUpperCase()), { method: 'POST' });
      let data;
      try {
        data = await resp.json();
      } catch (err) {
        throw new Error('服务器返回异常格式');
      }
      if (!resp.ok) {
        setBacktestResult('请求失败: ' + (data.error || ('网络错误：' + resp.status)));
        return;
      }

      // 组装结果文本
      let output = `股票: ${data.symbol}\n最终资产: ${data.final_value}\n\n信号:\n`;
      if (data.signals && data.prices) {
        for (let i = 0; i < data.signals.length; i++) {
          output += `Day ${i + 1}: Price=${data.prices[i]}, Signal=${data.signals[i]}\n`;
        }
      } else {
        output += JSON.stringify(data, null, 2);
      }
      setBacktestResult(output);
    } catch (e) {
      setBacktestResult('请求失败: ' + e.message);
    }
  };

  const showPriceChart = async () => {
    if (!symbol.trim()) {
      setChartData([]);
      setShowChart(true);
      return;
    }
    setChartLoading(true);
    setShowChart(true);
    try {
      const resp = await fetch('/api/chart?symbol=' + encodeURIComponent(symbol.trim().toUpperCase()));
      let data;
      try {
        data = await resp.json();
      } catch (err) {
        throw new Error('服务器返回异常格式');
      }
      if (!resp.ok) {
        setChartData([]);
        setChartLoading(false);
        return;
      }
      if (Array.isArray(data.chart)) {
        setChartData(data.chart);
      } else {
        setChartData([]);
      }
    } catch (e) {
      setChartData([]);
    }
    setChartLoading(false);
  };

  const priceChartOption = {
    title: { text: '收盘价折线图' },
    tooltip: { trigger: 'axis' },
    xAxis: {
      type: 'category',
      data: chartData.map((item, i) => `Day ${i + 1}`),
      name: '日期'
    },
    yAxis: { type: 'value', name: '收盘价' },
    series: [
      {
        data: chartData.map(item => item.close),
        type: 'line',
        smooth: true,
        name: '收盘价'
      }
    ]
  };

  return (
    <div style={{ padding: '1rem', fontFamily: 'sans-serif' }}>
      <h1>欢迎使用 Golang 量化交易平台</h1>

      <div style={{ marginBottom: '1rem' }}>
        <label htmlFor="symbol-input">股票代码：</label>
        <input
          id="symbol-input"
          type="text"
          placeholder="请输入股票代码"
          value={symbol}
          onChange={(e) => setSymbol(e.target.value)}
        />
      </div>

      <div style={{ marginBottom: '1rem' }}>
        <button onClick={getPrice}>获取价格</button>
        <pre>{priceResult}</pre>
      </div>

      <div style={{ marginBottom: '1rem' }}>
        <button onClick={runBacktest}>运行回测</button>
        <pre>{backtestResult}</pre>
      </div>

      <div style={{ marginBottom: '1rem' }}>
        <button onClick={showPriceChart}>显示收盘价折线图</button>
        {showChart && (
          chartLoading ? (
            <div>加载中...</div>
          ) : (
            chartData.length > 0 ? (
              <ReactECharts option={priceChartOption} style={{ height: 400, marginTop: 20 }} />
            ) : (
              <div style={{ marginTop: 20, color: 'red' }}>暂无数据</div>
            )
          )
        )}
      </div>
    </div>
  );
}

export default App;