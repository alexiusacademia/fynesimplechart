# Tutorial 2: Basic Scatter Plot

## Overview

Learn how to create and customize scatter plots, including:
- Point styling and sizing
- Data organization
- Best practices for scatter plots
- When to use scatter plots

**Time to complete:** 15 minutes

## Prerequisites

- Completed [Tutorial 1: Getting Started](01-getting-started.md)
- Understanding of X/Y coordinate systems

## What is a Scatter Plot?

A scatter plot displays individual data points as dots or markers on a chart. It's ideal for:
- Showing relationships between variables
- Identifying patterns or clusters
- Displaying individual measurements
- Comparing discrete data points

## Creating a Basic Scatter Plot

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Scatter Plot")
    w.Resize(fyne.NewSize(700, 500))

    // Create scattered data points
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1.5, 3.2),
        *fynesimplechart.NewNode(2.8, 5.1),
        *fynesimplechart.NewNode(3.2, 4.8),
        *fynesimplechart.NewNode(4.5, 7.2),
        *fynesimplechart.NewNode(5.1, 6.8),
        *fynesimplechart.NewNode(6.3, 8.5),
        *fynesimplechart.NewNode(7.2, 9.1),
        *fynesimplechart.NewNode(8.5, 10.2),
    }

    plot := fynesimplechart.NewPlot(nodes, "Sample Data")

    // Customize point appearance
    plot.ShowPoints = true    // Show data points (default)
    plot.PointSize = 4        // Larger points for visibility

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Temperature vs Time")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Customizing Point Size

Point size affects visibility and emphasis:

```go
// Small points - for dense data
plot.PointSize = 2

// Medium points - general use (default)
plot.PointSize = 3

// Large points - for emphasis or sparse data
plot.PointSize = 5

// Extra large points - for presentations
plot.PointSize = 7
```

### Example: Different Point Sizes

```go
// Create three plots with different point sizes
smallPoints := fynesimplechart.NewPlot(data1, "Small (2px)")
smallPoints.PointSize = 2

mediumPoints := fynesimplechart.NewPlot(data2, "Medium (4px)")
mediumPoints.PointSize = 4

largePoints := fynesimplechart.NewPlot(data3, "Large (6px)")
largePoints.PointSize = 6

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
    *smallPoints, *mediumPoints, *largePoints,
})
```

## Organizing Your Data

### Method 1: Manual Creation
```go
nodes := []fynesimplechart.Node{
    *fynesimplechart.NewNode(1, 2),
    *fynesimplechart.NewNode(2, 4),
}
```

### Method 2: Loop-Based Creation
```go
nodes := []fynesimplechart.Node{}
for i := 0; i < 10; i++ {
    x := float32(i)
    y := float32(i * 2)
    nodes = append(nodes, *fynesimplechart.NewNode(x, y))
}
```

### Method 3: From Data Source
```go
// Example: Reading from a slice of measurements
type Measurement struct {
    Time  float32
    Value float32
}

measurements := []Measurement{
    {1.0, 23.5},
    {2.0, 24.1},
    {3.0, 23.8},
}

nodes := []fynesimplechart.Node{}
for _, m := range measurements {
    nodes = append(nodes, *fynesimplechart.NewNode(m.Time, m.Value))
}
```

## Practical Example: Sensor Data

```go
package main

import (
    "math/rand"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Sensor Readings")
    w.Resize(fyne.NewSize(800, 600))

    // Simulate sensor readings with some variation
    nodes := []fynesimplechart.Node{}
    baseTemp := float32(20.0)

    for hour := 0; hour < 24; hour++ {
        // Add random variation to simulate real sensor data
        variation := float32(rand.Float64()*4 - 2) // -2 to +2
        temperature := baseTemp + float32(hour)*0.5 + variation

        nodes = append(nodes, *fynesimplechart.NewNode(
            float32(hour),
            temperature,
        ))
    }

    plot := fynesimplechart.NewPlot(nodes, "Temperature °C")
    plot.ShowPoints = true
    plot.PointSize = 4

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("24-Hour Temperature Readings")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Best Practices for Scatter Plots

### 1. Choose Appropriate Point Sizes
- **Dense data (>50 points):** Use size 2-3
- **Normal data (10-50 points):** Use size 3-4
- **Sparse data (<10 points):** Use size 4-6
- **Presentation mode:** Use size 5-7

### 2. Data Sorting
Sorting is optional for scatter plots, but can help if you plan to add lines later:

```go
import "sort"

// Sort by X coordinate
sort.Slice(nodes, func(i, j int) bool {
    return nodes[i].X < nodes[j].X
})
```

### 3. Handle Overlapping Points
If points overlap, consider:
- Reducing point size
- Using semi-transparent colors (Tutorial 8)
- Adding jitter (small random offset)

### 4. Grid Lines
Keep grid lines enabled for data analysis:
```go
chart.ShowGrid = true  // Helps read values
```

## Common Patterns

### Pattern 1: Correlation Analysis
```go
// Perfect positive correlation
for i := 0; i < 10; i++ {
    x := float32(i)
    y := x * 2  // Perfect linear relationship
    nodes = append(nodes, *fynesimplechart.NewNode(x, y))
}
```

### Pattern 2: Scattered Data
```go
// Random scatter
for i := 0; i < 20; i++ {
    x := float32(rand.Float64() * 10)
    y := float32(rand.Float64() * 10)
    nodes = append(nodes, *fynesimplechart.NewNode(x, y))
}
```

### Pattern 3: Clustered Data
```go
// Two clusters
// Cluster 1: bottom-left
for i := 0; i < 10; i++ {
    x := float32(rand.Float64()*3 + 1)
    y := float32(rand.Float64()*3 + 1)
    nodes = append(nodes, *fynesimplechart.NewNode(x, y))
}

// Cluster 2: top-right
for i := 0; i < 10; i++ {
    x := float32(rand.Float64()*3 + 7)
    y := float32(rand.Float64()*3 + 7)
    nodes = append(nodes, *fynesimplechart.NewNode(x, y))
}
```

## Troubleshooting

### Points Too Small
```go
plot.PointSize = 5  // Increase size
```

### Points Overlapping
```go
plot.PointSize = 2  // Decrease size
// Or spread data more
```

### Can't See Points
```go
plot.ShowPoints = true  // Ensure this is set
```

## Exercise

Create a scatter plot showing the relationship between study hours and test scores:

```go
studyHours := []float32{1, 2, 3, 4, 5, 6, 7, 8}
testScores := []float32{45, 52, 58, 65, 70, 78, 85, 92}

// Your code here...
```

<details>
<summary>Solution</summary>

```go
nodes := []fynesimplechart.Node{}
for i := 0; i < len(studyHours); i++ {
    nodes = append(nodes, *fynesimplechart.NewNode(
        studyHours[i],
        testScores[i],
    ))
}

plot := fynesimplechart.NewPlot(nodes, "Test Results")
plot.PointSize = 5

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
chart.SetChartTitle("Study Hours vs Test Scores")
```
</details>

## Summary

You learned:
- ✅ How to create scatter plots
- ✅ How to customize point sizes
- ✅ Different data organization methods
- ✅ Best practices for scatter plots
- ✅ Common data patterns

## Next Steps

Continue to [Tutorial 3: Line Charts](03-line-charts.md) to learn about connecting your data points with lines.
