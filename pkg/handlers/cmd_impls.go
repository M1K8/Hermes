package handlers

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/m1k8/hermes/pkg/aries"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"stock": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		gen := aries.NewAriesGenerator()

		// As you can see, names of subcommands (nested, top-level)
		// and subcommand groups are provided through the arguments.
		switch options[0].Name {
		case "limit":
			var (
				ticker string
				close  bool
				lp     string
				lpFl   float64
				dur    int
				short  bool
				err    error
			)
			limitOpts := options[0].Options

			for _, v := range limitOpts {
				switch v.Name {
				case "limit-price":
					lp = v.StringValue()
					lpFl, err = strconv.ParseFloat(lp, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = v.StringValue()
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(ticker, short, aries.ContractBuyOrSell(!close), aries.Limit, aries.OrderDuration(dur), lpFl, 0, 0, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
					},
				})
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: res,
				},
			})
		}
	},
}
