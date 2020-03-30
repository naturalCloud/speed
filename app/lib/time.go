package lib

import "time"

func MsTime()  int64 {
	return time.Now().UnixNano() / 1e6
}