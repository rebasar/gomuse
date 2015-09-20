package gomuse

import (
	"time"
)

// Shamelessly stolen from: https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format

type YMuseTime struct {
	time.Time
}

const ytLayout = "2006-01-02"

var nilTime = (time.Time{}).UnixNano()

func (yt *YMuseTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	timestring := string(b)
	if timestring == "null" {
		return nil
	}
	yt.Time, err = time.Parse(ytLayout, timestring)
	return err
}

func (yt *YMuseTime) MarshalJSON() ([]byte, error) {
	return []byte(yt.Time.Format(ytLayout)), nil
}

func (yt *YMuseTime) IsSet() bool {
	return yt.UnixNano() != nilTime
}
