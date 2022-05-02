package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"generate-stock": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var (
			ticker    string
			pt        string
			stop      string
			tstop     string
			alertType int
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
		if v, ok := argValMap["ticker"]; !ok {
			return
		} else {
			ticker = strings.ToUpper(v)
		}
		log.Println("Calling stock for " + ticker)
		if v, ok := argValMap["pt"]; ok {
			pt = v
		}
		if v, ok := argValMap["stop"]; ok {
			stop = v
		}
		if v, ok := argValMap["t-stop"]; ok {
			tstop = v
		}

	},
}
