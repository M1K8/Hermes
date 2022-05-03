package handlers

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/m1k8/hermes/pkg/aries"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"generate-stock": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var (
			ticker    string
			pt        float64
			stop      float64
			tstop     float64
			alertType int
			err       error
		)
		argValMap := make(map[string]string)

		for _, v := range i.ApplicationCommandData().Options {
			if v.Name == "type" {
				if v != nil {
					alertType = int(v.IntValue())
				}
			} else {
				argValMap[v.Name] = v.StringValue()
			}
		}

		log.Println(alertType)
		if v, ok := argValMap["ticker"]; !ok {
			return
		} else {
			ticker = strings.ToUpper(v)
		}
		log.Println("Calling stock for " + ticker)
		if v, ok := argValMap["pt"]; ok {
			pt, err = strconv.ParseFloat(v, 64)
			if err != nil {
				pt = 0
			}
		}
		if v, ok := argValMap["stop"]; ok {
			stop, err = strconv.ParseFloat(v, 64)
			if err != nil {
				pt = 0
			}
		}
		if v, ok := argValMap["t-stop"]; ok {
			tstop, err = strconv.ParseFloat(v, 64)
			if err != nil {
				pt = 0
			}
		}

		gen := aries.NewAriesGenerator()

		res, _ := gen.GetStockUrl(ticker, stop, tstop, pt)

		log.Println(res)

	},
	"generate-options": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var (
			ticker     string
			expiry     string
			strike     float64
			strikeType string
			pt         float64
			stop       float64
			tstop      float64
			alertType  int
			err        error
		)
		argValMap := make(map[string]string)

		for _, v := range i.ApplicationCommandData().Options {
			if v.Name == "type" {
				if v != nil {
					alertType = int(v.IntValue())
				}
			} else {
				argValMap[v.Name] = v.StringValue()
			}
		}

		log.Println(alertType)
		if v, ok := argValMap["ticker"]; !ok {
			return
		} else {
			ticker = strings.ToUpper(v)
		}
		log.Println("Calling stock for " + ticker)
		if v, ok := argValMap["expiry"]; !ok {
			return
		} else {
			expiry = v
		}

		if v, ok := argValMap["strike"]; !ok {
			return
		} else {
			strikePrice := v[:len(v)]
			strike, err = strconv.ParseFloat(strikePrice, 64)
			if err != nil {
				return
			}
			strikeType = v[len(v)-1:]
		}

		if v, ok := argValMap["pt"]; ok {
			pt, err = strconv.ParseFloat(v, 64)
			if err != nil {
				pt = 0
			}
		}
		if v, ok := argValMap["stop"]; ok {
			stop, err = strconv.ParseFloat(v, 64)
			if err != nil {
				pt = 0
			}
		}
		if v, ok := argValMap["t-stop"]; ok {
			tstop, err = strconv.ParseFloat(v, 64)
			if err != nil {
				pt = 0
			}
		}

		gen := aries.NewAriesGenerator()

		res, _ := gen.GetOptionsUrl(ticker, expiry, strikeType, strike, stop, tstop, pt)

		embed := &discordgo.MessageEmbed{}

		log.Println(res)

	},
}
