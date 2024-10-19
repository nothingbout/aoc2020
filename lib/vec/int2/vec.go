package int2

import "strings"

type single = int64

type Vec struct {
	X single
	Y single
}

func Zero() Vec            { return Vec{X: 0, Y: 0} }
func Make(X, Y single) Vec { return Vec{X, Y} }
func Add(a, b Vec) Vec     { return Make(a.X+b.X, a.Y+b.Y) }
func Sub(a, b Vec) Vec     { return Make(a.X-b.X, a.Y-b.Y) }

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
	for y := int64(0); y < gridSize.X; y++ {
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
