# Tutorial 15: Enhanced Features

## Overview

Learn to use the advanced features that take your charts to the next level: axis titles, flexible legend positioning, manual axis ranges, and data labels.

**Time to complete:** 25 minutes

**Demo Application:** Run the complete example:
```bash
cd examples
go run tutorial15_example.go
```

## What You'll Learn

- Adding descriptive axis titles
- Positioning legends for optimal layout
- Setting manual axis ranges for consistency
- Displaying data labels on chart elements
- Combining all features for professional charts

## Prerequisites

- Tutorial 1: Getting Started
- Tutorial 2: Basic Scatter Plot
- Basic understanding of Go and Fyne

---

## Part 1: Axis Titles

### Why Use Axis Titles?

Axis titles provide context and make your charts self-documenting. Instead of relying on the legend or external documentation, viewers immediately understand what each axis represents.

### Basic Usage

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Axis Titles Demo")

    // Temperature data over time
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(0, 20.5),
        *fynesimplechart.NewNode(1, 21.2),
        *fynesimplechart.NewNode(2, 22.8),
        *fynesimplechart.NewNode(3, 23.5),
        *fynesimplechart.NewNode(4, 22.1),
        *fynesimplechart.NewNode(5, 21.0),
    }

    plot := fynesimplechart.NewPlot(nodes, "Temperature")
    plot.ShowLine = true
    plot.ShowPoints = true

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Lab Temperature Monitoring")

    // Add axis titles for clarity
    chart.XAxisTitle = "Time (hours)"
    chart.YAxisTitle = "Temperature (°C)"

    chart.Resize(fyne.NewSize(600, 400))
    w.SetContent(chart)
    w.ShowAndRun()
}
```

### Best Practices

- **Be Specific**: "Revenue ($K)" is better than just "Revenue"
- **Include Units**: Always show units in parentheses
- **Keep it Concise**: Max 3-4 words
- **Use Standard Abbreviations**: "sec" for seconds, "m" for meters, etc.

### Common Patterns

```go
// Time series
chart.XAxisTitle = "Time (hours)"
chart.YAxisTitle = "Value"

// Scientific data
chart.XAxisTitle = "Distance (m)"
chart.YAxisTitle = "Force (N)"

// Business metrics
chart.XAxisTitle = "Month"
chart.YAxisTitle = "Revenue ($K)"

// Performance data
chart.XAxisTitle = "Iteration"
chart.YAxisTitle = "Execution Time (ms)"
```

---

## Part 2: Flexible Legend Positioning

### Why Position Matters

Different chart layouts work better with legends in different positions:
- Wide charts → Bottom legend
- Tall charts → Side legend
- Simple single-series → No legend needed

### Available Positions

```go
import "github.com/alexiusacademia/fynesimplechart"

// Five position options
chart.LegendPosition = fynesimplechart.LegendRight   // Default
chart.LegendPosition = fynesimplechart.LegendBottom  // Horizontal below
chart.LegendPosition = fynesimplechart.LegendTop     // Horizontal above
chart.LegendPosition = fynesimplechart.LegendLeft    // Vertical on left
chart.LegendPosition = fynesimplechart.LegendNone    // No legend

// Or completely hide it
chart.ShowLegend = false
```

### Example: Bottom Legend for Wide Charts

```go
func createBottomLegendChart() *fynesimplechart.ScatterPlot {
    // Create three series
    series1 := []fynesimplechart.Node{}
    series2 := []fynesimplechart.Node{}
    series3 := []fynesimplechart.Node{}

    for i := 0; i <= 20; i++ {
        x := float32(i)
        series1 = append(series1, *fynesimplechart.NewNode(x, 10+x*0.5))
        series2 = append(series2, *fynesimplechart.NewNode(x, 15+x*0.3))
        series3 = append(series3, *fynesimplechart.NewNode(x, 8+x*0.7))
    }

    plot1 := fynesimplechart.NewPlot(series1, "Product A")
    plot1.ShowLine = true
    plot1.LineWidth = 2

    plot2 := fynesimplechart.NewPlot(series2, "Product B")
    plot2.ShowLine = true
    plot2.LineWidth = 2

    plot3 := fynesimplechart.NewPlot(series3, "Product C")
    plot3.ShowLine = true
    plot3.LineWidth = 2

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot1, *plot2, *plot3})
    chart.SetChartTitle("Sales Trends by Product")

    // Bottom legend for wide chart
    chart.LegendPosition = fynesimplechart.LegendBottom

    return chart
}
```

### When to Use Each Position

| Position | Best For | Avoid When |
|----------|----------|------------|
| **Right** | Standard charts, portrait layout | Very wide charts |
| **Bottom** | Wide charts, dashboards | Many series (>4) |
| **Top** | Reports, presentations | Chart has main title |
| **Left** | When right side needed for notes | Left margin is tight |
| **None** | Single series, clean look | Multiple series |

---

## Part 3: Manual Axis Ranges

### Why Manual Ranges?

Automatic scaling can be misleading. Manual ranges ensure:
- **Consistent comparison** across multiple charts
- **Zero-based visualization** for percentages/growth
- **Focus on specific range** for detailed analysis
- **Prevent misleading scales** in presentations

### Basic Usage

```go
chart := fynesimplechart.NewGraphWidget(plots)

