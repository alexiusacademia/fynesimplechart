package fynesimplechart

import "errors"

type Node struct {
	X float32
	Y float32
}

func NewNode(x float32, y float32) *Node {
	return &Node{X: x, Y: y}
}

func MinY(plots []Plot) (float32, error) {
	nodes := []Node{}

	for _, p := range plots {
		nodes = append(nodes, p.Nodes...)
	}

	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	minimum := nodes[0].Y

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Y < minimum {
			minimum = nodes[i].Y
		}
	}

	return minimum, nil
}

func MaxY(plots []Plot) (float32, error) {
	nodes := []Node{}

	for _, p := range plots {
		nodes = append(nodes, p.Nodes...)
	}

	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	maximum := nodes[0].Y

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Y > maximum {
			maximum = nodes[i].Y
		}
	}

	return maximum, nil
}

func MinX(plots []Plot) (float32, error) {
	nodes := []Node{}

	for _, p := range plots {
		nodes = append(nodes, p.Nodes...)
	}

	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	minimum := nodes[0].X

	for i := 0; i < len(nodes); i++ {
		if nodes[i].X < minimum {
			minimum = nodes[i].X
		}
	}

	return minimum, nil
}

func MaxX(plots []Plot) (float32, error) {
	nodes := []Node{}

	for _, p := range plots {
		nodes = append(nodes, p.Nodes...)
	}

	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	maximum := nodes[0].X

	for i := 0; i < len(nodes); i++ {
		if nodes[i].X > maximum {
			maximum = nodes[i].X
		}
	}

	return maximum, nil
}
