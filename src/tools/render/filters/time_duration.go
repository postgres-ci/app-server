package filters

import (
	"fmt"
	"time"
)

func duration(diff time.Duration) string {

	if diff < time.Second {

		return diff.String()
	}

	var (
		value        string
		milliseconds = diff.Nanoseconds() / 1000000
		seconds      = (milliseconds / 1000) % 60
		minutes      = (milliseconds / (1000 * 60)) % 60
		hours        = (milliseconds / (1000 * 60 * 60)) % 24
	)

	switch true {
	case hours > 1:
		value += fmt.Sprintf("%d hours ", hours)
	case hours == 1:
		value += fmt.Sprintf("%d hour ", hours)
	}

	switch true {
	case minutes > 1:
		value += fmt.Sprintf("%d minutes ", minutes)
	case minutes == 1:
		value += fmt.Sprintf("%d minute ", minutes)
	}

	switch true {
	case seconds > 1:
		value += fmt.Sprintf("%d seconds", seconds)
	case seconds == 1:
		value += fmt.Sprintf("%d second", seconds)
	}

	return value
}
