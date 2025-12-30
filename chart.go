package fynesimplechart

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	defaultMarginTop    float32 = 50
	defaultMarginBottom float32 = 60
	defaultMarginLeft   float32 = 80
	defaultMarginRight  float32 = 100
	gridLineWidth       float32 = 0.5
	axisLineWidth       float32 = 1.5
	tickLength          float32 = 5
	minTickSpacing      float32 = 40 // Minimum pixels between ticks
)

type ScatterPlot struct {
	widget.BaseWidget

	Plots      []Plot
	ChartTitle string
	ShowGrid   bool

	mTop    float32
	mBottom float32
	mLeft   float32
	mRight  float32
}

// Constructor
func NewGraphWidget(plots []Plot) *ScatterPlot {
	w := &ScatterPlot{
		Plots:      plots,
		ChartTitle: "",
		ShowGrid:   true,
		mTop:       defaultMarginTop,
		mBottom:    defaultMarginBottom,
		mLeft:      defaultMarginLeft,
		mRight:     defaultMarginRight,
	}
	w.ExtendBaseWidget(w)
	return w
}

// SetChartTitle sets the main title for the chart
func (v *ScatterPlot) SetChartTitle(title string) {
	v.ChartTitle = title
	v.Refresh()
}

// Generates a new renderer for the ScatterPlot.
func (v *ScatterPlot) CreateRenderer() fyne.WidgetRenderer {
	v.ExtendBaseWidget(v)
	return &scatterChartRenderer{widget: v}
}

// Responsible for rendering the ScatterPlot.
type scatterChartRenderer struct {
	widget  *ScatterPlot
	objects []fyne.CanvasObject
}

// Calculates the minimum size of the graph.
func (r *scatterChartRenderer) MinSize() fyne.Size {
	return r.widget.Size()
}

// Layout the components.
func (r *scatterChartRenderer) Layout(size fyne.Size) {
	r.render()
}

// Called when the theme changes.
func (r *scatterChartRenderer) ApplyTheme() {
	r.render()
}

// Updates the widget's rendering.
func (r *scatterChartRenderer) Refresh() {
	r.render()
	canvas.Refresh(r.widget)
}

