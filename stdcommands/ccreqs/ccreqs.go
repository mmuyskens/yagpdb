package ccreqs

import (
	"fmt"

	"github.com/jonas747/dcmd"
	"github.com/mmuyskens/yagpdb/commands"
	"github.com/mmuyskens/yagpdb/common"
	"github.com/mmuyskens/yagpdb/stdcommands/util"
)

var Command = &commands.YAGCommand{
	CmdCategory:          commands.CategoryDebug,
	HideFromCommandsPage: true,
	Name:                 "ccreqs",
	Description:          "Returns the number of concurrent requests currently going on",
	HideFromHelp:         true,
	RunFunc: util.RequireBotAdmin(func(data *dcmd.Data) (interface{}, error) {
		return fmt.Sprintf("`%d`", common.BotSession.Ratelimiter.CurrentConcurrentLocks()), nil
	}),
}
