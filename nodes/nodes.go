package nodes

import (
	"eric-create/aoc_2021/vectors"
	"strconv"
)

type Node struct {
	Position vectors.Vector
	Neigbors [3][3]*Node
	Symbol   string
	Tags     []string
}

func (n *Node) SetNeighbor(neighbor *Node, direction vectors.Vector) {
	n.Neigbors[direction.Y+1][direction.X+1] = neighbor
}

func (n *Node) SymbolToInt() int {
	i, _ := strconv.Atoi(n.Symbol)
	return i
}

func NewNode(x, y int, symbol string) *Node {
	return &Node{
		Position: *vectors.NewVector(x, y),
		Neigbors: [3][3]*Node{},
		Symbol:   symbol,
	}
}

func IntsToField(ints [][]int) [][]*Node {
	nodes := [][]*Node{}

	for y := range ints {
		nodes = append(nodes, []*Node{})

		for x := range ints[y] {
			newNode := NewNode(x, y, strconv.Itoa(ints[y][x]))
			nodes[y] = append(nodes[y], newNode)
		}
	}

	DiscoverNeighbors(&nodes)

	return nodes
}

func DiscoverNeighbors(field *[][]*Node) {
	for y := range *field {
		for x := range (*field)[y] {
			// Left
			if x > 0 {
				(*field)[y][x].SetNeighbor((*field)[y][x-1], vectors.Left())
			}

			// Up
			if y > 0 {
				(*field)[y][x].SetNeighbor((*field)[y-1][x], vectors.Up())
			}

			// Right
			if x < len((*field)[y])-1 {
				(*field)[y][x].SetNeighbor((*field)[y][x+1], vectors.Left())
			}

			// Down
			if y < len(*field)-1 {
				(*field)[y][x].SetNeighbor((*field)[y+1][x], vectors.Left())
			}
		}
	}
}
