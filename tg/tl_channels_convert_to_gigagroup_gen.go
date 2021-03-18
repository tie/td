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

// ChannelsConvertToGigagroupRequest represents TL type `channels.convertToGigagroup#b290c69`.
//
// See https://core.telegram.org/method/channels.convertToGigagroup for reference.
type ChannelsConvertToGigagroupRequest struct {
	// Channel field of ChannelsConvertToGigagroupRequest.
	Channel InputChannelClass
}

// ChannelsConvertToGigagroupRequestTypeID is TL type id of ChannelsConvertToGigagroupRequest.
const ChannelsConvertToGigagroupRequestTypeID = 0xb290c69

func (c *ChannelsConvertToGigagroupRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsConvertToGigagroupRequest) String() string {
	if c == nil {
		return "ChannelsConvertToGigagroupRequest(nil)"
	}
	type Alias ChannelsConvertToGigagroupRequest
	return fmt.Sprintf("ChannelsConvertToGigagroupRequest%+v", Alias(*c))
}

// FillFrom fills ChannelsConvertToGigagroupRequest from given interface.
func (c *ChannelsConvertToGigagroupRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	c.Channel = from.GetChannel()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsConvertToGigagroupRequest) TypeID() uint32 {
	return ChannelsConvertToGigagroupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsConvertToGigagroupRequest) TypeName() string {
	return "channels.convertToGigagroup"
}

// TypeInfo returns info about TL type.
func (c *ChannelsConvertToGigagroupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.convertToGigagroup",
		ID:   ChannelsConvertToGigagroupRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelsConvertToGigagroupRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.convertToGigagroup#b290c69 as nil")
	}
	b.PutID(ChannelsConvertToGigagroupRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelsConvertToGigagroupRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.convertToGigagroup#b290c69 as nil")
	}
	if c.Channel == nil {
		return fmt.Errorf("unable to encode channels.convertToGigagroup#b290c69: field channel is nil")
	}
	if err := c.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.convertToGigagroup#b290c69: field channel: %w", err)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (c *ChannelsConvertToGigagroupRequest) GetChannel() (value InputChannelClass) {
	return c.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (c *ChannelsConvertToGigagroupRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return c.Channel.AsNotEmpty()
}

// Decode implements bin.Decoder.
func (c *ChannelsConvertToGigagroupRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.convertToGigagroup#b290c69 to nil")
	}
	if err := b.ConsumeID(ChannelsConvertToGigagroupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.convertToGigagroup#b290c69: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelsConvertToGigagroupRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.convertToGigagroup#b290c69 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.convertToGigagroup#b290c69: field channel: %w", err)
		}
		c.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsConvertToGigagroupRequest.
var (
	_ bin.Encoder     = &ChannelsConvertToGigagroupRequest{}
	_ bin.Decoder     = &ChannelsConvertToGigagroupRequest{}
	_ bin.BareEncoder = &ChannelsConvertToGigagroupRequest{}
	_ bin.BareDecoder = &ChannelsConvertToGigagroupRequest{}
)

// ChannelsConvertToGigagroup invokes method channels.convertToGigagroup#b290c69 returning error if any.
//
// See https://core.telegram.org/method/channels.convertToGigagroup for reference.
func ChannelsConvertToGigagroup(ctx context.Context, rpc Invoker, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsConvertToGigagroupRequest{
		Channel: channel,
	}
	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
