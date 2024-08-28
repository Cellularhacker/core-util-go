package core

import "time"

const (
	DateTimeFormatAddTimeZone = "2006-01-02 15:04:05 MST"
)

var (
	tzKST, _ = time.LoadLocation("Asia/Seoul")
	tzJST, _ = time.LoadLocation("Asia/Tokyo")
)

func ParseDatetimeInTimeZone(layout string, value string, TZ *time.Location) (time.Time, error) {
	t, e := time.ParseInLocation(layout, value, TZ)
	if e != nil {
		return time.Time{}, e
	}

	switch TZ {
	case tzKST, tzJST:
		return t.Add(9 * time.Hour), nil
	}

	return t.Local(), nil
}
func LocationKST() *time.Location {
	return tzKST
}
func LocationJST() *time.Location {
	return tzJST
}
