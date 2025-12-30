package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const longForm = "1/2/2006 15:04:05"
	t, _ := time.Parse(longForm, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedDate, _ := time.Parse("January 1, 2006 15:04:05", date)
	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate, _ := time.Parse("Monday, January 1, 2006 15:04:05", date)
	hour := parsedDate.Hour()
	if hour < 12 {
		return false
	}
	return true
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedDate, _ := time.Parse("1/2/2006 15:04:05", date)
	formattedDate := parsedDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	return formattedDate
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	thisYear := time.Now().Year()
	return time.Date(thisYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
