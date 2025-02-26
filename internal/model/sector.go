package model

type sector struct{
	Name		string 'json:"name"'		    //name of the sector
	SectorMarketCap	float64 'json:"sector_market_cap"'  //get the total market cap within the sector
}
