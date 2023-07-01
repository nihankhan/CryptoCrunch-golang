# CryptoCrunch

CryptoCrunch is a real-time cryptocurrency price aggregation and analysis tool.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

CryptoCrunch is a web-based application that collects real-time cryptocurrency price data from various sources and provides an aggregated view of the latest prices. It uses WebSockets to receive live data updates, and the data is processed and displayed in a table format on the user interface.

## Features

- Real-time cryptocurrency price updates
- Aggregated view of cryptocurrency prices in USD
- Timestamps for each price update
- User-friendly and responsive user interface

## Getting Started

To get started with CryptoCrunch, follow these steps:

1. Clone the repository:

git clone https://github.com/nihankhan/cryptocrunch.git


2. Install the required dependencies:

cd cryptocrunch
go get ./...


3. Build the project:

go build ./cmd/cryptocrunch

4. Run the application:

./cryptocrunch


5. Open your web browser and go to `http://localhost:8080` to access CryptoCrunch.

## Usage

Once you have the CryptoCrunch application running, you will see a table with the latest cryptocurrency prices in USD. The table will update in real-time as new price data is received.

## Contributing

Contributions to CryptoCrunch are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request. When contributing, please follow the existing code style and commit guidelines.

## License

This project is licensed under the [MIT License](LICENSE).





