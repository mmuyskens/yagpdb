package main

import (
	"github.com/mmuyskens/yagpdb/common/run"

	// Core yagpdb packages

	"github.com/mmuyskens/yagpdb/admin"
	"github.com/mmuyskens/yagpdb/bot/paginatedmessages"
	"github.com/mmuyskens/yagpdb/common/internalapi"
	"github.com/mmuyskens/yagpdb/common/scheduledevents2"

	// Plugin imports
	"github.com/mmuyskens/yagpdb/automod"
	"github.com/mmuyskens/yagpdb/automod_legacy"
	"github.com/mmuyskens/yagpdb/autorole"
	"github.com/mmuyskens/yagpdb/aylien"
	"github.com/mmuyskens/yagpdb/cah"
	"github.com/mmuyskens/yagpdb/commands"
	"github.com/mmuyskens/yagpdb/customcommands"
	"github.com/mmuyskens/yagpdb/discordlogger"
	"github.com/mmuyskens/yagpdb/logs"
	"github.com/mmuyskens/yagpdb/moderation"
	"github.com/mmuyskens/yagpdb/notifications"
	"github.com/mmuyskens/yagpdb/premium"
	"github.com/mmuyskens/yagpdb/premium/patreonpremiumsource"
	"github.com/mmuyskens/yagpdb/reddit"
	"github.com/mmuyskens/yagpdb/reminders"
	"github.com/mmuyskens/yagpdb/reputation"
	"github.com/mmuyskens/yagpdb/rolecommands"
	"github.com/mmuyskens/yagpdb/rsvp"
	"github.com/mmuyskens/yagpdb/safebrowsing"
	"github.com/mmuyskens/yagpdb/serverstats"
	"github.com/mmuyskens/yagpdb/soundboard"
	"github.com/mmuyskens/yagpdb/stdcommands"
	"github.com/mmuyskens/yagpdb/streaming"
	"github.com/mmuyskens/yagpdb/tickets"
	"github.com/mmuyskens/yagpdb/timezonecompanion"
	"github.com/mmuyskens/yagpdb/twitter"
	"github.com/mmuyskens/yagpdb/verification"
	"github.com/mmuyskens/yagpdb/youtube"
	// External plugins
)

func main() {

	run.Init()

	//BotSession.LogLevel = discordgo.LogInformational
	paginatedmessages.RegisterPlugin()

	// Setup plugins
	safebrowsing.RegisterPlugin()
	discordlogger.Register()
	commands.RegisterPlugin()
	stdcommands.RegisterPlugin()
	serverstats.RegisterPlugin()
	notifications.RegisterPlugin()
	customcommands.RegisterPlugin()
	reddit.RegisterPlugin()
	moderation.RegisterPlugin()
	reputation.RegisterPlugin()
	aylien.RegisterPlugin()
	streaming.RegisterPlugin()
	automod_legacy.RegisterPlugin()
	automod.RegisterPlugin()
	logs.RegisterPlugin()
	autorole.RegisterPlugin()
	reminders.RegisterPlugin()
	soundboard.RegisterPlugin()
	youtube.RegisterPlugin()
	rolecommands.RegisterPlugin()
	cah.RegisterPlugin()
	tickets.RegisterPlugin()
	verification.RegisterPlugin()
	premium.RegisterPlugin()
	patreonpremiumsource.RegisterPlugin()
	scheduledevents2.RegisterPlugin()
	twitter.RegisterPlugin()
	rsvp.RegisterPlugin()
	timezonecompanion.RegisterPlugin()
	admin.RegisterPlugin()
	internalapi.RegisterPlugin()

	run.Run()
}