// Fix Y-axis from 0 to 100
minY := float32(0)
maxY := float32(100)
chart.MinY = &minY
chart.MaxY = &maxY

// X-axis remains automatic (MinX and MaxX are nil)
```

### Example: Comparing Multiple Charts

```go
func createConsistentCharts() []*fynesimplechart.ScatterPlot {
    charts := make([]*fynesimplechart.ScatterPlot, 3)

    // Fixed range for all charts
    minY := float32(0)
    maxY := float32(100)

    // Chart 1: Q1 Data
    q1Data := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 45),
        *fynesimplechart.NewNode(2, 52),
        *fynesimplechart.NewNode(3, 48),
    }
    plot1 := fynesimplechart.NewPlot(q1Data, "Q1")
    plot1.ShowBars = true
    charts[0] = fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot1})
    charts[0].SetChartTitle("Q1 Performance")
    charts[0].MinY = &minY
    charts[0].MaxY = &maxY

    // Chart 2: Q2 Data
    q2Data := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 51),
        *fynesimplechart.NewNode(2, 58),
        *fynesimplechart.NewNode(3, 54),
    }
    plot2 := fynesimplechart.NewPlot(q2Data, "Q2")
    plot2.ShowBars = true
    charts[1] = fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot2})
    charts[1].SetChartTitle("Q2 Performance")
    charts[1].MinY = &minY
    charts[1].MaxY = &maxY

    // Chart 3: Q3 Data
    q3Data := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 61),
        *fynesimplechart.NewNode(2, 65),
        *fynesimplechart.NewNode(3, 62),
    }
    plot3 := fynesimplechart.NewPlot(q3Data, "Q3")
    plot3.ShowBars = true
    charts[2] = fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot3})
    charts[2].SetChartTitle("Q3 Performance")
    charts[2].MinY = &minY
    charts[2].MaxY = &maxY

    return charts
}
```

### Use Cases

**1. Dashboard Consistency**
```go
// All dashboard charts use same Y scale
minY, maxY := float32(0), float32(200)
chart1.MinY, chart1.MaxY = &minY, &maxY
chart2.MinY, chart2.MaxY = &minY, &maxY
chart3.MinY, chart3.MaxY = &minY, &maxY
```

**2. Percentage Charts**
```go
// Always 0-100 for percentages
minY, maxY := float32(0), float32(100)
chart.MinY = &minY
chart.MaxY = &maxY
chart.YAxisTitle = "Completion (%)"
```

**3. Zooming Into Range**
```go
// Focus on 90-110 range
minY, maxY := float32(90), float32(110)
chart.MinY = &minY
chart.MaxY = &maxY
```

---

## Part 4: Data Labels

### Why Data Labels?

Data labels show exact values without requiring users to interpolate from the axis. Perfect for:
- Presentation charts
- KPI dashboards
- Small datasets where precision matters
- When you want exact numbers visible

### Basic Usage

```go
plot := fynesimplechart.NewPlot(nodes, "Sales")
plot.ShowDataLabels = true              // Enable labels
plot.LabelFormat = "$%.0fK"             // Format string
plot.LabelSize = 10                     // Font size
plot.LabelColor = myColor               // Optional custom color
```

### Format String Examples

```go
// Decimals
plot.LabelFormat = "%.1f"      // "12.5"
plot.LabelFormat = "%.2f"      // "12.50"
plot.LabelFormat = "%.0f"      // "13"

// Currency
plot.LabelFormat = "$%.2f"     // "$12.50"
plot.LabelFormat = "$%.0fK"    // "$13K"
plot.LabelFormat = "€%.2f"     // "€12.50"

// Percentage
plot.LabelFormat = "%.0f%%"    // "13%"
plot.LabelFormat = "%.1f%%"    // "12.5%"

// Scientific
plot.LabelFormat = "%.2e"      // "1.25e+01"

