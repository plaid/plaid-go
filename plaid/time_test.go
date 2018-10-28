package plaid

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateMarshalJSON(t *testing.T) {
	testCases := []struct {
		desc     string
		data     Date
		expected []byte
	}{
		{
			desc:     "valid with no time portion",
			data:     Date(time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC)),
			expected: []byte("\"2016-01-05\""),
		},
		{
			desc:     "valid with time portion",
			data:     Date(time.Date(2016, time.January, 5, 1, 2, 3, 4, time.UTC)),
			expected: []byte("\"2016-01-05\""),
		},
	}
	for _, tc := range testCases {
		actual, err := json.Marshal(tc.data)
		assert.NoError(t, err, "test case: %s", tc.desc)
		assert.Exactly(t, tc.expected, actual, "test case: %s", tc.desc)
	}
}

func TestDateUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		desc     string
		data     []byte
		expected Date
	}{
		{
			desc:     "valid",
			data:     []byte("\"2016-01-05\""),
			expected: Date(time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, tc := range testCases {
		var actual Date
		err := json.Unmarshal(tc.data, &actual)
		assert.NoError(t, err, "test case: %s", tc.desc)
		assert.Exactly(t, tc.expected, actual, "test case: %s", tc.desc)
	}

	var actual Date
	assert.Error(t, json.Unmarshal([]byte("foobar"), &actual), "test case: invalid date")
}

func TestDateTimeMarshalJSON(t *testing.T) {
	testCases := []struct {
		desc     string
		data     DateTime
		expected []byte
	}{
		{
			desc:     "valid with no time portion",
			data:     DateTime(time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC)),
			expected: []byte("\"2016-01-05T00:00:00Z\""),
		},
		{
			desc:     "valid with time portion",
			data:     DateTime(time.Date(2016, time.January, 5, 1, 2, 3, 0, time.UTC)),
			expected: []byte("\"2016-01-05T01:02:03Z\""),
		},
		{
			desc:     "valid with nanoseconds portion",
			data:     DateTime(time.Date(2016, time.January, 5, 1, 2, 3, 4, time.UTC)),
			expected: []byte("\"2016-01-05T01:02:03Z\""),
		},
	}
	for _, tc := range testCases {
		actual, err := json.Marshal(tc.data)
		assert.NoError(t, err, "test case: %s", tc.desc)
		assert.Exactly(t, tc.expected, actual, "test case: %s", tc.desc)
	}
}

func TestDateTimeUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		desc     string
		data     []byte
		expected DateTime
	}{
		{
			desc:     "valid",
			data:     []byte("\"2016-01-05T00:00:00Z\""),
			expected: DateTime(time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC)),
		},
		{
			desc:     "valid",
			data:     []byte("\"2016-01-05T01:02:03Z\""),
			expected: DateTime(time.Date(2016, time.January, 5, 1, 2, 3, 0, time.UTC)),
		},
	}
	for _, tc := range testCases {
		var actual DateTime
		err := json.Unmarshal(tc.data, &actual)
		assert.NoError(t, err, "test case: %s", tc.desc)
		assert.Exactly(t, tc.expected, actual, "test case: %s", tc.desc)
	}

	var actual DateTime
	assert.Error(t, json.Unmarshal([]byte("foobar"), &actual), "test case: invalid datetime")
}
