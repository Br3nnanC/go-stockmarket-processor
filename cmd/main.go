package main

import(
	"log"
	"github.com/Br3nnanC/go-stockmarket-processor/internal/database"
	"fmt"
)

func main(){
	//test if database connected
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("database failed to connect. Error = ", err)
	}
	fmt.Println("database has connected to main")
	
	err = database.InsertStock("AAPL", "Apple", "Technology", 100000)
	if err != nil {
		log.Fatal("failed to insert the stock", err)
	}

	total, err := database.GetTotalMarketCap()
	if err != nil {
		log.Fatal("failed to calculate total market cap", err)
	}
	fmt.Println("the total current market cap of the sp500 is: ", total)
}
