package iso8601

import "time"

// ParseTime from string.
// a shortcut for time.Parse(time.RFC3339Nano, s)
func ParseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, s)
}

// FormatTime to string
// a shortcut for string(t.AppendFormat(([32]byte{})[:0], time.RFC3339Nano))
func FormatTime(t time.Time) string {
	var a [32]byte
	return string(t.AppendFormat(a[:0], time.RFC3339Nano))
}
