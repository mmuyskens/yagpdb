package tickets

import (
	"emperror.dev/errors"
	"github.com/mmuyskens/yagpdb/bot"
	"github.com/mmuyskens/yagpdb/bot/eventsystem"
	"github.com/mmuyskens/yagpdb/common"
	"github.com/mmuyskens/yagpdb/tickets/models"
)

var _ bot.BotInitHandler = (*Plugin)(nil)

func (p *Plugin) BotInit() {
	eventsystem.AddHandlerAsyncLast(p, p.handleChannelRemoved, eventsystem.EventChannelDelete)
}

func (p *Plugin) handleChannelRemoved(evt *eventsystem.EventData) (retry bool, err error) {
	del := evt.ChannelDelete()

	_, err = models.Tickets(
		models.TicketWhere.ChannelID.EQ(del.Channel.ID),
	).DeleteAll(evt.Context(), common.PQ)

	if err != nil {
		return true, errors.WithStackIf(err)
	}

	return false, nil
}
