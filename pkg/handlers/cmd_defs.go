package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "gen-stock",
		Description: "Generate a TradeAries link to the defined equity trade",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "ticker",
				Description: "The stock you want the option to be based on",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "type",
				Description: "Day or Long trade.",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Swing",
						Value: 0,
					},
					{
						Name:  "Day",
						Value: 1,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "stop",
				Description: "A stop loss that will expire the alert (this is currently non functional)",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "t-stop",
				Description: "A trailing stop loss %% that will expire the alert with the message STOP LOSS HIT",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "pt",
				Description: "A price target at which the trade will close",
				Required:    false,
			},
		},
	},
	{
		Name:        "gen-options",
		Description: "Generate a TradeAries link to the defined options trade",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "ticker",
				Description: "The stock you want the option to be based on",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "expiry",
				Description: "The expiry of the contract. If the expiry is this year, then mm/dd. Else, mm/ddd/yy.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "strike",
				Description: "The strike price of the contract + the contract type, e.g. 140C, 55.50P",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "type",
				Description: "Day or Long trade.",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Swing",
						Value: 0,
					},
					{
						Name:  "Day",
						Value: 1,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "stop",
				Description: "A stop loss that will expire the alert (this is currently non functional)",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "t-stop",
				Description: "A trailing stop loss %% that will expire the alert with the message STOP LOSS HIT",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "pt",
				Description: "A price target at which the trade will close",
				Required:    false,
			},
		},
	},
}
