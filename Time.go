package common

import "time"

type times struct {
	standard     string
	timestamp    int64
	dateTime     time.Time
	durationTime time.Duration
}

const DateFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"
const MsTimeFormat = "2006-01-02 15:04:05.999"

var Time = new(times)

func (t *times) getNowDateTime() {
	t.dateTime = time.Now()
}

func (t *times) getNowStandardTime() {
	t.standard = t.dateTime.Format(TimeFormat)
}

func (t *times) getNowTimeUnix() {
	toUnixTime, _ := time.Parse(TimeFormat, t.standard)
	t.timestamp = toUnixTime.Unix()
}

func (t *times) Now() *times {
	t.getNowDateTime()
	t.getNowStandardTime()
	t.getNowTimeUnix()

	return t
}

func (t *times) Tomorrow() *times {
	t.dateTime = time.Now().AddDate(0, 0, +1)
	return t
}

func (t *times) ToTimestamp() int64 {
	return t.timestamp
}

func (t *times) ToStandard() string {
	return t.standard
}

func (t *times) ToDateTime() time.Time {
	return t.dateTime
}

// 获取当天日期字符串
func (t *times) GetTodayDate() string {
	return t.dateTime.Format(DateFormat)
}

// 获取当天日期字符串: 毫秒级时间
func (t *times) GetTodayMsTime() string {
	return t.dateTime.Format(MsTimeFormat)
}

// 获取毫米级时间字符串
func (t *times) GetMsTime(d time.Time) string {
	return d.Format(MsTimeFormat)
}

// 通过格式化获取当天日期字符串
func (t *times) GetTodayDateString(str string) string {
	return t.dateTime.Format(str)
}

// 通过格式化获取当天日期字符串
func (t *times) GetDateString(str string) string {
	return t.dateTime.Format(str)
}

func (t *times) IntToSecond(second int) time.Duration {
	timeSecond := time.Duration(second)
	t.durationTime = timeSecond * time.Second
	return t.durationTime
}
