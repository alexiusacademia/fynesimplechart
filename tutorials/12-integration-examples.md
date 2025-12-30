# Tutorial 12: Integration Examples

## Overview

Real-world examples of integrating FyneSimpleChart into applications.

**Time to complete:** 20 minutes

## Example 1: Dashboard Application

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Business Dashboard")
    w.Resize(fyne.NewSize(1200, 800))

    // Revenue chart
    revenueData := []float32{45, 52, 48, 61, 58, 67}
    revenueNodes := createNodes(revenueData)
    revenuePlot := fynesimplechart.NewPlot(revenueNodes, "Revenue ($K)")
    revenuePlot.ShowLine = true
    revenuePlot.LineWidth = 2.5
    revenueChart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*revenuePlot})
    revenueChart.SetChartTitle("Monthly Revenue")
    revenueChart.Resize(fyne.NewSize(550, 350))

    // Users chart
    usersData := []float32{1200, 1350, 1420, 1560, 1650, 1780}
    usersNodes := createNodes(usersData)
    usersPlot := fynesimplechart.NewPlot(usersNodes, "Active Users")
    usersPlot.ShowLine = true
    usersPlot.LineWidth = 2.5
    usersChart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*usersPlot})
    usersChart.SetChartTitle("User Growth")
    usersChart.Resize(fyne.NewSize(550, 350))

    // Layout
    topRow := container.NewGridWithColumns(2, revenueChart, usersChart)

    // Add some metrics
    metrics := widget.NewLabel("Key Metrics: Revenue +15% | Users +48%")

    content := container.NewBorder(nil, metrics, nil, nil, topRow)

    w.SetContent(content)
    w.ShowAndRun()
}

func createNodes(data []float32) []fynesimplechart.Node {
    nodes := []fynesimplechart.Node{}
    for i, val := range data {
        nodes = append(nodes, *fynesimplechart.NewNode(
            float32(i+1),
            val,
        ))
    }
    return nodes
}
```

## Example 2: Data Analysis Tool

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Data Analyzer")
    w.Resize(fyne.NewSize(1000, 700))

    // Initial data
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 3),
        *fynesimplechart.NewNode(2, 5),
        *fynesimplechart.NewNode(3, 4),
    }

    plot := fynesimplechart.NewPlot(nodes, "Dataset")
    plot.ShowLine = true
    plot.ShowPoints = true

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Data Analysis")
    chart.Resize(fyne.NewSize(900, 550))

    // Controls
    showLineCheck := widget.NewCheck("Show Line", func(checked bool) {
        plot.ShowLine = checked
        chart.Refresh()
    })
    showLineCheck.SetChecked(true)

    showPointsCheck := widget.NewCheck("Show Points", func(checked bool) {
        plot.ShowPoints = checked
        chart.Refresh()
    })
    showPointsCheck.SetChecked(true)

    showGridCheck := widget.NewCheck("Show Grid", func(checked bool) {
        chart.ShowGrid = checked
        chart.Refresh()
    })
    showGridCheck.SetChecked(true)

    controls := container.NewHBox(
        showLineCheck,
        showPointsCheck,
        showGridCheck,
    )

    content := container.NewBorder(nil, controls, nil, nil, chart)

    w.SetContent(content)
    w.ShowAndRun()
}
```

## Example 3: IoT Monitoring System

