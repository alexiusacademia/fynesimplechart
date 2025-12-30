# Tutorial 4: Customizing Appearance

## Overview

Learn how to customize the visual appearance of your charts:
- Custom colors
- Chart titles
- Grid visibility
- Complete styling control

**Time to complete:** 20 minutes

## Prerequisites

- Completed [Tutorial 3: Line Charts](03-line-charts.md)
- Basic understanding of RGB colors

## Understanding Plot Colors

FyneSimpleChart uses professional color palettes by default, but you can customize any plot's color.

### Default Colors

When you don't specify a color, plots automatically get colors from a professional palette:

```go
plot1 := fynesimplechart.NewPlot(data1, "Series 1")  // Blue
plot2 := fynesimplechart.NewPlot(data2, "Series 2")  // Orange
plot3 := fynesimplechart.NewPlot(data3, "Series 3")  // Green
// ... continues with 10 professional colors
```

## Setting Custom Colors

### Basic Color Setting

```go
import "image/color"

plot := fynesimplechart.NewPlot(nodes, "My Data")

// Set a custom color (RGB + Alpha)
plot.PlotColor = color.RGBA{R: 220, G: 20, B: 60, A: 255}
```

### Common Color Examples

```go
// Red
plot.PlotColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}

// Green
plot.PlotColor = color.RGBA{R: 0, G: 255, B: 0, A: 255}

// Blue
plot.PlotColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}

// Purple
plot.PlotColor = color.RGBA{R: 128, G: 0, B: 128, A: 255}

// Orange
plot.PlotColor = color.RGBA{R: 255, G: 165, B: 0, A: 255}

// Cyan
plot.PlotColor = color.RGBA{R: 0, G: 255, B: 255, A: 255}
```

## Professional Color Palette

Here are professionally selected colors that work well together:

```go
// Professional Blue
color.RGBA{R: 31, G: 119, B: 180, A: 255}

// Professional Orange
color.RGBA{R: 255, G: 127, B: 14, A: 255}

// Professional Green
color.RGBA{R: 44, G: 160, B: 44, A: 255}

// Professional Red
color.RGBA{R: 214, G: 39, B: 40, A: 255}

// Professional Purple
color.RGBA{R: 148, G: 103, B: 189, A: 255}

// Professional Brown
color.RGBA{R: 140, G: 86, B: 75, A: 255}

// Professional Pink
color.RGBA{R: 227, G: 119, B: 194, A: 255}

// Professional Gray
color.RGBA{R: 127, G: 127, B: 127, A: 255}

// Professional Olive
color.RGBA{R: 188, G: 189, B: 34, A: 255}

// Professional Cyan
color.RGBA{R: 23, G: 190, B: 207, A: 255}
```

## Complete Styling Example

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
    w := a.NewWindow("Styled Chart")
    w.Resize(fyne.NewSize(800, 600))

    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 3),
        *fynesimplechart.NewNode(2, 5),
        *fynesimplechart.NewNode(3, 4),
        *fynesimplechart.NewNode(4, 7),
        *fynesimplechart.NewNode(5, 6),
    }

    plot := fynesimplechart.NewPlot(nodes, "Revenue")

    // Complete styling
    plot.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}  // Blue
    plot.ShowLine = true
    plot.ShowPoints = true
    plot.LineWidth = 2.5
    plot.PointSize = 5

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Q1 2024 Revenue Analysis")
    chart.ShowGrid = true

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Chart Titles

### Adding a Main Title

```go
chart := fynesimplechart.NewGraphWidget(plots)
chart.SetChartTitle("Monthly Sales Report")
```

### No Title

```go
// Don't call SetChartTitle() or set to empty string
chart.SetChartTitle("")
```

### Title Best Practices

```go
// Good titles - specific and descriptive
chart.SetChartTitle("Q1 2024 Revenue by Region")
chart.SetChartTitle("Customer Satisfaction Trend")
chart.SetChartTitle("Temperature Monitoring - Building A")

// Avoid - too generic
chart.SetChartTitle("Chart")
chart.SetChartTitle("Data")
```

## Grid Configuration

### Show Grid (Default)

```go
chart := fynesimplechart.NewGraphWidget(plots)
chart.ShowGrid = true
```

### Hide Grid

```go
chart.ShowGrid = false  // Clean look, harder to read values
```

