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
	"fmt"
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

func main() {

	if token == "" {
		log.Println("No token provided. Please set it as the DISCORD_API environment variable")
		return
	}

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
	} else {
		multi := io.MultiWriter(file, os.Stdout)
		log.SetOutput(multi)
	}

	// We need information about servers (which includes their channels),
	// messages and voice states.
	s.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = s.Open()
	if err != nil {
		log.Println("Error opening Discord session: ", err)
		return
	}

	for _, v := range handlers.Commands {
		c, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}

		log.Println("Created " + c.Name)

		defer func() {
			//cleanup
			cmds, _ := s.ApplicationCommands(s.State.User.ID, "")
			for _, cmd := range cmds {
				err = s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
				if err != nil {
					log.Println(fmt.Errorf("error removing %v: %w", cmd.Name, err))
				} else {
					log.Println("Removed " + cmd.Name)
				}
			}
		}()
	}

	log.Println("Ready!")

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
		<-sc
		s.Close()
	}()
}
