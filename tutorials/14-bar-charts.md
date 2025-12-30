# Tutorial 14: Bar Charts

## Overview

Learn how to create professional bar charts for categorical data and comparisons.

**Time to complete:** 20 minutes

## What Are Bar Charts?

Bar charts display data using rectangular bars where:
- **Height** represents the value
- **Width** is uniform across bars
- **Position** represents the category or data point

Perfect for:
- Sales by month/category
- Survey results
- Budget comparisons
- Performance metrics

## Basic Bar Chart

Create a simple vertical bar chart:

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/alexiusacademia/fynesimplechart"
)

func main() {
	a := app.New()
	w := a.NewWindow("Bar Chart Example")

	// Monthly sales data
	nodes := []fynesimplechart.Node{
		*fynesimplechart.NewNode(1, 45),  // January
		*fynesimplechart.NewNode(2, 52),  // February
		*fynesimplechart.NewNode(3, 48),  // March
		*fynesimplechart.NewNode(4, 61),  // April
		*fynesimplechart.NewNode(5, 58),  // May
		*fynesimplechart.NewNode(6, 67),  // June
	}

	plot := fynesimplechart.NewPlot(nodes, "Monthly Sales ($K)")

	// Enable bar chart
	plot.ShowBars = true
	plot.BarWidth = 0.8        // 80% of available space
	plot.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}

	// Disable points and lines
	plot.ShowPoints = false
	plot.ShowLine = false

	chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
	chart.SetChartTitle("2024 Sales Performance")
	chart.Resize(fyne.NewSize(800, 600))

	w.SetContent(chart)
	w.ShowAndRun()
}
```

### Key Properties

```go
plot.ShowBars = true       // Enable bar chart rendering
plot.BarWidth = 0.8        // Width as fraction (0.8 = 80% of space)
plot.ShowPoints = false    // Hide points
plot.ShowLine = false      // Hide lines
```

## Bar Width Configuration

Control how much space each bar occupies:

```go
// Narrow bars (50% of space)
plot.BarWidth = 0.5

// Standard bars (80% - recommended)
plot.BarWidth = 0.8

// Wide bars (100% - no gaps)
plot.BarWidth = 1.0

// Auto-width (use default)
plot.BarWidth = 0.8  // This is the default
```

**Tip:** Use `0.8` for single series, `0.25-0.35` for grouped bars.

## Adding Bar Borders

Make bars stand out with borders:

```go
plot := fynesimplechart.NewPlot(nodes, "Expenses")
plot.ShowBars = true
plot.PlotColor = color.RGBA{R: 44, G: 160, B: 44, A: 255}  // Green

// Add border
plot.BarBorderWidth = 2
plot.BarBorderColor = color.RGBA{R: 30, G: 110, B: 30, A: 255}  // Dark green
```

### Default Border Color

If you don't specify a border color, it auto-darkens the bar color:

```go
plot.BarBorderWidth = 1
// BarBorderColor is nil, so it uses darker version of PlotColor
```

## Grouped Bar Charts

Compare multiple series side-by-side by offsetting X values:

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/alexiusacademia/fynesimplechart"
)

func main() {
	a := app.New()
	w := a.NewWindow("Grouped Bar Chart")

	// Product A sales (regions 1-4)
	productA := []fynesimplechart.Node{
		*fynesimplechart.NewNode(1, 30),    // Region 1
		*fynesimplechart.NewNode(2, 35),    // Region 2
		*fynesimplechart.NewNode(3, 28),    // Region 3
		*fynesimplechart.NewNode(4, 42),    // Region 4
	}

	plotA := fynesimplechart.NewPlot(productA, "Product A")
	plotA.ShowBars = true
	plotA.BarWidth = 0.35  // Narrower for grouped display
	plotA.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}  // Blue
	plotA.ShowPoints = false
	plotA.ShowLine = false

	// Product B sales (offset X by 0.4 for grouping)
	productB := []fynesimplechart.Node{
		*fynesimplechart.NewNode(1.4, 25),  // Region 1 (offset)
		*fynesimplechart.NewNode(2.4, 32),  // Region 2 (offset)
		*fynesimplechart.NewNode(3.4, 30),  // Region 3 (offset)
		*fynesimplechart.NewNode(4.4, 38),  // Region 4 (offset)
	}

	plotB := fynesimplechart.NewPlot(productB, "Product B")
	plotB.ShowBars = true
	plotB.BarWidth = 0.35
	plotB.PlotColor = color.RGBA{R: 255, G: 127, B: 14, A: 255}  // Orange
	plotB.ShowPoints = false
	plotB.ShowLine = false

	chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plotA, *plotB})
	chart.SetChartTitle("Product Comparison by Region")
	chart.Resize(fyne.NewSize(800, 600))

	w.SetContent(chart)
	w.ShowAndRun()
}
```

