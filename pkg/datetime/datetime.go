package datetime

import (
	"fmt"
	"time"
)

var layout = "2006-01-02T15:04:05Z"

type datetime struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
	Sec   int
}

func (t *datetime) normalize(y1 int, M1 time.Month) {
	// Normalize negative values
	if t.Sec < 0 {
		t.Sec += 60
		t.Min--
	}
	if t.Min < 0 {
		t.Min += 60
		t.Hour--
	}
	if t.Hour < 0 {
		t.Hour += 24
		t.Day--
	}
	if t.Day < 0 {
		s := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		t.Day += 32 - s.Day()
		t.Month--
	}
	if t.Month < 0 {
		t.Month += 12
		t.Year--
	}
}

func diff(a, current time.Time) (out datetime, suffix string) {
	if a.Location() != current.Location() {
		current = current.In(a.Location())
	}
	suffix = "ago"
	if a.After(current) {
		a, current = current, a
		suffix = "left"
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := current.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := current.Clock()

	dt := datetime{
		Year:  int(y2 - y1),
		Month: int(M2 - M1),
		Day:   int(d2 - d1),
		Hour:  int(h2 - h1),
		Min:   int(m2 - m1),
		Sec:   int(s2 - s1),
	}

	dt.normalize(y1, M1)
	return dt, suffix
}

// DateTime - description
func DateTimeFromNow(val string) string {

	t1, err := time.Parse(layout, val)
	if err != nil {
		fmt.Println(err)
	}

	t2 := time.Now().UTC()
	dt, flag := diff(t1, t2)
	out := ""
	if dt.Year != 0 {
		if dt.Month != 0 {
			out = fmt.Sprintf("%vyr %vmonth %v", dt.Year, dt.Month, flag)
		} else if dt.Day != 0 {
			out = fmt.Sprintf("%vyr %vday %v", dt.Year, dt.Day, flag)
		} else {
			out = fmt.Sprintf("%vyr %v", dt.Year, flag)
		}
	} else if dt.Month != 0 {
		if dt.Day != 0 {
			out = fmt.Sprintf("%vmonth %vday %v", dt.Month, dt.Day, flag)
		} else if dt.Hour != 0 {
			out = fmt.Sprintf("%vmonth %vhour %v", dt.Month, dt.Hour, flag)
		} else {
			out = fmt.Sprintf("%vmonth %v", dt.Month, flag)
		}
	} else if dt.Day != 0 {
		if dt.Hour != 0 {
			out = fmt.Sprintf("%vday %vhour %v", dt.Day, dt.Hour, flag)
		} else if dt.Min != 0 {
			out = fmt.Sprintf("%vday %vmin %v", dt.Day, dt.Min, flag)
		} else {
			out = fmt.Sprintf("%vday %v", dt.Day, flag)
		}
	} else if dt.Hour != 0 {
		if dt.Min != 0 {
			out = fmt.Sprintf("%vhour %vmin %v", dt.Hour, dt.Min, flag)
		} else if dt.Sec != 0 {
			out = fmt.Sprintf("%vhour %vsec %v", dt.Hour, dt.Sec, flag)
		} else {
			out = fmt.Sprintf("%vhour %v", dt.Hour, flag)
		}
	} else if dt.Min != 0 {
		if dt.Sec != 0 {
			out = fmt.Sprintf("%vmin %vsec %v", dt.Min, dt.Sec, flag)
		} else {
			out = fmt.Sprintf("%vmin %v", dt.Min, flag)
		}
	} else {
		out = fmt.Sprintf("%vsec %v", dt.Sec, flag)
	}
	return out
}
