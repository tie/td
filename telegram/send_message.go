package telegram

import (
	"context"

	"github.com/gotd/td/tg"
)

// SendMessage sends message to peer.
func (c *Client) SendMessage(ctx context.Context, req *tg.MessagesSendMessageRequest) error {
	if req.RandomID == 0 {
		id, err := c.RandInt64()
		if err != nil {
			return err
		}
		req.RandomID = id
	}
	updates, err := tg.MessagesSendMessage(ctx, c, req)
	if err != nil {
		return err
	}
	return c.processUpdates(updates)
}
