package database

import "time"

func InsertStock(ticker, name, sector string, market_cap float64) error{
	_, err := DB.Exec(
	"INSERT INTO stocks(ticker, name, market_cap, sector, last_updated) VALUES($1, $2, $3, $4, $5) ON CONFLICT (ticker) DO UPDATE SET market_cap = $3, last_updated = $5",
	ticker,name,market_cap,sector,time.Now().UTC(),
	)
	return err
}

func GetTotalMarketCap()(float64, error){
	var total float64
	err := DB.QueryRow("SELECT SUM(market_cap) FROM stocks").Scan(&total)
	return total, err
}
