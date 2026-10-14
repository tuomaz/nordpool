package nordpool

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const BaseURI = "https://dataportal-api.nordpoolgroup.com/api/DayAheadPrices"

var DefaultAreas = []string{
	"SE1", "SE2", "SE3", "SE4",
	"NO1", "NO2", "NO3", "NO4", "NO5",
	"FI", "DK1", "DK2",
}

// GetDayAheadPrices fetches day-ahead prices for the given date, currency, and delivery areas.
// If date is zero, today's date is used.
// If currency is empty, "SEK" is used.
// If deliveryAreas is empty, DefaultAreas are queried.
func GetDayAheadPrices(date time.Time, currency string, deliveryAreas ...string) (*DayAheadData, error) {
	if currency == "" {
		currency = "SEK"
	}
	if len(deliveryAreas) == 0 {
		deliveryAreas = DefaultAreas
	}
	if date.IsZero() {
		date = time.Now()
	}

	dateStr := date.Format("2006-01-02")
	areasStr := strings.Join(deliveryAreas, ",")

	client := resty.New()
	data := &DayAheadData{}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"currency":     currency,
			"market":       "DayAhead",
			"deliveryArea": areasStr,
			"date":         dateStr,
		}).
		SetResult(data).
		Get(BaseURI)

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("nordpool api returned status %d: %s", resp.StatusCode(), resp.String())
	}

	return data, nil
}

// GetNordpoolData fetches today's day-ahead prices with default settings (SEK, all Nordic areas).
func GetNordpoolData() (*DayAheadData, error) {
	return GetDayAheadPrices(time.Now(), "SEK")
}
