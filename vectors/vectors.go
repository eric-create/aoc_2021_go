package vectors

import "eric-create/aoc_2021/utils"

type Vector struct {
	X int
	Y int
}

func (v *Vector) Add(other *Vector) *Vector {
	return &Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func NewVector(x, y int) *Vector {
	return &Vector{x, y}
}

func Up() Vector {
	return Vector{0, -1}
}

func RightUp() Vector {
	return Vector{1, -1}
}

func Right() Vector {
	return Vector{1, 0}
}

func RightDown() Vector {
	return Vector{1, 1}
}

func Down() Vector {
	return Vector{0, 1}
}

func LeftDown() Vector {
	return Vector{-1, 1}
}

func Left() Vector {
	return Vector{-1, 0}
}

func LeftUp() Vector {
	return Vector{-1, -1}
}

// Up, RightUp, Right, RightDown, Down, LeftDown, Left, LeftUp
func AllDirections() []Vector {
	return []Vector{
		Up(),
		RightUp(),
		Right(),
		RightDown(),
		Down(),
		LeftDown(),
		Left(),
		LeftUp(),
	}
}

// Up, Right, Down, Left
func ManhattanDirections() []Vector {
	return []Vector{
		Up(),
		Right(),
		Down(),
		Left(),
	}
}

// Right, Left
func Horizontal() []Vector {
	return []Vector{
		Right(),
		Left(),
	}
}

func Normalize(v *Vector) *Vector {
	x := utils.NormalizeInt(v.X)
	y := utils.NormalizeInt(v.Y)
	return NewVector(x, y)
}

// Returns the vector that points from a to b
func Difference(a, b *Vector) *Vector {
	return NewVector(b.X-a.X, b.Y-a.Y)
}
