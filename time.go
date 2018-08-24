package ext

import (
	"time"
)

func SleepUntil(ms int64) {
	time.Sleep(time.Unix(ms/1000, ms%1000*int64(time.Millisecond)).Sub(time.Now()))
}
