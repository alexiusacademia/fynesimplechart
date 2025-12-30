# Tutorial 6: Working with Negative Values

## Overview

Learn how to work with negative values and plot data across all four quadrants.

**Time to complete:** 10 minutes

## Negative Values Support

FyneSimpleChart fully supports negative values on both axes, automatically positioning the axes at zero when appropriate.

## Basic Negative Values

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Negative Values")
    w.Resize(fyne.NewSize(700, 500))

    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(-3, 2),
        *fynesimplechart.NewNode(-1, 4),
        *fynesimplechart.NewNode(0, 0),
        *fynesimplechart.NewNode(2, -3),
        *fynesimplechart.NewNode(4, -1),
    }

    plot := fynesimplechart.NewPlot(nodes, "Data")
    plot.ShowLine = true
    plot.ShowPoints = true

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("All Four Quadrants")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Practical Example: Profit/Loss

```go
// Monthly profit/loss tracking
profitLoss := []float32{-15, -8, 5, 12, 18, -3, 8, 15, 22, 18, 25, 30}

nodes := []fynesimplechart.Node{}
for month, value := range profitLoss {
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(month+1),
        value,
    ))
}

plot := fynesimplechart.NewPlot(nodes, "Profit/Loss ($K)")
plot.ShowLine = true
plot.LineWidth = 2.5

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
chart.SetChartTitle("Annual Profit/Loss Analysis")
```

## Summary

- ✅ Negative values work automatically
- ✅ Axes position at zero when appropriate
- ✅ All four quadrants supported

## Next Steps

Continue to [Tutorial 7: Grid and Axes Configuration](07-grid-and-axes.md).
