package time

import (
	"time"
)

func LastWorkdayOfMonth(t time.Time) time.Time {
	nm := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).AddDate(0, 1, 0)
	minus := -1
	
	if (nm.Weekday() == time.Monday) {
		minus -= 2
	} else if (nm.Weekday() == time.Sunday) {
		minus -= 1
	}
	
	return nm.AddDate(0, 0, minus)
}

func LastWorkdayOfMonthWithYM(year ,month int) time.Time {
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	
	return lastWorkdayOfMonth(date)
}

