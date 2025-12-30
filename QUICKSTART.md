# Quick Start Guide

## Installation

```bash
go get github.com/alexiusacademia/fynesimplechart
```

## Basic Usage

### Simple Scatter Plot

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("My Chart")

    // Create data points
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 2),
        *fynesimplechart.NewNode(2, 4),
        *fynesimplechart.NewNode(3, 3),
        *fynesimplechart.NewNode(4, 5),
    }

    // Create plot
    plot := fynesimplechart.NewPlot(nodes, "My Data")

    // Create chart
    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.Resize(fyne.NewSize(600, 400))

    w.SetContent(chart)
    w.ShowAndRun()
}
```

### Line Chart

```go
plot := fynesimplechart.NewPlot(nodes, "Temperature")
plot.ShowLine = true        // Connect points with lines
plot.ShowPoints = true      // Also show the points
plot.LineWidth = 2.0        // Thicker line
```

### Custom Colors

```go
import "image/color"

plot := fynesimplechart.NewPlot(nodes, "Sales")
plot.PlotColor = color.RGBA{R: 220, G: 20, B: 60, A: 255}  // Crimson
plot.PointSize = 5
```

### Multiple Series

```go
plot1 := fynesimplechart.NewPlot(data1, "Series A")
plot2 := fynesimplechart.NewPlot(data2, "Series B")
plot3 := fynesimplechart.NewPlot(data3, "Series C")

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
    *plot1, *plot2, *plot3,
})
chart.SetChartTitle("Comparison Chart")
```

### Negative Values

```go
// Negative values work automatically!
nodes := []fynesimplechart.Node{
    *fynesimplechart.NewNode(-5, 3),
    *fynesimplechart.NewNode(-2, -1),
    *fynesimplechart.NewNode(0, 0),
    *fynesimplechart.NewNode(3, 4),
}
```

## Configuration Options

### Plot Properties

```go
plot := fynesimplechart.NewPlot(nodes, "Title")

// Visibility
plot.ShowLine = true         // Show connecting lines (default: false)
plot.ShowPoints = true       // Show data points (default: true)

// Styling
plot.LineWidth = 2.5         // Line thickness (default: 1.5)
plot.PointSize = 4.0         // Point radius (default: 3.0)
plot.PlotColor = myColor     // Custom color (default: auto-assigned)
```

### Chart Properties

```go
chart := fynesimplechart.NewGraphWidget(plots)

chart.SetChartTitle("My Title")  // Add main chart title
chart.ShowGrid = true            // Show grid lines (default: true)
```

## Running the Examples

```bash
cd examples
go run main.go
```

This will launch a demo application with 6 different chart examples showcasing all features.

## Tips

1. **Point Size**: Use larger points (5-6) for sparse data, smaller (2-3) for dense data
2. **Line Width**: 1.5-2.5 works well for most cases
3. **Grid**: Keep grid enabled for data analysis, disable for presentations
4. **Colors**: Use contrasting colors for multiple series
5. **Negative Values**: The chart automatically positions axes at zero when appropriate

## Common Patterns

### Scatter Plot (Points Only)
```go
plot.ShowPoints = true
plot.ShowLine = false
plot.PointSize = 4
```

### Line Chart (Lines Only)
```go
plot.ShowPoints = false
plot.ShowLine = true
plot.LineWidth = 2
```

### Connected Scatter (Both)
```go
plot.ShowPoints = true
plot.ShowLine = true
plot.PointSize = 3
plot.LineWidth = 1.5
```

## Need More Help?

- See [examples/](examples/) for comprehensive examples
- Check [IMPROVEMENTS.md](IMPROVEMENTS.md) for detailed feature documentation
- Review [examples/README.md](examples/README.md) for API usage examples
