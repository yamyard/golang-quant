# Project Structure
This document describes the directory and file organization of the project.
```
golang-quant/
├── 🚀 cmd/
│   └── 🖥️ api/
│       └── main.go
├── 📊 data/
│   └── yahoo.go
├── 🧠 strategy/
│   └── rsi.go
├── 🔁 backtest/
│   └── engine.go
├── 💼 broker/
│   └── paper.go
├── 🌐 web/
│   ├── handlers.go
│   ├── router.go
│   └── views/
│       └── index.html
└── 📦 go.mod
```
## Description
```
| File/Folder            | Description                                              |
| ---------------------- | -------------------------------------------------------- |
| `cmd/api/main.go`      | Entry point, starts the API server                       |
| `data/yahoo.go`        | Simulated Yahoo Finance data fetching                    |
| `strategy/rsi.go`      | RSI (Relative Strength Index) signal generation strategy |
| `backtest/engine.go`   | Backtest engine, executes trades based on signals        |
| `broker/paper.go`      | Simulated paper trading account (virtual broker)         |
| `web/handlers.go`      | API route handlers for data and backtest operations      |
| `web/router.go`        | Gin router configuration                                 |
| `web/views/index.html` | Basic HTML frontend for interacting with the API         |
| `go.mod`               | Go module definition and dependency management           |
| `README.md`            | Project overview, structure, and usage instructions      |
```
