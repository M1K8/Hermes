package handlers

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/m1k8/hermes/pkg/aries"
	"github.com/m1k8/hermes/pkg/messages"
)

func getButton(res string) discordgo.Button {
	return discordgo.Button{
		Label: "Direct Trade Link",
		Style: discordgo.LinkButton,
		URL:   res,
	}
}

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
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Limit, aries.OrderDuration(dur), lpFl, 0, 0, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, lp, "", "", "", "", "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

			if err != nil {
				panic(err)
			}

		case "market":
			var (
				ticker string
				close  bool
				dur    int
				short  bool
				err    error
			)
			marketOpts := options[0].Options

			for _, v := range marketOpts {
				switch v.Name {
				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Market, aries.OrderDuration(dur), 0.0, 0, 0, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", "", "", "")
			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "stop-market":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
				dur    int
				short  bool
				err    error
			)
			limitOpts := options[0].Options

			for _, v := range limitOpts {
				switch v.Name {
				case "stop-price":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Stop_Market, aries.OrderDuration(dur), 0, 0, sFl, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", stop, "", "")
			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "stop-limit":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
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
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "stop-price":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Stop_Limit, aries.OrderDuration(dur), lpFl, 0, sFl, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, lp, "", stop, "", "", "")
			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})
		case "bracket":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
				lp     string
				lpFl   float64
				p      string
				pFl    float64
				dur    int
				short  bool
				err    error
			)
			bracketOpts := options[0].Options

			for _, v := range bracketOpts {
				switch v.Name {
				case "limit-price":
					lp = v.StringValue()
					lpFl, err = strconv.ParseFloat(lp, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "profit-limit":
					p = v.StringValue()
					pFl, err = strconv.ParseFloat(p, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "stop-loss":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Bracket, aries.OrderDuration(dur), lpFl, sFl, 0, 0, 0, pFl)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, lp, p, stop, "", "", "")
			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})
		case "bracket-pct":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
				lp     string
				lpFl   float64
				p      string
				pFl    float64
				dur    int
				short  bool
				err    error
			)
			bracketOpts := options[0].Options

			for _, v := range bracketOpts {
				switch v.Name {
				case "limit-pct":
					lp = v.StringValue()
					lpFl, err = strconv.ParseFloat(lp, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "profit-limit":
					p = v.StringValue()
					pFl, err = strconv.ParseFloat(p, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "stop-loss":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Bracket_Pct, aries.OrderDuration(dur), lpFl, sFl, 0, 0, 0, pFl)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, lp, p, stop, "", "", "")
			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "trailing-stop":
			var (
				ticker string
				close  bool
				ta     string
				taFl   float64
				dur    int
				short  bool
				err    error
			)
			limitOpts := options[0].Options

			for _, v := range limitOpts {
				switch v.Name {
				case "trail-amount":
					ta = v.StringValue()
					taFl, err = strconv.ParseFloat(ta, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}

				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Trailing_Stop, aries.OrderDuration(dur), 0, 0, 0, taFl, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", "", ta, "")
			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "trailing-percent":
			var (
				ticker string
				close  bool
				ta     string
				taFl   float64
				dur    int
				short  bool
				err    error
			)
			limitOpts := options[0].Options

			for _, v := range limitOpts {
				switch v.Name {
				case "trail-percent":
					ta = v.StringValue()
					taFl, err = strconv.ParseFloat(ta, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}

				case "duration":
					dur = int(v.IntValue())
				case "ticker":
					ticker = strings.ToUpper(v.StringValue())
				case "short":
					short = v.BoolValue()
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetStockUrl(strings.ToUpper(ticker), short, aries.ContractBuyOrSell(!close), aries.Trailing_Stop_Pct, aries.OrderDuration(dur), 0, 0, 0, 0, taFl, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetStockEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", "", "", ta)
			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		}
	},
	"options": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
					ticker = strings.Join(opts, ",")

				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Limit, aries.OrderDuration(dur), lpFl, 0, 0, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, lp, "", "", "", "", "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "market":
			var (
				ticker string
				close  bool
				dur    int
				err    error
			)
			marketOpts := options[0].Options

			for _, v := range marketOpts {
				switch v.Name {
				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
					ticker = strings.Join(opts, ",")

				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Market, aries.OrderDuration(dur), 0.0, 0, 0, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", "", "", "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "stop-market":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
				dur    int
				err    error
			)
			limitOpts := options[0].Options

			for _, v := range limitOpts {
				switch v.Name {
				case "stop-price":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
					ticker = strings.Join(opts, ",")
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Stop_Market, aries.OrderDuration(dur), 0, 0, sFl, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", stop, "", "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "stop-limit":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
				lp     string
				lpFl   float64
				dur    int
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
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "stop-price":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
					ticker = strings.Join(opts, ",")
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Stop_Limit, aries.OrderDuration(dur), lpFl, 0, sFl, 0, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, lp, "", "", stop, "", "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "bracket":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
				lp     string
				lpFl   float64
				p      string
				pFl    float64
				dur    int
				err    error
			)
			bracketOpts := options[0].Options

			for _, v := range bracketOpts {
				switch v.Name {
				case "limit-price":
					lp = v.StringValue()
					lpFl, err = strconv.ParseFloat(lp, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "profit-limit":
					p = v.StringValue()
					pFl, err = strconv.ParseFloat(p, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "stop-loss":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}

					ticker = opts[0]

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Bracket, aries.OrderDuration(dur), lpFl, sFl, 0, 0, 0, pFl)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, lp, p, stop, "", "", "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})
		case "bracket-pct":
			var (
				ticker string
				close  bool
				stop   string
				sFl    float64
				lp     string
				lpFl   float64
				p      string
				pFl    float64
				dur    int
				err    error
			)
			bracketOpts := options[0].Options

			for _, v := range bracketOpts {
				switch v.Name {
				case "limit-price":
					lp = v.StringValue()
					lpFl, err = strconv.ParseFloat(lp, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "limit-pct":
					p = v.StringValue()
					pFl, err = strconv.ParseFloat(p, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "stop-loss":
					stop = v.StringValue()
					sFl, err = strconv.ParseFloat(stop, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}

					ticker = opts[0]

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Bracket_Pct, aries.OrderDuration(dur), lpFl, sFl, 0, 0, 0, pFl)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, lp, p, stop, "", "", "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "trailing-stop":
			var (
				ticker string
				close  bool
				ta     string
				taFl   float64
				dur    int
				err    error
			)
			limitOpts := options[0].Options

			for _, v := range limitOpts {
				switch v.Name {
				case "trail-amount":
					ta = v.StringValue()
					taFl, err = strconv.ParseFloat(ta, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}

				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
					ticker = strings.Join(opts, ",")

				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Trailing_Stop, aries.OrderDuration(dur), 0, 0, 0, taFl, 0, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", "", ta, "")

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		case "trailing-percent":
			var (
				ticker string
				close  bool
				ta     string
				taFl   float64
				dur    int
				err    error
			)
			limitOpts := options[0].Options

			for _, v := range limitOpts {
				switch v.Name {
				case "trail-percent":
					ta = v.StringValue()
					taFl, err = strconv.ParseFloat(ta, 64)
					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}

				case "duration":
					dur = int(v.IntValue())
				case "contract-defs":
					ticker = strings.ToUpper(v.StringValue())
					opts, err := aries.ParseOptions(ticker)

					if err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: err.Error(),
								Flags:   1 << 6,
							},
						})
						return
					}
					ticker = strings.Join(opts, ",")

				case "close":
					close = v.BoolValue()
				}
			}

			res, err := gen.GetOptionsUrl(strings.ToUpper(ticker), aries.ContractBuyOrSell(!close), aries.Trailing_Stop_Pct, aries.OrderDuration(dur), 0, 0, 0, 0, taFl, 0)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: err.Error(),
						Flags:   1 << 6,
					},
				})
				return
			}

			embed := messages.GetOptionsEmbed(close, i.Interaction.Member.Mention(), ticker, "", "", "", "", "", ta)

			indicator := "BTO"

			if close {
				indicator = "STC"
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: indicator,
					Embeds: []*discordgo.MessageEmbed{
						embed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								getButton(res),
							},
						},
					},
				},
			})

		}
	},
}
