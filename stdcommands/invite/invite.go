package invite

import (
	"github.com/jonas747/dcmd"
	"github.com/mmuyskens/yagpdb/commands"
	"github.com/mmuyskens/yagpdb/common"
)

var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryGeneral,
	Name:        "Invite",
	Description: "Responds with bot invite link",
	RunInDM:     true,

	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		return "Please add the bot through the website\nhttps://" + common.ConfHost.GetString(), nil
	},
}
