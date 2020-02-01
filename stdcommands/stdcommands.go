package stdcommands

import (
	"github.com/mmuyskens/yagpdb/bot"
	"github.com/mmuyskens/yagpdb/bot/eventsystem"
	"github.com/mmuyskens/yagpdb/commands"
	"github.com/mmuyskens/yagpdb/common"
	"github.com/mmuyskens/yagpdb/stdcommands/advice"
	"github.com/mmuyskens/yagpdb/stdcommands/allocstat"
	"github.com/mmuyskens/yagpdb/stdcommands/banserver"
	"github.com/mmuyskens/yagpdb/stdcommands/calc"
	"github.com/mmuyskens/yagpdb/stdcommands/catfact"
	"github.com/mmuyskens/yagpdb/stdcommands/ccreqs"
	"github.com/mmuyskens/yagpdb/stdcommands/createinvite"
	"github.com/mmuyskens/yagpdb/stdcommands/currentshard"
	"github.com/mmuyskens/yagpdb/stdcommands/currenttime"
	"github.com/mmuyskens/yagpdb/stdcommands/customembed"
	"github.com/mmuyskens/yagpdb/stdcommands/dcallvoice"
	"github.com/mmuyskens/yagpdb/stdcommands/define"
	"github.com/mmuyskens/yagpdb/stdcommands/dogfact"
	"github.com/mmuyskens/yagpdb/stdcommands/findserver"
	"github.com/mmuyskens/yagpdb/stdcommands/globalrl"
	"github.com/mmuyskens/yagpdb/stdcommands/info"
	"github.com/mmuyskens/yagpdb/stdcommands/invite"
	"github.com/mmuyskens/yagpdb/stdcommands/leaveserver"
	"github.com/mmuyskens/yagpdb/stdcommands/listroles"
	"github.com/mmuyskens/yagpdb/stdcommands/memberfetcher"
	"github.com/mmuyskens/yagpdb/stdcommands/mentionrole"
	"github.com/mmuyskens/yagpdb/stdcommands/ping"
	"github.com/mmuyskens/yagpdb/stdcommands/poll"
	"github.com/mmuyskens/yagpdb/stdcommands/reverse"
	"github.com/mmuyskens/yagpdb/stdcommands/roll"
	"github.com/mmuyskens/yagpdb/stdcommands/setstatus"
	"github.com/mmuyskens/yagpdb/stdcommands/simpleembed"
	"github.com/mmuyskens/yagpdb/stdcommands/sleep"
	"github.com/mmuyskens/yagpdb/stdcommands/stateinfo"
	"github.com/mmuyskens/yagpdb/stdcommands/throw"
	"github.com/mmuyskens/yagpdb/stdcommands/toggledbg"
	"github.com/mmuyskens/yagpdb/stdcommands/topcommands"
	"github.com/mmuyskens/yagpdb/stdcommands/topevents"
	"github.com/mmuyskens/yagpdb/stdcommands/topgames"
	"github.com/mmuyskens/yagpdb/stdcommands/topic"
	"github.com/mmuyskens/yagpdb/stdcommands/topservers"
	"github.com/mmuyskens/yagpdb/stdcommands/unbanserver"
	"github.com/mmuyskens/yagpdb/stdcommands/undelete"
	"github.com/mmuyskens/yagpdb/stdcommands/viewperms"
	"github.com/mmuyskens/yagpdb/stdcommands/weather"
	"github.com/mmuyskens/yagpdb/stdcommands/wouldyourather"
	"github.com/mmuyskens/yagpdb/stdcommands/xkcd"
	"github.com/mmuyskens/yagpdb/stdcommands/yagstatus"
)

var (
	_ bot.BotInitHandler       = (*Plugin)(nil)
	_ commands.CommandProvider = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "Standard Commands",
		SysName:  "standard_commands",
		Category: common.PluginCategoryCore,
	}
}

func (p *Plugin) AddCommands() {
	commands.AddRootCommands(
		// Info
		info.Command,
		invite.Command,

		// Standard
		define.Command,
		reverse.Command,
		weather.Command,
		calc.Command,
		topic.Command,
		catfact.Command,
		dogfact.Command,
		advice.Command,
		ping.Command,
		throw.Command,
		roll.Command,
		customembed.Command,
		simpleembed.Command,
		currenttime.Command,
		mentionrole.Command,
		listroles.Command,
		wouldyourather.Command,
		poll.Command,
		undelete.Command,
		viewperms.Command,
		topgames.Command,
		xkcd.Command,

		// Maintenance
		stateinfo.Command,
		leaveserver.Command,
		banserver.Command,
		allocstat.Command,
		unbanserver.Command,
		topservers.Command,
		topcommands.Command,
		topevents.Command,
		currentshard.Command,
		memberfetcher.Command,
		yagstatus.Command,
		setstatus.Command,
		createinvite.Command,
		findserver.Command,
		dcallvoice.Command,
		ccreqs.Command,
		sleep.Command,
		toggledbg.Command,
		globalrl.Command,
	)

}

func (p *Plugin) BotInit() {
	eventsystem.AddHandlerAsyncLastLegacy(p, ping.HandleMessageCreate, eventsystem.EventMessageCreate)
	mentionrole.AddScheduledEventListener()
}

func RegisterPlugin() {
	common.RegisterPlugin(&Plugin{})
}
