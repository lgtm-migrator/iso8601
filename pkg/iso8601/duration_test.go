package iso8601

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDurationString(t *testing.T) {
	for _, c := range []struct {
		duration Duration
		expected string
	}{
		{duration: Duration{}, expected: "P0D"},
		{duration: Duration{Years: 1}, expected: "P1Y"},
		{duration: Duration{Weeks: 1}, expected: "P1W"},
		{duration: Duration{Months: 1}, expected: "P1M"},
		{duration: Duration{Days: 1}, expected: "P1D"},
		{duration: Duration{Hours: 1}, expected: "PT1H"},
		{duration: Duration{Hours: 24}, expected: "PT24H"},
		{duration: Duration{Minutes: 1}, expected: "PT1M"},
		{duration: Duration{Seconds: 1}, expected: "PT1S"},
		{duration: Duration{Nanoseconds: 1}, expected: "PT0.000000001S"},
		{duration: Duration{Minutes: 1, Seconds: 1}, expected: "PT1M1S"},
		{duration: Duration{Negative: true}, expected: "-P0D"},
		{duration: Duration{Hours: -1}, expected: "PT-1H"},
		{duration: Duration{Minutes: -1}, expected: "PT-1M"},
		{duration: Duration{Seconds: -1}, expected: "PT-1S"},
		{duration: Duration{Minutes: -1, Seconds: -1}, expected: "PT-1M-1S"},
		{duration: Duration{
			Years:       maxInt64,
			Months:      maxInt64,
			Weeks:       maxInt64,
			Days:        maxInt64,
			Hours:       maxInt64,
			Minutes:     maxInt64,
			Seconds:     maxInt64,
			Nanoseconds: int64(time.Second) - 1,
		}, expected: "P9223372036854775807Y9223372036854775807M9223372036854775807W9223372036854775807DT9223372036854775807H9223372036854775807M9223372036854775807.999999999S"},
		{duration: Duration{
			Years:       minInt64,
			Months:      minInt64,
			Weeks:       minInt64,
			Days:        minInt64,
			Hours:       minInt64,
			Minutes:     minInt64,
			Seconds:     minInt64,
			Nanoseconds: -int64(time.Second) + 1,
			Negative:    true,
		}, expected: "-P-9223372036854775808Y-9223372036854775808M-9223372036854775808W-9223372036854775808DT-9223372036854775808H-9223372036854775808M-9223372036854775808.999999999S"},
	} {
		t.Run(c.expected, func(t *testing.T) {
			assert.Equal(t, c.expected, c.duration.String())
		})
	}
}

