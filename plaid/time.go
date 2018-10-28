package plaid

import (
	"fmt"
	"time"
)

// DateFormat defines the date format used for client-facing date strings.
const DateFormat = "2006-01-02"

// DateTimeFormat defines the datetime format used for client-facing date
// strings.
const DateTimeFormat = time.RFC3339

// Date is a time.Time object whose default string format (including for json
// serialization) is a client-facing date string.
type Date time.Time

// MarshalJSON implements the json.Marshaler interface for Date. It outputs
// a string in the form "2006-01-02".
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Date. It expects
// a string of the form "2016-01-02".
func (t *Date) UnmarshalJSON(data []byte) error {
	v, err := time.Parse(`"`+DateFormat+`"`, string(data))
	if err != nil {
		return err
	}
	*t = Date(v)
	return nil
}

func (t Date) String() string {
	return fmt.Sprintf("\"%s\"", time.Time(t).Format(DateFormat))
}

// Format returns the date as an alternately formatted string. The format
// argument is the same as those accepted by time.Time.Format.
func (t Date) Format(format string) string {
	return time.Time(t).Format(format)
}

// DateTime is a time.Time object whose default string format (including for
// json serialization) is a client-facing datetime string.
type DateTime time.Time

// MarshalJSON implements the json.Marshaler interface for Date. It outputs
// a string in the form "2006-01-02T00:00:00Z".
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Date. It expects
// a string of the form "2016-01-02T00:00:00Z".
func (t *DateTime) UnmarshalJSON(data []byte) error {
	v, err := time.Parse(`"`+DateTimeFormat+`"`, string(data))
	if err != nil {
		return err
	}
	*t = DateTime(v)
	return nil
}

func (t DateTime) String() string {
	return fmt.Sprintf("\"%s\"", time.Time(t).Format(DateTimeFormat))
}

// Format returns the datetime as an alternately formatted string. The format
// argument is the same as those accepted by time.Time.Format.
func (t DateTime) Format(format string) string {
	return time.Time(t).Format(format)
}
