package createinvite

import (
	"github.com/jonas747/dcmd"
	"github.com/jonas747/discordgo"
	"github.com/mmuyskens/yagpdb/bot"
	"github.com/mmuyskens/yagpdb/commands"
	"github.com/mmuyskens/yagpdb/common"
	"github.com/mmuyskens/yagpdb/stdcommands/util"
)

var Command = &commands.YAGCommand{
	Cooldown:             2,
	CmdCategory:          commands.CategoryDebug,
	HideFromCommandsPage: true,
	Name:                 "createinvite",
	Description:          "Maintenance command, creates a invite for the specified server",
	HideFromHelp:         true,
	RequiredArgs:         1,
	Arguments: []*dcmd.ArgDef{
		{Name: "server", Type: dcmd.Int},
	},
	RunFunc: util.RequireBotAdmin(func(data *dcmd.Data) (interface{}, error) {
		channels, err := common.BotSession.GuildChannels(data.Args[0].Int64())
		if err != nil {
			return nil, err
		}

		channelID := int64(0)
		for _, v := range channels {
			if channelID == 0 || v.Type != discordgo.ChannelTypeGuildVoice {
				channelID = v.ID
				if v.Type != discordgo.ChannelTypeGuildVoice {
					break
				}
			}
		}

		if channelID == 0 {
			return "No possible channel :(", nil
		}

		invite, err := common.BotSession.ChannelInviteCreate(channelID, discordgo.Invite{
			MaxAge:  120,
			MaxUses: 1,
		})

		if err != nil {
			return nil, err
		}

		bot.SendDM(data.Msg.Author.ID, "discord.gg/"+invite.Code)
		return "Sent invite expiring in 120 seconds and with 1 use in DM", nil
	}),
}
