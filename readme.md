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
go get github.com/alexiusacademia/fynesimplechart
```

## Usage

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/alexiusacademia/fynesimplechart"
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

	plot := fynesimplechart.NewPlot(nodes)
	plot.ShowLine = true

	scatter := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})

	scatter.Resize(fyne.NewSize(400, 300))

	w.SetContent(scatter)
	w.ShowAndRun()
}

```

<img width="404" alt="Screenshot 2024-04-13 at 12 06 17 PM" src="https://github.com/alexiusacademia/fyne-simple-chart/assets/19258246/8d72061a-0d75-469e-b1fc-3cb6575e2d8f">
