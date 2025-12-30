# Tutorial 9: Area Fills

## Overview

Learn how to create professional area charts with fills to zero or between curves.

**Time to complete:** 20 minutes

## What Are Area Fills?

Area fills shade the region:
- **Below a curve** (fill to Y=0 axis)
- **Between two curves** (confidence intervals, ranges, etc.)

They're perfect for:
- Revenue/profit visualization
- Temperature ranges
- Statistical confidence intervals
- Stock price ranges

## Basic Area Fill to Zero

Fill the area from a curve down to the Y=0 axis.

### Example: Revenue Over Time

```go
package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/alexiusacademia/fynesimplechart"
)

func main() {
	a := app.New()
	w := a.NewWindow("Revenue Area Chart")

	// Create revenue data
	nodes := []fynesimplechart.Node{}
	for i := 0; i <= 20; i++ {
		x := float32(i)
		y := float32(math.Sin(float64(i)*0.3)*5 + 6)
		nodes = append(nodes, *fynesimplechart.NewNode(x, y))
	}

	// Create plot with area fill
	plot := fynesimplechart.NewPlot(nodes, "Revenue")
	plot.ShowLine = true
	plot.ShowPoints = false
	plot.LineWidth = 2.5
	plot.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}

	// Enable fill to zero
	plot.FillArea = true
	plot.FillToZero = true

	chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
	chart.SetChartTitle("Monthly Revenue ($K)")
	chart.Resize(fyne.NewSize(800, 600))

	w.SetContent(chart)
	w.ShowAndRun()
}
```

### Key Properties

```go
plot.FillArea = true      // Enable area fill
plot.FillToZero = true    // Fill from curve to Y=0
```

The fill will automatically use the plot color with 30% transparency.

## Custom Fill Color

Override the default fill color:

```go
plot.FillArea = true
plot.FillToZero = true
plot.FillColor = color.RGBA{R: 31, G: 119, B: 180, A: 100}  // Custom transparency
```

### Transparency Guidelines

```go
// Subtle (recommended for overlapping fills)
A: 50   // ~20% opacity

// Normal (good for single fills)
A: 76   // ~30% opacity (default)

// Strong (for emphasis)
A: 128  // ~50% opacity

// Very strong (may obscure grid)
A: 200  // ~78% opacity
```

## Fill Between Two Curves

Shade the area between two data series.

### Example: Temperature Range

```go
package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/alexiusacademia/fynesimplechart"
)

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Range")

	// Maximum temperature
	maxTemp := []fynesimplechart.Node{}
	for hour := 0; hour <= 24; hour++ {
		x := float32(hour)
		y := float32(20 + math.Sin((float64(hour)-6)*math.Pi/12)*8)
		maxTemp = append(maxTemp, *fynesimplechart.NewNode(x, y))
	}

	maxPlot := fynesimplechart.NewPlot(maxTemp, "Max Temperature")
	maxPlot.ShowLine = true
	maxPlot.ShowPoints = false
	maxPlot.LineWidth = 2
	maxPlot.PlotColor = color.RGBA{R: 244, G: 67, B: 54, A: 255} // Red

	// Minimum temperature
	minTemp := []fynesimplechart.Node{}
	for hour := 0; hour <= 24; hour++ {
		x := float32(hour)
		y := float32(15 + math.Sin((float64(hour)-6)*math.Pi/12)*5)
		minTemp = append(minTemp, *fynesimplechart.NewNode(x, y))
	}

	minPlot := fynesimplechart.NewPlot(minTemp, "Min Temperature")
	minPlot.ShowLine = true
	minPlot.ShowPoints = false
	minPlot.LineWidth = 2
	minPlot.PlotColor = color.RGBA{R: 33, G: 150, B: 243, A: 255} // Blue

	// Fill between min and max
	minPlot.FillArea = true
	minPlot.FillToPlotIdx = 0  // Fill to first plot (maxPlot)
	minPlot.FillColor = color.RGBA{R: 158, G: 158, B: 158, A: 80}

	chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
		*maxPlot,
		*minPlot,
	})
	chart.SetChartTitle("Daily Temperature Range")
	chart.Resize(fyne.NewSize(800, 600))

	w.SetContent(chart)
	w.ShowAndRun()
}
```

### How Fill Between Works

```go
// Plot order matters!
plots := []fynesimplechart.Plot{
	*upperPlot,  // Index 0
	*lowerPlot,  // Index 1
}

// Fill from lowerPlot to upperPlot
lowerPlot.FillArea = true
lowerPlot.FillToPlotIdx = 0  // Fill to index 0 (upperPlot)
```

**Important:** The plot index refers to the position in the plots array.

## Confidence Intervals

A common use case for fill between curves.

