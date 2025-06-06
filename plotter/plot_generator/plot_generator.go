package plot_generator

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"plotter/csv_process"
)

func GeneratePlot(priceData []csv_process.UsdPrice) (*plot.Plot, error) {
	points := make(plotter.XYs, len(priceData))

	for i, el := range priceData {
		points[i].X = float64(el.Date.Unix())
		points[i].Y = float64(el.Price)
	}

	pricePlot := plot.New()
	pricePlot.Title.Text = "USD Price Chart"
	pricePlot.X.Label.Text = "Date"
	pricePlot.Y.Label.Text = "Price"

	if err := plotutil.AddLinePoints(pricePlot, "USDT", points); err != nil {
		return nil, fmt.Errorf("error adding line to plot %v", err)
	}

	pricePlot.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}

	return pricePlot, nil
}
