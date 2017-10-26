package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//
type Ticker struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	Priceusd         string `json:"price_usd"`
	Pricebtc         string `json:"price_btc"`
	Hvolumeusd24     string `json:"24h_volume_usd"`
	Marketcapusd     string `json:"market_cap_usd"`
	Availablesupply  string `json:"available_supply"`
	Totalsupply      string `json:"total_supply"`
	Percentchange1h  string `json:"percent_change_1h"`
	Percentchange24h string `json:"percent_change_24h"`
	Percentchange7d  string `json:"percent_change_7d"`
	Lastupdated      string `json:"last_updated"`
	Priceeur         string `json:"price_eur"`
	Hvolumeeur24     string `json:"24h_volume_eur"`
	Marketcapeur     string `json:"market_cap_eur"`
}

func main() {
	res, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?convert=EUR&limit=10")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	println(res.ContentLength)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	//fmt.Println(string(data))
	var ticker []Ticker
	err = json.Unmarshal(data, &ticker)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
	//fmt.Println("unmashaled", ticker)

	for _, one := range ticker {
		fmt.Println("Name: ", one.Name)
		fmt.Println("Symbol: ", one.Symbol)
		fmt.Println("Price USD: ", one.Priceusd)
		fmt.Println("=============================")
	}
}
