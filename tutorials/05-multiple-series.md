# Tutorial 5: Multiple Data Series

## Overview

Learn how to display and compare multiple datasets on a single chart.

**Time to complete:** 15 minutes

## Prerequisites

- Completed [Tutorial 4: Customizing Appearance](04-customizing-appearance.md)

## Why Multiple Series?

Multiple series allow you to:
- Compare trends across datasets
- Show relationships between variables
- Display complementary data
- Analyze correlations

## Creating Multiple Series

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Multiple Series")
    w.Resize(fyne.NewSize(800, 600))

    // Series 1
    data1 := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 3),
        *fynesimplechart.NewNode(2, 5),
        *fynesimplechart.NewNode(3, 4),
        *fynesimplechart.NewNode(4, 7),
    }
    plot1 := fynesimplechart.NewPlot(data1, "Product A")
    plot1.ShowLine = true

    // Series 2
    data2 := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 2),
        *fynesimplechart.NewNode(2, 4),
        *fynesimplechart.NewNode(3, 6),
        *fynesimplechart.NewNode(4, 5),
    }
    plot2 := fynesimplechart.NewPlot(data2, "Product B")
    plot2.ShowLine = true

    // Combine into one chart
    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
        *plot1,
        *plot2,
    })
    chart.SetChartTitle("Product Comparison")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Differentiating Series

### By Color (Automatic)
```go
// Colors are assigned automatically
plot1 := fynesimplechart.NewPlot(data1, "Series 1")  // Blue
plot2 := fynesimplechart.NewPlot(data2, "Series 2")  // Orange
plot3 := fynesimplechart.NewPlot(data3, "Series 3")  // Green
```

### By Style
```go
import "image/color"

// Series 1: Line only
plot1 := fynesimplechart.NewPlot(data1, "Trend")
plot1.ShowLine = true
plot1.ShowPoints = false
plot1.LineWidth = 2

// Series 2: Points only
plot2 := fynesimplechart.NewPlot(data2, "Measurements")
plot2.ShowLine = false
plot2.ShowPoints = true
plot2.PointSize = 5

// Series 3: Both
plot3 := fynesimplechart.NewPlot(data3, "Combined")
plot3.ShowLine = true
plot3.ShowPoints = true
plot3.PlotColor = color.RGBA{R: 214, G: 39, B: 40, A: 255}
```

## Practical Example: Temperature Comparison

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Temperature Comparison")
    w.Resize(fyne.NewSize(900, 650))

    // Indoor temperature
    indoor := []fynesimplechart.Node{
        *fynesimplechart.NewNode(0, 22),
        *fynesimplechart.NewNode(4, 21),
        *fynesimplechart.NewNode(8, 22),
        *fynesimplechart.NewNode(12, 24),
        *fynesimplechart.NewNode(16, 25),
        *fynesimplechart.NewNode(20, 23),
        *fynesimplechart.NewNode(24, 22),
    }
    indoorPlot := fynesimplechart.NewPlot(indoor, "Indoor (°C)")
    indoorPlot.ShowLine = true
    indoorPlot.LineWidth = 2.5

    // Outdoor temperature
    outdoor := []fynesimplechart.Node{
        *fynesimplechart.NewNode(0, 15),
        *fynesimplechart.NewNode(4, 12),
        *fynesimplechart.NewNode(8, 16),
        *fynesimplechart.NewNode(12, 22),
        *fynesimplechart.NewNode(16, 25),
        *fynesimplechart.NewNode(20, 20),
        *fynesimplechart.NewNode(24, 16),
    }
    outdoorPlot := fynesimplechart.NewPlot(outdoor, "Outdoor (°C)")
    outdoorPlot.ShowLine = true
    outdoorPlot.LineWidth = 2.5

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
        *indoorPlot,
        *outdoorPlot,
    })
    chart.SetChartTitle("24-Hour Temperature Monitoring")

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Best Practices

### 1. Limit Series Count
```go
// Good: 2-5 series
chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
    *plot1, *plot2, *plot3,
})

// Avoid: Too many series (>6) becomes cluttered
```

### 2. Use Consistent Data Ranges
```go
// All series should cover similar X ranges
// Good:
series1: X from 1-10
series2: X from 1-10
series3: X from 1-10

// Avoid:
series1: X from 1-10
series2: X from 5-15  // Different range
```

### 3. Meaningful Titles
```go
// Good titles
plot1 := fynesimplechart.NewPlot(data1, "North Region")
plot2 := fynesimplechart.NewPlot(data2, "South Region")

// Avoid generic titles
plot1 := fynesimplechart.NewPlot(data1, "Data 1")
plot2 := fynesimplechart.NewPlot(data2, "Data 2")
```

## Common Patterns

### Before/After Comparison
```go
before := fynesimplechart.NewPlot(beforeData, "Before Optimization")
before.ShowLine = true
before.LineWidth = 2

after := fynesimplechart.NewPlot(afterData, "After Optimization")
after.ShowLine = true
after.LineWidth = 2

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
    *before, *after,
})
chart.SetChartTitle("Performance Improvement")
```

### Regional Comparison
```go
north := fynesimplechart.NewPlot(northData, "North")
south := fynesimplechart.NewPlot(southData, "South")
east := fynesimplechart.NewPlot(eastData, "East")
west := fynesimplechart.NewPlot(westData, "West")

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
    *north, *south, *east, *west,
})
chart.SetChartTitle("Regional Sales Performance")
```

### Actual vs Forecast
```go
import "image/color"

actual := fynesimplechart.NewPlot(actualData, "Actual")
actual.ShowLine = true
actual.ShowPoints = true
actual.LineWidth = 2.5
actual.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}

forecast := fynesimplechart.NewPlot(forecastData, "Forecast")
forecast.ShowLine = true
forecast.ShowPoints = false
forecast.LineWidth = 1.5
forecast.PlotColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
    *actual, *forecast,
})
```

## Exercise

Create a chart comparing website traffic from three sources:
- Organic: 100, 120, 135, 150
- Paid: 50, 55, 70, 85
- Social: 30, 45, 60, 65

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
    w := a.NewWindow("Traffic Sources")
    w.Resize(fyne.NewSize(800, 600))

    organic := []float32{100, 120, 135, 150}
    paid := []float32{50, 55, 70, 85}
    social := []float32{30, 45, 60, 65}

    createPlot := func(data []float32, title string) *fynesimplechart.Plot {
        nodes := []fynesimplechart.Node{}
        for week, value := range data {
            nodes = append(nodes, *fynesimplechart.NewNode(
                float32(week+1), value,
            ))
        }
        plot := fynesimplechart.NewPlot(nodes, title)
        plot.ShowLine = true
        plot.ShowPoints = true
        plot.LineWidth = 2
        return plot
    }

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
        *createPlot(organic, "Organic"),
        *createPlot(paid, "Paid"),
        *createPlot(social, "Social"),
    })
    chart.SetChartTitle("Weekly Traffic by Source")

    w.SetContent(chart)
    w.ShowAndRun()
}
```
</details>

## Summary

You learned:
- ✅ Creating multiple data series
- ✅ Differentiating series visually
- ✅ Comparing datasets effectively
- ✅ Best practices for multiple series

## Next Steps

Continue to [Tutorial 6: Working with Negative Values](06-negative-values.md) to learn about plotting data in all four quadrants.
