package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/m1k8/hermes/pkg/aries"
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
				Description: "Order Tyoe.",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Limit",
						Value: aries.Limit,
					},
					{
						Name:  "Market",
						Value: aries.Market,
					},
					{
						Name:  "Bracket",
						Value: aries.Bracket,
					},
					{
						Name:  "Stop_Market",
						Value: aries.Stop_Market,
					},
					{
						Name:  "Stop_Limit",
						Value: aries.Stop_Limit,
					},
					{
						Name:  "Trailing_Stop",
						Value: aries.Trailing_Stop,
					},
					{
						Name:  "Trailing_Stop_Pct",
						Value: aries.Trailing_Stop_Pct,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "duration",
				Description: "Order Duration.",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "UntilCancel",
						Value: aries.UntilCancel,
					},
					{
						Name:  "Day",
						Value: aries.Day,
					},
					{
						Name:  "Extended",
						Value: aries.Extended,
					},
					{
						Name:  "All_Hours",
						Value: aries.All_Hours,
					},
					{
						Name:  "Immediate",
						Value: aries.Immediate,
					},
					{
						Name:  "Fill_or_kill",
						Value: aries.Fill_or_kill,
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
				Description: "A trailing stop loss amount that will expire the alert with the message STOP LOSS HIT",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "t-stop-pct",
				Description: "A trailing stop loss %% that will expire the alert with the message STOP LOSS HIT",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "p-limit",
				Description: "Profit limit",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "l",
				Description: "Price limit",
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
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "Buy Or Sell (Options)",
				Description: "Whether to buy or sell the given contgract",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Buy",
						Value: aries.Buy,
					},
					{
						Name:  "Sell",
						Value: aries.Sell,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "type",
				Description: "Order Tyoe.",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Limit",
						Value: aries.Limit,
					},
					{
						Name:  "Market",
						Value: aries.Market,
					},
					{
						Name:  "Bracket",
						Value: aries.Bracket,
					},
					{
						Name:  "Stop_Market",
						Value: aries.Stop_Market,
					},
					{
						Name:  "Stop_Limit",
						Value: aries.Stop_Limit,
					},
					{
						Name:  "Trailing_Stop",
						Value: aries.Trailing_Stop,
					},
					{
						Name:  "Trailing_Stop_Pct",
						Value: aries.Trailing_Stop_Pct,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "duration",
				Description: "Order Duration.",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "UntilCancel",
						Value: aries.UntilCancel,
					},
					{
						Name:  "Day",
						Value: aries.Day,
					},
					{
						Name:  "Extended",
						Value: aries.Extended,
					},
					{
						Name:  "All_Hours",
						Value: aries.All_Hours,
					},
					{
						Name:  "Immediate",
						Value: aries.Immediate,
					},
					{
						Name:  "Fill_or_kill",
						Value: aries.Fill_or_kill,
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
				Description: "A trailing stop loss amount that will expire the alert with the message STOP LOSS HIT",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "t-stop-pct",
				Description: "A trailing stop loss %% that will expire the alert with the message STOP LOSS HIT",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "p-limit",
				Description: "Profit limit",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "l",
				Description: "Price limit",
				Required:    false,
			},
		},
	},
}
