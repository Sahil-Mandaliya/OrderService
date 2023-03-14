package util

import "time"

const (
	FormatDDMMYYY = "02/01/2006"
)

func FormatTimeToStringDDMMYYYY(t time.Time) string {
	return t.Format(FormatDDMMYYY)
}
