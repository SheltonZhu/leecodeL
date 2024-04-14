package reverseinteger

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	x := 123
	expected := 321
	ret := reverse(x)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	x := -123
	expected := -321
	ret := reverse(x)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	x := 120
	expected := 21
	ret := reverse(x)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test4(t *testing.T) {
	x := -2147483412
	expected := -2143847412
	ret := reverse(x)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
