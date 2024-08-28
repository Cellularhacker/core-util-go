package core

import "time"

const (
	DateTimeFormatAddTimeZone = "2006-01-02 15:04:05 MST"
)

var (
	TimeZoneKST, _ = time.LoadLocation("Asia/Seoul")
	TimeZoneJST, _ = time.LoadLocation("Asia/Tokyo")
	TimeZoneUTC, _ = time.LoadLocation("UTC")
)

func ParseDatetimeInTimeZone(layout string, value string, TZ *time.Location) (time.Time, error) {
	t, e := time.ParseInLocation(layout, value, TZ)
	if e != nil {
		return time.Time{}, e
	}

	switch TZ {
	case TimeZoneKST, TimeZoneJST:
		return t.Add(9 * time.Hour), nil
	case TimeZoneUTC:
		return t.UTC(), nil
	}

	return t.Local(), nil
}
