package nordpool

import (
	"testing"
	"time"
)

func TestGetDayAheadPrices_Live(t *testing.T) {
	data, err := GetDayAheadPrices(time.Now(), "SEK", "SE2")
	if err != nil {
		t.Fatalf("failed to fetch prices: %v", err)
	}

	if data.Currency != "SEK" {
		t.Errorf("expected currency SEK, got %s", data.Currency)
	}

	if len(data.MultiAreaEntries) == 0 {
		t.Fatal("expected multiAreaEntries to be non-empty")
	}

	t.Logf("Fetched %d entries for area SE2", len(data.MultiAreaEntries))

	// Test GetPriceForTime
	now := time.Now()
	price, err := data.GetPriceForTime(now, "SE2")
	if err != nil {
		t.Logf("Note: GetPriceForTime for current time returned err: %v (current time might be outside today's delivery window)", err)
	} else {
		t.Logf("Current price for SE2: %.2f SEK/MWh", price)
	}

	// Test GetSliceOfPricesForArea
	slice, err := data.GetSliceOfPricesForArea("SE2")
	if err != nil {
		t.Fatalf("failed to get slice for SE2: %v", err)
	}
	if len(slice) == 0 {
		t.Fatal("expected slice to be non-empty")
	}
	t.Logf("Slice of prices length: %d, first price: %.2f", len(slice), slice[0])
}
