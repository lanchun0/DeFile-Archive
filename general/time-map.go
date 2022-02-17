package general

import "time"

func GenerateTimeStamp() time.Time {
	timestamp := time.Now()
	tm := time.Unix(timestamp.Unix(), 0)
	return tm
}

func Timestamp2Str(t time.Time) string {
	str := t.Format("2006-Jan-02 15:04:05 pm Mon MST")
	return str
}

func Str2Timestamp(str string) (time.Time, error) {
	return time.Parse("2006-Jan-02 15:04:05 pm Mon MST", str)
}
