# Tutorial 8: Color Palettes

## Overview

Master color selection and create professional color schemes.

**Time to complete:** 15 minutes

## Default Color Palette

FyneSimpleChart includes a 10-color professional palette based on industry standards (D3.js/Plotly):

```go
// These colors are assigned automatically
1. Blue:    RGB(31, 119, 180)
2. Orange:  RGB(255, 127, 14)
3. Green:   RGB(44, 160, 44)
4. Red:     RGB(214, 39, 40)
5. Purple:  RGB(148, 103, 189)
6. Brown:   RGB(140, 86, 75)
7. Pink:    RGB(227, 119, 194)
8. Gray:    RGB(127, 127, 127)
9. Olive:   RGB(188, 189, 34)
10. Cyan:   RGB(23, 190, 207)
```

## Creating Custom Palettes

### Corporate Palette
```go
import "image/color"

// Define company colors
corporateBlue := color.RGBA{R: 0, G: 71, B: 171, A: 255}
corporateGold := color.RGBA{R: 255, G: 184, B: 28, A: 255}
corporateGray := color.RGBA{R: 100, G: 100, B: 100, A: 255}

plot1 := fynesimplechart.NewPlot(data1, "Primary")
plot1.PlotColor = corporateBlue

plot2 := fynesimplechart.NewPlot(data2, "Secondary")
plot2.PlotColor = corporateGold
```

### Heatmap Palette (Cool to Warm)
```go
coldBlue := color.RGBA{R: 0, G: 100, B: 200, A: 255}
neutral := color.RGBA{R: 100, G: 100, B: 100, A: 255}
hotRed := color.RGBA{R: 200, G: 0, B: 0, A: 255}
```

### Colorblind-Friendly Palette
```go
// Optimized for color vision deficiency
blue := color.RGBA{R: 0, G: 114, B: 178, A: 255}
orange := color.RGBA{R: 230, G: 159, B: 0, A: 255}
skyBlue := color.RGBA{R: 86, G: 180, B: 233, A: 255}
bluishGreen := color.RGBA{R: 0, G: 158, B: 115, A: 255}
```

## Color Selection Guidelines

### 1. Contrast
Ensure colors are distinguishable:
```go
// Good contrast
blue := color.RGBA{R: 31, G: 119, B: 180, A: 255}
orange := color.RGBA{R: 255, G: 127, B: 14, A: 255}
```

### 2. Semantic Meaning
```go
// Financial data
profit := color.RGBA{R: 0, G: 200, B: 0, A: 255}      // Green
loss := color.RGBA{R: 200, G: 0, B: 0, A: 255}        // Red

// Temperature data
cold := color.RGBA{R: 0, G: 100, B: 255, A: 255}      // Blue
hot := color.RGBA{R: 255, G: 100, B: 0, A: 255}       // Red-Orange
```

### 3. Brand Consistency
Match your organization's colors:
```go
brandPrimary := color.RGBA{R: yourR, G: yourG, B: yourB, A: 255}
plot.PlotColor = brandPrimary
```

## Summary

- ✅ Professional default palette
- ✅ Easy custom colors
- ✅ Semantic color usage
- ✅ Accessibility considerations

## Next Steps

Continue to [Tutorial 9: Mathematical Visualizations](09-mathematical-functions.md).
