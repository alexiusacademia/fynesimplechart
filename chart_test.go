package fynesimplechart

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
)

func renderChart(t *testing.T, chart *ScatterPlot) []fyne.CanvasObject {
	t.Helper()
	chart.Resize(fyne.NewSize(600, 400))
	r := test.WidgetRenderer(chart)
	r.Refresh()
	return r.Objects()
}

func countCircles(objs []fyne.CanvasObject) int {
	n := 0
	for _, o := range objs {
		if _, ok := o.(*canvas.Circle); ok {
			n++
		}
	}
	return n
}

// Issue #2: updating a plot's data after creating the chart must redraw it.
func TestSetPlotNodesUpdatesChart(t *testing.T) {
	test.NewApp()

	plot := NewPlot([]Node{{X: 0, Y: 1}}, "Live")
	chart := NewGraphWidget([]Plot{*plot})
	chart.ShowLegend = false

	if got := countCircles(renderChart(t, chart)); got != 1 {
		t.Fatalf("initial render: got %d points, want 1", got)
	}

	chart.SetPlotNodes(0, []Node{{X: 0, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 3}})
	if got := countCircles(renderChart(t, chart)); got != 3 {
		t.Fatalf("after SetPlotNodes: got %d points, want 3", got)
	}

	// Out-of-range indexes are ignored rather than panicking.
	chart.SetPlotNodes(5, nil)
	chart.SetPlotNodes(-1, nil)
}

// Issue #1: with negative data the X axis must sit at Y=0, not at the bottom
// of the plot area, and points must stay inside the plot area.
func TestNegativeValuesAxisAtZero(t *testing.T) {
	test.NewApp()

	plot := NewPlot([]Node{{X: -5, Y: -10}, {X: 5, Y: 10}}, "Signed")
	chart := NewGraphWidget([]Plot{*plot})
	objs := renderChart(t, chart)

	plotHeight := float32(400) - defaultMarginTop - defaultMarginBottom
	wantAxisY := defaultMarginTop + plotHeight/2 // data range is symmetric around 0

	found := false
	for _, o := range objs {
		l, ok := o.(*canvas.Line)
		if !ok || l.StrokeWidth != axisLineWidth {
			continue
		}
		if l.Position1.Y == l.Position2.Y && l.Position1.X == defaultMarginLeft &&
			abs32(l.Position1.Y-wantAxisY) < 0.5 {
			found = true
		}
	}
	if !found {
		t.Fatalf("no horizontal X axis at y=%.1f (zero line)", wantAxisY)
	}

	for _, o := range objs {
		c, ok := o.(*canvas.Circle)
		if !ok {
			continue
		}
		center := c.Position().Y + c.Size().Height/2
		if center < defaultMarginTop || center > defaultMarginTop+plotHeight {
			t.Errorf("point at y=%.1f is outside the plot area", center)
		}
	}
}

func abs32(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
