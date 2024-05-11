package flights

import (
	"arrivals_lounge/test_utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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

func TestGetArrivals(t *testing.T) {
	os.Setenv("API_KEY", "test")

	var server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"flights": [
				{
					"from": "London",
					"code": "testCode",
					"scheduledArrival": "2022-01-01T00:00:00Z",
					"status": {
						"arrived": "2022-01-01T00:00:00Z",
						"expected_at": "2022-01-01T00:00:00Z",
						"cancelled": false
					}
				}
			]
		}`))
	}))
	defer server.Close()
	flights, err := GetArrivals(server.URL, "testCode")

	require.NoError(t, err, "failed to get arrivals")
	require.Equal(t, 1, len(flights), "expected 1 flight")
	require.Equal(t, "London", flights[0].Origin, "origin does not match")
	require.Equal(t, "testCode", flights[0].Code, "arrival code does not match")

	testArrivalTime := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)
	require.Equal(t, testArrivalTime, flights[0].ArrivedAt.Local(), "arrival time does not match")

	testExpectedArrivalTime := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)
	require.Equal(t, testExpectedArrivalTime, flights[0].ExpectedAt.Local(), "expected arrival time does not match")
	require.Equal(t, false, flights[0].Cancelled, "cancellation status does not match")
	require.NoError(t, err, "failed to decode response body")
}
