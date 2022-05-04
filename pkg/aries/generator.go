package aries

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/uniplaces/carbon"
)

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
	order      = "order"
	close      = "close"
	optionsArg = "options/"

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

func (g *Generator) GetStockUrl(ticker string, short bool, buySell ContractBuyOrSell, order_type OrderType, duration OrderDuration, limit, stopLossP, stopPriceP, t_a, t_p, pl float64) (string, error) {
	var (
		buySellStr   string
		oType        string
		dur          string
		limitStr     string
		stopLossStr  string
		stopPriceStr string
		ta           string
		tp           string
		plStr        string
		shortStr     string
	)

	if buySell {
		buySellStr = order
	} else {
		buySellStr = close
	}

	if short {
		shortStr = "&sh=true"
	}

	switch order_type {
	case Limit:
		oType = "?" + orderType + "0"
	case Market:
		oType = "?" + orderType + "1"
	case Bracket:
		oType = "?" + orderType + "2"
	case Stop_Market:
		oType = "?" + orderType + "3"
	case Stop_Limit:
		oType = "?" + orderType + "4"
	case Trailing_Stop:
		oType = "?" + orderType + "5"
	case Trailing_Stop_Pct:
		oType = "?" + orderType + "6"
	default:
		oType = ""
	}

	switch duration {
	case UntilCancel:
		dur = "&" + orderDur + "0"
	case Day:
		dur = "&" + orderDur + "1"
	case Extended:
		dur = "&" + orderDur + "2"
	case All_Hours:
		dur = "&" + orderDur + "3"
	case Immediate:
		dur = "&" + orderDur + "4"
	case Fill_or_kill:
		dur = "&" + orderDur + "5"
	default:
		dur = ""
	}

	if limit == 0.0 {
		limitStr = ""
	} else {
		limitStr = fmt.Sprintf("&%v%.2f", limitPrice, limit)
	}

	if t_a == 0.0 {
		ta = ""
	} else {
		ta = fmt.Sprintf("&%v%.2f", trailAmount, t_a)
	}

	if stopLossP == 0.0 {
		stopLossStr = ""
	} else {
		stopLossStr = fmt.Sprintf("&%v%.2f", stopLoss, stopLossP)
	}

	if stopPriceP == 0.0 {
		stopPriceStr = ""
	} else {
		stopPriceStr = fmt.Sprintf("&%v%.2f", stopPrice, stopPriceP)
	}

	if t_p == 0.0 {
		tp = ""
	} else {
		tp = fmt.Sprintf("&%v%.2f", trailPercent, t_p)
	}

	if pl == 0.0 {
		plStr = ""
	} else {
		plStr = fmt.Sprintf("&%v%.2f", profitLimit, pl)
	}

	return g.baseUrl + buySellStr + "/" + ticker + oType + dur + limitStr + ta + stopLossStr + stopPriceStr + tp + plStr + shortStr, nil
}

func (g *Generator) GetOptionsUrl(ticker string, buySell ContractBuyOrSell, order_type OrderType, duration OrderDuration, limit, stopLossP, stopPriceP, t_a, t_p, pl float64) (string, error) {
	var (
		buySellStr   string
		oType        string
		dur          string
		limitStr     string
		stopLossStr  string
		stopPriceStr string
		ta           string
		tp           string
		plStr        string
	)

	if buySell {
		buySellStr = order
	} else {
		buySellStr = close
	}

	switch order_type {
	case Limit:
		oType = "&" + orderType + "0"
	case Market:
		oType = "&" + orderType + "1"
	case Bracket:
		oType = "&" + orderType + "2"
	case Stop_Market:
		oType = "&" + orderType + "3"
	case Stop_Limit:
		oType = "&" + orderType + "4"
	case Trailing_Stop:
		oType = "&" + orderType + "5"
	case Trailing_Stop_Pct:
		oType = "&" + orderType + "6"
	default:
		oType = ""
	}

	switch duration {
	case UntilCancel:
		dur = "&" + orderDur + "0"
	case Day:
		dur = "&" + orderDur + "1"
	case Extended:
		dur = "&" + orderDur + "2"
	case All_Hours:
		dur = "&" + orderDur + "3"
	case Immediate:
		dur = "&" + orderDur + "4"
	case Fill_or_kill:
		dur = "&" + orderDur + "5"
	default:
		dur = ""
	}

	if limit == 0.0 {
		limitStr = ""
	} else {
		limitStr = fmt.Sprintf("&%v%.2f", limitPrice, limit)
	}

	if t_a == 0.0 {
		ta = ""
	} else {
		ta = fmt.Sprintf("&%v%.2f", trailAmount, t_a)
	}

	if stopLossP == 0.0 {
		stopLossStr = ""
	} else {
		stopLossStr = fmt.Sprintf("&%v%.2f", stopLoss, stopLossP)
	}

	if stopPriceP == 0.0 {
		stopPriceStr = ""
	} else {
		stopPriceStr = fmt.Sprintf("&%v%.2f", stopPrice, stopPriceP)
	}

	if t_p == 0.0 {
		tp = ""
	} else {
		tp = fmt.Sprintf("&%v%.2f", trailPercent, t_p)
	}

	if pl == 0.0 {
		plStr = ""
	} else {
		plStr = fmt.Sprintf("&%v%.2f", profitLimit, pl)
	}

	return g.baseUrl + optionsArg + buySellStr + "?" + options + ticker + oType + dur + limitStr + ta + stopLossStr + stopPriceStr + tp + plStr, nil
}

func ParseOption(opt string) (option string, err error) {
	res := strings.Split(opt, " ")

	if len(res) != 4 {
		err = errors.New("Invalid contract definition: " + opt)
		return
	}
	buySell := ""
	switch strings.ToUpper(res[0]) {
	case "B":
		buySell = "B!"
	case "S":
		buySell = "S!"
	default:
		err = errors.New("Inavlid Buy/Sell indicator: " + res[0])
		return
	}

	ticker := strings.ToUpper(res[1])

	m, d, y, err := parseDate(res[2])
	if err != nil {
		return
	}

	strike := res[3]

	var putCall string

	switch pc := strings.ToUpper(strike[len(strike)-1:]); pc {
	case "P":
		putCall = "P"
	case "C":
		putCall = "C"
	default:
		err = errors.New("Invalid Put/Call Indicator: " + strike[len(strike)-2:])
		return
	}

	option = buySell + ticker + "_" + y + m + d + putCall + strike[:len(strike)-1]

	return
}

func ParseOptions(opts string) ([]string, error) {
	res := strings.Split(opts, ",")

	if len(res) == 0 {
		return nil, errors.New("No contracts provided")
	}

	resArr := make([]string, 0)

	for _, v := range res {
		opt, err := ParseOption(v)
		if err != nil {
			return nil, err
		}

		resArr = append(resArr, opt)
	}

	return resArr, nil
}

func parseDate(in string) (m, d, y string, err error) {
	strs := strings.Split(in, "/")

	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()

	if len(strs) != 2 && len(strs) != 3 {
		return "", "", "", errors.New("invalid date string provided")
	}
	m = strs[0]
	if len(m) == 1 {
		m = "0" + m
	}

	d = strs[1]
	if len(d) == 1 {
		d = "0" + d
	}

	if len(strs) == 3 {
		y = strs[2]
	} else {
		nowC := carbon.NewCarbon(time.Now()).Year()
		y = fmt.Sprintf("%v", nowC)
		y = y[2:]
	}

	return
}
