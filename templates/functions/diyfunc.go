package functions

import "time"

func PYear(t time.Time) int {
	return t.Year()
}

func PMonth(t time.Time) string {
	return t.Month().String()
}

func PDay(t time.Time) int {
	return t.Day()
}
