// DateNow returns the current date formatted for Telegram messages.

package internal

import (
	"fmt"
	"time"
)

func DateNow() string {
	weekdays := [7]string{"неділя", "понеділок", "вівторок", "середа",
		"четвер", "п'ятниця", "субота"}
	week := weekdays[int(time.Now().Weekday())]
	return fmt.Sprintf("%s, %v", week, time.Now().Format("02.01.2006"))
}
