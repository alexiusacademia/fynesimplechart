# Tutorial 9: Mathematical Visualizations

## Overview

Learn to visualize mathematical functions and equations.

**Time to complete:** 20 minutes

## Prerequisites

- Basic understanding of mathematical functions

## Plotting Functions

### Sine Wave
```go
import "math"

nodes := []fynesimplechart.Node{}
for i := 0; i <= 100; i++ {
    x := float32(i) / 10.0  // 0 to 10
    y := float32(math.Sin(float64(x)))
    nodes = append(nodes, *fynesimplechart.NewNode(x, y))
}

plot := fynesimplechart.NewPlot(nodes, "sin(x)")
plot.ShowLine = true
plot.ShowPoints = false
plot.LineWidth = 2
```

### Polynomial Functions
```go
// Quadratic: y = x²
nodes := []fynesimplechart.Node{}
for x := -5.0; x <= 5.0; x += 0.1 {
    y := x * x
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(x),
        float32(y),
    ))
}

plot := fynesimplechart.NewPlot(nodes, "y = x²")
```

### Exponential Functions
```go
// Exponential growth: y = e^x
nodes := []fynesimplechart.Node{}
for x := 0.0; x <= 5.0; x += 0.1 {
    y := math.Exp(x)
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(x),
        float32(y),
    ))
}

plot := fynesimplechart.NewPlot(nodes, "e^x")
```

## Complete Example: Multiple Functions

```go
package main

import (
    "math"
    "image/color"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Mathematical Functions")
    w.Resize(fyne.NewSize(900, 650))

    // Sine function
    sineNodes := []fynesimplechart.Node{}
    for i := 0; i <= 60; i++ {
        x := float32(i) / 10.0
        y := float32(math.Sin(float64(x)))
        sineNodes = append(sineNodes, *fynesimplechart.NewNode(x, y))
    }
    sinePlot := fynesimplechart.NewPlot(sineNodes, "sin(x)")
    sinePlot.ShowLine = true
    sinePlot.ShowPoints = false
    sinePlot.LineWidth = 2
    sinePlot.PlotColor = color.RGBA{R: 255, G: 99, B: 71, A: 255}

    // Cosine function
    cosineNodes := []fynesimplechart.Node{}
    for i := 0; i <= 60; i++ {
        x := float32(i) / 10.0
        y := float32(math.Cos(float64(x)))
        cosineNodes = append(cosineNodes, *fynesimplechart.NewNode(x, y))
    }
    cosinePlot := fynesimplechart.NewPlot(cosineNodes, "cos(x)")
    cosinePlot.ShowLine = true
    cosinePlot.ShowPoints = false
    cosinePlot.LineWidth = 2
    cosinePlot.PlotColor = color.RGBA{R: 65, G: 105, B: 225, A: 255}

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
        *sinePlot,
        *cosinePlot,
    })
    chart.SetChartTitle("Trigonometric Functions")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Tips for Smooth Curves

1. **Use small increments**: 0.05 to 0.1 for smooth curves
2. **Hide points**: Too many points creates clutter
3. **Appropriate range**: Choose X range that shows interesting behavior

## Common Functions

### Logarithmic
```go
for x := 0.1; x <= 10.0; x += 0.1 {
    y := math.Log(x)  // Natural log
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(x), float32(y),
    ))
}
```

### Damped Oscillation
```go
for x := 0.0; x <= 10.0; x += 0.1 {
    y := math.Sin(x) * math.Exp(-x/10.0)
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(x), float32(y),
    ))
}
```

## Summary

- ✅ Plotting mathematical functions
- ✅ Creating smooth curves
- ✅ Multiple function comparison
- ✅ Common function types

## Next Steps

Continue to [Tutorial 11: Real-time Data](11-realtime-data.md).
