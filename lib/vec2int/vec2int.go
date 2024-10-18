package vec2int

type single = int64

type Vec struct {
	X single
	Y single
}

func Make(X, Y single) Vec { return Vec{X, Y} }
func Add(a, b Vec) Vec     { return Make(a.X+b.X, a.Y+b.Y) }
func Sub(a, b Vec) Vec     { return Make(a.X-b.X, a.Y-b.Y) }
