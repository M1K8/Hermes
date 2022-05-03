package aries

import "sync"

/*
Base:
https://app.tradearies.com/

Open equity position:
/order/{SYMBOL}

Close equity position:
/close/{SYMBOL}

Open options position:
/options/order

Close options position:
/options/close

—

URL Queries:

o: List of options contracts for option trades separated by commas. Pre-fixed B! for buy and S! for sell. Spaces are replaced by _ . Example: o=S!AAPL_220520C152.5 means sell AAPL May 20th 2022 152.50 Call

t: Order type. Limit = 0, Market = 1, Bracket = 2 (does not work for multi-leg option orders), Stop Market = 3, Stop Limit = 4, Trailing Stop = 5, Trailing Stop % = 6

d: Order duration. Until Cancel = 0, Day = 1, Extended = 2, All Hours = 3, Immediate = 4, Fill or Kill = 5

l: Limit price. Example: 152.50

s: Stop price.

pl: Profit Limit (for bracket orders)

sl: Stop Loss (for bracket orders)

ta: Trail Amount (for trailing stop)

tp: Trail percent (for trailing stop %)

Example:
https://app.tradearies.com/options/order?o=B!AAPL_220520C152.5,S!AAPL_220520C155&t=0&l=1.55

*/

type Generator struct {
	baseUrl string
}

var gen *Generator
var once sync.Once = sync.Once{}

func NewAriesGenerator() *Generator {
	once.Do(func() {
		gen = &Generator{
			baseUrl: "https://app.tradearies.com/",
		}
	})

	return gen

}

type OrderDuration int
type OrderType int
type ContractBuyOrSell bool

const (
	UntilCancel  OrderDuration = 0
	Day          OrderDuration = 1
	Extended     OrderDuration = 2
	All_Hours    OrderDuration = 3
	Immediate    OrderDuration = 4
	Fill_or_kill OrderDuration = 5

	Buy  ContractBuyOrSell = true
	Sell ContractBuyOrSell = false

	Limit             OrderType = 0
	Market            OrderType = 1
	Bracket           OrderType = 2
	Stop_Market       OrderType = 3
	Stop_Limit        OrderType = 4
	Trailing_Stop     OrderType = 5
	Trailing_Stop_Pct OrderType = 6 //TODO - Verification based on the type

	// Params
	order = "/order/"
	close = "/close"

	options      = "o="
	orderType    = "t="
	orderDur     = "d="
	limitPrice   = "l="
	stopPrice    = "s="
	profitLimit  = "pl="
	stopLoss     = "sl="
	trailAmount  = "ta="
	trailPercent = "tp="
)

func (g *Generator) GetBaseUrl() string {
	return g.baseUrl
}

func (g *Generator) GetStockUrl(ticker string, buySell OrderType, duration OrderDuration, limit, stop, t_a, t_p, pl, pt float64) (string, error) {
	return g.baseUrl + "", nil
}

func (g *Generator) GetOptionsUrl(ticker, expiry, strikeTyoe string, buySell OrderType, duration OrderDuration, strike, limit, stop, t_a, t_p, pl, pt float64) (string, error) {
	return g.baseUrl + "options/", nil
}
