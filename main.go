package main

import (
	"fmt"
	"maxRTE/calendar/weekday"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Enter a valid date")
		os.Exit(0)
	}
	testDate := strings.Trim(os.Args[1], " ")

	weekday, err := weekday.WeekDay(testDate)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(weekday)

}
