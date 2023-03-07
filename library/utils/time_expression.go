package utils

/*

绝对参考:
year  -> 今年起始时间
month -> 本月第一天
week  -> 本周周一
today -> 今天起始时间
now   -> 当前时间, 精确到秒

增量单位:
Y  -> 年
M  -> 月
w  -> 周
d  -> 天
m  -> 分钟
h  -> 小时
s  -> 秒

例如:
year+1M+2w-3d+4h+1m
today-1d

*/

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var baseFuncMap = make(map[string]func() time.Time)

var deltaFuncMap = make(map[string]func(time2 time.Time, d int) time.Time)

func init() {
	baseFuncMap["year"] = YearBegin
	baseFuncMap["month"] = MonthBegin
	baseFuncMap["week"] = WeekBegin
	baseFuncMap["today"] = TodayBegin
	baseFuncMap["now"] = NowSec

	//
	deltaFuncMap["Y"] = func(time2 time.Time, d int) time.Time {
		return time2.AddDate(d, 0, 0)
	}
	deltaFuncMap["M"] = func(time2 time.Time, d int) time.Time {
		return time2.AddDate(0, d, 0)
	}
	deltaFuncMap["w"] = func(time2 time.Time, d int) time.Time {
		return time2.AddDate(0, 0, 7*d)
	}
	deltaFuncMap["d"] = func(time2 time.Time, d int) time.Time {
		return time2.AddDate(0, 0, d)
	}
	deltaFuncMap["h"] = func(time2 time.Time, d int) time.Time {
		return time2.Add(time.Duration(d) * time.Hour)
	}
	deltaFuncMap["m"] = func(time2 time.Time, d int) time.Time {
		return time2.Add(time.Duration(d) * time.Minute)
	}
	deltaFuncMap["s"] = func(time2 time.Time, d int) time.Time {
		return time2.Add(time.Duration(d) * time.Second)
	}

}

// YearBegin return the 1st day begin time of this year
func YearBegin() time.Time {
	now := time.Now()
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
}

// MonthBegin return the 1st day begin time of this month
func MonthBegin() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// WeekBegin return begin time of monday of this week
func WeekBegin() time.Time {
	today := TodayBegin()
	wd := today.Weekday()
	if wd == time.Monday {
		return today
	}
	if wd == time.Sunday {
		wd = 7
	}
	wd--
	return today.AddDate(0, 0, -int(wd))
}

// TodayBegin return begin time of today
func TodayBegin() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// NowSec return current Time, trim nanoseconds
func NowSec() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())
}

// see ParseTimeExp
func ParseMultiTimeExp(exps ...string) ([]*time.Time, error) {
	times := make([]*time.Time, len(exps))
	for i, e := range exps {
		t, err := ParseTimeExp(e)
		if err != nil {
			return nil, err
		}
		times[i] = t
	}
	return times, nil
}

// ParseTimeExp parse expression like "today-1h" to time.Time
//
// the expresion should be in form like [basepart][segment]*
//
// basepart is required, should be one of year,month,week,today,now.
//
// segment is optional, no quantity limit. should be in form [-+]num?unit
//
// sign flag is required, should be - or +.
// num is optional, default 1.
// unit is required, should one of Y,M,w,d,h,m,s
func ParseTimeExp(exp string) (*time.Time, error) {
	segs := parseSegments(exp)
	if len(segs) == 0 {
		return nil, fmt.Errorf("invalid expression: %q. no segment found", exp)
	}
	base := "now"

	if _, ok := baseFuncMap[segs[0]]; ok {
		base = segs[0]
	} else {
		return nil, fmt.Errorf("invalid expression: %q. should start with one of: year, month, week, today, now", exp)
	}
	baseDate := baseFuncMap[base]()
	for i := 1; i < len(segs); i++ {
		d, s := decodeSegment(segs[i])
		if f, ok := deltaFuncMap[s]; ok {
			baseDate = f(baseDate, d)
		} else {
			return nil, fmt.Errorf("invalid segment: %q in %q. should be one of: Y, M, w, d, h, m, s", s, segs[i])
		}

	}

	return &baseDate, nil
}

func parseSegments(exp string) []string {
	l := len(exp)
	var segs []string
	lastI := 0
	for i := 0; i < l; i++ {
		if exp[i] == '+' || exp[i] == '-' {
			s := strings.TrimSpace(exp[lastI:i])
			if len(s) > 1 {
				segs = append(segs, s)
			}
			lastI = i
		}
	}
	s := strings.TrimSpace(exp[lastI:l])
	if len(s) > 1 {
		segs = append(segs, s)
	}

	return segs
}

func decodeSegment(seg string) (int, string) {
	d := 1
	s := seg[1:]
	for i, r := range seg {
		if i == 0 {
			continue
		}
		if r < '0' || r > '9' {
			ds := seg[1:i]
			if ds != "" {
				d, _ = strconv.Atoi(ds)
			}
			s = seg[i:]
			break
		}
	}
	if seg[0] == '-' {
		d = -d
	}
	return d, s
}
