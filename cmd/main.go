package main

import (
	"fmt"
	"github.com/alknopfler/ztpfw-bot-slack/pkg/eventHandler"
	"github.com/alknopfler/ztpfw-bot-slack/pkg/utils"
	"github.com/slack-go/slack"
	"os"
	"regexp"
	"strings"
)

func main() {

	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Println("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				if ev.SubType == "message_changed" {
					fmt.Println("Message edited... ignoring")
					continue
				}

				// Get conversation information
				channelInformation, err := rtm.GetConversationInfo(ev.Channel, true)
				if err != nil {
					fmt.Println("Ignoring message since we cannot get channel information")
					continue
				}

				channelName := channelInformation.NameNormalized
				text := strings.TrimSpace(strings.ToLower(ev.Text))

				// Commands are implemented using a keyword rather than using slash commands to avoid
				// having to publish the bot in order to receive webhooks
				r := regexp.MustCompile("^(ztpfw) (get|set|help) (status|motd|help)(.*)$")
				if r.MatchString(text) {
					captureGroups := r.FindStringSubmatch(text)
					operation := captureGroups[2]
					operationGroup := captureGroups[3]
					operationArgs := captureGroups[4]
					userID := strings.ToLower(ev.User)
					if operation == "help" || operationGroup == "help" {
						fmt.Println("Printing help on channel %s", channelName)
						utils.PrintCommandsUsage(rtm, ev)
						break
					}
					// add user that fires the command to the args
					commandOutput := eventHandler.ProcessCommand(channelName, userID, operation, operationGroup, operationArgs)
					rtm.SendMessage(rtm.NewOutgoingMessage(commandOutput, ev.Channel))

				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
				break

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break

			default:
				//Take no action
			}
		}
	}

}
