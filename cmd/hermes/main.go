/*
 * Copyright 2022 M1K
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/m1k8/hermes/pkg/handlers"
)

var token string
var s *discordgo.Session

func init() {
	var err error
	token = "OTcxNTE4Mjc3MDczMzgzNDI0.YnLq5w.5T2S7TZSLlpcIlY0XUHdZzcZeYY"
	s, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		//case discordgo.InteractionMessageComponent:
		//	if h, ok := handlers.ComponentHandlers[i.MessageComponentData().CustomID]; ok {
		//		h(s, i)
		//	}

		case discordgo.InteractionApplicationCommand:
			if h, ok := handlers.CommandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		}
	})
}

func main() {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
	} else {
		multi := io.MultiWriter(file, os.Stdout)
		log.SetOutput(multi)
	}

	//s.LogLevel = discordgo.LogInformational

	// We need information about servers (which includes their channels),
	// messages and voice states.
	s.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	// Open the websocket and begin listening.
	err = s.Open()
	if err != nil {
		log.Println("Error opening Discord session: ", err)
		return
	}
	for _, v := range handlers.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		log.Println("Created " + v.Name)
	}

	log.Println("Ready!")

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
		<-sc
		s.Close()
	}()
}
