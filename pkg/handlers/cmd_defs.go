package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/m1k8/hermes/pkg/aries"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "stock",
		Description: "Generate a direct link to a ticker with the supplied parameters",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "limit",
				Description: "Limit order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-price",
						Description: "Limit Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "market",
				Description: "Market order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "stop-market",
				Description: "Stop Market order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-price",
						Description: "Stop Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "stop-limit",
				Description: "Stop limit order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-price",
						Description: "Stop Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-price",
						Description: "Limit Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "bracket",
				Description: "Bracket order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-price",
						Description: "Limit price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "profit-limit",
						Description: "Profit Limit.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-loss",
						Description: "Stop Loss.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "trailing-stop",
				Description: "Trailing Stop.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "trail-amount",
						Description: "Trail Amount.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "trailing-percent",
				Description: "Trailing Stop Percent.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "trail-percent",
						Description: "Trail Percent.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "bracket_pct",
				Description: "Bracket order based on percentages",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "ticker",
						Description: "The stock you want the option to be based on.",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "short",
						Description: "Short.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-pct",
						Description: "Limit percentage gain.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "profit-limit",
						Description: "Profit Limit.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-loss",
						Description: "Stop Loss (%).",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
		},
	},

	{
		Name:        "options",
		Description: "Generate a direct link to one or more contracts with the supplied parameters",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "limit",
				Description: "Limit order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A comma seperated list of 1 or more contracts, for example B AAPL 4/20 142.5C",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-price",
						Description: "Limit Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "market",
				Description: "Market order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A comma seperated list of 1 or more contracts, for example B AAPL 4/20 142.5C",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "stop-market",
				Description: "Stop Market order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A comma seperated list of 1 or more contracts, for example B AAPL 4/20 142.5C",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-price",
						Description: "Stop Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "stop-limit",
				Description: "Stop limit order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A comma seperated list of 1 or more contracts, for example B AAPL 4/20 142.5C",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-price",
						Description: "Stop Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-price",
						Description: "Limit Price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "bracket",
				Description: "Bracket order.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A single contract, for example B AAPL 4/20 142.5C; where B is buy (S for sell)",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-price",
						Description: "Limit price.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "profit-limit",
						Description: "Profit Limit.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-loss",
						Description: "Stop Loss.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "bracket_pct",
				Description: "Bracket order based on percentages",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A single contract, for example B AAPL 4/20 142.5C; where B is buy (S for sell)",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "limit-pct",
						Description: "Limit percentage gain.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "profit-limit",
						Description: "Profit Limit.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stop-loss",
						Description: "Stop Loss (%).",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "trailing-stop",
				Description: "Trailing Stop.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A comma seperated list of 1 or more contracts, for example B AAPL 4/20 142.5C",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "trail-amount",
						Description: "Trail Amount.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
			{
				Name:        "trailing-percent",
				Description: "Trailing Stop Percent.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "contract-defs",
						Description: "A comma seperated list of 1 or more contracts, for example B AAPL 4/20 142.5C",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "trail-percent",
						Description: "Trail Percent.",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "close",
						Description: "Is this a close order?",
						Required:    false,
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
				},
			},
		},
	},
}
