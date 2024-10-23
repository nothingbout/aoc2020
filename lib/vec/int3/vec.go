package int3

import (
	"aoc2020/lib"
	"math"
	"strings"
	"sync"
)

type single = int64

type Vec struct {
	X single
	Y single
	Z single
}

// func (v Vec) String() string {
//     return fmt.Sprintf("%v", v)
// }

func Zero() Vec                { return Vec{X: 0, Y: 0, Z: 0} }
func MinValue() Vec            { return Make(math.MinInt64, math.MinInt64, math.MinInt64) }
func MaxValue() Vec            { return Make(math.MaxInt64, math.MaxInt64, math.MaxInt64) }
func Make(X, Y, Z single) Vec  { return Vec{X, Y, Z} }
func Add(a, b Vec) Vec         { return Make(a.X+b.X, a.Y+b.Y, a.Z+b.Z) }
func Sub(a, b Vec) Vec         { return Make(a.X-b.X, a.Y-b.Y, a.Z-b.Z) }
func Scale(a Vec, s int64) Vec { return Make(a.X*s, a.Y*s, a.Z*s) }
func Abs(a Vec) Vec            { return Make(lib.Abs(a.X), lib.Abs(a.Y), lib.Abs(a.Z)) }
func Min(a, b Vec) Vec         { return Make(lib.Min(a.X, b.X), lib.Min(a.Y, b.Y), lib.Min(a.Z, b.Z)) }
func Max(a, b Vec) Vec         { return Make(lib.Max(a.X, b.X), lib.Max(a.Y, b.Y), lib.Max(a.Z, b.Z)) }

var mutexAdjOffsets26 sync.Mutex
var adjOffsets26 []Vec

func AdjOffsets26() []Vec {
	mutexAdjOffsets26.Lock()
	defer mutexAdjOffsets26.Unlock()
	if len(adjOffsets26) == 0 {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if z == 0 && y == 0 && x == 0 {
						continue
					}
					adjOffsets26 = append(adjOffsets26, Make(single(x), single(y), single(z)))
				}
			}
		}
	}
	return adjOffsets26
}

func SprintGrid(gridSize Vec, sprintCell func(pos Vec) string) string {
	var result strings.Builder
	for z := int64(0); z < gridSize.Z; z++ {
		for y := int64(0); y < gridSize.Y; y++ {
			for x := int64(0); x < gridSize.X; x++ {
				pos := Make(x, y, z)
				result.WriteString(sprintCell(pos))
			}
			result.WriteString("\n")
		}
		result.WriteString("\n")
	}
	return result.String()
}

type Bounds struct {
	Pos  Vec
	Size Vec
}

func MakeBounds(pos, size Vec) Bounds { return Bounds{Pos: pos, Size: size} }

func (b Bounds) MinPos() Vec { return b.Pos }
func (b Bounds) MaxPos() Vec { return Add(b.Pos, Sub(b.Size, Make(1, 1, 1))) }

func (b Bounds) ContainsPoint(pos Vec) bool {
	if pos.X < b.Pos.X || pos.X >= b.Pos.X+b.Size.X {
		return false
	}
	if pos.Y < b.Pos.Y || pos.Y >= b.Pos.Y+b.Size.Y {
		return false
	}
	if pos.Z < b.Pos.Z || pos.Z >= b.Pos.Z+b.Size.Z {
		return false
	}
	return true
}
