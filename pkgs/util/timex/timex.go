package timex

import (
	"time"
)

func GetDateForAMonthToUnixMilliSecond(y, m int) (startDay, endDay int64) {
	startDayOfMonth, endDayOfMonth := GetDateForAMonth(y, m)

	startDay = startDayOfMonth.UnixMilli()
	endDay = endDayOfMonth.UnixMilli()

	return
}

func GetDateForAMonth(y, m int) (startDay, endDay time.Time) {
	startDay = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
	endDay = time.Date(y, time.Month(m+1), 1, 0, 0, 0, -1, time.Local)
	return
}

func GetDateForADay(y, m, d int) (startOfDay, endOfDay time.Time) {
	startOfDay = time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	endOfDay = time.Date(y, time.Month(m), d+1, 0, 0, 0, -1, time.Local)
	return
}

func GetDateForADayUnixMillisecond(milli int64) (startDay, endDay int64) {
	y, m, d := time.UnixMilli(milli).Date()
	start, end := GetDateForADay(y, int(m), d)

	startDay = start.UnixMilli()
	endDay = end.UnixMilli()
	return
}

func GetDaysInMonth(year int, month time.Month) (days []int32) {
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 1-1)
	for d := firstOfMonth; d.Before(lastOfMonth); d = d.AddDate(0, 0, 1) {
		days = append(days, int32(d.Day()))
	}
	return
}
