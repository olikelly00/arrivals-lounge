package flights

import (
	"arrivals_lounge/test_utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestFlightStructConstructs(t *testing.T) {
	TestFlight := Flight{
		"BA 341",
		"London",
		time.Date(2024, 5, 15, 14, 00, 00, 000, time.Local),
		time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local),
		false,
		time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local),
	}
	resultCode, expectedCode := TestFlight.Code, "BA 341"
	if resultCode != expectedCode {
		t.Errorf("Result is %v when %v is expected", resultCode, expectedCode)
	}
	resultOrigin, expectedOrigin := TestFlight.Origin, "London"
	if resultOrigin != expectedOrigin {
		t.Errorf("Result is %v when %v is expected", resultOrigin, expectedOrigin)
	}
	resultDueTime, expectedDueTime := TestFlight.DueTime, time.Date(2024, 5, 15, 14, 00, 00, 000, time.Local)
	if resultDueTime != expectedDueTime {
		t.Errorf("Result is %v when %v is expected", resultDueTime, expectedDueTime)
	}
}

func TestTerminalOutput(t *testing.T) {
	recording := test_utils.StartRecording()
	fmt.Println("Flight BA 341 from London is expected at 14:00")
	result := test_utils.EndRecording(recording)

	expected := "Flight BA 341 from London is expected at 14:00"

	if result != expected {
		t.Errorf("result is %v but %v was expected", result, expected)
	}
}

// func TestDisplay(t *testing.T) {
// 	//If I have three flights and I want to call the Display function on them,I would expect for the table to have 7 columns,4 rows.

// 	TestFlights := []Flight{
// 		{"BA 114", "London", time.Date(2024, 5, 15, 16, 23, 00, 000, time.Local), time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local), false, time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local)},
// 		{"LH 888", "Berlin", time.Date(2024, 5, 15, 17, 24, 00, 000, time.Local), time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local), false, time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local)},
// 		{"JA 903", "Tokyo", time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local), time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local), false, time.Date(2024, 5, 15, 18, 20, 00, 000, time.Local)},
// 	}

// 	result := Display(TestFlights)
// 	expected := "Time From Code Status\n16:23 London BA 114 18:20 false 18:20\n17:24 Berlin LH 888 18:20 false 18:20\n18:20 Tokyo JA 903 18:20 false 18:20"
// 	if result != expected {
// 		t.Errorf("result is %v but %v was expected", result, expected)
// 	}
// }

// func TestReadfromJSON(t *testing.T) {
// 	result := ReadfromJSON("/Users/olikelly/Documents/Coding/arrivals-lounge/testFile.json")
// 	expected := []Flight{
// 		{"BA 114", "London", time.Date(2024, 5, 15, 16, 23, 00, 000, time.UTC), time.Time{}, false, time.Date(2024, 3, 14, 13, 8, 10, 000, time.UTC)},
// 		{"LH 888", "Berlin", time.Date(2024, 5, 15, 17, 24, 00, 000, time.UTC), time.Time{}, true, time.Time{}},
// 		{"JA 903", "Tokyo", time.Date(2024, 5, 15, 18, 20, 00, 000, time.UTC), time.Time{}, false, time.Date(2024, 3, 14, 13, 3, 40, 000, time.UTC)},
// 	}
// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("result is %v but %v was expected", result, expected)
// 	}
// }

// func TestGetArrivals(t *testing.T) {
// 	req, err := http.NewRequest("GET", "https://go-flights-api.onrender.com/flights?code=JFK", nil)
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	GetArrivals.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
//         t.Errorf("handler returned wrong status code: got %v want %v",
//             status, http.StatusOK)
//     }
// }

func TestGetArrivals(t *testing.T) {
	os.Setenv("API_KEY", "test")
	// Create a mock HTTP server
	//var receivedReq *http.Request
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//Store the received request
		receivedReq = req
		// Print the received request
		fmt.Printf("Received request: %v\n", req)
		// Send response to be tested

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`OK`))
	}))

	// Close the server when test finishes
	defer server.Close()

	// Call your function
	GetArrivals(server.URL)

	// if receivedReq.URL.String() != server.URL {
	// 	t.Errorf("wrong URL, got: %s, want: %s", receivedReq.URL.String(), server.URL)
	// }

	// Check if function's http response was valid
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code, got: %d, want: %d", resp.StatusCode, http.StatusOK)
	}
}
