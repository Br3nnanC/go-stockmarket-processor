package model

import "time"

type wholeMarket struct{
	Time		time.Time 'json:"time"'		//time of snapshot
	marketCap	float64 'json:"market_cap"'	//get the entire market cap of the s&p
}
