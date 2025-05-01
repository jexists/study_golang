package main

import (
	"fmt"
)

type MarketMain struct {
	MarketNo       int    `bson:"market_no" json:"market_no"`
	MarketNickname string `bson:"market_nickname" json:"market_nickname"`
	MarketTitle    string `bson:"market_title" json:"market_title"`
}
type MarketDB struct {
	MarketMain
	MarketUsername string `bson:"market_username" json:"market_username"`
	MarketContent  string `bson:"market_content" json:"market_content"`
}

type MarketInList struct {
	MarketMain
}

func main() {
	var market MarketDB
	fmt.Println(market.MarketNo)
}

// type  MarketDB struct {
// 	ID                    primitive.ObjectID       `bson:"_id,omitempty" json:"_id"`
// 	MarketNo              int                      `bson:"market_no" json:"market_no"`
// 	MarketUsername        string                   `bson:"market_username" json:"market_username"`
// 	MarketNickname        string                   `bson:"market_nickname" json:"market_nickname"`
// 	MarketTitle           string                   `bson:"market_title" json:"market_title"`
// 	MarketContent         string                   `bson:"market_content" json:"market_content"`
// }

// type  MarketInList struct {
// 	ID                    primitive.ObjectID       `bson:"_id,omitempty" json:"_id"`
// 	MarketNo              int                      `bson:"market_no" json:"market_no"`
// 	MarketNickname        string                   `bson:"market_nickname" json:"market_nickname"`
// 	MarketTitle           string                   `bson:"market_title" json:"market_title"`
// }