// Custom
plot.LabelFormat = "%.1f°C"    // "12.5°C"
plot.LabelFormat = "%.0f units" // "13 units"
```

### Example: Bar Chart with Labels

```go
func createLabeledBarChart() *fynesimplechart.ScatterPlot {
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 45),
        *fynesimplechart.NewNode(2, 62),
        *fynesimplechart.NewNode(3, 38),
        *fynesimplechart.NewNode(4, 71),
        *fynesimplechart.NewNode(5, 55),
        *fynesimplechart.NewNode(6, 49),
    }

    plot := fynesimplechart.NewPlot(nodes, "Revenue")
    plot.ShowBars = true
    plot.BarWidth = 0.7
    plot.ShowPoints = false
    plot.ShowLine = false

    // Add data labels
    plot.ShowDataLabels = true
    plot.LabelFormat = "$%.0fK"
    plot.LabelSize = 10

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Monthly Revenue")
    chart.XAxisTitle = "Month"
    chart.YAxisTitle = "Revenue ($K)"

    return chart
}
```

### Smart Label Positioning

The library automatically positions labels intelligently:

**For Points:**
- Labels appear above the point
- Offset by point size + 3 pixels
- Centered horizontally

**For Bars:**
- Labels centered on top of bar
- For negative bars: labels appear below
- Prevents overlap with bar

**Custom Colors:**
```go
import "image/color"

// Match label color to plot color
plot.LabelColor = color.RGBA{R: 44, G: 160, B: 44, A: 255}

// Or use contrasting color
plot.LabelColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}  // Black

// Default (nil) uses theme foreground color
plot.LabelColor = nil
```

---

## Part 5: Combining All Features

### Professional Dashboard Example

Here's a complete example using all enhanced features:

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
    w := a.NewWindow("Professional Dashboard")

    // Create sales data for two quarters
    q1Sales := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 45),
        *fynesimplechart.NewNode(2, 52),
        *fynesimplechart.NewNode(3, 48),
        *fynesimplechart.NewNode(4, 61),
    }

    q2Sales := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1.3, 48),
        *fynesimplechart.NewNode(2.3, 55),
        *fynesimplechart.NewNode(3.3, 51),
        *fynesimplechart.NewNode(4.3, 64),
    }

    // Q1 Plot with all features
    q1 := fynesimplechart.NewPlot(q1Sales, "Q1 2024")
    q1.ShowBars = true
    q1.BarWidth = 0.25
    q1.ShowDataLabels = true
    q1.LabelFormat = "$%.0fK"
    q1.LabelSize = 9
    q1.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}

    // Q2 Plot with all features
    q2 := fynesimplechart.NewPlot(q2Sales, "Q2 2024")
    q2.ShowBars = true
    q2.BarWidth = 0.25
    q2.ShowDataLabels = true
    q2.LabelFormat = "$%.0fK"
    q2.LabelSize = 9
    q2.PlotColor = color.RGBA{R: 255, G: 127, B: 14, A: 255}

    // Create chart with all enhancements
    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*q1, *q2})

    // 1. Chart title
    chart.SetChartTitle("Quarterly Sales Comparison")

    // 2. Axis titles
    chart.XAxisTitle = "Region"
    chart.YAxisTitle = "Sales ($K)"

    // 3. Legend positioning
    chart.LegendPosition = fynesimplechart.LegendTop

    // 4. Manual axis ranges for consistency
    minY := float32(0)
    maxY := float32(80)
    chart.MinY = &minY
    chart.MaxY = &maxY

    chart.Resize(fyne.NewSize(800, 600))
    w.SetContent(chart)
    w.ShowAndRun()
}
```

### Result

This creates a professional chart with:
- ✅ Clear chart title
- ✅ Descriptive axis titles with units
- ✅ Legend positioned at top (saves horizontal space)
- ✅ Fixed Y-axis (0-80) for fair comparison
- ✅ Data labels showing exact values
- ✅ Grouped bars for quarterly comparison

---

## Common Patterns

### Pattern 1: Time Series Dashboard

```go
// Temperature monitoring over 24 hours
chart.SetChartTitle("Lab Temperature - Last 24h")
chart.XAxisTitle = "Time (hours)"
chart.YAxisTitle = "Temperature (°C)"
chart.LegendPosition = fynesimplechart.LegendBottom

// Fix range to expected temperature bounds
minY, maxY := float32(18), float32(26)
chart.MinY, chart.MaxY = &minY, &maxY

plot.ShowDataLabels = true
plot.LabelFormat = "%.1f°C"
```

### Pattern 2: KPI Chart

