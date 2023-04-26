package nordpool

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type NordpoolData struct {
	Data struct {
		Rows []struct {
			Columns []struct {
				Index                            int    `json:"Index,omitempty"`
				Scale                            int    `json:"Scale,omitempty"`
				SecondaryValue                   any    `json:"SecondaryValue,omitempty"`
				IsDominatingDirection            bool   `json:"IsDominatingDirection,omitempty"`
				IsValid                          bool   `json:"IsValid,omitempty"`
				IsAdditionalData                 bool   `json:"IsAdditionalData,omitempty"`
				Behavior                         int    `json:"Behavior,omitempty"`
				Name                             string `json:"Name,omitempty"`
				Value                            string `json:"Value,omitempty"`
				GroupHeader                      string `json:"GroupHeader,omitempty"`
				DisplayNegativeValueInBlue       bool   `json:"DisplayNegativeValueInBlue,omitempty"`
				CombinedName                     string `json:"CombinedName,omitempty"`
				DateTimeForData                  string `json:"DateTimeForData,omitempty"`
				DisplayName                      string `json:"DisplayName,omitempty"`
				DisplayNameOrDominatingDirection string `json:"DisplayNameOrDominatingDirection,omitempty"`
				IsOfficial                       bool   `json:"IsOfficial,omitempty"`
				UseDashDisplayStyle              bool   `json:"UseDashDisplayStyle,omitempty"`
			} `json:"Columns,omitempty"`
			Name            string `json:"Name,omitempty"`
			StartTime       string `json:"StartTime,omitempty"`
			EndTime         string `json:"EndTime,omitempty"`
			DateTimeForData string `json:"DateTimeForData,omitempty"`
			DayNumber       int    `json:"DayNumber,omitempty"`
			StartTimeDate   string `json:"StartTimeDate,omitempty"`
			IsExtraRow      bool   `json:"IsExtraRow,omitempty"`
			IsNtcRow        bool   `json:"IsNtcRow,omitempty"`
			EmptyValue      string `json:"EmptyValue,omitempty"`
			Parent          any    `json:"Parent,omitempty"`
		} `json:"Rows,omitempty"`
		IsDivided                 bool     `json:"IsDivided,omitempty"`
		SectionNames              []string `json:"SectionNames,omitempty"`
		EntityIDs                 []string `json:"EntityIDs,omitempty"`
		DataStartdate             string   `json:"DataStartdate,omitempty"`
		DataEnddate               string   `json:"DataEnddate,omitempty"`
		MinDateForTimeScale       string   `json:"MinDateForTimeScale,omitempty"`
		AreaChanges               []any    `json:"AreaChanges,omitempty"`
		Units                     []string `json:"Units,omitempty"`
		LatestResultDate          string   `json:"LatestResultDate,omitempty"`
		ContainsPreliminaryValues bool     `json:"ContainsPreliminaryValues,omitempty"`
		ContainsExchangeRates     bool     `json:"ContainsExchangeRates,omitempty"`
		ExchangeRateOfficial      string   `json:"ExchangeRateOfficial,omitempty"`
		ExchangeRatePreliminary   string   `json:"ExchangeRatePreliminary,omitempty"`
		ExchangeUnit              any      `json:"ExchangeUnit,omitempty"`
		DateUpdated               string   `json:"DateUpdated,omitempty"`
		CombinedHeadersEnabled    bool     `json:"CombinedHeadersEnabled,omitempty"`
		DataType                  int      `json:"DataType,omitempty"`
		TimeZoneInformation       int      `json:"TimeZoneInformation,omitempty"`
	} `json:"data,omitempty"`
	CacheKey string `json:"cacheKey,omitempty"`
	Conf     struct {
		ID               string `json:"Id,omitempty"`
		Name             any    `json:"Name,omitempty"`
		Published        string `json:"Published,omitempty"`
		ShowGraph        bool   `json:"ShowGraph,omitempty"`
		ResolutionPeriod struct {
			ID           string `json:"Id,omitempty"`
			Resolution   int    `json:"Resolution,omitempty"`
			Unit         int    `json:"Unit,omitempty"`
			PeriodNumber int    `json:"PeriodNumber,omitempty"`
		} `json:"ResolutionPeriod,omitempty"`
		ResolutionPeriodY struct {
			ID           string `json:"Id,omitempty"`
			Resolution   int    `json:"Resolution,omitempty"`
			Unit         int    `json:"Unit,omitempty"`
			PeriodNumber int    `json:"PeriodNumber,omitempty"`
		} `json:"ResolutionPeriodY,omitempty"`
		Entities []struct {
			ProductType struct {
				ID         string `json:"Id,omitempty"`
				Attributes []struct {
					ID       string   `json:"Id,omitempty"`
					Name     string   `json:"Name,omitempty"`
					Role     string   `json:"Role,omitempty"`
					HasRoles bool     `json:"HasRoles,omitempty"`
					Values   []string `json:"Values,omitempty"`
				} `json:"Attributes,omitempty"`
				Name        string `json:"Name,omitempty"`
				DisplayName string `json:"DisplayName,omitempty"`
			} `json:"ProductType,omitempty"`
			SecondaryProductType struct {
				ID          string `json:"Id,omitempty"`
				Attributes  any    `json:"Attributes,omitempty"`
				Name        string `json:"Name,omitempty"`
				DisplayName string `json:"DisplayName,omitempty"`
			} `json:"SecondaryProductType,omitempty"`
			SecondaryProductBehavior int    `json:"SecondaryProductBehavior,omitempty"`
			ID                       string `json:"Id,omitempty"`
			Name                     string `json:"Name,omitempty"`
			GroupHeader              string `json:"GroupHeader,omitempty"`
			DataUpdated              string `json:"DataUpdated,omitempty"`
			Attributes               []struct {
				ID       string `json:"Id,omitempty"`
				Name     string `json:"Name,omitempty"`
				HasRoles bool   `json:"HasRoles,omitempty"`
				Role     string `json:"Role,omitempty"`
				Value    string `json:"Value,omitempty"`
			} `json:"Attributes,omitempty"`
			Drillable                   bool   `json:"Drillable,omitempty"`
			DateRanges                  []any  `json:"DateRanges,omitempty"`
			Index                       int    `json:"Index,omitempty"`
			IndexForColumn              int    `json:"IndexForColumn,omitempty"`
			MinMaxDisabled              bool   `json:"MinMaxDisabled,omitempty"`
			DisableNumberGroupSeparator int    `json:"DisableNumberGroupSeparator,omitempty"`
			TimeserieID                 any    `json:"TimeserieID,omitempty"`
			SecondaryTimeserieID        string `json:"SecondaryTimeserieID,omitempty"`
			HasPreliminary              bool   `json:"HasPreliminary,omitempty"`
			TimeseriePreliminaryID      any    `json:"TimeseriePreliminaryID,omitempty"`
			Scale                       int    `json:"Scale,omitempty"`
			SecondaryScale              int    `json:"SecondaryScale,omitempty"`
			DataType                    int    `json:"DataType,omitempty"`
			SecondaryDataType           int    `json:"SecondaryDataType,omitempty"`
			LastUpdate                  string `json:"LastUpdate,omitempty"`
			Unit                        any    `json:"Unit,omitempty"`
			IsDominatingDirection       bool   `json:"IsDominatingDirection,omitempty"`
			DisplayAsSeparatedColumn    bool   `json:"DisplayAsSeparatedColumn,omitempty"`
			EnableInChart               bool   `json:"EnableInChart,omitempty"`
			BlueNegativeValues          bool   `json:"BlueNegativeValues,omitempty"`
		} `json:"Entities,omitempty"`
		TableType int `json:"TableType,omitempty"`
		ExtraRows []struct {
			ID             string   `json:"Id,omitempty"`
			Header         string   `json:"Header,omitempty"`
			ColumnProducts []string `json:"ColumnProducts,omitempty"`
		} `json:"ExtraRows,omitempty"`
		Filters []struct {
			ID            string   `json:"Id,omitempty"`
			AttributeName string   `json:"AttributeName,omitempty"`
			Values        []string `json:"Values,omitempty"`
			DefaultValue  string   `json:"DefaultValue,omitempty"`
		} `json:"Filters,omitempty"`
		IsDrillDownEnabled bool   `json:"IsDrillDownEnabled,omitempty"`
		DrillDownMode      int    `json:"DrillDownMode,omitempty"`
		IsMinValueEnabled  bool   `json:"IsMinValueEnabled,omitempty"`
		IsMaxValueEnabled  bool   `json:"IsMaxValueEnabled,omitempty"`
		ValidYearsBack     int    `json:"ValidYearsBack,omitempty"`
		TimeScaleUnit      string `json:"TimeScaleUnit,omitempty"`
		IsNtcEnabled       bool   `json:"IsNtcEnabled,omitempty"`
		NtcProductType     struct {
			ID          string `json:"Id,omitempty"`
			Attributes  any    `json:"Attributes,omitempty"`
			Name        string `json:"Name,omitempty"`
			DisplayName string `json:"DisplayName,omitempty"`
		} `json:"NtcProductType,omitempty"`
		NtcHeader                string `json:"NtcHeader,omitempty"`
		ShowTimelineGraph        int    `json:"ShowTimelineGraph,omitempty"`
		ExchangeMode             int    `json:"ExchangeMode,omitempty"`
		IsPivotTable             int    `json:"IsPivotTable,omitempty"`
		IsCombinedHeadersEnabled int    `json:"IsCombinedHeadersEnabled,omitempty"`
		NtcFormat                int    `json:"NtcFormat,omitempty"`
		DisplayHourAlsoInUKTime  bool   `json:"DisplayHourAlsoInUKTime,omitempty"`
	} `json:"conf,omitempty"`
	Header struct {
		Title              string `json:"title,omitempty"`
		Description        string `json:"description,omitempty"`
		QuestionMarkInfo   string `json:"questionMarkInfo,omitempty"`
		HideDownloadButton string `json:"hideDownloadButton,omitempty"`
	} `json:"header,omitempty"`
	EndDate  any `json:"endDate,omitempty"`
	Currency any `json:"currency,omitempty"`
	PageID   int `json:"pageId,omitempty"`
}

func (n *NordpoolData) GetSliceOfPricesForArea(area string) ([]float64, error) {
	log.Printf("len rows = %v", len(n.Data.Rows))

	result := make([]float64, 0)
	index := n.findIndexForArea(area)

	if index > 0 {
		for _, row := range n.Data.Rows {
			log.Printf("Trying to parse %s", row.Columns[index].Value)
			price, err := strconv.ParseFloat(strings.Replace(row.Columns[index].Value, ",", ".", 1), 64)
			if err != nil {
				return nil, err
			}
			result = append(result, price)
		}
		return result, nil
	}
	return nil, fmt.Errorf("nordpool: could not find area %s in data", area)
}

func (n *NordpoolData) findIndexForArea(area string) int {
	for _, column := range n.Data.Rows[0].Columns {
		if column.Name == area {
			return column.Index
		}
	}
	return -1
}
