# Tutorial 3: Line Charts

## Overview

Learn how to create line charts by connecting data points:
- Enabling line connections
- Customizing line appearance
- Combining lines and points
- When to use line charts

**Time to complete:** 15 minutes

## Prerequisites

- Completed [Tutorial 2: Basic Scatter Plot](02-basic-scatter-plot.md)

## What is a Line Chart?

A line chart connects data points with lines to show:
- Trends over time
- Continuous relationships
- Sequential data progression
- Rate of change

## Creating Your First Line Chart

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Line Chart")
    w.Resize(fyne.NewSize(700, 500))

    // Create sequential data points
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 3),
        *fynesimplechart.NewNode(2, 5),
        *fynesimplechart.NewNode(3, 4),
        *fynesimplechart.NewNode(4, 7),
        *fynesimplechart.NewNode(5, 6),
        *fynesimplechart.NewNode(6, 8),
    }

    plot := fynesimplechart.NewPlot(nodes, "Sales")

    // Enable line connections
    plot.ShowLine = true
    plot.LineWidth = 2

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Monthly Sales Trend")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Line Chart Variations

### 1. Line Only (No Points)
Best for smooth trends and continuous data:

```go
plot := fynesimplechart.NewPlot(nodes, "Temperature")
plot.ShowLine = true     // Show lines
plot.ShowPoints = false  // Hide points
plot.LineWidth = 2.5     // Thicker line
```

### 2. Line with Points
Best for discrete measurements with trends:

```go
plot := fynesimplechart.NewPlot(nodes, "Measurements")
plot.ShowLine = true    // Show lines
plot.ShowPoints = true  // Show points
plot.LineWidth = 2
plot.PointSize = 4
```

### 3. Points Only (Scatter)
Best for non-sequential or independent data:

```go
plot := fynesimplechart.NewPlot(nodes, "Scattered")
plot.ShowLine = false   // Hide lines
plot.ShowPoints = true  // Show points
plot.PointSize = 5
```

## Customizing Line Width

```go
// Thin line - subtle emphasis
plot.LineWidth = 1.0

// Normal line - default (good for most cases)
plot.LineWidth = 1.5

// Medium line - clear visibility
plot.LineWidth = 2.5

// Thick line - strong emphasis
plot.LineWidth = 3.5

// Very thick - presentation mode
plot.LineWidth = 5.0
```

## Practical Example: Temperature Monitoring

```go
package main

import (
    "math"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Temperature Monitoring")
    w.Resize(fyne.NewSize(800, 600))

    // Generate 24-hour temperature data
    nodes := []fynesimplechart.Node{}

    for hour := 0; hour <= 24; hour++ {
        // Simulate temperature variation using sine wave
        // Lowest at 4 AM, highest at 4 PM
        x := float32(hour)
        baseline := float32(20.0)
        variation := float32(math.Sin((float64(hour)-4)*math.Pi/12) * 8)
        y := baseline + variation

        nodes = append(nodes, *fynesimplechart.NewNode(x, y))
    }

    plot := fynesimplechart.NewPlot(nodes, "Temperature (°C)")
    plot.ShowLine = true
    plot.ShowPoints = false  // Smooth line without points
    plot.LineWidth = 2.5

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("24-Hour Temperature Profile")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Creating Smooth Curves

For smooth, continuous data, use many points:

```go
import "math"

// Create smooth sine wave
nodes := []fynesimplechart.Node{}

// Use small increments for smoothness
for i := 0; i <= 100; i++ {
    x := float32(i) / 10.0  // 0 to 10 in 0.1 increments
    y := float32(math.Sin(float64(x)))

    nodes = append(nodes, *fynesimplechart.NewNode(x, y))
}

plot := fynesimplechart.NewPlot(nodes, "sin(x)")
plot.ShowLine = true
plot.ShowPoints = false  // Too many points to show
plot.LineWidth = 2
```

## Multi-Segment Line Charts

Create line charts with multiple connected segments:

```go
// Stock price example
nodes := []fynesimplechart.Node{
    *fynesimplechart.NewNode(1, 100),   // Day 1
    *fynesimplechart.NewNode(2, 105),   // Day 2 - increase
    *fynesimplechart.NewNode(3, 103),   // Day 3 - decrease
    *fynesimplechart.NewNode(4, 108),   // Day 4 - increase
    *fynesimplechart.NewNode(5, 110),   // Day 5 - increase
    *fynesimplechart.NewNode(6, 107),   // Day 6 - decrease
    *fynesimplechart.NewNode(7, 112),   // Day 7 - increase
}

