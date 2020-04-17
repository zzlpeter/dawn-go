package libs

import "time"

func StampSecond() int64 {
	return time.Now().Unix()
}

func StampNanoSecond() int64 {
	return time.Now().UnixNano()
}

func StampMillSecond() int64 {
	return time.Now().UnixNano() / 1e6
}