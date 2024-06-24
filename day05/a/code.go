package main

import (
	"eric-create/aoc_2021/nodes"
	"eric-create/aoc_2021/utils"
	"eric-create/aoc_2021/vectors"
)

func main() {
	textLines, _ := utils.ReadLines("input.txt")
	lines, xMax, yMax := GetLines((*textLines)[:len(*textLines)-1])

	field := nodes.EmptyField(xMax, yMax)

	for _, line := range *lines {
		line.Walk(&field)
	}
	nodes.Print(field)
}

// func SetLines(lines []*Line, field *[][]nodes.Node) {
// 	for _, line := range lines {

// 	}
// }

func GetLines(textLines []string) (*[]*Line, int, int) {
	lines := []*Line{}

	var line *Line = nil
	xMax := 0
	yMax := 0

	for _, textLine := range textLines {
		line, xMax, yMax = GetLine(textLine, xMax, yMax)
		lines = append(lines, line)
	}

	return &lines, xMax + 1, yMax + 1
}

func GetLine(textLine string, xMax, yMax int) (*Line, int, int) {
	numbers := utils.ExtractInts(textLine)

	xStart := numbers[0]
	yStart := numbers[1]
	xEnd := numbers[2]
	yEnd := numbers[3]
	xMax, yMax = GetMax(xStart, yStart, xEnd, yEnd, xMax, yMax)

	start := vectors.NewVector(xStart, yStart)
	end := vectors.NewVector(xEnd, yEnd)
	line := NewLine(start, end)

	return line, xMax, yMax
}

func GetMax(xStart, yStart, xEnd, yEnd, xMax, yMax int) (int, int) {
	if xStart > xMax {
		xMax = xStart
	}
	if xEnd > xMax {
		xMax = xEnd
	}

	if yStart > yMax {
		yMax = yStart
	}
	if yEnd > yMax {
		yMax = yEnd
	}

	return xMax, yMax
}

type Line struct {
	Start      *vectors.Vector
	End        *vectors.Vector
	Difference *vectors.Vector
	Direction  *vectors.Vector
	Positions  []*nodes.Node
}

func NewLine(start, end *vectors.Vector) *Line {
	difference := vectors.Difference(start, end)
	direction := vectors.Normalize(difference)
	return &Line{start, end, difference, direction, []*nodes.Node{}}
}

func (l *Line) Walk(field *[][]*nodes.Node) {
	cursor := l.Start
	finish := false

	for {

		// Create node if never been there.
		if (*field)[cursor.Y][cursor.X] == nil {
			(*field)[cursor.Y][cursor.X] = nodes.NewNode(cursor.X, cursor.Y, "x")
		}

		// Always tag.
		(*field)[cursor.Y][cursor.X].Tags = append((*field)[cursor.Y][cursor.X].Tags, "1")

		// Also tag the end position.
		if finish {
			break
		}
		cursor = cursor.Add(l.Direction)
		if *cursor == *l.End {
			finish = true
		}
	}
}
