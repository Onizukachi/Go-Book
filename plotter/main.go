package main

import (
	"fmt"
	"os"
	"plotter/csv_process"
	"plotter/plot_generator"

	"gonum.org/v1/plot/vg"
)

func main() {
	priceData, err := csv_process.LoadFromFile("./prices.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading file: %v", err)
		os.Exit(1)
	}

	plot, err := plot_generator.GeneratePlot(priceData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plot: %v", err)
		os.Exit(1)
	}

	if err := plot.Save(15*vg.Inch, 4*vg.Inch, "prices.png"); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving plot: %v", err)
		os.Exit(1)
	}
}
