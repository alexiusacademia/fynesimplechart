package fynesimplechart

import (
	"image/color"
	"log"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ScatterPlot struct {
	widget.BaseWidget

	Plots []Plot

	mTop    float32
	mBottom float32
	mLeft   float32
	mRight  float32
}

// Constructor
func NewGraphWidget(plots []Plot) *ScatterPlot {
	w := &ScatterPlot{
		Plots:   plots,
		mTop:    30,
		mBottom: 30,
		mLeft:   60,
		mRight:  30,
	}
	w.ExtendBaseWidget(w)
	return w
}

// Generates a new renderer for the RatingCurveView.
func (v *ScatterPlot) CreateRenderer() fyne.WidgetRenderer {
	v.ExtendBaseWidget(v) // Ensure the base widget is extended properly

	return &scatterChartRenderer{widget: v}
}

// Responsible for rendering the RatingCurveView.
type scatterChartRenderer struct {
	widget  *ScatterPlot
	objects []fyne.CanvasObject
}

// Calculates the minimum size of the graph.
func (r *scatterChartRenderer) MinSize() fyne.Size {
	return r.widget.Size() // Provide a default value
}

// Layout the components.
func (r *scatterChartRenderer) Layout(size fyne.Size) {
	r.drawNodes()
	r.drawBorder()
}

// Called when the theme changes.
func (r *scatterChartRenderer) ApplyTheme() {
	// Update any theme dependent properties here
}

// Updates the widget's rendering.
func (r *scatterChartRenderer) Refresh() {
	r.drawNodes()
	r.drawBorder()
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
	// Perform necessary cleanup
}

func (r *scatterChartRenderer) getMargins() (float32, float32, float32, float32) {
	return r.widget.mTop, r.widget.mBottom, r.widget.mLeft, r.widget.mRight
}

// Draw the nodes of the graph.
func (r *scatterChartRenderer) drawNodes() {
	// Removed all objects to redraw
	r.objects = []fyne.CanvasObject{}

	mTop, mBottom, mLeft, mRight := r.getMargins()

	widgetSize := r.widget.Size()
	widgetWidth := widgetSize.Width

	// Get the minimums and maximums of coordinates
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

	for i := 0; i < len(r.widget.Plots); i++ {
		// Generate randome color
		red := rand.Intn(255)
		green := rand.Intn(255)
		blue := rand.Intn(255)

		plotColor := color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: 255}

		nodes := r.widget.Plots[i].Nodes

		chartUnScaledWidth := maxX - minX
		chartUnScaledHeight := maxY - minY
		if chartUnScaledHeight == 0 {
			log.Fatal("Nothing to plot in y axis")
			return
		}

		// Get the ratio of chart max width and max height plots
		plotAreaWidth := widgetWidth - mLeft - mRight
		plotAreaHeight := widgetSize.Height - mTop - mBottom

		scaleX := plotAreaWidth / (chartUnScaledWidth + 2)
		scaleY := plotAreaHeight / (chartUnScaledHeight + 2)

		// originPosition := fyne.NewPos(mLeft, widgetSize.Height-mBottom)

		// Start the plotting of nodes
		for j := 0; j < len(nodes); j++ {
			// Create a canvas circle
			c := canvas.NewCircle(plotColor)
			c.FillColor = plotColor
			c.StrokeColor = plotColor
			c.StrokeWidth = 1
			radius := 2

			x := (nodes[j].X+1)*scaleX + mLeft - float32(radius)                 // Always add 1 for the clearance
			y := mTop + plotAreaHeight - (nodes[j].Y+1)*scaleY - float32(radius) // Always add 1 for the clearance

			c.Resize(fyne.NewSize(float32(radius)*2, float32(radius)*2))
			c.Move(fyne.NewPos(x, y))
			r.objects = append(r.objects, c)
		}

		// Connect nodes
		if r.widget.Plots[i].ShowLine {
			for k := 0; k < (len(nodes) - 1); k++ {
				x1 := (nodes[k].X+1)*scaleX + mLeft
				y1 := mTop + plotAreaHeight - (nodes[k].Y+1)*scaleY
				x2 := (nodes[k+1].X+1)*scaleX + mLeft
				y2 := mTop + plotAreaHeight - (nodes[k+1].Y+1)*scaleY

				l := canvas.NewLine(plotColor)
				l.StrokeWidth = 1.5
				l.StrokeColor = plotColor
				l.Position1 = fyne.NewPos(x1, y1)
				l.Position2 = fyne.NewPos(x2, y2)
				r.objects = append(r.objects, l)
			}
		}

		// Draw axes
		xAxis := canvas.NewLine(theme.ForegroundColor())
		xAxis.StrokeWidth = 0.5
		xAxis.StrokeColor = theme.ForegroundColor()
		xAxis.Position1.X = mLeft
		xAxis.Position1.Y = mTop + plotAreaHeight - scaleY
		xAxis.Position2.X = mLeft + plotAreaWidth
		xAxis.Position2.Y = xAxis.Position1.Y
		r.objects = append(r.objects, xAxis)

		yAxis := canvas.NewLine(theme.ForegroundColor())
		yAxis.StrokeWidth = 0.5
		yAxis.StrokeColor = theme.ForegroundColor()
		yAxis.Position1.X = mLeft + scaleX
		yAxis.Position1.Y = mTop
		yAxis.Position2.X = mLeft + scaleX
		yAxis.Position2.Y = mTop + plotAreaHeight
		r.objects = append(r.objects, yAxis)

		xAxisArrow := canvas.NewText(">", theme.ForegroundColor())
		xAxisArrow.TextSize = 18
		xAxisArrow.Move(fyne.NewPos(mLeft+plotAreaWidth-xAxisArrow.MinSize().Width,
			mTop+plotAreaHeight-xAxisArrow.MinSize().Height/2-scaleY,
		))
		r.objects = append(r.objects, xAxisArrow)

		yAxisArrow := canvas.NewText("^", theme.ForegroundColor())
		yAxisArrow.TextSize = 18
		yAxisArrow.Move(fyne.NewPos(mLeft+scaleX-yAxisArrow.MinSize().Width/2, mTop))
		r.objects = append(r.objects, yAxisArrow)

		axisDirectionTextY := canvas.NewText("Y", theme.ForegroundColor())
		axisDirectionTextY.TextSize = 16
		axisDirectionTextY.Move(fyne.NewPos(
			mLeft+scaleX-axisDirectionTextY.MinSize().Width,
			mTop+mTop/2))
		r.objects = append(r.objects, axisDirectionTextY)

		axisDirectionTextX := canvas.NewText("X", theme.ForegroundColor())
		axisDirectionTextX.TextSize = 16
		axisDirectionTextX.Move(fyne.NewPos(
			mLeft+plotAreaWidth-mRight/2,
			mTop+plotAreaHeight-scaleY,
		))
		r.objects = append(r.objects, axisDirectionTextX)

		// Display the axes titles
		// xAxisTitle := canvas.NewText(r.widget.Plots[i].XAxisTitle, theme.ForegroundColor())
		// xAxisTitle.TextStyle.Bold = true
		// xAxisTitle.TextSize = 14
		// xAxisTitle.Move(
		// 	fyne.NewPos(
		// 		mLeft,
		// 		r.widget.Size().Height-mBottom/2-xAxisTitle.MinSize().Height/2,
		// 	))
		// r.objects = append(r.objects, xAxisTitle)
	}

}

func (r *scatterChartRenderer) drawBorder() {
	size := r.widget.Size()
	width := size.Width
	height := size.Height
	width = width - r.widget.mLeft - r.widget.mRight
	height = height - r.widget.mTop - r.widget.mBottom

	border := canvas.NewRectangle(theme.ForegroundColor())
	border.Resize(fyne.NewSize(width, height))
	border.Move(fyne.NewPos(r.widget.mLeft, r.widget.mTop))
	border.FillColor = nil
	border.StrokeColor = theme.ForegroundColor()
	border.StrokeWidth = 2
	r.objects = append(r.objects, border)
}
