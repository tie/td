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

// ChannelsReadHistoryRequest represents TL type `channels.readHistory#cc104937`.
// Mark channel/supergroup¹ history as read
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.readHistory for reference.
type ChannelsReadHistoryRequest struct {
	// Channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// ID of message up to which messages should be marked as read
	MaxID int
}

// ChannelsReadHistoryRequestTypeID is TL type id of ChannelsReadHistoryRequest.
const ChannelsReadHistoryRequestTypeID = 0xcc104937

func (r *ChannelsReadHistoryRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Channel == nil) {
		return false
	}
	if !(r.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ChannelsReadHistoryRequest) String() string {
	if r == nil {
		return "ChannelsReadHistoryRequest(nil)"
	}
	type Alias ChannelsReadHistoryRequest
	return fmt.Sprintf("ChannelsReadHistoryRequest%+v", Alias(*r))
}

// FillFrom fills ChannelsReadHistoryRequest from given interface.
func (r *ChannelsReadHistoryRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetMaxID() (value int)
}) {
	r.Channel = from.GetChannel()
	r.MaxID = from.GetMaxID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsReadHistoryRequest) TypeID() uint32 {
	return ChannelsReadHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsReadHistoryRequest) TypeName() string {
	return "channels.readHistory"
}

// TypeInfo returns info about TL type.
func (r *ChannelsReadHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.readHistory",
		ID:   ChannelsReadHistoryRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ChannelsReadHistoryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.readHistory#cc104937 as nil")
	}
	b.PutID(ChannelsReadHistoryRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ChannelsReadHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.readHistory#cc104937 as nil")
	}
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.readHistory#cc104937: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.readHistory#cc104937: field channel: %w", err)
	}
	b.PutInt(r.MaxID)
	return nil
}

// GetChannel returns value of Channel field.
func (r *ChannelsReadHistoryRequest) GetChannel() (value InputChannelClass) {
	return r.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (r *ChannelsReadHistoryRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return r.Channel.AsNotEmpty()
}

// GetMaxID returns value of MaxID field.
func (r *ChannelsReadHistoryRequest) GetMaxID() (value int) {
	return r.MaxID
}

// Decode implements bin.Decoder.
func (r *ChannelsReadHistoryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.readHistory#cc104937 to nil")
	}
	if err := b.ConsumeID(ChannelsReadHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.readHistory#cc104937: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ChannelsReadHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.readHistory#cc104937 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.readHistory#cc104937: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.readHistory#cc104937: field max_id: %w", err)
		}
		r.MaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsReadHistoryRequest.
var (
	_ bin.Encoder     = &ChannelsReadHistoryRequest{}
	_ bin.Decoder     = &ChannelsReadHistoryRequest{}
	_ bin.BareEncoder = &ChannelsReadHistoryRequest{}
	_ bin.BareDecoder = &ChannelsReadHistoryRequest{}
)

// ChannelsReadHistory invokes method channels.readHistory#cc104937 returning error if any.
// Mark channel/supergroup¹ history as read
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/channels.readHistory for reference.
func ChannelsReadHistory(ctx context.Context, rpc Invoker, request *ChannelsReadHistoryRequest) (bool, error) {
	var result BoolBox

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
