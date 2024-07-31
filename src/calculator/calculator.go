package calculator

import (
	"fmt"
	"time"
)

func Run(fromYear int, toYear int, weekDay time.Weekday, calendarId string, headline string, importDates bool) {
	months := monthsWithTheFiveWeekdays(fromYear, toYear, weekDay)

	fmt.Printf("Found: %v\n", months)

	if !importDates {
		fmt.Println("Done")
		return
	} else {
		addDaysToCalendar(calendarId, headline, &months)
	}
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

func addDaysToCalendar(calendarId string, headline string, dates *[]time.Time) {
	fmt.Printf("Not implemented yet: %s, %s, %d\n", calendarId, headline, len(*dates))
}
