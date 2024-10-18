package lib

import (
	"reflect"
	"strconv"
	"testing"
)

func Must[A any](x A, err error) A {
	if err != nil {
		panic(err)
	}
	return x
}

func ParseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func AssertEqual[A any](t *testing.T, got, want A) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertGotError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected error but got nil")
	}
}
