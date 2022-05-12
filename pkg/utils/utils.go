package utils

import "github.com/slack-go/slack"

const (
	ZTPFW_HELP = "*ZTPFW Bot Commands*:\n" +
		"- Get the hosts status for the ztpfw lab: `ztpfw get status <host>`\n" +
		"- Set the status for one of the ztpfw lab host: `ztpfw set motd <host> <user_id> <pr>`\n" +
		"- help: `ztpfw help`\n"
)

func PrintCommandsUsage(rtm *slack.RTM, ev *slack.MessageEvent) {
	rtm.SendMessage(rtm.NewOutgoingMessage(ZTPFW_HELP, ev.Channel))
}
