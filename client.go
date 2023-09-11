package nordpool

import (
	"github.com/go-resty/resty/v2"
)

const URI = "https://www.nordpoolspot.com/api/marketdata/page/10"

func GetNordpoolData() (*NordpoolData, error) {
	client := resty.New()
	nordpoolData := &NordpoolData{}
	_, err := client.R().SetResult(nordpoolData).Get(URI + "?currency=SEK,SEK,EUR,EUR")

	return nordpoolData, err
}