func TestDurationTimeDuration(t *testing.T) {
	for _, c := range []struct {
		duration Duration
		expected time.Duration
		error    error
	}{
		{duration: Duration{}, expected: 0},
		{duration: Duration{Months: 1}, expected: Month},
		{duration: Duration{Years: 1}, expected: Year},
		{duration: Duration{Weeks: 1}, expected: Week},
		{duration: Duration{Days: 1}, expected: Day},
		{duration: Duration{Hours: 24}, expected: 24 * time.Hour},
		{duration: Duration{Hours: 1}, expected: time.Hour},
		{duration: Duration{Minutes: 1}, expected: time.Minute},
		{duration: Duration{Seconds: 1}, expected: time.Second},
		{duration: Duration{Nanoseconds: 1}, expected: time.Nanosecond},
		{duration: Duration{Minutes: 1, Seconds: 1}, expected: time.Minute + time.Second},
		{duration: Duration{Negative: true, Hours: 1}, expected: -time.Hour},
		{duration: Duration{Negative: true}, expected: 0},
		{duration: Duration{Hours: -1}, expected: -time.Hour},
		{duration: Duration{Minutes: -1}, expected: -time.Minute},
		{duration: Duration{Seconds: -1}, expected: -time.Second},
		{duration: Duration{Minutes: -1, Seconds: -1}, expected: -time.Minute - time.Second},
		{duration: Duration{Years: maxInt64/int64(Year) + 1}, error: ErrOverflow},
		{duration: Duration{Years: maxInt64 / int64(Year), Months: 13}, error: ErrOverflow},
		{duration: Duration{Years: -maxInt64/int64(Year) - 1}, error: ErrOverflow},
		{duration: Duration{Years: -maxInt64 / int64(Year), Months: -13}, error: ErrOverflow},
		{duration: Duration{Months: maxInt64/int64(Month) + 1}, error: ErrOverflow},
		{duration: Duration{Months: maxInt64 / int64(Month), Weeks: 5}, error: ErrOverflow},
		{duration: Duration{Months: -maxInt64/int64(Month) - 1}, error: ErrOverflow},
		{duration: Duration{Months: -maxInt64 / int64(Month), Weeks: -5}, error: ErrOverflow},
		{duration: Duration{Weeks: maxInt64/int64(Week) + 1}, error: ErrOverflow},
		{duration: Duration{Weeks: maxInt64 / int64(Week), Days: 8}, error: ErrOverflow},
		{duration: Duration{Weeks: -maxInt64/int64(Week) - 1}, error: ErrOverflow},
		{duration: Duration{Weeks: -maxInt64 / int64(Week), Days: -8}, error: ErrOverflow},
		{duration: Duration{Days: maxInt64/int64(Day) + 1}, error: ErrOverflow},
		{duration: Duration{Days: maxInt64 / int64(Day), Hours: 25}, error: ErrOverflow},
		{duration: Duration{Days: -maxInt64/int64(Day) - 1}, error: ErrOverflow},
		{duration: Duration{Days: -maxInt64 / int64(Day), Hours: -25}, error: ErrOverflow},
		{duration: Duration{Hours: maxInt64/int64(time.Hour) + 1}, error: ErrOverflow},
		{duration: Duration{Hours: maxInt64 / int64(time.Hour), Minutes: 61}, error: ErrOverflow},
		{duration: Duration{Hours: -maxInt64/int64(time.Hour) - 1}, error: ErrOverflow},
		{duration: Duration{Hours: -maxInt64 / int64(time.Hour), Minutes: -61}, error: ErrOverflow},
		{duration: Duration{Minutes: maxInt64/int64(time.Minute) + 1}, error: ErrOverflow},
		{duration: Duration{Minutes: maxInt64 / int64(time.Minute), Seconds: 61}, error: ErrOverflow},
		{duration: Duration{Minutes: -maxInt64/int64(time.Minute) - 1}, error: ErrOverflow},
		{duration: Duration{Minutes: -maxInt64 / int64(time.Minute), Seconds: -61}, error: ErrOverflow},
		{duration: Duration{Seconds: maxInt64/int64(time.Second) + 1}, error: ErrOverflow},
		{duration: Duration{Seconds: maxInt64 / int64(time.Second), Nanoseconds: 1e9 + 1}, error: ErrOverflow},
		{duration: Duration{Seconds: -maxInt64/int64(time.Second) - 1}, error: ErrOverflow},
		{duration: Duration{Seconds: -maxInt64 / int64(time.Second), Nanoseconds: -1e9 - 1}, error: ErrOverflow},
	} {
		t.Run("", func(t *testing.T) {
			v, err := c.duration.TimeDuration()
			require.Equal(t, c.error, err)
			assert.Equal(t, c.expected, v)
		})
	}
}

