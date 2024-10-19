package lib

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/exp/constraints"
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

func SplitLines(lines []string, separatorLine string) (groups [][]string) {
	curGroup := []string{}
	for _, line := range lines {
		if line == separatorLine {
			groups = append(groups, curGroup)
			curGroup = []string{}
		} else {
			curGroup = append(curGroup, line)
		}
	}
	groups = append(groups, curGroup)
	return groups
}

func CloneSlice[A any](slice []A) []A {
	return append(slice[:0:0], slice...)
}

func CloneMap[A comparable, B any](orig map[A]B) map[A]B {
	clone := make(map[A]B)
	for k, v := range orig {
		clone[k] = v
	}
	return clone
}

func Min[A constraints.Ordered](x, y A) A {
	if x < y {
		return x
	}
	return y
}

func Max[A constraints.Ordered](x, y A) A {
	if x > y {
		return x
	}
	return y
}
