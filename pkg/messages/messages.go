package messages

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var footerText = "There is risk involved in all trade ideas, please do your own due diligence. This is not financial advice"

func prettifyOptionsStr(opts string) string {
	optsSlice := strings.Split(opts, ",")

	//S!F_220112C11,
	resStr := ""

	for _, v := range optsSlice {
		// remove first 2 chars
		tmpV := v[2:]
		tickerSplit := strings.Split(tmpV, "_")
		//F|_ |220112C11

		ticker := tickerSplit[0]

		// date format is consistent, so plit like this: 220112 | C11

		date := tickerSplit[1][:6]
		m := date[2:4]
		d := date[4:]
		y := "20" + date[0:2]

		strike := tickerSplit[1][6:]

		if resStr != "" {
			resStr += ", "
		}
		resStr += ticker + " " + m + "/" + d + "/" + y + " " + strike
	}

	return resStr
}

func sendEmbed(s *discordgo.Session, thisGuild, channelToSendTo, roleToSendTo string, e *discordgo.MessageEmbed) (msg *discordgo.Message, err error) {
	ms := &discordgo.MessageSend{
		Embed: e,
	}

	msg, err = s.ChannelMessageSendComplex(channelToSendTo, ms)
	return
}

func GetStockEmbed(close bool, alerter, ticker, limit, profit_limit, stop_loss, stop_price, trail_amt, trail_pct string) *discordgo.MessageEmbed {
	var colour int
	if close {
		colour = 0xff0000
	} else {
		colour = 0x00ff00
	}
	embed := &discordgo.MessageEmbed{
		Title:     "Stock alert for **" + ticker + "**",
		Author:    &discordgo.MessageEmbedAuthor{},
		Color:     colour,
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text:    footerText,
			IconURL: "https://i.imgur.com/4RzvZL1.jpg",
		},
	}

	if limit != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "💰 Limit Price: *__$" + limit + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if profit_limit != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "💹 Profit Limit: *__$" + profit_limit + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if stop_loss != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "📉 Stop Loss: *__$" + stop_loss + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if stop_price != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "🚫 Stop Price: *__$" + stop_price + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if trail_amt != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "🔢 Stop Trailing Amount: *__$" + trail_amt + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if trail_pct != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "🔢 Stop Trailing Percent: *__" + trail_pct + "%" + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	return embed
}

func GetOptionsEmbed(close bool, alerter, con, limit, profit_limit, stop_loss, stop_price, trail_amt, trail_pct string) *discordgo.MessageEmbed {
	var colour int
	if close {
		colour = 0xff0000
	} else {
		colour = 0x00ff00
	}

	con = prettifyOptionsStr(con)
	embed := &discordgo.MessageEmbed{
		Title:     "Contract alert for **" + con + "**",
		Author:    &discordgo.MessageEmbedAuthor{},
		Color:     colour,
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text:    footerText,
			IconURL: "https://i.imgur.com/4RzvZL1.jpg",
		},
	}

	if limit != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "💰 Limit Price: *__$" + limit + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if profit_limit != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "💹 Profit Limit: *__$" + profit_limit + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if stop_loss != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "📉 Stop Loss: *__$" + stop_loss + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if stop_price != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "🚫 Stop Price: *__$" + stop_price + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if trail_amt != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "🔢 Stop Trailing Amount: *__$" + trail_amt + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	if trail_pct != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "🔢 Stop Trailing Percent: *__" + trail_pct + "%" + "__*",
			Value:  "-------------------------------",
			Inline: false,
		})
	}

	return embed
}
