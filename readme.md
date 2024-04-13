# fyne-simple-chart

A simple chart implementation that an be used in fyne applications.

## Warning

As of now, negative values are not yet implemented. Can be used in plotting to the first quadrant of the cartessian plane.

## Features

- [x] Scatter Plot
- [ ] Bar Chart
- [ ] Pie Chart
- [ ] Column Chart

## Installation

```sh
go get github.com/alexiusacademia/fyne-simple-chart
```

## Usage

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	fynesimplechart "github.com/alexiusacademia/fyne-simple-chart/package"
)

func main() {
	nodes := []fynesimplechart.Node{
		*fynesimplechart.NewNode(1, 3),
		*fynesimplechart.NewNode(2, 0.5),
		*fynesimplechart.NewNode(5, 0.5),
		*fynesimplechart.NewNode(6, 3),
	}

	ch := fynesimplechart.NewGraphWidget(nodes, 10, "Horizontal", "Vertical")

	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(400, 300))
	w.SetContent(ch)
	w.ShowAndRun()
}

```