func TestNewDuration(t *testing.T) {
	for _, c := range []struct {
		nano     int64
		expected Duration
	}{
		{nano: 0, expected: Duration{}},
		{nano: 1, expected: Duration{Nanoseconds: 1}},
		{nano: int64(time.Hour), expected: Duration{Hours: 1}},
		{nano: int64(time.Minute), expected: Duration{Minutes: 1}},
		{nano: int64(time.Second), expected: Duration{Seconds: 1}},
		{nano: int64(time.Millisecond), expected: Duration{Nanoseconds: 1e6}},
		{nano: int64(time.Microsecond), expected: Duration{Nanoseconds: 1e3}},
		{nano: int64(time.Nanosecond), expected: Duration{Nanoseconds: 1}},
		{nano: -int64(time.Hour), expected: Duration{Hours: 1, Negative: true}},
		{nano: -int64(time.Minute), expected: Duration{Minutes: 1, Negative: true}},
		{nano: -int64(time.Second), expected: Duration{Seconds: 1, Negative: true}},
		{nano: -int64(time.Millisecond), expected: Duration{Nanoseconds: 1e6, Negative: true}},
		{nano: -int64(time.Microsecond), expected: Duration{Nanoseconds: 1e3, Negative: true}},
		{nano: -int64(time.Nanosecond), expected: Duration{Nanoseconds: 1, Negative: true}},
		{nano: 1000 * int64(time.Hour), expected: Duration{Hours: 1000}},
		{
			nano: int64(time.Nanosecond +
				time.Microsecond +
				time.Millisecond +
				time.Second +
				time.Minute +
				time.Hour),
			expected: Duration{Hours: 1, Minutes: 1, Seconds: 1, Nanoseconds: 1001001},
		},
	} {
		t.Run(c.expected.String(), func(t *testing.T) {
			v := NewDuration(c.nano)
			assert.Equal(t, &c.expected, v)
		})
	}
}

func TestParseDuration(t *testing.T) {
	for _, c := range []struct {
		s        string
		expected Duration
		err      error
	}{
		{
			s: "P1Y1M1W1DT1H1M1.001S",
			expected: Duration{Years: 1, Months: 1, Weeks: 1, Days: 1,
				Hours: 1, Minutes: 1, Seconds: 1, Nanoseconds: 1e6},
		},
		{s: "P0D", expected: Duration{}},
		{s: "P1D", expected: Duration{Days: 1}},
		{s: "P-1D", expected: Duration{Days: -1}},
		{s: "P1M", expected: Duration{Months: 1}},
		{s: "P-1M", expected: Duration{Months: -1}},
		{s: "P1Y", expected: Duration{Years: 1}},
		{s: "P-1Y", expected: Duration{Years: -1}},
		{s: "PT1S", expected: Duration{Seconds: 1}},
		{s: "PT-1S", expected: Duration{Seconds: -1}},
		{s: "PT1H", expected: Duration{Hours: 1}},
		{s: "PT-1H", expected: Duration{Hours: -1}},
		{s: "PT1M", expected: Duration{Minutes: 1}},
		{s: "PT-1M", expected: Duration{Minutes: -1}},
		{s: "PT1H1M1S", expected: Duration{Hours: 1, Minutes: 1, Seconds: 1}},
		{s: "PT0.5H", expected: Duration{Minutes: 30}},
		{s: "PT0.001S", expected: Duration{Nanoseconds: 1e6}},
		{s: "-PT1H", expected: Duration{Hours: 1, Negative: true}},
		{s: "+PT1H", expected: Duration{Hours: 1}},
		{s: "", err: ErrInvalidDuration{String: ""}},
		{s: "-", expected: Duration{Negative: true}, err: ErrInvalidDuration{String: "-"}},
		{s: "+", err: ErrInvalidDuration{String: "+"}},
		{s: "1D", err: ErrInvalidDuration{String: "1D"}},
	} {
		t.Run(c.s, func(t *testing.T) {
			v, err := ParseDuration(c.s)
			require.Equal(t, c.err, err)
			assert.Equal(t, c.expected, v)
		})
	}
}

func BenchmarkDurationString(b *testing.B) {
	x := Duration{
		Years:       1234,
		Months:      -1234,
		Weeks:       1234,
		Days:        -1234,
		Hours:       1234,
		Minutes:     -1234,
		Nanoseconds: 1234,
		Negative:    true,
	}
	for i := 0; i < b.N; i++ {
		_ = x.String()
	}
}

func BenchmarkParseDuration(b *testing.B) {
	x := "P1Y23M34W56DT78H90M12.3456789S"
	for i := 0; i < b.N; i++ {
		_, err := ParseDuration(x)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewDuration(int64(Year))
	}
}

func BenchmarkDurationTimeDuration(b *testing.B) {
	x := Duration{
		Years:       100,
		Weeks:       100,
		Months:      100,
		Days:        100,
		Hours:       100,
		Minutes:     100,
		Seconds:     100,
		Nanoseconds: 1000}
	for i := 0; i < b.N; i++ {
		_ = x.MustTimeDuration()
	}
}
