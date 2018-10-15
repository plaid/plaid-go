package plaid

import "time"

// DateLayout is the time format expected by the api
const DateLayout = "2012-11-01"

// Date returns a time from the specified date
func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
