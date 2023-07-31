package weekday

import (
	"errors"
	"fmt"
	"time"
)

func WeekDay(input string) (string, error) {
	testDate, error := time.Parse("01-02-2006", input)
	if error != nil {
		return "", errors.New("invalid date")
	}
	fmt.Println(testDate)

	curDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println(curDate)
	total := 0

	fmt.Println(testDate.Before(curDate))
	for testDate.Before(curDate) {
		if total == 7 {
			total = 0
		}
		total += 1
		curDate = curDate.AddDate(0, 0, -1)
	}

	fmt.Println(total)

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	return days[len(days)-total], nil
}
