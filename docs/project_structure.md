# Project Structure
This document describes the directory and file organization of the project.
```
golang-quant/
├── 🔁 backtest/
│   └── engine.go
├── 💼 broker/
│   └── paper.go
├── 🚀 cmd/
│   └── 🖥️ api/
│       └── main.go
├── 📊 data/
│   └── yahoo.go
├── 📚 docs/
│   └── project_structure.md
├── 🧠 strategy/
│   └── rsi.go
├── 🌐 web/
│   ├── handlers.go
│   ├── router.go
│   └── views/
│       └── index.html
├── 📦 go.mod
├── 🏃 Makefile
└── 📄 README.md
```
## Description
Descriptions for each part:
- backtest/  
Contains the core backtesting engine logic, responsible for simulating trading strategies over historical data.

    backtest/engine.go
    Implements the main functions to run backtests on trading signals and price data.

    broker/
    Provides broker simulation logic, such as paper trading accounts that track orders and balances.

    broker/paper.go
    Implements a paper broker for simulating order execution and account management.

    cmd/
    Entry points for running the application or its components.

    cmd/api/
    Contains the main entry for the web API server.

    cmd/api/main.go
    Starts the API server that exposes the application's functionality.

    data/
    Contains modules for obtaining and managing external market data.

    data/yahoo.go
    Implements data fetching from Yahoo Finance or similar sources.

    docs/
    Documentation files for the project.

    docs/project_structure.md
    Documentation describing the structure and organization of the project.

    strategy/
    Houses different trading strategies and their implementations.

    strategy/rsi.go
    Implements an RSI-based trading strategy.

    web/
    The web interface and API logic for the application.

    web/handlers.go
    Defines HTTP handlers for API endpoints.

    web/router.go
    Sets up routing for the web server.

    web/views/
    Contains frontend templates and static files.

    web/views/index.html
    The main HTML page for the web interface.

    go.mod
    The Go module definition file, declaring dependencies.

    Makefile
    Automation script for building, running, and managing the project.

    README.md
    The main documentation file, providing an overview and usage instructions for the project.
