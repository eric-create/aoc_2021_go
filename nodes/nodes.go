package nodes

import (
	"eric-create/aoc_2021/vectors"
	"fmt"
	"slices"
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

func (n *Node) GetNeighbor(direction vectors.Vector) *Node {
	return n.Neigbors[direction.Y+1][direction.X+1]
}

func (n *Node) GetNeighbors(directions []vectors.Vector) []*Node {
	neighbors := []*Node{}

	for _, direction := range directions {
		if neighbor := n.GetNeighbor(direction); neighbor != nil {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
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
	for _, row := range *field {
		for _, node := range row {
			for _, direction := range vectors.AllDirections() {
				if position := Navigate(*field, node.Position, direction); position != nil {
					node.SetNeighbor((*field)[position.Y][position.X], direction)
				}
			}
		}
	}
}

// Returns `nil` if there is no neighbor in the specified `direction`, that means that an
// edge of `field` was reached.
func Navigate[T any](field [][]T, position, direction vectors.Vector) *vectors.Vector {
	xMax := len(field[0]) - 1
	yMax := len(field) - 1

	xNew := position.X + direction.X
	yNew := position.Y + direction.Y

	if xNew < 0 || xNew > xMax || yNew < 0 || yNew > yMax {
		return nil
	}

	new := vectors.Vector{X: xNew, Y: yNew}

	return &new
}

func EmptyField(xMax, yMax int) [][]*Node {
	field := [][]*Node{}

	for y := 0; y < yMax; y++ {
		field = append(field, []*Node{})

		for x := 0; x < xMax; x++ {
			field[y] = append(field[y], nil)
		}
	}

	return field
}

func Print(field [][]*Node) {
	for _, nodes := range field {
		for _, node := range nodes {
			if node != nil {
				symbol := node.Symbol
				if len(symbol) == 1 {
					symbol = symbol + " "
				}
				fmt.Print(symbol, " ")
			} else {
				fmt.Print(".. ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintField(field [][]*Node) {
	for _, nodes := range field {
		for _, node := range nodes {
			symbol := node.Symbol
			if len(symbol) == 1 {
				symbol = symbol + " "
			}
			fmt.Print(symbol, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintFieldWithTag(field [][]*Node, tag string) {
	for _, nodes := range field {
		for _, node := range nodes {
			if slices.Contains(node.Tags, tag) {
				symbol := node.Symbol
				if len(symbol) == 1 {
					symbol = symbol + " "
				}
				fmt.Print(symbol, " ")
			} else {
				fmt.Print(".. ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
