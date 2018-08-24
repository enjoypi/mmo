package ext

import (
	"log"
	"os"
	"runtime"
)

type Trace bool

const (
	DEBUG = 2
	ERROR = 3
	FATAL = 4
)

var (
	tl          *log.Logger = log.New(os.Stdout, "[TRACE] ", log.LstdFlags)
	dl          *log.Logger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	il          *log.Logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	el          *log.Logger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	TraceSwitch Trace
)

func init() {
}

func (t Trace) T() string {
	if t {
		p, _, _, ok := runtime.Caller(2)
		if ok {
			s := runtime.FuncForPC(p).Name()
			tl.Printf("--> [%s]\n", s)
			return s
		}
	}
	return ""
}

func (t Trace) UT(s string) {
	if t && len(s) > 0 {
		tl.Printf("<-- [%s]\n", s)
	}
}

func T() string {
	return TraceSwitch.T()
}

func UT(s string) {
	TraceSwitch.UT(s)
}

func LogDebug(format string, v ...interface{}) {
	dl.Printf(format, v...)
}

func LogInfo(format string, v ...interface{}) {
	il.Printf(format, v...)
}

func LogError(err error) error {
	el.Printf("%s\n%s", err.Error(), Stack())
	return err
}