```go
// Mean value
meanNodes := []fynesimplechart.Node{}
for i := 0; i <= 30; i++ {
	x := float32(i)
	y := 5 + float32(i)*0.3
	meanNodes = append(meanNodes, *fynesimplechart.NewNode(x, y))
}

meanPlot := fynesimplechart.NewPlot(meanNodes, "Mean")
meanPlot.ShowLine = true
meanPlot.LineWidth = 2.5
meanPlot.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}

// Upper confidence bound
upperNodes := []fynesimplechart.Node{}
for i := 0; i <= 30; i++ {
	x := float32(i)
	y := 5 + float32(i)*0.3 + 2  // Mean + 2
	upperNodes = append(upperNodes, *fynesimplechart.NewNode(x, y))
}

upperPlot := fynesimplechart.NewPlot(upperNodes, "Upper 95% CI")
upperPlot.ShowLine = true
upperPlot.LineWidth = 1
upperPlot.PlotColor = color.RGBA{R: 150, G: 150, B: 150, A: 200}

// Lower confidence bound
lowerNodes := []fynesimplechart.Node{}
for i := 0; i <= 30; i++ {
	x := float32(i)
	y := 5 + float32(i)*0.3 - 2  // Mean - 2
	lowerNodes = append(lowerNodes, *fynesimplechart.NewNode(x, y))
}

lowerPlot := fynesimplechart.NewPlot(lowerNodes, "Lower 95% CI")
lowerPlot.ShowLine = true
lowerPlot.LineWidth = 1
lowerPlot.PlotColor = color.RGBA{R: 150, G: 150, B: 150, A: 200}

// Fill the confidence interval
lowerPlot.FillArea = true
lowerPlot.FillToPlotIdx = 1  // Fill to upperPlot (index 1)
lowerPlot.FillColor = color.RGBA{R: 31, G: 119, B: 180, A: 50}

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
	*meanPlot,   // Index 0
	*upperPlot,  // Index 1
	*lowerPlot,  // Index 2
})
```

## Multiple Fills

You can have multiple area fills in one chart.

```go
// Plot 1: Fill to zero
plot1.FillArea = true
plot1.FillToZero = true

// Plot 2: Fill to another plot
plot2.FillArea = true
plot2.FillToPlotIdx = 0

// Both fills will render correctly
chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
	*plot1,
	*plot2,
})
```

## Design Best Practices

### 1. Color Selection for Fills

```go
// Good: Use color related to the line
lineColor := color.RGBA{R: 31, G: 119, B: 180, A: 255}
fillColor := color.RGBA{R: 31, G: 119, B: 180, A: 76}  // Same RGB, low alpha

// Good: Use neutral color for ranges
fillColor := color.RGBA{R: 158, G: 158, B: 158, A: 80}  // Gray

// Avoid: Unrelated colors
lineColor := color.RGBA{R: 31, G: 119, B: 180, A: 255}  // Blue line
fillColor := color.RGBA{R: 244, G: 67, B: 54, A: 76}    // Red fill (confusing!)
```

### 2. Transparency Levels

```go
// Single fill: Higher transparency OK
plot.FillColor = color.RGBA{R: 31, G: 119, B: 180, A: 100}

// Multiple overlapping fills: Lower transparency
fill1.FillColor = color.RGBA{R: 31, G: 119, B: 180, A: 50}
fill2.FillColor = color.RGBA{R: 244, G: 67, B: 54, A: 50}
```

### 3. Line Width with Fills

```go
// Area fill present: Use thinner lines
plot.LineWidth = 2.0

// No fill: Can use thicker lines
plot.LineWidth = 3.0
```

## Common Use Cases

### Revenue/Profit Charts

```go
plot.FillArea = true
plot.FillToZero = true
plot.PlotColor = color.RGBA{R: 76, G: 175, B: 80, A: 255}  // Green
```

### Stock Price Ranges

```go
// High-Low range
lowPlot.FillArea = true
lowPlot.FillToPlotIdx = 0  // Fill to highPlot
```

### Weather Data

```go
// Temperature min/max
minPlot.FillArea = true
minPlot.FillToPlotIdx = 0  // Fill to maxPlot
minPlot.FillColor = color.RGBA{R: 255, G: 193, B: 7, A: 100}  // Amber
```

### Statistical Analysis

```go
// Confidence intervals
lowerCI.FillArea = true
lowerCI.FillToPlotIdx = 1  // Fill to upperCI
lowerCI.FillColor = color.RGBA{R: 31, G: 119, B: 180, A: 50}  // Light blue
```

## Troubleshooting

### Fill Not Appearing

Check these:
```go
// 1. FillArea must be true
plot.FillArea = true

// 2. For fill to zero:
plot.FillToZero = true

// 3. For fill between:
plot.FillToPlotIdx = 0  // Valid index (>= 0)

// 4. Plots must be added to chart
chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
	*targetPlot,  // Index 0
	*sourcePlot,  // This one has FillToPlotIdx = 0
})
```

### Fill in Wrong Place

```go
// Ensure correct plot index
plots := []fynesimplechart.Plot{
	*upperPlot,  // Index 0
	*lowerPlot,  // Index 1
}

// Fill from lower to upper
lowerPlot.FillToPlotIdx = 0  // Correct!
```

### Fill Too Dark/Light

```go
// Adjust alpha channel (0-255)
plot.FillColor = color.RGBA{
	R: 31,
	G: 119,
	B: 180,
	A: 76,  // Increase for darker, decrease for lighter
}
```

## Complete Example

See [examples/area_fill_demo.go](../examples/area_fill_demo.go) for 4 complete area fill examples:
1. Fill to Zero (Revenue)
2. Fill Between Curves
3. Temperature Range
4. Confidence Interval

## Summary

Area fills in FyneSimpleChart:
- ✅ Fill to zero axis
- ✅ Fill between two curves
- ✅ Custom fill colors
- ✅ Automatic transparency
- ✅ Multiple fills per chart
- ✅ Smooth rendering with 500 interpolation points

## Exercises

1. **Basic Fill**: Create a chart showing daily website traffic with area fill to zero
2. **Range Visualization**: Plot stock high/low prices with fill between them
3. **Multiple Ranges**: Show three different product sales with overlapping fills
4. **Custom Colors**: Create a temperature chart with warm colors (red/orange) for the fill

## Next Steps

Continue to [Tutorial 10: Mathematical Visualizations](10-mathematical-functions.md) to learn how to plot mathematical functions.
