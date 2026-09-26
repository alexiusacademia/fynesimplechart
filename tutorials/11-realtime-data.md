# Tutorial 10: Real-time Data

## Overview

Learn how to update charts dynamically with real-time data.

**Time to complete:** 15 minutes

## Understanding Widget Refresh

`NewGraphWidget` stores its own copy of every `Plot` you pass in. Changing the
original `plot` afterwards (for example `plot.Nodes = nodes`) does **not**
affect the chart. Update the chart's data with `SetPlotNodes`, which also
redraws it:

```go
chart.SetPlotNodes(0, nodes)  // Replace the data of the first plot and redraw
```

You can also edit `chart.Plots[i]` directly and then call `chart.Refresh()`.

### Updating from a goroutine

Since Fyne v2.6, any change to the UI made from another goroutine must be
wrapped in `fyne.Do`. Otherwise Fyne logs threading warnings. Pass a copy of
your slice so the goroutine can keep changing its own data while the chart draws:

```go
snapshot := append([]fynesimplechart.Node(nil), nodes...)
fyne.Do(func() {
    chart.SetPlotNodes(0, snapshot)
})
```

On Fyne versions before v2.6, `fyne.Do` does not exist; call
`chart.SetPlotNodes(0, snapshot)` directly.

## Basic Real-time Example

```go
package main

import (
    "time"
    "math/rand"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Real-time Data")
    w.Resize(fyne.NewSize(800, 600))

    nodes := []fynesimplechart.Node{}
    plot := fynesimplechart.NewPlot(nodes, "Live Data")
    plot.ShowLine = true
    plot.LineWidth = 2

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Real-time Monitoring")

    // Update data every second
    go func() {
        x := float32(0)
        for {
            // Add new data point
            y := float32(rand.Float64()*10 + 50)
            nodes = append(nodes, *fynesimplechart.NewNode(x, y))

            // Keep only last 20 points
            if len(nodes) > 20 {
                nodes = nodes[1:]
                // Adjust X coordinates
                for i := range nodes {
                    nodes[i].X = float32(i)
                }
            }

            // Update the chart (on the UI goroutine) with a copy of the data
            snapshot := append([]fynesimplechart.Node(nil), nodes...)
            fyne.Do(func() {
                chart.SetPlotNodes(0, snapshot)
            })

            x++
            time.Sleep(1 * time.Second)
        }
    }()

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Streaming Data Pattern

```go
type DataStream struct {
    nodes []fynesimplechart.Node
    maxPoints int
}

func (ds *DataStream) AddPoint(x, y float32) {
    ds.nodes = append(ds.nodes, *fynesimplechart.NewNode(x, y))

    if len(ds.nodes) > ds.maxPoints {
        ds.nodes = ds.nodes[1:]
    }
}

func (ds *DataStream) GetNodes() []fynesimplechart.Node {
    return ds.nodes
}
```

## Best Practices

### 1. Limit Data Points
```go
// Keep only recent data
maxPoints := 50
if len(nodes) > maxPoints {
    nodes = nodes[len(nodes)-maxPoints:]
}
```

### 2. Update Frequency
```go
// Good: 0.5-2 seconds
time.Sleep(1 * time.Second)

// Avoid: Too frequent (<100ms) causes flicker
// Avoid: Too slow (>5s) not "real-time"
```

### 3. Goroutine Management
```go
// Use context for clean shutdown
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Update data
            time.Sleep(1 * time.Second)
        }
    }
}()
```

## Practical Example: Temperature Monitor

```go
package main

import (
    "time"
    "math"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("Temperature Monitor")
    w.Resize(fyne.NewSize(900, 650))

    nodes := []fynesimplechart.Node{}
    plot := fynesimplechart.NewPlot(nodes, "Temperature (°C)")
    plot.ShowLine = true
    plot.ShowPoints = true
    plot.LineWidth = 2
    plot.PointSize = 3

    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.SetChartTitle("Live Temperature Monitoring")

    // Simulate temperature sensor
    go func() {
        t := 0.0
        for {
            // Simulate temperature with sine wave + noise
            baseTemp := 22.0 + math.Sin(t/10.0)*3
            noise := (math.Sin(t*7) + math.Sin(t*13)) * 0.5
            temp := float32(baseTemp + noise)

            nodes = append(nodes, *fynesimplechart.NewNode(
                float32(len(nodes)),
                temp,
            ))

            // Keep last 30 readings
            if len(nodes) > 30 {
                nodes = nodes[1:]
                for i := range nodes {
                    nodes[i].X = float32(i)
                }
            }

            snapshot := append([]fynesimplechart.Node(nil), nodes...)
            fyne.Do(func() {
                chart.SetPlotNodes(0, snapshot)
            })

            t += 0.5
            time.Sleep(500 * time.Millisecond)
        }
    }()

    w.SetContent(chart)
    w.ShowAndRun()
}
```

## Summary

- ✅ Real-time updates with goroutines
- ✅ Chart refresh mechanism
- ✅ Data window management
- ✅ Update frequency control

## Next Steps

Continue to [Tutorial 12: Best Practices](12-best-practices.md).
