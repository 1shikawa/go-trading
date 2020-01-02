// package main

// import (
// 	// "fmt"
// 	// "go-trading/app/models"
// 	"go-trading/app/controllers"
// 	"go-trading/config"
// 	"go-trading/utils"

// 	// "time"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	utils.LoggingSettings(config.Config.LogFile)
// apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

// tickerChannel := make(chan bitflyer.Ticker)
// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
// for ticker := range tickerChannel {
// 	fmt.Println(ticker)
// 	fmt.Println(ticker.GetMidPrice())
// 	fmt.Println(ticker.DateTime())
// 	fmt.Println(ticker.TruncateDateTime(time.Second))
// 	fmt.Println(ticker.TruncateDateTime(time.Minute))
// 	fmt.Println(ticker.TruncateDateTime(time.Hour))
// }

// 実際の注文
// order := &bitflyer.Order{
// 	ProductCode:     config.Config.ProductCode,
// 	ChildOrderType:  "MARKET",
// 	// ChildOrderType:  "LIMIT",
// 	Side:            "BUY",
// 	// Price:           7000,
// 	Size:            0.001,
// 	MinuteToExpires: 1,
// 	TimeInForce:     "GTC",
// }
// res, _ := apiClient.SendOrder(order)
// fmt.Println(res.ChildOrderAcceptanceID)

//i := "JRF20181012-144016-140584"
//params := map[string]string{
//	"product_code": config.Config.ProductCode,
//	"child_order_acceptance_id": i,
//}
//r, _ := apiClient.ListOrder(params)
//fmt.Println(r)
// fmt.Println(models.DbConnection)
// 	controllers.StreamIngestionData()
// 	controllers.StartWebServer()

// }

//import (
//	"gotrading/app/controllers"
//	"gotrading/config"
//	"gotrading/utils"
//)
//
//func main(){
//	utils.LoggingSettings(config.Config.LogFile)
//	controllers.StreamIngestionData()
//	controllers.StartWebServer()
//}
package main

import (
	"go-trading/app/controllers"
	"go-trading/config"
	"go-trading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
