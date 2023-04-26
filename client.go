package nordpool

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const URI = "https://www.nordpoolspot.com/api/marketdata/page/10"

func GetNordpoolData() (*NordpoolData, error) {
	client := resty.New()
	nordpoolData := &NordpoolData{}
	resp, err := client.R().EnableTrace().SetResult(nordpoolData).Get(URI + "?currency=SEK,SEK,EUR,EUR")

	ti := resp.Request.TraceInfo()
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	return nordpoolData, err
}
