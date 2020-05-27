package iso8601

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTime(t *testing.T) {
	for _, c := range []struct {
		s        string
		expected time.Time
	}{
		{s: "2001-02-03T04:05:06.07Z", expected: time.Date(2001, 2, 3, 4, 5, 6, 70e6, time.UTC)},
		{s: "2001-02-03T04:05:06Z", expected: time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)},
		// {s: "2001-02-03", expected: time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)},
	} {
		t.Run(c.s, func(t *testing.T) {
			v, err := ParseTime(c.s)
			require.NoError(t, err)
			assert.Equal(t, c.expected, v)
		})
	}
}

func BenchmarkParseTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ParseTime("2001-02-03T04:05:06.07Z")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestFormatTime(t *testing.T) {
	for _, c := range []struct {
		t        time.Time
		expected string
	}{
		{t: time.Date(2001, 2, 3, 4, 5, 6, 70e6, time.UTC), expected: "2001-02-03T04:05:06.07Z"},
		{t: time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC), expected: "2001-02-03T04:05:06Z"},
	} {
		t.Run(c.expected, func(t *testing.T) {
			s := FormatTime(c.t)
			assert.Equal(t, c.expected, s)
		})
	}
}
