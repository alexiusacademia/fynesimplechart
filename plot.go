package fynesimplechart

type Plot struct {
	Nodes      []Node
	Ticks      int
	XAxisTitle string
	YAxisTitle string

	ShowLine bool
}

func NewPlot(nodes []Node) *Plot {
	plot := &Plot{
		Nodes: nodes,
	}

	return plot
}
