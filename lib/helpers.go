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

func FilterSlice[A any](slice []A, predicate func(A) bool) []A {
	var result []A
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func MapSlice[A, B any](slice []A, f func(A) B) []B {
	result := make([]B, 0, len(slice))
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}

func SliceRemoveAt[A any](slice []A, index int) []A {
	return append(slice[:index], slice[index+1:]...)
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

func Abs[A constraints.Float | constraints.Integer](x A) A {
	if x < 0 {
		return -x
	}
	return x
}

// GCD: greatest common divisor (euclidean algorithm)
func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCD: least common denominator
func LCD(a, b int64) int64 {
	return (a * b) / GCD(a, b)
}