### When to Hide Grid

- Presentation mode (cleaner look)
- When exact values aren't important
- Minimalist design preference

### When to Show Grid

- Data analysis
- Reading specific values
- Professional reports
- Scientific visualization

## Creating Color Themes

### Corporate Theme Example

```go
// Company colors: Blue and Gold
corporateBlue := color.RGBA{R: 0, G: 71, B: 171, A: 255}
corporateGold := color.RGBA{R: 255, G: 184, B: 28, A: 255}

plot1 := fynesimplechart.NewPlot(data1, "Division A")
plot1.PlotColor = corporateBlue

plot2 := fynesimplechart.NewPlot(data2, "Division B")
plot2.PlotColor = corporateGold
```

### Scientific Theme Example

```go
// Cool colors for scientific data
scientificBlue := color.RGBA{R: 65, G: 105, B: 225, A: 255}
scientificGreen := color.RGBA{R: 46, G: 139, B: 87, A: 255}
scientificPurple := color.RGBA{R: 138, G: 43, B: 226, A: 255}

temp := fynesimplechart.NewPlot(tempData, "Temperature")
temp.PlotColor = scientificBlue

pressure := fynesimplechart.NewPlot(pressureData, "Pressure")
pressure.PlotColor = scientificGreen

humidity := fynesimplechart.NewPlot(humidityData, "Humidity")
humidity.PlotColor = scientificPurple
```

### Traffic Light Theme Example

```go
// Good/Warning/Critical states
good := color.RGBA{R: 76, G: 175, B: 80, A: 255}      // Green
warning := color.RGBA{R: 255, G: 193, B: 7, A: 255}   // Amber
critical := color.RGBA{R: 244, G: 67, B: 54, A: 255}  // Red

plotGood := fynesimplechart.NewPlot(goodData, "Normal")
plotGood.PlotColor = good

plotWarning := fynesimplechart.NewPlot(warningData, "Warning")
plotWarning.PlotColor = warning

plotCritical := fynesimplechart.NewPlot(criticalData, "Critical")
plotCritical.PlotColor = critical
```

## Complete Styling Reference

```go
plot := fynesimplechart.NewPlot(nodes, "Title")

// Color
plot.PlotColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}

// Points
plot.ShowPoints = true   // true/false
plot.PointSize = 4       // 1.0 to 10.0 (recommended: 2-6)

// Lines
plot.ShowLine = true     // true/false
plot.LineWidth = 2.5     // 0.5 to 5.0 (recommended: 1.5-3.0)

// Chart-level settings
chart.SetChartTitle("My Chart")  // String or ""
chart.ShowGrid = true            // true/false
```

## Practical Examples

### Example 1: Financial Dashboard

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
    w := a.NewWindow("Financial Dashboard")
    w.Resize(fyne.NewSize(900, 650))

    // Revenue data
    revenue := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 45),
        *fynesimplechart.NewNode(2, 52),
        *fynesimplechart.NewNode(3, 48),
        *fynesimplechart.NewNode(4, 61),
    }
    revPlot := fynesimplechart.NewPlot(revenue, "Revenue ($K)")
    revPlot.PlotColor = color.RGBA{R: 76, G: 175, B: 80, A: 255}  // Green (positive)
    revPlot.ShowLine = true
    revPlot.ShowPoints = true
    revPlot.LineWidth = 3
    revPlot.PointSize = 5

    // Expenses data
    expenses := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 32),
        *fynesimplechart.NewNode(2, 35),
        *fynesimplechart.NewNode(3, 33),
        *fynesimplechart.NewNode(4, 38),
    }
    expPlot := fynesimplechart.NewPlot(expenses, "Expenses ($K)")
    expPlot.PlotColor = color.RGBA{R: 244, G: 67, B: 54, A: 255}  // Red (negative)
    expPlot.ShowLine = true
    expPlot.ShowPoints = true
    expPlot.LineWidth = 3
    expPlot.PointSize = 5

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*revPlot, *expPlot})
    chart.SetChartTitle("Q1 2024 Financial Overview")
    chart.ShowGrid = true

    w.SetContent(chart)
    w.ShowAndRun()
}
```

### Example 2: Scientific Measurements

```go
// Minimal, professional scientific style
plot := fynesimplechart.NewPlot(measurements, "Experiment A")
plot.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}
plot.ShowLine = true
plot.ShowPoints = true
plot.LineWidth = 1.5  // Thin, precise line
plot.PointSize = 3    // Small, discrete points

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
chart.SetChartTitle("Phase Transition Temperature Analysis")
chart.ShowGrid = true  // Essential for reading values
```

### Example 3: Marketing Presentation

```go
// Bold, presentation-ready style
plot := fynesimplechart.NewPlot(engagement, "User Engagement")
plot.PlotColor = color.RGBA{R: 255, G: 87, B: 34, A: 255}  // Bright orange
plot.ShowLine = true
plot.ShowPoints = false  // Clean lines only
plot.LineWidth = 4       // Thick, visible from distance

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
chart.SetChartTitle("2024 User Engagement Growth")
chart.ShowGrid = false  // Clean, minimal look
```

## Color Selection Tips

### 1. Contrast is Key
```go
// Good contrast - easy to distinguish
blue := color.RGBA{R: 31, G: 119, B: 180, A: 255}
orange := color.RGBA{R: 255, G: 127, B: 14, A: 255}

