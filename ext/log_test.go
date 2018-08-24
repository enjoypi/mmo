package ext

import (
	"testing"
)

var (
	tt Trace = true
)

func TestTrace(t *testing.T) {
	if testing.Verbose() {
		defer tt.UT(tt.T())
	}
}
