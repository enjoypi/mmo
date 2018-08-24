package ext

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type child struct {
	I int
	S string
	C *child
}

type parent struct {
	Pi int
	Ps string
	C  *child
}

func TestMergeMapStruct(t *testing.T) {
	v := make(map[string]interface{})
	v["Pi"] = RandomInt(1234567890123456789)
	v["Ps"] = RandomString(20)
	v["C"] = make(map[string]interface{})
	c := v["C"].(map[string]interface{})
	c["I"] = RandomInt(1234567890123456789)
	c["S"] = RandomString(20)
	c["C"] = make(map[string]interface{})
	cc := c["C"].(map[string]interface{})
	cc["I"] = RandomInt(1234567890123456789)
	cc["S"] = RandomString(20)

	var p parent
	MergeMapStruct(v, &p)

	require.Equal(t, v["Pi"], p.Pi)
	//require.Equal(t, v["Ps"], p.Ps)
	//require.NotNil(t, p.C)
	//require.Equal(t, c["I"], p.C.I)
	//require.Equal(t, c["S"], p.C.S)
	//require.NotNil(t, p.C.C)
	//require.Equal(t, cc["I"], p.C.C.I)
	//require.Equal(t, cc["S"], p.C.C.S)
}
