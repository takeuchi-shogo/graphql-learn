package value

import "time"

func FormatToString(t time.Time) string {
	return t.Format(time.RFC3339)
}
