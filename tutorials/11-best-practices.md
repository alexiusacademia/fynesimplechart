# Tutorial 11: Best Practices

## Overview

Learn professional tips and best practices for creating effective charts.

**Time to complete:** 15 minutes

## Performance Best Practices

### 1. Optimize Data Points
```go
// Good: Reasonable number of points
for i := 0; i < 100; i++ {  // 100 points
    // ...
}

// Avoid: Too many points (slow rendering)
for i := 0; i < 10000; i++ {  // 10,000 points
    // ...
}
```

**Recommendation:** Keep points under 1000 for smooth performance.

### 2. Efficient Updates
```go
// Good: Update plot data, then refresh once
plot.Nodes = newNodes
plot.LineWidth = 2.5
plot.PointSize = 4
chart.Refresh()  // Single refresh

// Avoid: Multiple refreshes
plot.Nodes = newNodes
chart.Refresh()
plot.LineWidth = 2.5
chart.Refresh()  // Unnecessary extra refresh
```

### 3. Data Preparation
```go
// Good: Prepare data before creating plot
nodes := prepareData(rawData)
plot := fynesimplechart.NewPlot(nodes, "Data")

// Avoid: Processing during render
// (Don't do complex calculations in render loops)
```

## Visual Design Best Practices

### 1. Choose Appropriate Chart Type

| Data Type | Best Chart | Why |
|-----------|-----------|-----|
| Time series | Line chart | Shows trends |
| Discrete points | Scatter plot | Shows distribution |
| Comparison | Multiple series | Shows relationships |
| Progress | Line with points | Shows steps and trend |

### 2. Color Selection
```go
// Good: High contrast, meaningful colors
profit := color.RGBA{R: 76, G: 175, B: 80, A: 255}   // Green
loss := color.RGBA{R: 244, G: 67, B: 54, A: 255}     // Red

// Avoid: Similar colors
color1 := color.RGBA{R: 100, G: 100, B: 150, A: 255}
color2 := color.RGBA{R: 100, G: 150, B: 100, A: 255}  // Too similar
```

### 3. Point and Line Sizing

```go
// Dense data (>50 points)
plot.PointSize = 2
plot.LineWidth = 1.5

// Normal data (10-50 points)
plot.PointSize = 3
plot.LineWidth = 2.0

// Sparse data (<10 points)
plot.PointSize = 5
plot.LineWidth = 2.5

// Presentation
plot.PointSize = 6
plot.LineWidth = 3.5
```

### 4. Grid Usage
```go
// Show grid for:
// - Data analysis
// - Reading exact values
// - Scientific reports
chart.ShowGrid = true

// Hide grid for:
// - Executive presentations
// - Marketing materials
// - Artistic visualization
chart.ShowGrid = false
```

## Data Management Best Practices

### 1. Validate Data
```go
func validateNodes(nodes []fynesimplechart.Node) bool {
    if len(nodes) == 0 {
        return false
    }

    for _, node := range nodes {
        if math.IsNaN(float64(node.X)) || math.IsNaN(float64(node.Y)) {
            return false
        }
        if math.IsInf(float64(node.X), 0) || math.IsInf(float64(node.Y), 0) {
            return false
        }
    }
    return true
}
```

### 2. Handle Missing Data
```go
// Skip nil/invalid values
nodes := []fynesimplechart.Node{}
for i, value := range data {
    if value != nil && !math.IsNaN(float64(*value)) {
        nodes = append(nodes, *fynesimplechart.NewNode(
            float32(i),
            *value,
        ))
    }
}
```

### 3. Sort Data for Line Charts
```go
import "sort"

// Sort by X coordinate
sort.Slice(nodes, func(i, j int) bool {
    return nodes[i].X < nodes[j].X
})
```

## Chart Organization Best Practices

