package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	fynesimplechart "github.com/alexiusacademia/fyne-simple-chart/package"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	nodes := []fynesimplechart.Node{
		*fynesimplechart.NewNode(1, 1),
		*fynesimplechart.NewNode(1, 3),
		*fynesimplechart.NewNode(4, 0),
		*fynesimplechart.NewNode(5, 2),
	}

	scatter := fynesimplechart.NewGraphWidget(nodes, 5, "Horizontal", "Vertical")
	scatter.Resize(fyne.NewSize(400, 300))

	w.SetContent(scatter)
	w.ShowAndRun()
}
