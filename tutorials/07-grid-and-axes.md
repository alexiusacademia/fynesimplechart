# Tutorial 7: Grid and Axes Configuration

## Overview

Learn about grid lines, axis labels, and chart layout.

**Time to complete:** 10 minutes

## Grid Configuration

### Show Grid (Default)
```go
chart := fynesimplechart.NewGraphWidget(plots)
chart.ShowGrid = true  // Shows grid lines
```

### Hide Grid
```go
chart.ShowGrid = false  // Clean, minimal look
```

## Understanding Grid Lines

The grid automatically:
- Calculates optimal tick spacing
- Aligns with axis labels
- Uses professional "nice numbers" (1, 2, 5, 10 multiples)
- Adjusts based on data range

## Axis System

The chart automatically provides:
- **Numeric labels** on both axes
- **Tick marks** at regular intervals
- **Zero-line positioning** for negative values
- **Arrow indicators** at axis ends
- **X/Y markers** for axis identification

## When to Show/Hide Grid

### Show Grid For:
- Data analysis and reading values
- Scientific presentations
- Financial reports
- Technical documentation

### Hide Grid For:
- Marketing presentations
- Minimal aesthetic
- When exact values aren't important
- Print materials (ink saving)

## Example Comparison

```go
// Analytical style
chart1 := fynesimplechart.NewGraphWidget(plots)
chart1.ShowGrid = true
chart1.SetChartTitle("Detailed Analysis")

// Presentation style
chart2 := fynesimplechart.NewGraphWidget(plots)
chart2.ShowGrid = false
chart2.SetChartTitle("Executive Summary")
```

## Summary

- ✅ Grid is configurable (on/off)
- ✅ Axes are automatic and smart
- ✅ Labels adjust to data range
- ✅ Choose style based on use case

## Next Steps

Continue to [Tutorial 8: Color Palettes](08-color-palettes.md).