### 1. Meaningful Titles
```go
// Good: Descriptive and specific
chart.SetChartTitle("Q1 2024 Revenue by Region (in $K)")
plot := fynesimplechart.NewPlot(data, "North America")

// Avoid: Generic and vague
chart.SetChartTitle("Chart")
plot := fynesimplechart.NewPlot(data, "Data 1")
```

### 2. Limit Series Count
```go
// Good: 2-4 series (easy to read)
chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
    *plot1, *plot2, *plot3,
})

// Avoid: Too many series (>6)
// Becomes cluttered and hard to read
```

### 3. Consistent Scales
```go
// Good: All series on similar scale
series1: Y from 0-100
series2: Y from 0-100

// Avoid: Vastly different scales
series1: Y from 0-100
series2: Y from 0-10000  // Makes series1 invisible
```

## Code Organization Best Practices

### 1. Reusable Functions
```go
// Create helper functions for common operations
func createPlotFromData(data []float32, title string, showLine bool) *fynesimplechart.Plot {
    nodes := []fynesimplechart.Node{}
    for i, val := range data {
        nodes = append(nodes, *fynesimplechart.NewNode(
            float32(i),
            val,
        ))
    }

    plot := fynesimplechart.NewPlot(nodes, title)
    plot.ShowLine = showLine
    plot.LineWidth = 2
    return plot
}
```

### 2. Configuration Structs
```go
type ChartConfig struct {
    Title      string
    ShowGrid   bool
    WindowSize fyne.Size
}

func createChart(plots []fynesimplechart.Plot, config ChartConfig) *fynesimplechart.ScatterPlot {
    chart := fynesimplechart.NewGraphWidget(plots)
    chart.SetChartTitle(config.Title)
    chart.ShowGrid = config.ShowGrid
    chart.Resize(config.WindowSize)
    return chart
}
```

### 3. Error Handling
```go
func createSafeChart(data []float32) (*fynesimplechart.ScatterPlot, error) {
    if len(data) == 0 {
        return nil, fmt.Errorf("no data provided")
    }

    nodes := []fynesimplechart.Node{}
    for i, val := range data {
        nodes = append(nodes, *fynesimplechart.NewNode(
            float32(i),
            val,
        ))
    }

    plot := fynesimplechart.NewPlot(nodes, "Data")
    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})

    return chart, nil
}
```

## Common Pitfalls to Avoid

### 1. Unsorted Line Chart Data
```go
// Problem: Random X order creates zigzag lines
// Solution: Sort by X before creating plot
sort.Slice(nodes, func(i, j int) bool {
    return nodes[i].X < nodes[j].X
})
```

### 2. Too Many Updates
```go
// Problem: Refreshing too frequently (causes flicker)
// Solution: Batch updates
for _, point := range newPoints {
    nodes = append(nodes, point)
}
plot.Nodes = nodes
chart.Refresh()  // Single refresh
```

### 3. Memory Leaks in Real-time
```go
// Problem: Unlimited data accumulation
// Solution: Limit buffer size
if len(nodes) > maxPoints {
    nodes = nodes[1:]  // Remove oldest
}
```

### 4. Poor Color Contrast
```go
// Problem: Similar colors hard to distinguish
// Solution: Use contrasting colors from palette
```

## Checklist for Professional Charts

Before deploying your chart, verify:

- [ ] Data is validated and clean
- [ ] Appropriate chart type selected
- [ ] Colors have good contrast
- [ ] Point/line sizes match data density
- [ ] Grid setting appropriate for use case
- [ ] Title is descriptive
- [ ] Series labels are meaningful
- [ ] Performance tested with actual data volume
- [ ] Code is organized and maintainable
- [ ] Error cases handled

## Summary

- ✅ Performance optimization
- ✅ Visual design principles
- ✅ Data management
- ✅ Code organization
- ✅ Common pitfalls

## Next Steps

Continue to [Tutorial 12: Integration Examples](12-integration-examples.md) for real-world application examples.
