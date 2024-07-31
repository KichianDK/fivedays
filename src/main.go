package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/KichianDK/fivedays/src/calculator"
)

func main() {

	currentYear := time.Now().Year()
	fromYear := flag.Int("from", currentYear, "Wich year to calculate months for")
	toYear := flag.Int("until", currentYear, "Last year to calculate months for")
	weekDay := flag.Int("weekday", 1, "Weekday to look for. Defaults to Monday")

	googleCalendarId := flag.String("calendarId", "", "Your google calander ID")
	importDates := flag.Bool("import", false, "Import dates to calendar")
	calendarText := flag.String("text", "", "The headlie for the calendar item")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("use -h for parameters")
		os.Exit(0)
	}

	calculator.Run(*fromYear, *toYear, time.Weekday(*weekDay), *googleCalendarId, *calendarText, *importDates)
}
