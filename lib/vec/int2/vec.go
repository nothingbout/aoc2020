package int2

import (
	"aoc2020/lib"
	"math"
	"strings"
)

type single = int64

type Vec struct {
	X single
	Y single
}

func Zero() Vec                { return Vec{X: 0, Y: 0} }
func MinValue() Vec            { return Make(math.MinInt64, math.MinInt64) }
func MaxValue() Vec            { return Make(math.MaxInt64, math.MaxInt64) }
func Make(X, Y single) Vec     { return Vec{X, Y} }
func Add(a, b Vec) Vec         { return Make(a.X+b.X, a.Y+b.Y) }
func Sub(a, b Vec) Vec         { return Make(a.X-b.X, a.Y-b.Y) }
func Scale(a Vec, s int64) Vec { return Make(a.X*s, a.Y*s) }
func Abs(a Vec) Vec            { return Make(lib.Abs(a.X), lib.Abs(a.Y)) }
func Min(a, b Vec) Vec         { return Make(lib.Min(a.X, b.X), lib.Min(a.Y, b.Y)) }
func Max(a, b Vec) Vec         { return Make(lib.Max(a.X, b.X), lib.Max(a.Y, b.Y)) }

var AdjOffsets8 = []Vec{
	{X: -1, Y: -1},
	{X: 0, Y: -1},
	{X: 1, Y: -1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: -1, Y: 1},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
}

func SprintGrid(gridSize Vec, sprintCell func(pos Vec) string) string {
	var result strings.Builder
	for y := int64(0); y < gridSize.Y; y++ {
		for x := int64(0); x < gridSize.X; x++ {
			pos := Make(x, y)
			result.WriteString(sprintCell(pos))
		}
		result.WriteString("\n")
	}
	return result.String()
}

func IsInBounds(pos, bPos, bSize Vec) bool {
	if pos.X < bPos.X || pos.X >= bPos.X+bSize.X {
		return false
	}
	if pos.Y < bPos.Y || pos.Y >= bPos.Y+bSize.Y {
		return false
	}
	return true
}
