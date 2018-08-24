package ext

import (
	"bytes"
	"runtime"
)

func Stack() string {
	b := bytes.NewBuffer(make([]byte, 4096))
	runtime.Stack(b.Bytes(), false)
	return b.String()
}
