package flights

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type JSONFlight struct {
	Flights []struct {
		From             string `json:"from"`
		To               string `json:"to"`
		Code             string `json:"code"`
		ScheduledArrival string `json:"scheduled_arrival"`
		Status           struct {
			Arrived    string `json:"arrived"`
			Cancelled  bool   `json:"cancelled"`
			ExpectedAt string `json:"expected_at"`
		} `json:"status"`
	} `json:"flights"`
}

func GetArrivals(code string) []Flight {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	var flights []Flight
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API key is not set. Please set the API_KEY environment variable.")
	}
	apiURL := "https://go-flights-api.onrender.com/flights?code=" + code
	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case <-ticker.C:
			request.Header.Set("x-makers-password", apiKey)
			request.Header.Set("Content-Type", "application/json; charset=utf-8")

			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				log.Fatalf("Unable to send http request due to %s", err)
			}
			defer response.Body.Close()

			responseBody, responsebodyError := io.ReadAll(response.Body)
			if responsebodyError != nil {
				log.Fatalf("Unable to read response body due to %s", responsebodyError)
			}
			var wrapper JSONFlight
			unmarshallingError := json.Unmarshal(responseBody, &wrapper)
			if unmarshallingError != nil {
				log.Fatalf("Unable to unmarshal JSON due to %s", unmarshallingError)
			}

			for _, flight := range wrapper.Flights {
				dueTime, _ := time.Parse(time.RFC3339, flight.ScheduledArrival)
				arrivedAt, _ := time.Parse(time.RFC3339, flight.Status.Arrived)
				expectedAt, _ := time.Parse(time.RFC3339, flight.Status.ExpectedAt)
				flights = append(flights, Flight{
					Origin:     flight.From,
					Code:       flight.Code,
					DueTime:    dueTime.Local(),
					ArrivedAt:  arrivedAt.Local(),
					Cancelled:  flight.Status.Cancelled,
					ExpectedAt: expectedAt.Local(),
				})
			}
			return flights
		}
	}
}

// ```
// Your code is generally well-structured, but there are a few areas that could be improved:

// Error Handling: You're ignoring errors returned by time.Parse and http.NewRequest. It's generally a good practice to handle these errors instead of ignoring them. If these functions fail, they will return a zero value and an error. Ignoring these errors could lead to unexpected behavior.

// Infinite Loop: Your for loop will run indefinitely because there's no condition to break the loop. This could lead to high CPU usage and other issues. You should add a condition to break the loop when necessary.

// HTTP Client: You're creating a new HTTP client in each iteration of the loop. It's more efficient to create the client once outside the loop and reuse it.

// Ticker Usage: You're using a ticker to delay the execution of the code inside the loop. However, the return flights statement is inside the loop, so the function will return during the first iteration of the loop, and the ticker won't have any effect. You should move the return flights statement outside the loop.

// Error Logging: Instead of using fmt.Println to print errors, consider using log.Println or log.Fatalf. This will print the errors along with other useful information like the date and time.

// HTTP Request Error Handling: After making the HTTP request, you should check the HTTP status code to make sure the request was successful. If the status code is not 200, you should handle the error.
// ```
