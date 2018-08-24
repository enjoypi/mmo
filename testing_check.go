package ext

import (
	"errors"
	"reflect"
	"testing"
)

func CheckEqual(t *testing.T, l interface{}, r interface{}) {

	if reflect.DeepEqual(l, r) {
		return
	}

	t.Fatalf("%T:%+v != %T:%+v\n%s", l, l, r, r, Stack())
}

func CheckLesser(t *testing.T, lhs interface{}, rhs interface{}) {

	lvalue := reflect.ValueOf(lhs)
	rvalue := reflect.ValueOf(rhs)

	if compare(lvalue, rvalue) < 0 {
		return
	}
	t.Fatalf("%T:%+v >= %T:%+v\n%s", lhs, lhs, rhs, rhs, Stack())
}

func CheckLE(t *testing.T, lhs interface{}, rhs interface{}) {

	lvalue := reflect.ValueOf(lhs)
	rvalue := reflect.ValueOf(rhs)

	if compare(lvalue, rvalue) <= 0 {
		return
	}
	t.Fatalf("%T:%+v > %T:%+v\n%s", lhs, lhs, rhs, rhs, Stack())
}

func CheckGreater(t *testing.T, lhs interface{}, rhs interface{}) {

	lvalue := reflect.ValueOf(lhs)
	rvalue := reflect.ValueOf(rhs)

	if compare(lvalue, rvalue) > 0 {
		return
	}
	t.Fatalf("%T:%+v <= %T:%+v\n%s", lhs, lhs, rhs, rhs, Stack())
}

func CheckGE(t *testing.T, lhs interface{}, rhs interface{}) {

	lvalue := reflect.ValueOf(lhs)
	rvalue := reflect.ValueOf(rhs)

	if compare(lvalue, rvalue) >= 0 {
		return
	}
	t.Fatalf("%T:%+v < %T:%+v\n%s", lhs, lhs, rhs, rhs, Stack())
}

func compare(lhs, rhs reflect.Value) int {
	kind := lhs.Kind()

	if kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 {
		lv := lhs.Int()
		rv := rhs.Int()
		if lv > rv {
			return 1
		}
		if lv == rv {
			return 0
		}
		return -1
	}
	if kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 {
		lv := lhs.Uint()
		rv := rhs.Uint()
		if lv > rv {
			return 1
		}
		if lv == rv {
			return 0
		}
		return -1
	}
	//	reflect.Float32:
	//	reflect.Float64:
	//	reflect.Complex64:
	//	reflect.Complex128:
	//	reflect.String:

	panic(errors.New(""))
}

func AssertNoError(t *testing.T, err error, what string) {
	if err != nil {
		t.Fatalf("failed to %s: %s", what, err)
	}
}