// Returns the background color of the widget.
func (r *scatterChartRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

// Return the objects contained in the widget.
func (r *scatterChartRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

// Called when the widget is destroyed.
func (r *scatterChartRenderer) Destroy() {
}

// Main render function
func (r *scatterChartRenderer) render() {
	r.objects = []fyne.CanvasObject{}

	if len(r.widget.Plots) == 0 {
		return
	}

	mTop, mBottom, mLeft, mRight := r.widget.mTop, r.widget.mBottom, r.widget.mLeft, r.widget.mRight
	widgetSize := r.widget.Size()

	// Get data bounds
	maxX, err := MaxX(r.widget.Plots)
	if err != nil {
		return
	}
	minX, err := MinX(r.widget.Plots)
	if err != nil {
		return
	}
	maxY, err := MaxY(r.widget.Plots)
	if err != nil {
		return
	}
	minY, err := MinY(r.widget.Plots)
	if err != nil {
		return
	}

	// Add 10% padding to the data range
	rangeX := maxX - minX
	rangeY := maxY - minY

	if rangeX == 0 {
		rangeX = 1
		minX -= 0.5
		maxX += 0.5
	} else {
		padding := rangeX * 0.1
		minX -= padding
		maxX += padding
	}

	if rangeY == 0 {
		rangeY = 1
		minY -= 0.5
		maxY += 0.5
	} else {
		padding := rangeY * 0.1
		minY -= padding
		maxY += padding
	}

	plotAreaWidth := widgetSize.Width - mLeft - mRight
	plotAreaHeight := widgetSize.Height - mTop - mBottom

	// Draw chart title if present
	if r.widget.ChartTitle != "" {
		titleText := canvas.NewText(r.widget.ChartTitle, theme.ForegroundColor())
		titleText.TextSize = 16
		titleText.TextStyle.Bold = true
		titleText.Alignment = fyne.TextAlignCenter
		titleWidth := titleText.MinSize().Width
		titleText.Move(fyne.NewPos((widgetSize.Width-titleWidth)/2, 10))
		r.objects = append(r.objects, titleText)
	}

	// Draw grid and axes
	if r.widget.ShowGrid {
		r.drawGrid(minX, maxX, minY, maxY, plotAreaWidth, plotAreaHeight, mLeft, mTop)
	}

	r.drawAxes(minX, maxX, minY, maxY, plotAreaWidth, plotAreaHeight, mLeft, mTop, mBottom)

	// Generate colors for plots
	colors := r.generateColors(len(r.widget.Plots))

	// Draw area fills first (so they appear behind everything)
	for i, plot := range r.widget.Plots {
		plotColor := colors[i]
		if plot.PlotColor != nil {
			plotColor = plot.PlotColor
		}

		if plot.FillArea {
			r.drawAreaFill(i, plot, plotColor, minX, maxX, minY, maxY, plotAreaWidth, plotAreaHeight, mLeft, mTop)
		}
	}

	// Draw each plot (lines and points on top of fills)
	for i, plot := range r.widget.Plots {
		plotColor := colors[i]
		if plot.PlotColor != nil {
			plotColor = plot.PlotColor
		}

		r.drawPlot(plot, plotColor, minX, maxX, minY, maxY, plotAreaWidth, plotAreaHeight, mLeft, mTop)
	}

	// Draw legend
	r.drawLegend(colors, widgetSize.Width-mRight, mTop)

	// Draw border
	r.drawBorder(plotAreaWidth, plotAreaHeight, mLeft, mTop)
}

// Draw a single plot
func (r *scatterChartRenderer) drawPlot(plot Plot, plotColor color.Color, minX, maxX, minY, maxY, plotWidth, plotHeight, mLeft, mTop float32) {
	nodes := plot.Nodes
	if len(nodes) == 0 {
		return
	}

	rangeX := maxX - minX
	rangeY := maxY - minY

	// Transform function from data coordinates to screen coordinates
	dataToScreenX := func(x float32) float32 {
		return mLeft + ((x-minX)/rangeX)*plotWidth
	}
	dataToScreenY := func(y float32) float32 {
		return mTop + plotHeight - ((y-minY)/rangeY)*plotHeight
	}

	// Draw bars first (so they appear behind lines and points)
	if plot.ShowBars {
		r.drawBars(plot, plotColor, minX, maxX, minY, maxY, plotWidth, plotHeight, mLeft, mTop, dataToScreenX, dataToScreenY)
	}

	// Draw lines (so they appear behind points)
	if plot.ShowLine && len(nodes) > 1 {
		for k := 0; k < len(nodes)-1; k++ {
			x1 := dataToScreenX(nodes[k].X)
			y1 := dataToScreenY(nodes[k].Y)
			x2 := dataToScreenX(nodes[k+1].X)
			y2 := dataToScreenY(nodes[k+1].Y)

			line := canvas.NewLine(plotColor)
			line.StrokeWidth = plot.LineWidth
			line.Position1 = fyne.NewPos(x1, y1)
			line.Position2 = fyne.NewPos(x2, y2)
			r.objects = append(r.objects, line)
		}
	}

	// Draw points
	if plot.ShowPoints {
		for j := 0; j < len(nodes); j++ {
			x := dataToScreenX(nodes[j].X)
			y := dataToScreenY(nodes[j].Y)

			circle := canvas.NewCircle(plotColor)
			circle.FillColor = plotColor
			circle.StrokeColor = plotColor
			circle.StrokeWidth = 1

			radius := plot.PointSize
			circle.Resize(fyne.NewSize(radius*2, radius*2))
			circle.Move(fyne.NewPos(x-radius, y-radius))
			r.objects = append(r.objects, circle)
		}
	}
}

// Draw bars for a bar chart
func (r *scatterChartRenderer) drawBars(plot Plot, plotColor color.Color, minX, maxX, minY, maxY, plotWidth, plotHeight, mLeft, mTop float32, dataToScreenX, dataToScreenY func(float32) float32) {
	nodes := plot.Nodes
	if len(nodes) == 0 {
		return
	}

	// Calculate bar width
	// If BarWidth is 0, use auto-sizing based on data spacing
	barWidthData := plot.BarWidth
	if barWidthData == 0 {
		barWidthData = 0.8
	}

	// Calculate the spacing between data points
	var spacing float32
	if len(nodes) > 1 {
		// Average spacing between consecutive X values
		totalSpacing := float32(0)
		for i := 0; i < len(nodes)-1; i++ {
			totalSpacing += nodes[i+1].X - nodes[i].X
		}
		spacing = totalSpacing / float32(len(nodes)-1)
	} else {
		// Single bar - use a reasonable default
		spacing = (maxX - minX) / 10
		if spacing == 0 {
			spacing = 1
		}
	}

	// Convert spacing to screen coordinates
	barWidthScreen := spacing * barWidthData * (plotWidth / (maxX - minX))

	// Get zero line position
	zeroY := dataToScreenY(0)
	// Clamp to plot area
	if zeroY < mTop {
		zeroY = mTop
	}
	if zeroY > mTop+plotHeight {
		zeroY = mTop + plotHeight
	}

	// Draw each bar
	for _, node := range nodes {
		centerX := dataToScreenX(node.X)
		topY := dataToScreenY(node.Y)

		// Calculate bar position and size
		barX := centerX - barWidthScreen/2
		barY := topY
		barHeight := zeroY - topY

		// Handle negative values
		if barHeight < 0 {
			barY = zeroY
			barHeight = -barHeight
		}

		// Skip bars with zero or negligible height
		if barHeight < 0.5 {
			continue
		}

		// Create the bar rectangle
		bar := canvas.NewRectangle(plotColor)
		bar.Move(fyne.NewPos(barX, barY))
		bar.Resize(fyne.NewSize(barWidthScreen, barHeight))
		r.objects = append(r.objects, bar)

		// Draw border if specified
		if plot.BarBorderWidth > 0 {
			borderColor := plot.BarBorderColor
			if borderColor == nil {
				// Default to darker version of bar color
				if rgba, ok := plotColor.(color.RGBA); ok {
					borderColor = color.RGBA{
						R: uint8(float32(rgba.R) * 0.7),
						G: uint8(float32(rgba.G) * 0.7),
						B: uint8(float32(rgba.B) * 0.7),
						A: rgba.A,
					}
				} else {
					borderColor = plotColor
				}
			}

			// Draw four border lines
			// Top
			topLine := canvas.NewLine(borderColor)
			topLine.StrokeWidth = plot.BarBorderWidth
			topLine.Position1 = fyne.NewPos(barX, barY)
			topLine.Position2 = fyne.NewPos(barX+barWidthScreen, barY)
			r.objects = append(r.objects, topLine)

			// Bottom
			bottomLine := canvas.NewLine(borderColor)
			bottomLine.StrokeWidth = plot.BarBorderWidth
			bottomLine.Position1 = fyne.NewPos(barX, barY+barHeight)
			bottomLine.Position2 = fyne.NewPos(barX+barWidthScreen, barY+barHeight)
			r.objects = append(r.objects, bottomLine)

			// Left
			leftLine := canvas.NewLine(borderColor)
			leftLine.StrokeWidth = plot.BarBorderWidth
			leftLine.Position1 = fyne.NewPos(barX, barY)
			leftLine.Position2 = fyne.NewPos(barX, barY+barHeight)
			r.objects = append(r.objects, leftLine)

			// Right
			rightLine := canvas.NewLine(borderColor)
			rightLine.StrokeWidth = plot.BarBorderWidth
			rightLine.Position1 = fyne.NewPos(barX+barWidthScreen, barY)
			rightLine.Position2 = fyne.NewPos(barX+barWidthScreen, barY+barHeight)
			r.objects = append(r.objects, rightLine)
		}
	}
}

// Draw area fill for a plot using smooth polygon rendering
func (r *scatterChartRenderer) drawAreaFill(plotIdx int, plot Plot, plotColor color.Color, minX, maxX, minY, maxY, plotWidth, plotHeight, mLeft, mTop float32) {
	nodes := plot.Nodes
	if len(nodes) < 2 {
		return
	}

	rangeX := maxX - minX
	rangeY := maxY - minY

	// Transform function from data coordinates to screen coordinates
	dataToScreenX := func(x float32) float32 {
		return mLeft + ((x-minX)/rangeX)*plotWidth
	}
	dataToScreenY := func(y float32) float32 {
		return mTop + plotHeight - ((y-minY)/rangeY)*plotHeight
	}

	// Determine fill color (use custom or derive from plot color with transparency)
	fillColor := plot.FillColor
	if fillColor == nil {
		// Use plot color with 30% opacity
		if rgba, ok := plotColor.(color.RGBA); ok {
			fillColor = color.RGBA{R: rgba.R, G: rgba.G, B: rgba.B, A: 76} // 76 = 30% of 255
		} else {
			r, g, b, _ := plotColor.RGBA()
			fillColor = color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: 76}
		}
	}

	// Fill to zero (Y-axis)
	if plot.FillToZero {
		zeroY := dataToScreenY(0)
		// Clamp to plot area
		if zeroY < mTop {
			zeroY = mTop
		}
		if zeroY > mTop+plotHeight {
			zeroY = mTop + plotHeight
		}

		// Draw vertical rectangles from curve to zero line
		// Sample at many points for smoothness
		steps := 500
		for step := 0; step < steps; step++ {
			t := float32(step) / float32(steps-1)
			dataX := nodes[0].X + t*(nodes[len(nodes)-1].X-nodes[0].X)
			dataY := interpolateY(nodes, dataX)

			screenX := dataToScreenX(dataX)
			screenY := dataToScreenY(dataY)

			// Draw thin vertical rectangle from curve to zero
			rectY1 := screenY
			rectY2 := zeroY
			if rectY1 > rectY2 {
				rectY1, rectY2 = rectY2, rectY1
			}

			rectHeight := rectY2 - rectY1
			if rectHeight > 0 {
				rect := canvas.NewRectangle(fillColor)
				rect.Move(fyne.NewPos(screenX, rectY1))
				rect.Resize(fyne.NewSize(plotWidth/float32(steps)+0.5, rectHeight))
				r.objects = append(r.objects, rect)
			}
		}
	}

	// Fill between two plots
	if plot.FillToPlotIdx >= 0 && plot.FillToPlotIdx < len(r.widget.Plots) {
		otherPlot := r.widget.Plots[plot.FillToPlotIdx]
		otherNodes := otherPlot.Nodes

		if len(otherNodes) < 2 {
			return
		}

		// Find common X range
		minCommonX := math.Max(float64(nodes[0].X), float64(otherNodes[0].X))
		maxCommonX := math.Min(float64(nodes[len(nodes)-1].X), float64(otherNodes[len(otherNodes)-1].X))

		if minCommonX >= maxCommonX {
			return
		}

		// Draw vertical rectangles between the two curves
		// Sample at many points for smoothness
		steps := 500
		for step := 0; step < steps; step++ {
			t := float32(step) / float32(steps-1)
			dataX := float32(minCommonX) + t*float32(maxCommonX-minCommonX)

			// Get Y values for both curves at this X
			dataY1 := interpolateY(nodes, dataX)
			dataY2 := interpolateY(otherNodes, dataX)

			// Convert to screen coordinates
			screenX := dataToScreenX(dataX)
			screenY1 := dataToScreenY(dataY1)
			screenY2 := dataToScreenY(dataY2)

			// Draw thin vertical rectangle between the two curves
			rectY1 := screenY1
			rectY2 := screenY2
			if rectY1 > rectY2 {
				rectY1, rectY2 = rectY2, rectY1
			}

			rectHeight := rectY2 - rectY1
			if rectHeight > 0 {
				rect := canvas.NewRectangle(fillColor)
				rect.Move(fyne.NewPos(screenX, rectY1))
				rect.Resize(fyne.NewSize(plotWidth/float32(steps)+0.5, rectHeight))
				r.objects = append(r.objects, rect)
			}
		}
	}
}

// Interpolate Y value for a given X in a set of nodes
func interpolateY(nodes []Node, x float32) float32 {
	if len(nodes) == 0 {
		return 0
	}

	// If x is before first node, return first Y
	if x <= nodes[0].X {
		return nodes[0].Y
	}

	// If x is after last node, return last Y
	if x >= nodes[len(nodes)-1].X {
		return nodes[len(nodes)-1].Y
	}

	// Find the two nodes to interpolate between
	for i := 0; i < len(nodes)-1; i++ {
		if x >= nodes[i].X && x <= nodes[i+1].X {
			// Linear interpolation
			t := (x - nodes[i].X) / (nodes[i+1].X - nodes[i].X)
			return nodes[i].Y + t*(nodes[i+1].Y-nodes[i].Y)
		}
	}

	return nodes[len(nodes)-1].Y
}

// Draw grid lines
func (r *scatterChartRenderer) drawGrid(minX, maxX, minY, maxY, plotWidth, plotHeight, mLeft, mTop float32) {
	gridColor := color.RGBA{R: 128, G: 128, B: 128, A: 50}

	rangeX := maxX - minX
	rangeY := maxY - minY

	// Calculate nice tick intervals
	numXTicks := int(plotWidth / minTickSpacing)
	numYTicks := int(plotHeight / minTickSpacing)

	if numXTicks < 2 {
		numXTicks = 2
	}
	if numYTicks < 2 {
		numYTicks = 2
	}

	xTickInterval := calculateNiceInterval(rangeX, numXTicks)
	yTickInterval := calculateNiceInterval(rangeY, numYTicks)

	// Draw vertical grid lines
	xStart := math.Ceil(float64(minX/xTickInterval)) * float64(xTickInterval)
	for x := float32(xStart); x <= maxX; x += xTickInterval {
		screenX := mLeft + ((x-minX)/rangeX)*plotWidth
		line := canvas.NewLine(gridColor)
		line.StrokeWidth = gridLineWidth
		line.Position1 = fyne.NewPos(screenX, mTop)
		line.Position2 = fyne.NewPos(screenX, mTop+plotHeight)
		r.objects = append(r.objects, line)
	}

	// Draw horizontal grid lines
	yStart := math.Ceil(float64(minY/yTickInterval)) * float64(yTickInterval)
	for y := float32(yStart); y <= maxY; y += yTickInterval {
		screenY := mTop + plotHeight - ((y-minY)/rangeY)*plotHeight
		line := canvas.NewLine(gridColor)
		line.StrokeWidth = gridLineWidth
		line.Position1 = fyne.NewPos(mLeft, screenY)
		line.Position2 = fyne.NewPos(mLeft+plotWidth, screenY)
		r.objects = append(r.objects, line)
	}
}

// Draw axes with labels
func (r *scatterChartRenderer) drawAxes(minX, maxX, minY, maxY, plotWidth, plotHeight, mLeft, mTop, mBottom float32) {
	foregroundColor := theme.ForegroundColor()
	rangeX := maxX - minX
	rangeY := maxY - minY

	// Calculate tick intervals
	numXTicks := int(plotWidth / minTickSpacing)
	numYTicks := int(plotHeight / minTickSpacing)

	if numXTicks < 2 {
		numXTicks = 2
	}
	if numYTicks < 2 {
		numYTicks = 2
	}

	xTickInterval := calculateNiceInterval(rangeX, numXTicks)
	yTickInterval := calculateNiceInterval(rangeY, numYTicks)

	// Draw X axis
	xAxisY := mTop + plotHeight
	if minY < 0 && maxY > 0 {
		// Zero line is visible
		xAxisY = mTop + plotHeight - ((-minY)/rangeY)*plotHeight
	}

	xAxis := canvas.NewLine(foregroundColor)
	xAxis.StrokeWidth = axisLineWidth
	xAxis.Position1 = fyne.NewPos(mLeft, xAxisY)
	xAxis.Position2 = fyne.NewPos(mLeft+plotWidth, xAxisY)
	r.objects = append(r.objects, xAxis)

	// Draw X axis ticks and labels
	xStart := math.Ceil(float64(minX/xTickInterval)) * float64(xTickInterval)
	for x := float32(xStart); x <= maxX; x += xTickInterval {
		screenX := mLeft + ((x-minX)/rangeX)*plotWidth

		// Tick mark
		tick := canvas.NewLine(foregroundColor)
		tick.StrokeWidth = axisLineWidth
		tick.Position1 = fyne.NewPos(screenX, xAxisY)
		tick.Position2 = fyne.NewPos(screenX, xAxisY+tickLength)
		r.objects = append(r.objects, tick)

		// Label
		labelText := formatAxisLabel(x)
		label := canvas.NewText(labelText, foregroundColor)
		label.TextSize = 10
		labelWidth := label.MinSize().Width
		label.Move(fyne.NewPos(screenX-labelWidth/2, xAxisY+tickLength+2))
		r.objects = append(r.objects, label)
	}

	// Draw Y axis
	yAxisX := mLeft
	if minX < 0 && maxX > 0 {
		// Zero line is visible
		yAxisX = mLeft + ((-minX)/rangeX)*plotWidth
	}

	yAxis := canvas.NewLine(foregroundColor)
	yAxis.StrokeWidth = axisLineWidth
	yAxis.Position1 = fyne.NewPos(yAxisX, mTop)
	yAxis.Position2 = fyne.NewPos(yAxisX, mTop+plotHeight)
	r.objects = append(r.objects, yAxis)

	// Draw Y axis ticks and labels
	yStart := math.Ceil(float64(minY/yTickInterval)) * float64(yTickInterval)
	for y := float32(yStart); y <= maxY; y += yTickInterval {
		screenY := mTop + plotHeight - ((y-minY)/rangeY)*plotHeight

		// Tick mark
		tick := canvas.NewLine(foregroundColor)
		tick.StrokeWidth = axisLineWidth
		tick.Position1 = fyne.NewPos(yAxisX-tickLength, screenY)
		tick.Position2 = fyne.NewPos(yAxisX, screenY)
		r.objects = append(r.objects, tick)

		// Label
		labelText := formatAxisLabel(y)
		label := canvas.NewText(labelText, foregroundColor)
		label.TextSize = 10
		labelWidth := label.MinSize().Width
		labelHeight := label.MinSize().Height
		label.Move(fyne.NewPos(yAxisX-tickLength-labelWidth-5, screenY-labelHeight/2))
		r.objects = append(r.objects, label)
	}

	// Draw axis arrows
	arrowSize := float32(8)

	// X axis arrow
	xArrowY := xAxisY
	xArrowTip := mLeft + plotWidth
	xArrow1 := canvas.NewLine(foregroundColor)
	xArrow1.StrokeWidth = axisLineWidth
	xArrow1.Position1 = fyne.NewPos(xArrowTip, xArrowY)
	xArrow1.Position2 = fyne.NewPos(xArrowTip-arrowSize, xArrowY-arrowSize/2)
	r.objects = append(r.objects, xArrow1)

	xArrow2 := canvas.NewLine(foregroundColor)
	xArrow2.StrokeWidth = axisLineWidth
	xArrow2.Position1 = fyne.NewPos(xArrowTip, xArrowY)
	xArrow2.Position2 = fyne.NewPos(xArrowTip-arrowSize, xArrowY+arrowSize/2)
	r.objects = append(r.objects, xArrow2)

	// Y axis arrow
	yArrowX := yAxisX
	yArrowTip := mTop
	yArrow1 := canvas.NewLine(foregroundColor)
	yArrow1.StrokeWidth = axisLineWidth
	yArrow1.Position1 = fyne.NewPos(yArrowX, yArrowTip)
	yArrow1.Position2 = fyne.NewPos(yArrowX-arrowSize/2, yArrowTip+arrowSize)
	r.objects = append(r.objects, yArrow1)

	yArrow2 := canvas.NewLine(foregroundColor)
	yArrow2.StrokeWidth = axisLineWidth
	yArrow2.Position1 = fyne.NewPos(yArrowX, yArrowTip)
	yArrow2.Position2 = fyne.NewPos(yArrowX+arrowSize/2, yArrowTip+arrowSize)
	r.objects = append(r.objects, yArrow2)

	// Axis labels (X and Y markers)
	xLabel := canvas.NewText("X", foregroundColor)
	xLabel.TextSize = 14
	xLabel.TextStyle.Bold = true
	xLabel.Move(fyne.NewPos(xArrowTip+5, xArrowY-xLabel.MinSize().Height/2))
	r.objects = append(r.objects, xLabel)

	yLabel := canvas.NewText("Y", foregroundColor)
	yLabel.TextSize = 14
	yLabel.TextStyle.Bold = true
	yLabel.Move(fyne.NewPos(yArrowX-yLabel.MinSize().Width/2, yArrowTip-yLabel.MinSize().Height-5))
	r.objects = append(r.objects, yLabel)
}

// Draw legend
func (r *scatterChartRenderer) drawLegend(colors []color.Color, x, y float32) {
	if len(r.widget.Plots) == 0 {
		return
	}

	foregroundColor := theme.ForegroundColor()

	legendTitle := canvas.NewText("LEGEND", foregroundColor)
	legendTitle.TextSize = 11
	legendTitle.TextStyle.Bold = true
	legendTitle.Move(fyne.NewPos(x+5, y))
	r.objects = append(r.objects, legendTitle)

	currentY := y + 18

	for i, plot := range r.widget.Plots {
		plotColor := colors[i]
		if plot.PlotColor != nil {
			plotColor = plot.PlotColor
		}

		// Draw indicator based on plot style
		if plot.ShowLine && !plot.ShowPoints {
			// Line only - draw a short line
			line := canvas.NewLine(plotColor)
			line.StrokeWidth = plot.LineWidth
			line.Position1 = fyne.NewPos(x+10, currentY+5)
			line.Position2 = fyne.NewPos(x+25, currentY+5)
			r.objects = append(r.objects, line)
		} else if plot.ShowPoints && !plot.ShowLine {
			// Points only - draw a circle
			circle := canvas.NewCircle(plotColor)
			circle.FillColor = plotColor
			circle.StrokeColor = plotColor
			circle.Resize(fyne.NewSize(6, 6))
			circle.Move(fyne.NewPos(x+14, currentY+2))
			r.objects = append(r.objects, circle)
		} else {
			// Both - draw line with circle
			line := canvas.NewLine(plotColor)
			line.StrokeWidth = plot.LineWidth
			line.Position1 = fyne.NewPos(x+10, currentY+5)
			line.Position2 = fyne.NewPos(x+25, currentY+5)
			r.objects = append(r.objects, line)

			circle := canvas.NewCircle(plotColor)
			circle.FillColor = plotColor
			circle.StrokeColor = plotColor
			circle.Resize(fyne.NewSize(6, 6))
			circle.Move(fyne.NewPos(x+14, currentY+2))
			r.objects = append(r.objects, circle)
		}

		// Label
		label := canvas.NewText(plot.Title, foregroundColor)
		label.TextSize = 10
		label.Move(fyne.NewPos(x+30, currentY))
		r.objects = append(r.objects, label)

		currentY += 20
	}
}

// Draw border around plot area
func (r *scatterChartRenderer) drawBorder(width, height, x, y float32) {
	border := canvas.NewRectangle(color.Transparent)
	border.Resize(fyne.NewSize(width, height))
	border.Move(fyne.NewPos(x, y))
	border.StrokeColor = theme.ForegroundColor()
	border.StrokeWidth = 1.5
	r.objects = append(r.objects, border)
}

// Generate colors using a better color palette
func (r *scatterChartRenderer) generateColors(count int) []color.Color {
	// Professional color palette with good contrast
	predefinedColors := []color.Color{
		color.RGBA{R: 31, G: 119, B: 180, A: 255},   // Blue
		color.RGBA{R: 255, G: 127, B: 14, A: 255},   // Orange
		color.RGBA{R: 44, G: 160, B: 44, A: 255},    // Green
		color.RGBA{R: 214, G: 39, B: 40, A: 255},    // Red
		color.RGBA{R: 148, G: 103, B: 189, A: 255},  // Purple
		color.RGBA{R: 140, G: 86, B: 75, A: 255},    // Brown
		color.RGBA{R: 227, G: 119, B: 194, A: 255},  // Pink
		color.RGBA{R: 127, G: 127, B: 127, A: 255},  // Gray
		color.RGBA{R: 188, G: 189, B: 34, A: 255},   // Olive
		color.RGBA{R: 23, G: 190, B: 207, A: 255},   // Cyan
	}

	colors := make([]color.Color, count)
	for i := 0; i < count; i++ {
		colors[i] = predefinedColors[i%len(predefinedColors)]
	}
	return colors
}

// Calculate nice interval for axis ticks
func calculateNiceInterval(dataRange float32, numTicks int) float32 {
	if dataRange == 0 {
		return 1
	}

	roughInterval := dataRange / float32(numTicks)

	// Find the magnitude of the interval
	magnitude := math.Pow(10, math.Floor(math.Log10(float64(roughInterval))))

	// Normalize the interval
	normalized := float64(roughInterval) / magnitude

	var niceInterval float64
	if normalized <= 1 {
		niceInterval = 1
	} else if normalized <= 2 {
		niceInterval = 2
	} else if normalized <= 5 {
		niceInterval = 5
	} else {
		niceInterval = 10
	}

	return float32(niceInterval * magnitude)
}

// Format axis label with appropriate precision
func formatAxisLabel(value float32) string {
	absValue := float32(math.Abs(float64(value)))

	if absValue == 0 {
		return "0"
	} else if absValue >= 1000 {
		return fmt.Sprintf("%.0f", value)
	} else if absValue >= 100 {
		return fmt.Sprintf("%.1f", value)
	} else if absValue >= 10 {
		return fmt.Sprintf("%.1f", value)
	} else if absValue >= 1 {
		return fmt.Sprintf("%.2f", value)
	} else {
		return fmt.Sprintf("%.3f", value)
	}
}
