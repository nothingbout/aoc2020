package int3

import (
	"aoc2020/lib"
	"testing"
)

func TestAdjOffsets26(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		lib.AssertEqual(t, len(AdjOffsets26()), 26)
	})
}
