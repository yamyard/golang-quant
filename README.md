# Golang Quant
A lightweight quantitative trading platform implemented in Golang, featuring price display, signal generation, backtesting engine, and a web visualization interface.

## Install Dependencies
```
go mod tidy
```

## Start the Project
To start in the terminal:
```
make run
```
After starting, visit the following address in your browser:
```
http://localhost:8080
```
Then follow the instructions to enter the stock symbol in the textbox.

## Project Features
- Display historical stock prices
- Generate trading signals based on the RSI strategy
- Execute simple signal-based backtesting
- View prices, RSI trading signals, and final account value on the web page

## Example Interaction
- After clicking the button "Get Price" on the webpage, you will see:
  - Daily stock prices
- After clicking the button "Run Backtest" on the webpage, you will see:
  - Final account value
  - Daily RSI trading signals (BUY / SELL / HOLD)
