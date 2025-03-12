# Crypto Masters

Crypto Masters is a Go application that fetches and displays the current exchange rates for various cryptocurrencies using the CEX.IO API.

## Features

- Fetches real-time exchange rates for cryptocurrencies.
- Supports multiple cryptocurrencies including BTC, ETH, and BCH.
- Concurrently fetches data using goroutines and sync.WaitGroup for efficient performance.

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/crypto-masters.git
    cd crypto-masters
    ```

2. Install dependencies (if any):

    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:

    ```sh
    go run main.go
    ```

2. The application will fetch and display the current exchange rates for BTC, ETH, and BCH.

## Project Structure

- `main.go`: The main entry point of the application. It initializes the list of cryptocurrencies and fetches their exchange rates concurrently.
- `api/cex.go`: Contains the `GetRate` function which fetches the exchange rate for a given cryptocurrency from the CEX.IO API.
- `api/responses.go`: Defines the `CEXResponse` struct which represents the JSON response from the CEX.IO API.
- `util/jsonfn.go`: Contains the `ParseJson` function which parses JSON-encoded data.

## Example Output