```go
package main

import (
    "time"
    "math"
    "math/rand"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/alexiusacademia/fynesimplechart"
)

type Sensor struct {
    name  string
    nodes []fynesimplechart.Node
    plot  *fynesimplechart.Plot
}

func main() {
    a := app.New()
    w := a.NewWindow("IoT Monitoring")
    w.Resize(fyne.NewSize(1000, 800))

    // Create sensors
    tempSensor := &Sensor{name: "Temperature (°C)"}
    tempSensor.nodes = []fynesimplechart.Node{}
    tempSensor.plot = fynesimplechart.NewPlot(tempSensor.nodes, tempSensor.name)
    tempSensor.plot.ShowLine = true
    tempSensor.plot.LineWidth = 2

    humiditySensor := &Sensor{name: "Humidity (%)"}
    humiditySensor.nodes = []fynesimplechart.Node{}
    humiditySensor.plot = fynesimplechart.NewPlot(humiditySensor.nodes, humiditySensor.name)
    humiditySensor.plot.ShowLine = true
    humiditySensor.plot.LineWidth = 2

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
        *tempSensor.plot,
        *humiditySensor.plot,
    })
    chart.SetChartTitle("Environmental Monitoring")
    chart.Resize(fyne.NewSize(900, 600))

    // Status label
    status := widget.NewLabel("Status: Monitoring...")

    // Simulate sensor readings
    go func() {
        t := 0.0
        for {
            // Generate sensor data
            temp := float32(22 + math.Sin(t/10)*3 + rand.Float64()*2)
            humidity := float32(60 + math.Cos(t/8)*10 + rand.Float64()*5)

            // Add readings
            tempSensor.nodes = append(tempSensor.nodes, *fynesimplechart.NewNode(
                float32(len(tempSensor.nodes)),
                temp,
            ))
            humiditySensor.nodes = append(humiditySensor.nodes, *fynesimplechart.NewNode(
                float32(len(humiditySensor.nodes)),
                humidity,
            ))

            // Keep last 30 readings
            if len(tempSensor.nodes) > 30 {
                tempSensor.nodes = tempSensor.nodes[1:]
                for i := range tempSensor.nodes {
                    tempSensor.nodes[i].X = float32(i)
                }
            }
            if len(humiditySensor.nodes) > 30 {
                humiditySensor.nodes = humiditySensor.nodes[1:]
                for i := range humiditySensor.nodes {
                    humiditySensor.nodes[i].X = float32(i)
                }
            }

            // Update plots
            tempSensor.plot.Nodes = tempSensor.nodes
            humiditySensor.plot.Nodes = humiditySensor.nodes
            chart.Refresh()

            status.SetText("Status: Last update " + time.Now().Format("15:04:05"))

            t += 0.5
            time.Sleep(1 * time.Second)
        }
    }()

    content := container.NewBorder(nil, status, nil, nil, chart)

    w.SetContent(content)
    w.ShowAndRun()
}
```

## Example 4: Financial Portfolio Tracker

```go
package main

import (
    "image/color"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Portfolio Tracker")
    w.Resize(fyne.NewSize(1100, 750))

    // Stock A
    stockA := []float32{100, 105, 103, 108, 110, 107, 112}
    nodesA := createDailyNodes(stockA)
    plotA := fynesimplechart.NewPlot(nodesA, "Stock A")
    plotA.ShowLine = true
    plotA.ShowPoints = true
    plotA.LineWidth = 2
    plotA.PlotColor = color.RGBA{R: 76, G: 175, B: 80, A: 255}

    // Stock B
    stockB := []float32{50, 48, 52, 54, 53, 56, 58}
    nodesB := createDailyNodes(stockB)
    plotB := fynesimplechart.NewPlot(nodesB, "Stock B")
    plotB.ShowLine = true
    plotB.ShowPoints = true
    plotB.LineWidth = 2
    plotB.PlotColor = color.RGBA{R: 33, G: 150, B: 243, A: 255}

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{
        *plotA, *plotB,
    })
    chart.SetChartTitle("Portfolio Performance (Last 7 Days)")
    chart.Resize(fyne.NewSize(1000, 600))

    // Summary
    summary := widget.NewLabel("Stock A: +12% | Stock B: +16%")

    content := container.NewBorder(nil, summary, nil, nil, chart)

    w.SetContent(content)
    w.ShowAndRun()
}

func createDailyNodes(prices []float32) []fynesimplechart.Node {
    nodes := []fynesimplechart.Node{}
    for day, price := range prices {
        nodes = append(nodes, *fynesimplechart.NewNode(
            float32(day+1),
            price,
        ))
    }
    return nodes
}
```

## Integration Tips

### 1. Layout Integration
Use Fyne containers to combine charts with other widgets:
```go
// Border layout
container.NewBorder(top, bottom, left, right, center)

// Grid layout
container.NewGridWithColumns(2, chart1, chart2)

// VBox/HBox
container.NewVBox(title, chart, controls)
```

### 2. Responsive Sizing
```go
// Fixed size
chart.Resize(fyne.NewSize(800, 600))

// Responsive (fills container)
// Don't call Resize() - let container manage size
```

### 3. State Management
```go
type AppState struct {
    data   []float32
    plot   *fynesimplechart.Plot
    chart  *fynesimplechart.ScatterPlot
}

func (s *AppState) UpdateData(newData []float32) {
    s.data = newData
    s.plot.Nodes = createNodes(newData)
    s.chart.Refresh()
}
```

## Summary

- ✅ Dashboard integration
- ✅ Interactive controls
- ✅ Real-time monitoring
- ✅ Multi-chart layouts
- ✅ State management

## Congratulations!

You've completed all tutorials and are now ready to create professional charts with FyneSimpleChart!

## Additional Resources

- [Quick Start Guide](../QUICKSTART.md)
- [Examples Directory](../examples/)
- [Improvements Documentation](../IMPROVEMENTS.md)
- Main library: [github.com/alexiusacademia/fynesimplechart](https://github.com/alexiusacademia/fynesimplechart)
