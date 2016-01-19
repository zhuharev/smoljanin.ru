package since

import (
	"fmt"
	"github.com/Unknwon/i18n"
	"math"
	"strings"
	"time"
)

func Since(t time.Time) string {
	//now := time.Now()

	s := since(time.Since(t))

	if s <= 0 {
		return i18n.Tr("ru-RU", "now")
	}

	if ms := int(s.Minutes()); ms <= 45 {
		return fmt.Sprintf("%d %s назад", ms, plural("минута_минуты_минут", ms))
	} else if ms <= 90 {
		return fmt.Sprintf("час назад")
	}

	if hs := int(s.Hours()); hs < 22 {
		return i18n.Tr("ru-RU", "%d %s назад", hs, plural("час_часа_часов", hs))
	} else {
		return fmt.Sprintf("вчера")
	}

	return s.String()
}

func plural(word string, num int) string {
	var forms = strings.Split(word, "_")
	if num%10 == 1 && num%100 != 11 {
		return forms[0]
	} else if num%10 >= 2 && num%10 <= 4 && (num%100 < 10 || num%100 >= 20) {
		return forms[1]
	} else {
		return forms[2]
	}
}

type since time.Duration

func (s since) Seconds() int {
	return int(time.Duration(s).Seconds())
}

func (s since) Minutes() int {
	return int(time.Duration(s).Minutes())
}

func (s since) Hours() int {
	return int(time.Duration(s).Hours())
}

func (s since) Days() int {
	return int(math.Floor(float64(s.Seconds()) / 86400))
}

func (s since) Weeks() int {
	return int(math.Floor(float64(s.Days() / 7)))
}

func (s since) Months() int {
	return int(math.Floor(float64(s.Weeks()) / 4))
}

func (s since) Years() int {
	return int(math.Floor(float64(s.Days()) / 365))
}

func (s since) String() string {
	return time.Duration(s).String()
}
