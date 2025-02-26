package main

import(
	"log"
	"github.com/Br3nnanC/go-stockmarkt-processor/internal/database"
)

func main(){
	//test if database connected
	err := database.connectDB()
	if err != nil {
		log.Fatal("database failed to connect. Error = ", err)
	}
	fmt.Println("database has connected to main")
	
}
