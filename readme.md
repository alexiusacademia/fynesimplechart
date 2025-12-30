# FyneSimpleChart

A professional, feature-rich charting library for [Fyne](https://fyne.io/) applications with industry-standard visualization capabilities.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22.0-blue)](https://go.dev/)
[![Fyne Version](https://img.shields.io/badge/fyne-v2.4.4-blue)](https://fyne.io/)

## ✨ Features

### Chart Types
- ✅ **Scatter Plots** - Data point visualization
- ✅ **Line Charts** - Continuous data trends
- ✅ **Bar Charts** - Categorical data comparison with grouped bars support
- ✅ **Area Fills** - Shaded regions (fill to zero or between curves)
- ⬜ Pie Charts (planned)
- ⬜ Stacked Bars (planned)

### Professional Features
- ✅ **Smart Grid System** - Automatic tick intervals with "nice numbers" algorithm
- ✅ **Axis Labels & Titles** - Numeric labels with dynamic precision plus custom axis titles
- ✅ **Negative Values** - Full support for all four quadrants
- ✅ **Multiple Series** - Compare unlimited datasets with auto-colors
- ✅ **Custom Styling** - Colors, line widths, point sizes, bar borders
- ✅ **Flexible Legends** - Positionable legends (top/bottom/left/right) or hide completely
- ✅ **Data Labels** - Show values directly on points/bars with custom formatting
- ✅ **Manual Axis Ranges** - Override auto-scaling for consistent comparisons
- ✅ **Chart Titles** - Main title and series legends
- ✅ **Real-time Updates** - Dynamic data visualization
- ✅ **Professional Color Palette** - D3.js/Plotly-inspired 10-color system

## 📦 Installation

```bash
go get github.com/alexiusacademia/fynesimplechart
```

**Requirements:**
- Go 1.22.0 or later
- Fyne v2.4.4 or later

## 🚀 Quick Start

### Simple Scatter Plot

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/alexiusacademia/fynesimplechart"
)

func main() {
    a := app.New()
    w := a.NewWindow("My First Chart")

    // Create data points
    nodes := []fynesimplechart.Node{
        *fynesimplechart.NewNode(1, 2),
        *fynesimplechart.NewNode(2, 4),
        *fynesimplechart.NewNode(3, 3),
        *fynesimplechart.NewNode(4, 5),
    }

    // Create plot
    plot := fynesimplechart.NewPlot(nodes, "My Data")

    // Create chart
    chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*plot})
    chart.Resize(fyne.NewSize(600, 400))

    w.SetContent(chart)
    w.ShowAndRun()
}
```

### Line Chart

```go
plot := fynesimplechart.NewPlot(nodes, "Temperature")
plot.ShowLine = true
plot.LineWidth = 2.0
plot.PlotColor = color.RGBA{R: 31, G: 119, B: 180, A: 255}
```

### Bar Chart

```go
plot := fynesimplechart.NewPlot(nodes, "Monthly Sales")
plot.ShowBars = true
plot.BarWidth = 0.8
plot.ShowPoints = false
plot.ShowLine = false

// Optional: Add borders
plot.BarBorderWidth = 1
```

### Grouped Bar Chart

```go
// Q1 data
q1 := fynesimplechart.NewPlot(data1, "Q1")
q1.ShowBars = true
q1.BarWidth = 0.25

// Q2 data (offset X by 0.3)
q2 := fynesimplechart.NewPlot(data2, "Q2")
q2.ShowBars = true
q2.BarWidth = 0.25

