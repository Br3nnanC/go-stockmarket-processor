package model

import "time"

type stock struct{
	Ticker		string 'json:"ticker"'		//stocks ticker symbol
	Name		string 'json:"name"'		//company name
	MarketCap	float64 'json:"market_cap"'	//market cap of company
	Sector		string 'json:"sector"'		//sector the company is in
	LastUpdated	time.Time 'json:"last_updated"'	//the time between the last pull of this stock
}
