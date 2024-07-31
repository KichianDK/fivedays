package calculator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLastDayInMonth(t *testing.T) {

	// Given
	januar := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	februar2024 := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC) // Leap year
	februar2023 := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)
	september := time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC)
	december := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)

	// When
	lastInJanuar := lastDayInMonth(januar.Year(), int(januar.Month()))
	lastInFebrua2024 := lastDayInMonth(februar2024.Year(), int(februar2024.Month()))
	lastInFebrua2023 := lastDayInMonth(februar2023.Year(), int(februar2023.Month()))
	lastInSeptember := lastDayInMonth(december.Year(), int(september.Month()))
	lastinDecember := lastDayInMonth(december.Year(), int(december.Month()))

	// Then
	assert.Equal(t, 31, lastInJanuar)
	assert.Equal(t, 29, lastInFebrua2024)
	assert.Equal(t, 28, lastInFebrua2023)
	assert.Equal(t, 30, lastInSeptember)
	assert.Equal(t, 31, lastinDecember)
}

func TestFirstWeekDayInMonth(t *testing.T) {

	// Given
	august2024 := time.Date(2024, time.Month(8), 1, 0, 0, 0, 0, time.Local)

	// When
	firstTuesday := firstWeekDayInMonth(august2024.Year(), int(august2024.Month()), time.Tuesday)

	// Then
	assert.Equal(t, 6, firstTuesday.Day())
}

func TestMonthsWithTheFiveWeekdays(t *testing.T) {
	// Given
	year := 2024

	// When
	months := monthsWithTheFiveWeekdays(year, year, time.Tuesday)

	// Then
	assert.Equal(t, 5, len(months))
	assert.Equal(t, time.July, months[2].Month())
	assert.Equal(t, time.December, months[4].Month())
}
