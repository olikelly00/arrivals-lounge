# Arrivals Lounge 🛬✈️

Welcome to the README for Arrivals Lounge, a terminal-based application for displaying live flight arrivals data.

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [Contact](#contact)
- [Testing](#testing)

## Description

Arrivals Lounge is a Go-based terminal application that fetches live flight arrivals data from the FlightAware API, formats it into an arrivals board with anticipated delays, and displays it to the user. Built from scratch by Oli Kelly, this application provides real-time flight information for any chosen airport.

## Features

- Fetch live flight arrivals data for any airport using its IATA code
- Display formatted arrivals board in the terminal
- Show anticipated delays for incoming flights
- Easy-to-use command-line interface

## Technologies Used

- **Backend**: Go (Golang)
- **External API**: FlightAware API
- **Version Control**: Git

## Installation

To run Arrivals Lounge locally, you'll need to install:

- [Go (Golang) - Official Go Installation Guide](https://golang.org/doc/install)

### Steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/olikelly00/arrivals-lounge.git
    ```

2. Navigate to the project directory:
    ```sh
    cd arrivals-lounge
    ```

3. Install dependencies (if any are not included in go.mod):
    ```sh
    go mod tidy
    ```

## Usage

1. Ensure you have set up your FlightAware API credentials in the `.env` file.
2. Run the application with the IATA code for your chosen airport:
    ```sh
    go run main.go LHR
    ```
    Replace `LHR` with the IATA code of any airport you want to check.

## Project Structure

The Arrivals Lounge project is organized as follows:

- `main.go`: Entry point of the program
- `go.mod`: Lists all the modules (external packages) the project depends on
- `go.sum`: Contains checksums for all the dependencies
- `README.md`: Project documentation (this file)
- `.env`: Environment variables file (not tracked in git)
- `.gitignore`: Specifies files to be ignored by git
- `flights/`: Package containing flight-related functionality
  - `flights.go`: Defines flight struct and display function
  - `json.go`: Functions for fetching flight data from API
  - `flight_test.go`: Unit tests for flight package
- `test_utils/`: Utility functions for testing
  - `recording.go`: Possibly for recording API responses for tests
- `flightData.json`: Sample flight data for testing or development
- `testFile.json`: Additional test data file

## Contributing

Contributions to Arrivals Lounge are welcome! Please feel free to submit a Pull Request.

## Contact

For any questions or feedback, please contact Oli Kelly:

- GitHub: [@olikelly00](https://github.com/olikelly00)

## Testing

Arrivals Lounge uses Go's built-in testing framework. To run the tests, follow these steps:

1. Ensure you are in the project root directory.
2. Run all tests with:
    ```sh
    go test ./...
    ```

3. To run tests for a specific package:
    ```sh
    go test ./flights
    ```

4. For verbose output, add the `-v` flag:
    ```sh
    go test -v ./...
    ```

5. To run a specific test function:
    ```sh
    go test -v ./flights -run TestFunctionName
    ```
    Replace `TestFunctionName` with the name of the test function you want to run.

For more information on Go testing, refer to the [official Go testing documentation](https://golang.org/pkg/testing/).

