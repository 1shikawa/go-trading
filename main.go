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

package main

import (
	"fmt"
	"go-trading/app/models"
	"time"
)

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

func main() {
	s := models.NewSignalEvents()
	df, _ := models.GetAllCandle("BTC_USD", time.Minute, 10)
	c1 := df.Candles[0]
	c2 := df.Candles[5]
	s.Buy("BTC_USD", c1.Time.UTC(), c1.Close, 1.0, true)
	s.Sell("BTC_USD", c2.Time.UTC(), c2.Close, 1.0, true)
	fmt.Println(models.GetSignalEventsByCount(1))
	fmt.Println(models.GetSignalEventsAfterTime(c1.Time))
	fmt.Println(s.CollectAfter(time.Now().UTC()))
	fmt.Println(s.CollectAfter(c1.Time))
}