### Grouping Strategy

For grouped bars:
1. **Use narrow bars**: `BarWidth = 0.25` to `0.35`
2. **Offset X values**: Add `0.3` to `0.4` to X coordinates
3. **Keep series count low**: 2-3 series work best

```go
// Series 1: X = 1, 2, 3, 4
// Series 2: X = 1.3, 2.3, 3.3, 4.3
// Series 3: X = 1.6, 2.6, 3.6, 4.6

// All use BarWidth = 0.25
```

## Three-Series Grouped Bars

```go
// Q1 at X = 1, 2, 3, 4
q1 := fynesimplechart.NewPlot(data1, "Q1 2024")
q1.ShowBars = true
q1.BarWidth = 0.25
q1.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}
q1.ShowPoints = false
q1.ShowLine = false

// Q2 at X = 1.3, 2.3, 3.3, 4.3
q2 := fynesimplechart.NewPlot(data2, "Q2 2024")
q2.ShowBars = true
q2.BarWidth = 0.25
q2.PlotColor = color.RGBA{R: 255, G: 127, B: 14, A: 255}
q2.ShowPoints = false
q2.ShowLine = false

// Q3 at X = 1.6, 2.6, 3.6, 4.6
q3 := fynesimplechart.NewPlot(data3, "Q3 2024")
q3.ShowBars = true
q3.BarWidth = 0.25
q3.PlotColor = color.RGBA{R: 44, G: 160, B: 44, A: 255}
q3.ShowPoints = false
q3.ShowLine = false

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*q1, *q2, *q3})
```

## Handling Negative Values

Bar charts automatically handle negative values:

```go
// Profit/loss data
nodes := []fynesimplechart.Node{
	*fynesimplechart.NewNode(1, 15),   // Profit
	*fynesimplechart.NewNode(2, -8),   // Loss (negative)
	*fynesimplechart.NewNode(3, 12),   // Profit
	*fynesimplechart.NewNode(4, -5),   // Loss
	*fynesimplechart.NewNode(5, 20),   // Profit
}

plot := fynesimplechart.NewPlot(nodes, "Profit/Loss")
plot.ShowBars = true
plot.BarWidth = 0.8

// Negative bars automatically extend downward from Y=0
```

Negative bars:
- Extend **downward** from the zero line
- Use the same color as positive bars
- Work with borders and all other features

## Combining Bars with Lines

Mix bar charts with line charts in one visualization:

```go
// Bar chart for actual sales
actual := fynesimplechart.NewPlot(actualData, "Actual")
actual.ShowBars = true
actual.BarWidth = 0.6
actual.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}
actual.ShowPoints = false
actual.ShowLine = false

// Line chart for target
target := fynesimplechart.NewPlot(targetData, "Target")
target.ShowLine = true
target.LineWidth = 3
target.PlotColor = color.RGBA{R: 244, G: 67, B: 54, A: 255}  // Red
target.ShowPoints = true
target.PointSize = 5
target.ShowBars = false

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*actual, *target})
chart.SetChartTitle("Sales: Actual vs Target")
```

## Color Recommendations

### Single Series
```go
// Professional blue
color.RGBA{R: 31, G: 119, B: 180, A: 255}

// Success green
color.RGBA{R: 44, G: 160, B: 44, A: 255}

// Warning orange
color.RGBA{R: 255, G: 127, B: 14, A: 255}
```

