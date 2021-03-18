// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// ChannelsEditTitleRequest represents TL type `channels.editTitle#566decd0`.
// Edit the name of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.editTitle for reference.
type ChannelsEditTitleRequest struct {
	// Channel/supergroup
	Channel InputChannelClass
	// New name
	Title string
}

// ChannelsEditTitleRequestTypeID is TL type id of ChannelsEditTitleRequest.
const ChannelsEditTitleRequestTypeID = 0x566decd0

func (e *ChannelsEditTitleRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsEditTitleRequest) String() string {
	if e == nil {
		return "ChannelsEditTitleRequest(nil)"
	}
	type Alias ChannelsEditTitleRequest
	return fmt.Sprintf("ChannelsEditTitleRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsEditTitleRequest from given interface.
func (e *ChannelsEditTitleRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetTitle() (value string)
}) {
	e.Channel = from.GetChannel()
	e.Title = from.GetTitle()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsEditTitleRequest) TypeID() uint32 {
	return ChannelsEditTitleRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsEditTitleRequest) TypeName() string {
	return "channels.editTitle"
}

// TypeInfo returns info about TL type.
func (e *ChannelsEditTitleRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.editTitle",
		ID:   ChannelsEditTitleRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ChannelsEditTitleRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editTitle#566decd0 as nil")
	}
	b.PutID(ChannelsEditTitleRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChannelsEditTitleRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editTitle#566decd0 as nil")
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editTitle#566decd0: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editTitle#566decd0: field channel: %w", err)
	}
	b.PutString(e.Title)
	return nil
}

// GetChannel returns value of Channel field.
func (e *ChannelsEditTitleRequest) GetChannel() (value InputChannelClass) {
	return e.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsEditTitleRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// GetTitle returns value of Title field.
func (e *ChannelsEditTitleRequest) GetTitle() (value string) {
	return e.Title
}

// Decode implements bin.Decoder.
func (e *ChannelsEditTitleRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editTitle#566decd0 to nil")
	}
	if err := b.ConsumeID(ChannelsEditTitleRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editTitle#566decd0: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChannelsEditTitleRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editTitle#566decd0 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editTitle#566decd0: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editTitle#566decd0: field title: %w", err)
		}
		e.Title = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditTitleRequest.
var (
	_ bin.Encoder     = &ChannelsEditTitleRequest{}
	_ bin.Decoder     = &ChannelsEditTitleRequest{}
	_ bin.BareEncoder = &ChannelsEditTitleRequest{}
	_ bin.BareDecoder = &ChannelsEditTitleRequest{}
)

// ChannelsEditTitle invokes method channels.editTitle#566decd0 returning error if any.
// Edit the name of a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  400 CHAT_TITLE_EMPTY: No chat title provided
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//
// See https://core.telegram.org/method/channels.editTitle for reference.
// Can be used by bots.
func ChannelsEditTitle(ctx context.Context, rpc Invoker, request *ChannelsEditTitleRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
