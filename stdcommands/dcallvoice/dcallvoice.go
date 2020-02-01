package dcallvoice

import (
	"fmt"

	"github.com/jonas747/dcmd"
	"github.com/jonas747/discordgo"
	"github.com/mmuyskens/yagpdb/bot"
	"github.com/mmuyskens/yagpdb/commands"
	"github.com/mmuyskens/yagpdb/common"
	"github.com/mmuyskens/yagpdb/stdcommands/util"
)

var Command = &commands.YAGCommand{
	CmdCategory:          commands.CategoryDebug,
	HideFromCommandsPage: true,
	Name:                 "dcallvoice",
	Description:          "Disconnects from all the voice channels the bot is in",
	HideFromHelp:         true,
	RunFunc: util.RequireBotAdmin(func(data *dcmd.Data) (interface{}, error) {

		vcs := make([]*discordgo.VoiceState, 0)

		guilds := bot.State.GuildsSlice(true)

		for _, g := range guilds {
			vc := g.VoiceState(true, common.BotUser.ID)
			if vc != nil {
				vcs = append(vcs, vc)
				go bot.ShardManager.SessionForGuild(g.ID).GatewayManager.ChannelVoiceLeave(g.ID)
			}
		}

		return fmt.Sprintf("Leaving %d voice channels...", len(vcs)), nil
	}),
}
