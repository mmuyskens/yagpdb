package toggledbg

import (
	"github.com/mmuyskens/yagpdb/common"
	"github.com/sirupsen/logrus"

	"github.com/jonas747/dcmd"
	"github.com/mmuyskens/yagpdb/commands"
	"github.com/mmuyskens/yagpdb/stdcommands/util"
)

var Command = &commands.YAGCommand{
	CmdCategory:          commands.CategoryDebug,
	HideFromCommandsPage: true,
	Name:                 "toggledbg",
	Description:          "Toggles Debug Logging",
	HideFromHelp:         true,
	RunFunc: util.RequireOwner(func(data *dcmd.Data) (interface{}, error) {
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			common.SetLoggingLevel(logrus.InfoLevel)
			return "Disabled debug logging", nil
		}

		common.SetLoggingLevel(logrus.DebugLevel)
		return "Enabled debug logging", nil

	}),
}
