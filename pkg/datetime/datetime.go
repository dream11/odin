package main

import (
	"fmt"
	"time"
)

var layout = "2006-01-02T15:04:05.000000Z"

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

func diff(a, b time.Time) (out datetime) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	dt := datetime{
		Year:  int(y2 - y1),
		Month: int(M2 - M1),
		Day:   int(d2 - d1),
		Hour:  int(h2 - h1),
		Min:   int(m2 - m1),
		Sec:   int(s2 - s1),
	}

	dt.normalize(y1, M1)
	return dt
}

func dateTime(val string) string {

	t1, err := time.Parse(layout, val)
	if err != nil {
		fmt.Println(err)
	}

	t2 := time.Now()
	dt := diff(t1, t2)
	out := ""
	if dt.Year != 0 {
		if dt.Month != 0 {
			out = fmt.Sprintf("%syr%smonth ago", dt.Year, dt.Month)
		} else if dt.Day != 0 {
			out = fmt.Sprintf("%syr%sday ago", dt.Year, dt.Day)
		} else {
			out = fmt.Sprintf("%syr ago", dt.Year)
		}
	} else if dt.Month != 0 {
		if dt.Day != 0 {
			out = fmt.Sprintf("%smonth%sday ago", dt.Month, dt.Day)
		} else if dt.Hour != 0 {
			out = fmt.Sprintf("%smonth%shour ago", dt.Month, dt.Hour)
		} else {
			out = fmt.Sprintf("%smonth ago", dt.Month)
		}
	} else if dt.Day != 0 {
		if dt.Hour != 0 {
			out = fmt.Sprintf("%sday%shour ago", dt.Day, dt.Hour)
		} else if dt.Min != 0 {
			out = fmt.Sprintf("%sday%smin ago", dt.Day, dt.Min)
		} else {
			out = fmt.Sprintf("%sday ago", dt.Day)
		}
	} else if dt.Hour != 0 {
		if dt.Min != 0 {
			out = fmt.Sprintf("%hour%smin ago", dt.Hour, dt.Min)
		} else if dt.Sec != 0 {
			out = fmt.Sprintf("%hour%ssec ago", dt.Hour, dt.Sec)
		} else {
			out = fmt.Sprintf("%hour ago", dt.Hour)
		}
	} else if dt.Min != 0 {
		if dt.Sec != 0 {
			out = fmt.Sprintf("%min%ssec ago", dt.Min, dt.Sec)
		} else {
			out = fmt.Sprintf("%min ago", dt.Min)
		}
	} else {
		out = fmt.Sprintf("%ssec ago", dt.Sec)
	}
	return out
}