# Tutorial 1: Getting Started

## Overview

In this tutorial, you'll learn how to:
- Install FyneSimpleChart
- Set up your development environment
- Create your first simple chart
- Run and view your chart

**Time to complete:** 10 minutes

## Prerequisites

- Go 1.22.0 or later installed
- Basic Go programming knowledge
- Terminal/command line familiarity

## Step 1: Install Go and Fyne

First, ensure you have Go installed:

```bash
go version
```

You should see output like: `go version go1.22.0 darwin/amd64`

## Step 2: Create a New Project

Create a new directory for your project:

```bash
mkdir my-chart-app
cd my-chart-app
go mod init my-chart-app
```

## Step 3: Install FyneSimpleChart

Install the library and its dependencies:

```bash
go get github.com/alexiusacademia/fynesimplechart
go get fyne.io/fyne/v2
```

## Step 4: Create Your First Chart

Create a file named `main.go`:

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    // Create a new Fyne application
    a := app.New()
    w := a.NewWindow("My First Chart")
    w.Resize(fyne.NewSize(600, 400))

    // Create some data points
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 2),
        *fynesimplechart.NewNode(2, 4),
        *fynesimplechart.NewNode(3, 3),
        *fynesimplechart.NewNode(4, 5),
        *fynesimplechart.NewNode(5, 7),
    }

    // Create a plot
    plot := fynesimplechart.NewPlot(nodes, "My Data")

    // Create the chart widget
    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})

    // Display the chart
    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Step 5: Run Your Chart

Run your application:

```bash
go run main.go
```

You should see a window with your first chart displaying 5 data points with:
- Grid lines
- Axis labels with numeric values
- A legend showing "My Data"
- A border around the plot area

## Understanding the Code

Let's break down what each part does:

### Creating Data Points
```go
nodes := []fynesimplechart.Node{
    *fynesimplechart.NewNode(1, 2),  // X=1, Y=2
    *fynesimplechart.NewNode(2, 4),  // X=2, Y=4
    // ...
}
```
Each `Node` represents a point with X and Y coordinates.

### Creating a Plot
```go
plot := fynesimplechart.NewPlot(nodes, "My Data")
```
A `Plot` contains your data points and a title that appears in the legend.

### Creating the Chart
```go
chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
```
The `GraphWidget` can contain multiple plots (we'll cover this in Tutorial 5).

## Common Issues

### Issue: "no required module provides package"
**Solution:** Run `go mod tidy` to download dependencies.

### Issue: Window doesn't appear
**Solution:** Make sure you have the required system dependencies for Fyne:
- **macOS:** Xcode command line tools
- **Linux:** gcc, libgl1-mesa-dev, xorg-dev
- **Windows:** gcc via TDM-GCC or similar

### Issue: Chart is too small
**Solution:** Use `chart.Resize()` or `w.Resize()` to set a larger size:
```go
chart.Resize(fyne.NewSize(800, 600))
```

## Congratulations!

You've created your first chart with FyneSimpleChart!

## What You Learned

- ✅ How to install FyneSimpleChart
- ✅ How to create data points
- ✅ How to create a plot
- ✅ How to display a chart in a window

## Next Steps

Continue to [Tutorial 2: Basic Scatter Plot](02-basic-scatter-plot.md) to learn more about scatter plot features and customization.

## Quick Reference

```go
// Basic chart template
nodes := []fynesimplechart.Node{
    *fynesimplechart.NewNode(x, y),
}
plot := fynesimplechart.NewPlot(nodes, "Title")
chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
```
