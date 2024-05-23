package date

import "time"

func GetThisHourTime() time.Time {
	now := time.Now()
	year, month, day := now.Date()
	hour, _, _ := now.Clock()
	return time.Date(year, month, day, hour, 0, 0, 0, now.Location())
}

func GetThisMinuteTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
}
func GetLastHourTime() time.Time {
	now := time.Now()
	year, month, day := now.Date()
	hour, _, _ := now.Clock()
	t := time.Date(year, month, day, hour, 0, 0, 0, now.Location())
	return time.Unix(t.Unix()-3600, 0)
}
func GetThisDayStartTime() time.Time {
	now := time.Now()
	year, month, day := now.Date()
	t := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return time.Unix(t.Unix(), 0)
}
