package csv_process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type UsdPrice struct {
	Price float64
	Date  time.Time
}

func LoadFromFile(path string) ([]UsdPrice, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file %v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var usdPrices []UsdPrice

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading file %v", err)
		}

		date, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing date %v", err)
		}

		price, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing price %v", err)
		}

		usdPrices = append(usdPrices, UsdPrice{
			Price: price,
			Date:  date,
		})
	}

	return usdPrices, nil
}
