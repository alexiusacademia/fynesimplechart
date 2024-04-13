package fynesimplechart

import "errors"

type Node struct {
	X float32
	Y float32
}

func NewNode(x float32, y float32) *Node {
	return &Node{X: x, Y: y}
}

func MinY(nodes []Node) (float32, error) {
	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	minimum := nodes[0].Y
	hasNegative := false

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Y < minimum {
			minimum = nodes[i].Y
		}
		if nodes[i].Y < 0 {
			hasNegative = true
		}
	}

	if hasNegative {
		minimum -= 1 // Add 1 unit for clearance
	}

	return minimum, nil
}

func MaxY(nodes []Node) (float32, error) {
	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	maximum := nodes[0].Y

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Y > maximum {
			maximum = nodes[i].Y
		}
	}

	maximum += 1 // Add 1 unit for clearance

	return maximum, nil
}

func MinX(nodes []Node) (float32, error) {
	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	minimum := nodes[0].X
	hasNegative := false

	for i := 0; i < len(nodes); i++ {
		if nodes[i].X < minimum {
			minimum = nodes[i].X
		}
		if nodes[i].X < 0 {
			hasNegative = true
		}
	}

	if hasNegative {
		minimum -= 1 // Add 1 unit for clearance
	} else {
		minimum = 0
	}

	// Check if all nodes are in the right quadrant

	return minimum, nil
}

func MaxX(nodes []Node) (float32, error) {
	if len(nodes) == 0 {
		return 0, errors.New("No nodes to iterate.")
	}

	maximum := nodes[0].X

	for i := 0; i < len(nodes); i++ {
		if nodes[i].X > maximum {
			maximum = nodes[i].X
		}
	}

	maximum += 1 // Add 1 unit for clearance

	return maximum, nil
}
