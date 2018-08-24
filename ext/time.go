package ext

import (
	"math/rand"
	"time"
)

func SleepUntil(ms int64) {
	time.Sleep(time.Unix(ms/1000, ms%1000*int64(time.Millisecond)).Sub(time.Now()))
}

func SleepRandMS(min, max int64) {
	time.Sleep(time.Duration(min+rand.Int63n(max-min)) * time.Millisecond)
}
