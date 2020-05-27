package iso8601

import "time"

// ParseTime from string.
// a shortcut for time.Parse(time.RFC3339Nano, s)
func ParseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, s)
}

// FormatTime to string
// a shortcut for string(t.AppendFormat(make([]byte, 0, 32), time.RFC3339Nano))
func FormatTime(t time.Time) string {
	return string(t.AppendFormat(make([]byte, 0, 32), time.RFC3339Nano))
}
