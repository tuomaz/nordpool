package nordpool

import (
	"fmt"
	"time"
)

type DayAheadData struct {
	Currency             string                `json:"currency"`
	ExchangeRate         float64               `json:"exchangeRate"`
	AreaStates           []AreaState           `json:"areaStates"`
	AreaAverages         []AreaAverage         `json:"areaAverages"`
	MultiAreaEntries     []MultiAreaEntry      `json:"multiAreaEntries"`
	BlockPriceAggregates []BlockPriceAggregate `json:"blockPriceAggregates"`
}

// NordpoolData is an alias for DayAheadData for backwards compatibility.
type NordpoolData = DayAheadData

type AreaState struct {
	State string   `json:"state"`
	Areas []string `json:"areas"`
}

type AreaAverage struct {
	AreaCode string  `json:"areaCode"`
	Price    float64 `json:"price"`
}

type MultiAreaEntry struct {
	DeliveryStart time.Time          `json:"deliveryStart"`
	DeliveryEnd   time.Time          `json:"deliveryEnd"`
	EntryPerArea  map[string]float64 `json:"entryPerArea"`
}

type BlockPriceAggregate struct {
	BlockName           string                       `json:"blockName"`
	DeliveryStart       time.Time                    `json:"deliveryStart"`
	DeliveryEnd         time.Time                    `json:"deliveryEnd"`
	AveragePricePerArea map[string]BlockPriceDetails `json:"averagePricePerArea"`
}

type BlockPriceDetails struct {
	Average float64 `json:"average"`
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
}

// GetPriceForTime returns the price for the specified time and delivery area.
func (d *DayAheadData) GetPriceForTime(t time.Time, area string) (float64, error) {
	for _, entry := range d.MultiAreaEntries {
		if (t.After(entry.DeliveryStart) || t.Equal(entry.DeliveryStart)) && t.Before(entry.DeliveryEnd) {
			if val, ok := entry.EntryPerArea[area]; ok {
				return val, nil
			}
			return 0, fmt.Errorf("nordpool: area %s not found in entry for time %s", area, t.Format(time.RFC3339))
		}
	}
	return 0, fmt.Errorf("nordpool: no entry found for time %s", t.Format(time.RFC3339))
}

// GetSliceOfPricesForArea returns all prices for the given area across the dataset.
func (d *DayAheadData) GetSliceOfPricesForArea(area string) ([]float64, error) {
	if len(d.MultiAreaEntries) == 0 {
		return nil, fmt.Errorf("nordpool: no entries in dataset")
	}

	result := make([]float64, 0, len(d.MultiAreaEntries))
	for _, entry := range d.MultiAreaEntries {
		if val, ok := entry.EntryPerArea[area]; ok {
			result = append(result, val)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("nordpool: could not find area %s in data", area)
	}

	return result, nil
}
