package lib

import (
	"strconv"
	"testing"
)

func TestMust(t *testing.T) {
	t.Run("return value on success", func(t *testing.T) {
		AssertEqual(t, Must(strconv.ParseInt("1337", 10, 64)), 1337)
	})

	t.Run("panic on error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, but got none")
			}
		}()

		_ = Must(strconv.ParseInt("not an integer", 10, 64))
	})
}

func TestParseInt64(t *testing.T) {
	t.Run("valid integer", func(t *testing.T) {
		value, err := ParseInt64("1337")
		AssertEqual(t, value, 1337)
		AssertEqual(t, err, nil)
	})
	t.Run("invalid integer", func(t *testing.T) {
		_, err := ParseInt64("not an integer")
		AssertGotError(t, err)
	})
	t.Run("overflow", func(t *testing.T) {
		_, err := ParseInt64("23094820394820398402948203984")
		AssertGotError(t, err)
	})
}

func TestSplitLines(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		lines := []string{
			"a",
			"b",
			"",
			"c",
			"",
			"e",
			"f",
			"g",
		}
		got := SplitLines(lines, "")
		want := [][]string{
			{"a", "b"},
			{"c"},
			{"e", "f", "g"},
		}
		AssertEqual(t, got, want)
	})
}

func TestLCD(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		AssertEqual(t, LCD(12, 15), 60)
	})
}