plot := fynesimplechart.NewPlot(nodes, "Stock Price ($)")
plot.ShowLine = true
plot.ShowPoints = true  // Show daily closing points
plot.LineWidth = 2
plot.PointSize = 3
```

## Best Practices

### 1. Data Order Matters
Line charts connect points in order, so sort your data:

```go
import "sort"

// Sort by X coordinate
sort.Slice(nodes, func(i, j int) bool {
    return nodes[i].X < nodes[j].X
})
```

### 2. Choose Line Width Based on Context

| Context | Recommended Width | Reason |
|---------|------------------|---------|
| Dense data | 1.0 - 1.5 | Prevents visual clutter |
| Normal data | 1.5 - 2.5 | Good balance |
| Sparse data | 2.5 - 3.5 | Clear visibility |
| Presentation | 3.0 - 5.0 | High visibility |

### 3. When to Show Points

```go
// Show points when:
// - You have <30 data points
// - Emphasizing individual measurements
// - Data is discrete (not continuous)

plot.ShowPoints = true

// Hide points when:
// - You have >50 data points
// - Showing smooth trends
// - Data is continuous

plot.ShowPoints = false
```

### 4. Smooth vs Jagged Lines

```go
// Smooth line: many points, small increments
for i := 0; i <= 100; i++ {
    x := float32(i) / 10.0
    // ...
}

// Jagged line: few points, large increments
for i := 0; i <= 10; i++ {
    x := float32(i)
    // ...
}
```

## Common Use Cases

### Time Series Data
```go
// Example: Website traffic over time
traffic := []int{1200, 1350, 1280, 1420, 1560, 1490, 1630}
nodes := []fynesimplechart.Node{}

for day, visits := range traffic {
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(day+1),
        float32(visits),
    ))
}

plot := fynesimplechart.NewPlot(nodes, "Daily Visitors")
plot.ShowLine = true
plot.ShowPoints = true
plot.LineWidth = 2
```

### Progress Tracking
```go
// Example: Weight loss tracking
weights := []float32{85.5, 84.8, 84.2, 83.9, 83.5, 83.0}
nodes := []fynesimplechart.Node{}

for week, weight := range weights {
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(week+1),
        weight,
    ))
}

plot := fynesimplechart.NewPlot(nodes, "Weight (kg)")
plot.ShowLine = true
plot.LineWidth = 2.5
```

### Mathematical Functions
```go
// Example: Exponential growth
nodes := []fynesimplechart.Node{}

for x := 0.0; x <= 5.0; x += 0.2 {
    y := math.Exp(x / 2)  // e^(x/2)
    nodes = append(nodes, *fynesimplechart.NewNode(
        float32(x),
        float32(y),
    ))
}

plot := fynesimplechart.NewPlot(nodes, "e^(x/2)")
plot.ShowLine = true
plot.ShowPoints = false
plot.LineWidth = 2
```

## Styling Tips

### Professional Look
```go
plot.ShowLine = true
plot.ShowPoints = true
plot.LineWidth = 2.0
plot.PointSize = 3.5
```

### Minimal Look
```go
plot.ShowLine = true
plot.ShowPoints = false
plot.LineWidth = 1.5
```

### Bold Look
```go
plot.ShowLine = true
plot.ShowPoints = true
plot.LineWidth = 3.0
plot.PointSize = 5.0
```

## Exercise

Create a line chart showing company revenue growth over 6 quarters:

```
Q1: $45,000
Q2: $52,000
Q3: $48,000
Q4: $61,000
Q5: $58,000
Q6: $67,000
```

<details>
<summary>Solution</summary>

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Revenue Growth")
    w.Resize(fyne.NewSize(700, 500))

    revenue := []float32{45000, 52000, 48000, 61000, 58000, 67000}
    nodes := []fynesimplechart.Node{}

    for quarter, amount := range revenue {
        nodes = append(nodes, *fynesimplechart.NewNode(
            float32(quarter+1),
            amount,
        ))
    }

    plot := fynesimplechart.NewPlot(nodes, "Revenue ($)")
    plot.ShowLine = true
    plot.ShowPoints = true
    plot.LineWidth = 2.5
    plot.PointSize = 4

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Quarterly Revenue Growth")

    w.SetContent(chart)
    w.ShowAndRun()
}
```
</details>

## Summary

You learned:
- ✅ How to create line charts
- ✅ How to customize line width
- ✅ When to show/hide points
- ✅ Creating smooth curves
- ✅ Best practices for line charts

## Next Steps

Continue to [Tutorial 4: Customizing Appearance](04-customizing-appearance.md) to learn about colors, styles, and visual customization.
