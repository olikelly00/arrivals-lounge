package flights

import (
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Flight struct {
	Code       string    `json:"code"`
	Origin     string    `json:"from"`
	DueTime    time.Time `json:"scheduled_arrival"`
	ArrivedAt  time.Time `json:"arrived"`
	Cancelled  bool      `json:"cancelled"`
	ExpectedAt time.Time `json:"expected_at"`
}

func Display(flights []Flight) {
	headerFmt := color.New(color.FgYellow, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Time", "From", "Code", "Arrived", "Status", "Expected").WithPadding(6)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithHeaderSeparatorRow('\u2708')

	for _, flight := range flights {
		tbl.AddRow(
			flight.DueTime.Format("15:04"),
			flight.Origin,
			flight.Code,
			func() string {
				if !flight.ArrivedAt.IsZero() {
					return flight.ArrivedAt.Format("15:04")
				} else {
					return ""
				}
			}(),
			func() string {
				if !flight.Cancelled && !flight.ArrivedAt.IsZero() {
					return "Arrived"
				} else if !flight.Cancelled && flight.ExpectedAt.After(flight.DueTime.Add(2*time.Hour)) {
					return "Severe delay (2hr+)"
				} else if !flight.Cancelled && flight.ExpectedAt.After(flight.DueTime.Add(1*time.Hour)) {
					return "Moderate delay (1hr+)"
				} else if !flight.Cancelled && flight.ExpectedAt.After(flight.DueTime) {
					return "Minor delay (up to 1hr)"
				} else if !flight.Cancelled {
					return "En route"
				} else {
					return "Cancelled"
				}
			}(),
			func() string {
				if flight.ArrivedAt.IsZero() {
					return flight.ExpectedAt.Format("15:04")
				} else {
					return ""
				}
			}())
	}
	tbl.Print()
}
