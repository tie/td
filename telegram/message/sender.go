package message

import (
	"context"
	"crypto/rand"
	"io"

	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
)

// Sender is a message sending helper.
type Sender struct {
	raw  tg.Invoker
	rand io.Reader

	uploader Uploader
	resolver peer.Resolver
}

// NewSender creates a new Sender.
func NewSender(raw tg.Invoker) *Sender {
	return &Sender{
		raw:      raw,
		rand:     rand.Reader,
		uploader: uploader.NewUploader(raw),
		resolver: peer.DefaultResolver(raw),
	}
}

// WithUploader sets file uploader to use.
func (s *Sender) WithUploader(u Uploader) *Sender {
	s.uploader = u
	return s
}

// WithResolver sets peer resolver to use.
func (s *Sender) WithResolver(resolver peer.Resolver) *Sender {
	s.resolver = resolver
	return s
}

// WithRand sets random ID source.
func (s *Sender) WithRand(r io.Reader) *Sender {
	s.rand = r
	return s
}

// ClearAllDrafts clears all drafts in all peers.
func (s *Sender) ClearAllDrafts(ctx context.Context) error {
	_, err := tg.MessagesClearAllDrafts(ctx, s.raw)
	return err
}

// sendMessage sends message to peer.
func (s *Sender) sendMessage(ctx context.Context, req *tg.MessagesSendMessageRequest) (tg.UpdatesClass, error) {
	if req.RandomID == 0 {
		id, err := crypto.RandInt64(s.rand)
		if err != nil {
			return nil, xerrors.Errorf("generate random_id: %w", err)
		}
		req.RandomID = id
	}

	return tg.MessagesSendMessage(ctx, s.raw, req)
}

// sendMedia sends message with single media to peer.
func (s *Sender) sendMedia(ctx context.Context, req *tg.MessagesSendMediaRequest) (tg.UpdatesClass, error) {
	if req.RandomID == 0 {
		id, err := crypto.RandInt64(s.rand)
		if err != nil {
			return nil, xerrors.Errorf("generate random_id: %w", err)
		}
		req.RandomID = id
	}

	return tg.MessagesSendMedia(ctx, s.raw, req)
}

// sendMultiMedia sends message with multiple media to peer.
func (s *Sender) sendMultiMedia(ctx context.Context, req *tg.MessagesSendMultiMediaRequest) (tg.UpdatesClass, error) {
	for i := range req.MultiMedia {
		id, err := crypto.RandInt64(s.rand)
		if err != nil {
			return nil, xerrors.Errorf("generate random_id: %w", err)
		}
		req.MultiMedia[i].RandomID = id
	}

	return tg.MessagesSendMultiMedia(ctx, s.raw, req)
}

// editMessage edits message.
func (s *Sender) editMessage(ctx context.Context, req *tg.MessagesEditMessageRequest) (tg.UpdatesClass, error) {
	return tg.MessagesEditMessage(ctx, s.raw, req)
}

// forwardMessages forwards message to peer.
func (s *Sender) forwardMessages(ctx context.Context, req *tg.MessagesForwardMessagesRequest) (tg.UpdatesClass, error) {
	req.RandomID = make([]int64, len(req.ID))
	for i := range req.RandomID {
		id, err := crypto.RandInt64(s.rand)
		if err != nil {
			return nil, xerrors.Errorf("generate random_id: %w", err)
		}
		req.RandomID[i] = id
	}

	return tg.MessagesForwardMessages(ctx, s.raw, req)
}

// startBot starts a conversation with a bot using a deep linking parameter.
func (s *Sender) startBot(ctx context.Context, req *tg.MessagesStartBotRequest) (tg.UpdatesClass, error) {
	if req.RandomID == 0 {
		id, err := crypto.RandInt64(s.rand)
		if err != nil {
			return nil, xerrors.Errorf("generate random_id: %w", err)
		}
		req.RandomID = id
	}

	return tg.MessagesStartBot(ctx, s.raw, req)
}

// sendInlineBotResult sends inline query result message to peer.
func (s *Sender) sendInlineBotResult(
	ctx context.Context,
	req *tg.MessagesSendInlineBotResultRequest,
) (tg.UpdatesClass, error) {
	if req.RandomID == 0 {
		id, err := crypto.RandInt64(s.rand)
		if err != nil {
			return nil, xerrors.Errorf("generate random_id: %w", err)
		}
		req.RandomID = id
	}

	return tg.MessagesSendInlineBotResult(ctx, s.raw, req)
}