### Multiple Series (use distinct colors)
```go
// Series 1: Blue
color.RGBA{R: 31, G: 119, B: 180, A: 255}

// Series 2: Orange
color.RGBA{R: 255, G: 127, B: 14, A: 255}

// Series 3: Green
color.RGBA{R: 44, G: 160, B: 44, A: 255}
```

## Use Cases

### Monthly Sales
```go
// Perfect for showing sales trends by month
plot.ShowBars = true
plot.BarWidth = 0.8
```

### Survey Results
```go
// Great for displaying response counts
plot.ShowBars = true
plot.BarBorderWidth = 1  // Add borders for clarity
```

### Budget Allocation
```go
// Excellent for comparing department budgets
// Use grouped bars for multi-year comparison
```

### Performance Metrics
```go
// Ideal for KPI dashboards
// Combine with target lines
```

## Common Patterns

### Department Budget Chart
```go
departments := []fynesimplechart.Node{
	*fynesimplechart.NewNode(1, 120),  // Marketing
	*fynesimplechart.NewNode(2, 95),   // Sales
	*fynesimplechart.NewNode(3, 150),  // Engineering
	*fynesimplechart.NewNode(4, 80),   // HR
	*fynesimplechart.NewNode(5, 65),   // Operations
}

plot := fynesimplechart.NewPlot(departments, "Budget ($K)")
plot.ShowBars = true
plot.BarWidth = 0.7
plot.PlotColor = color.RGBA{R: 44, G: 160, B: 44, A: 255}
plot.BarBorderWidth = 2
plot.ShowPoints = false
plot.ShowLine = false
```

### Quarterly Comparison
```go
// Use grouped bars with 3-4 series
// BarWidth = 0.25
// Offset X values by 0.3
```

## Troubleshooting

### Bars Too Wide
```go
// Reduce BarWidth
plot.BarWidth = 0.6  // Instead of 0.8
```

### Bars Overlapping (Grouped Charts)
```go
// Increase X offset or reduce BarWidth
// Try offset of 0.4 with BarWidth of 0.3
```

### Bars Not Showing
```go
// Check these:
plot.ShowBars = true         // Must be true
plot.ShowPoints = false      // Should be false
plot.ShowLine = false        // Should be false
```

### Border Too Thick
```go
// Reduce border width
plot.BarBorderWidth = 1  // Instead of 2 or 3
```

## Design Best Practices

### 1. Bar Width
```go
// Single series: 0.7-0.8
plot.BarWidth = 0.8

// Grouped bars: 0.25-0.35
plot.BarWidth = 0.3

// Many categories: 0.5-0.6
plot.BarWidth = 0.6
```

### 2. Colors
```go
// Use meaningful colors
// Green for profit, Red for loss
// Blue for neutral data
```

### 3. Borders
```go
// Use borders for:
// - Small bars
// - Printed reports
// - Presentations

// Skip borders for:
// - Large datasets
// - Screen-only viewing
```

### 4. Grouping Limit
```go
// Recommended: 2-3 series
// Maximum: 4 series
// Beyond 4: Consider stacked bars or separate charts
```

## Complete Example

See [examples/bar_chart_demo.go](../examples/bar_chart_demo.go) for:
1. Simple sales bar chart
2. Product comparison (grouped)
3. Bordered bars
4. Three-series quarterly comparison

## Summary

Bar charts in FyneSimpleChart:
- ✅ Vertical bars with auto-sizing
- ✅ Configurable bar width
- ✅ Optional borders with custom colors
- ✅ Grouped bars support
- ✅ Negative value handling
- ✅ Mix with lines and points
- ✅ Professional appearance

## Exercises

1. **Basic Bar Chart**: Create a chart showing website traffic by day of week
2. **Grouped Bars**: Compare product sales across 3 different stores
3. **Bordered Bars**: Create a department expense chart with thick borders
4. **Mixed Chart**: Combine bars (actual) with a line (target goal)

## Next Steps

Continue to [Tutorial 15: Best Practices](12-best-practices.md) to learn performance tips and design guidelines.

**Note:** Bar charts are a powerful tool for categorical data. Experiment with width, borders, and grouping to find what works best for your data!