chart := fynesimplechart.NewGraphWidget([]fynesimplechart.Plot{*q1, *q2})
```

### Area Fill

```go
plot := fynesimplechart.NewPlot(nodes, "Revenue")
plot.ShowLine = true
plot.FillArea = true
plot.FillToZero = true  // Fill from curve to Y=0
```

## 📚 Documentation

- **[Quick Start Guide](QUICKSTART.md)** - Fast reference for common tasks
- **[Tutorials](tutorials/)** - 15 comprehensive step-by-step tutorials
- **[API Documentation](IMPROVEMENTS.md)** - Complete feature reference

### Tutorials

Complete tutorial series covering everything from basics to advanced features:

1. [Getting Started](tutorials/01-getting-started.md) - Installation and first chart
2. [Basic Scatter Plot](tutorials/02-basic-scatter-plot.md) - Point visualization
3. [Line Charts](tutorials/03-line-charts.md) - Connected data trends
4. [Customizing Appearance](tutorials/04-customizing-appearance.md) - Colors and styling
5. [Multiple Data Series](tutorials/05-multiple-series.md) - Comparing datasets
6. [Working with Negative Values](tutorials/06-negative-values.md) - All quadrants
7. [Grid and Axes Configuration](tutorials/07-grid-and-axes.md) - Grid control
8. [Color Palettes](tutorials/08-color-palettes.md) - Color schemes
9. [Area Fills](tutorials/09-area-fills.md) - Shaded regions
10. [Mathematical Visualizations](tutorials/10-mathematical-functions.md) - Function plotting
11. [Real-time Data](tutorials/11-realtime-data.md) - Dynamic updates
12. [Best Practices](tutorials/12-best-practices.md) - Performance & design
13. [Integration Examples](tutorials/13-integration-examples.md) - Real applications
14. [Bar Charts](tutorials/14-bar-charts.md) - Categorical data
15. [Enhanced Features](tutorials/15-enhanced-features.md) - Axis titles, legends, ranges, labels

## 🎨 Configuration Options

### Plot Properties

```go
plot := fynesimplechart.NewPlot(nodes, "Title")

// Visibility
plot.ShowLine = true
plot.ShowPoints = true
plot.ShowBars = false

// Styling
plot.LineWidth = 2.5
plot.PointSize = 4.0
plot.PlotColor = myColor

// Bar Charts
plot.ShowBars = true
plot.BarWidth = 0.8
plot.BarBorderWidth = 1
plot.BarBorderColor = borderColor

// Area Fill
plot.FillArea = true
plot.FillToZero = true
plot.FillToPlotIdx = 0
plot.FillColor = fillColor

// Data Labels
plot.ShowDataLabels = true
plot.LabelFormat = "%.1f"  // or "$%.0fK", "%.0f%%", etc.
plot.LabelSize = 10
plot.LabelColor = labelColor
```

### Chart Properties

```go
chart := fynesimplechart.NewGraphWidget(plots)

// Titles
chart.SetChartTitle("My Chart Title")
chart.XAxisTitle = "Time (seconds)"
chart.YAxisTitle = "Value (units)"

// Display
chart.ShowGrid = true   // Toggle grid visibility
chart.ShowLegend = true // Toggle legend visibility

// Legend Position
chart.LegendPosition = fynesimplechart.LegendBottom
// Options: LegendRight (default), LegendBottom, LegendTop, LegendLeft, LegendNone

// Manual Axis Ranges (optional)
minY := float32(0)
maxY := float32(100)
chart.MinY = &minY
chart.MaxY = &maxY
```

## 🎯 Use Cases

- **Business Analytics** - Sales trends, revenue charts, KPI dashboards
- **Scientific Visualization** - Data analysis, experimental results
- **Financial Applications** - Stock charts, profit/loss, budget tracking
- **IoT Monitoring** - Sensor data, real-time metrics
- **Survey Results** - Response analysis, demographic data
- **Performance Metrics** - System monitoring, benchmarks

## 🏗️ Architecture

Built with industry-standard charting practices:

- **Nice Numbers Algorithm** - Human-friendly tick intervals (1, 2, 5, 10...)
- **Automatic Scaling** - Intelligent axis range calculation
- **Coordinate Transformation** - Clean data-to-screen mapping
- **Modular Rendering** - Separated concerns (grid, axes, plots, legend)
- **Professional Colors** - Accessible, high-contrast palette

## 🤝 Contributing

Contributions are welcome! Please feel free to submit pull requests, report bugs, or suggest features.

## 📝 License

This project is licensed under the MIT License.

## 🙏 Acknowledgments

- Built with [Fyne](https://fyne.io/) - A beautiful cross-platform GUI toolkit for Go
- Color palette inspired by [D3.js](https://d3js.org/) and [Plotly](https://plotly.com/)
- Nice numbers algorithm based on industry-standard charting libraries

## 📸 Screenshots

### Basic Charts
<img width="1198" height="827" alt="Scatter and Line Charts" src="https://github.com/user-attachments/assets/4ead8bb9-cb52-46f6-9d02-03ad2f0cbb0b" />

### Multiple Features
The library supports scatter plots, line charts, bar charts, area fills, multiple series, negative values, custom colors, and more. See the [tutorials](tutorials/) for detailed examples and usage.

---

**Made with ❤️ using Go and Fyne**
