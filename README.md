Welcome to the README for Arrivals Lounge. This repo contains the codebase for Arrivals Lounge, built from scratch by Oli Kelly. 

## Description

Arrivals Lounge is a terminal-based application that, when run, will fetch live flight arrivals data from the FlightAware API, format them into an arrivals board along with any anticipated delays, and display it to the user. 

To run the application, simply type the following into your CLI:

```go

go run main.go //followed by the IATA code for your chosen airport, eg. LHR for London Heathrow

```

## Installation

To run Arrivals Lounge locally, you'll need to install [Golang](https://go.dev/) 


## Usage
1. Clone the repository `git clone https://github.com/olikelly00/arrivals-lounge` 

2. Go to the project directory `cd arrivals-lounge`

3. Install dependencies 

4. Run the backend server `go run main.go + IATA code`



## Project Structure 

**Root directory**

main.go
- The entry point of our program - this is the file you run in CLI that all other files feed into.

go.mod
- Lists all the modules (external packages) your project depends on.

go.sum
- Contains checksums for all the dependencies listed in go.mod. This ensures you're download the exact same versions of dependencies every time you build the project.

README.md
- The one you're reading!

**Flights package**

flights.go
- Houses the flight 'model' (struct), which defines the properties of the flights allowing for manipulaton within the application. 
- Also includes the Display function, which takes the flight data received from the API and formats it into a board for the user.

json.go
- Holds functions for fetching flight data from an external API.

flight_test.go
- Holds the unit tests written for the application