// Bad contrast - hard to distinguish
lightBlue := color.RGBA{R: 150, G: 180, B: 200, A: 255}
lightGreen := color.RGBA{R: 150, G: 200, B: 180, A: 255}
```

### 2. Semantic Colors
```go
// Use colors that match meaning
positive := color.RGBA{R: 0, G: 255, B: 0, A: 255}     // Green
negative := color.RGBA{R: 255, G: 0, B: 0, A: 255}     // Red
neutral := color.RGBA{R: 128, G: 128, B: 128, A: 255}  // Gray
```

### 3. Accessibility
```go
// Colorblind-friendly palette
blue := color.RGBA{R: 0, G: 114, B: 178, A: 255}
orange := color.RGBA{R: 230, G: 159, B: 0, A: 255}
skyBlue := color.RGBA{R: 86, G: 180, B: 233, A: 255}
```

## Exercise

Create a styled chart comparing three products' sales with:
- Product A (leader) - Green, thick line
- Product B (competitive) - Orange, medium line
- Product C (struggling) - Red, thin line

<details>
<summary>Solution</summary>

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
    w := a.NewWindow("Product Comparison")
    w.Resize(fyne.NewSize(800, 600))

    // Product A - Leader
    dataA := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 85),
        *fynesimplechart.NewNode(2, 90),
        *fynesimplechart.NewNode(3, 92),
        *fynesimplechart.NewNode(4, 95),
    }
    plotA := fynesimplechart.NewPlot(dataA, "Product A")
    plotA.PlotColor = color.RGBA{R: 76, G: 175, B: 80, A: 255}  // Green
    plotA.ShowLine = true
    plotA.LineWidth = 3.5

    // Product B - Competitive
    dataB := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 70),
        *fynesimplechart.NewNode(2, 72),
        *fynesimplechart.NewNode(3, 75),
        *fynesimplechart.NewNode(4, 78),
    }
    plotB := fynesimplechart.NewPlot(dataB, "Product B")
    plotB.PlotColor = color.RGBA{R: 255, G: 152, B: 0, A: 255}  // Orange
    plotB.ShowLine = true
    plotB.LineWidth = 2.5

    // Product C - Struggling
    dataC := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 45),
        *fynesimplechart.NewNode(2, 42),
        *fynesimplechart.NewNode(3, 40),
        *fynesimplechart.NewNode(4, 38),
    }
    plotC := fynesimplechart.NewPlot(dataC, "Product C")
    plotC.PlotColor = color.RGBA{R: 244, G: 67, B: 54, A: 255}  // Red
    plotC.ShowLine = true
    plotC.LineWidth = 1.5

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plotA, *plotB, *plotC})
    chart.SetChartTitle("Quarterly Product Sales Comparison")

    w.SetContent(chart)
    w.ShowAndRun()
}
```
</details>

## Summary

You learned:
- ✅ How to set custom colors
- ✅ Professional color palettes
- ✅ Adding and styling chart titles
- ✅ Controlling grid visibility
- ✅ Creating themed charts
- ✅ Complete styling control

## Next Steps

Continue to [Tutorial 5: Multiple Data Series](05-multiple-series.md) to learn how to compare multiple datasets on one chart.
