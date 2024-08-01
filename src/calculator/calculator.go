package calculator

import (
	"fmt"
	"time"
)

// Returns the first searched for weekday for each month with 5 occurences of that weekday.
func Run(fromYear int, toYear int, weekDay time.Weekday, calendarId string, headline string, importDates bool) ([]time.Time, error) {
	months := monthsWithTheFiveWeekdays(fromYear, toYear, weekDay)

	if importDates {
		return months, addDaysToCalendar(calendarId, headline, &months)
	}

	return months, nil
}

func monthsWithTheFiveWeekdays(fromYear int, toYear int, weekDay time.Weekday) []time.Time {
	var foundMonths []time.Time

	for curYear := fromYear; curYear <= toYear; curYear++ {
		for curMonth := 1; curMonth <= 12; curMonth++ {
			firstDate := firstWeekDayInMonth(curYear, curMonth, weekDay)
			if firstDate.Day()+28 <= lastDayInMonth(curYear, curMonth) {
				foundMonths = append(foundMonths, firstDate)
			}
		}
	}
	return foundMonths
}

func firstWeekDayInMonth(year int, month int, weekDay time.Weekday) time.Time {
	calcDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	for calcDate.Weekday() != weekDay {
		calcDate = calcDate.AddDate(0, 0, 1)
	}

	return calcDate
}

func lastDayInMonth(year int, month int) int {
	switch month {
	case 12: // December
		return 31
	default:
		return time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()
	}
}

func addDaysToCalendar(calendarId string, headline string, dates *[]time.Time) error {
	fmt.Printf("Not implemented yet: %s, %s, %d\n", calendarId, headline, len(*dates))
	return nil
}