// uploadMedia uploads file and associate it to a chat (without actually sending it to the chat).
func (s *Sender) uploadMedia(ctx context.Context, req *tg.MessagesUploadMediaRequest) (tg.MessageMediaClass, error) {
	return tg.MessagesUploadMedia(ctx, s.raw, req)
}

// getDocumentByHash finds document by hash, MIME type and size.
func (s *Sender) getDocumentByHash(
	ctx context.Context,
	req *tg.MessagesGetDocumentByHashRequest,
) (tg.DocumentClass, error) {
	return tg.MessagesGetDocumentByHash(ctx, s.raw, req)
}

// saveDraft saves a message draft associated to a chat.
func (s *Sender) saveDraft(ctx context.Context, req *tg.MessagesSaveDraftRequest) error {
	_, err := tg.MessagesSaveDraft(ctx, s.raw, req)
	return err
}

// sendVote votes in a poll.
func (s *Sender) sendVote(ctx context.Context, req *tg.MessagesSendVoteRequest) (tg.UpdatesClass, error) {
	return tg.MessagesSendVote(ctx, s.raw, req)
}

// setTyping sends a typing event to a conversation partner or group.
func (s *Sender) setTyping(ctx context.Context, req *tg.MessagesSetTypingRequest) error {
	_, err := tg.MessagesSetTyping(ctx, s.raw, req)
	return err
}

// report reports a message in a chat for violation of Telegram's Terms of Service.
func (s *Sender) report(ctx context.Context, req *tg.MessagesReportRequest) (bool, error) {
	return tg.MessagesReport(ctx, s.raw, req)
}

// reportSpam reports a new incoming chat for spam, if the peer settings of the chat allow us to do that.
func (s *Sender) reportSpam(ctx context.Context, p tg.InputPeerClass) (bool, error) {
	return tg.MessagesReportSpam(ctx, s.raw, p)
}

// getPeerSettings returns peer settings.
func (s *Sender) getPeerSettings(ctx context.Context, p tg.InputPeerClass) (*tg.PeerSettings, error) {
	return tg.MessagesGetPeerSettings(ctx, s.raw, p)
}

// sendScreenshotNotification sends notification about screenshot to peer.
func (s *Sender) sendScreenshotNotification(
	ctx context.Context,
	req *tg.MessagesSendScreenshotNotificationRequest,
) (tg.UpdatesClass, error) {
	if req.RandomID == 0 {
		id, err := crypto.RandInt64(s.rand)
		if err != nil {
			return nil, xerrors.Errorf("generate random_id: %w", err)
		}
		req.RandomID = id
	}

	return tg.MessagesSendScreenshotNotification(ctx, s.raw, req)
}

// sendScheduledMessages sends scheduled messages using given ids.
func (s *Sender) sendScheduledMessages(
	ctx context.Context,
	req *tg.MessagesSendScheduledMessagesRequest,
) (tg.UpdatesClass, error) {
	return tg.MessagesSendScheduledMessages(ctx, s.raw, req)
}

// deleteScheduledMessages deletes scheduled messages using given ids.
func (s *Sender) deleteScheduledMessages(
	ctx context.Context,
	req *tg.MessagesDeleteScheduledMessagesRequest,
) (tg.UpdatesClass, error) {
	return tg.MessagesDeleteScheduledMessages(ctx, s.raw, req)
}

// getScheduledHistory gets scheduled messages history.
func (s *Sender) getScheduledHistory(
	ctx context.Context,
	req *tg.MessagesGetScheduledHistoryRequest,
) (tg.MessagesMessagesClass, error) {
	return tg.MessagesGetScheduledHistory(ctx, s.raw, req)
}

// getScheduledMessages gets scheduled messages using given ids.
func (s *Sender) getScheduledMessages(
	ctx context.Context,
	req *tg.MessagesGetScheduledMessagesRequest,
) (tg.MessagesMessagesClass, error) {
	return tg.MessagesGetScheduledMessages(ctx, s.raw, req)
}