```go
// Single KPI with target line
chart.SetChartTitle("Monthly Target Achievement")
chart.XAxisTitle = "Month"
chart.YAxisTitle = "Achievement (%)"
chart.ShowLegend = false  // Single metric, no legend needed

// Always 0-100 for percentages
minY, maxY := float32(0), float32(100)
chart.MinY, chart.MaxY = &minY, &maxY

plot.ShowDataLabels = true
plot.LabelFormat = "%.0f%%"
plot.LabelSize = 11
```

### Pattern 3: Multi-Series Comparison

```go
// Multiple products over time
chart.SetChartTitle("Product Performance Q1-Q4")
chart.XAxisTitle = "Quarter"
chart.YAxisTitle = "Units Sold (K)"
chart.LegendPosition = fynesimplechart.LegendTop

// All products use same scale
minY, maxY := float32(0), float32(200)
chart.MinY, chart.MaxY = &minY, &maxY

// Enable labels on all plots
for _, plot := range plots {
    plot.ShowDataLabels = true
    plot.LabelFormat = "%.0fK"
    plot.LabelSize = 8
}
```

---

## Exercises

### Exercise 1: Temperature Chart
Create a temperature monitoring chart with:
- Hourly temperature readings (0-23 hours)
- Y-axis title "Temperature (°C)"
- X-axis title "Hour of Day"
- Fixed Y range from 15°C to 30°C
- Data labels showing one decimal place
- Legend at bottom

### Exercise 2: Sales Dashboard
Create a grouped bar chart showing:
- 4 regions (North, South, East, West)
- 2 quarters (Q1, Q2)
- Fixed Y range 0-100
- Currency labels ($XXK format)
- Legend at top
- Axis titles with units

### Exercise 3: No-Legend Chart
Create a simple line chart with:
- Single data series
- No legend (ShowLegend = false)
- Data labels at each point
- Custom axis ranges
- Professional title and axis labels

---

## Common Issues

### Issue 1: Labels Overlap

**Problem**: Too many data points cause labels to overlap

**Solution**: Either reduce label size or disable labels for dense data
```go
if len(nodes) > 10 {
    plot.LabelSize = 8  // Smaller size
    // Or skip labels entirely
    plot.ShowDataLabels = false
}
```

### Issue 2: Manual Range Too Tight

**Problem**: Data points clipped because range is too small

**Solution**: Add padding to your manual range
```go
// Instead of exact data range
minY := float32(10)  // Minimum data value
maxY := float32(50)  // Maximum data value

// Add 10% padding
padding := (maxY - minY) * 0.1
minY -= padding
maxY += padding
chart.MinY, chart.MaxY = &minY, &maxY
```

### Issue 3: Legend Overlaps with Title

**Problem**: Top legend overlaps with chart title

**Solution**: Use bottom legend or increase top margin
```go
// Use bottom instead
chart.LegendPosition = fynesimplechart.LegendBottom

// Or use left/right for titled charts
chart.LegendPosition = fynesimplechart.LegendRight
```

---

## Best Practices

### 1. Axis Titles
- ✅ Always include units
- ✅ Keep concise (3-4 words max)
- ✅ Use standard abbreviations
- ❌ Don't duplicate information from chart title

### 2. Legend Positioning
- ✅ Bottom for wide charts (aspect > 2:1)
- ✅ Right/Left for standard charts
- ✅ None for single-series
- ❌ Don't use top if you have a chart title

### 3. Manual Ranges
- ✅ Use for dashboard consistency
- ✅ Use for percentage/ratio data (0-100)
- ✅ Add padding to prevent clipping
- ❌ Don't use if data range is unpredictable

### 4. Data Labels
- ✅ Use for small datasets (<15 points)
- ✅ Match format to data type (currency, %)
- ✅ Use readable font size (9-12)
- ❌ Don't use with dense data (causes overlap)

---

## Summary

You've learned how to use all four enhanced features:

1. **Axis Titles** - Add context with `XAxisTitle` and `YAxisTitle`
2. **Legend Positioning** - Control layout with `LegendPosition`
3. **Manual Ranges** - Fix scales with `MinX`, `MaxX`, `MinY`, `MaxY`
4. **Data Labels** - Show values with `ShowDataLabels` and `LabelFormat`

These features transform basic charts into professional, presentation-ready visualizations.

## Next Steps

- Practice combining features in different scenarios
- Experiment with different legend positions
- Try various label formats for your data type
- Build a complete dashboard using manual ranges for consistency
- Review [Tutorial 12: Best Practices](12-best-practices.md) for design tips

---

**Need Help?**
- [Quick Reference](../QUICKSTART.md)
- [Complete API Documentation](../IMPROVEMENTS.md)
- [Feature Summary](../FEATURES_SUMMARY.md)
