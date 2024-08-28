package core

import "time"

var (
	TimezoneKST, _ = time.LoadLocation("Asia/Seoul")
	TimezoneJST, _ = time.LoadLocation("Asia/Tokyo")
	TimezoneUTC, _ = time.LoadLocation("UTC")
	TimezoneGMT, _ = time.LoadLocation("GMT")
)
