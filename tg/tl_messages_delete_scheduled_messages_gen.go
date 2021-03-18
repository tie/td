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

// MessagesDeleteScheduledMessagesRequest represents TL type `messages.deleteScheduledMessages#59ae2b16`.
// Delete scheduled messages
//
// See https://core.telegram.org/method/messages.deleteScheduledMessages for reference.
type MessagesDeleteScheduledMessagesRequest struct {
	// Peer
	Peer InputPeerClass
	// Scheduled message IDs
	ID []int
}

// MessagesDeleteScheduledMessagesRequestTypeID is TL type id of MessagesDeleteScheduledMessagesRequest.
const MessagesDeleteScheduledMessagesRequestTypeID = 0x59ae2b16

func (d *MessagesDeleteScheduledMessagesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteScheduledMessagesRequest) String() string {
	if d == nil {
		return "MessagesDeleteScheduledMessagesRequest(nil)"
	}
	type Alias MessagesDeleteScheduledMessagesRequest
	return fmt.Sprintf("MessagesDeleteScheduledMessagesRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteScheduledMessagesRequest from given interface.
func (d *MessagesDeleteScheduledMessagesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
}) {
	d.Peer = from.GetPeer()
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteScheduledMessagesRequest) TypeID() uint32 {
	return MessagesDeleteScheduledMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteScheduledMessagesRequest) TypeName() string {
	return "messages.deleteScheduledMessages"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteScheduledMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteScheduledMessages",
		ID:   MessagesDeleteScheduledMessagesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteScheduledMessagesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteScheduledMessages#59ae2b16 as nil")
	}
	b.PutID(MessagesDeleteScheduledMessagesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteScheduledMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteScheduledMessages#59ae2b16 as nil")
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode messages.deleteScheduledMessages#59ae2b16: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteScheduledMessages#59ae2b16: field peer: %w", err)
	}
	b.PutVectorHeader(len(d.ID))
	for _, v := range d.ID {
		b.PutInt(v)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *MessagesDeleteScheduledMessagesRequest) GetPeer() (value InputPeerClass) {
	return d.Peer
}

// GetID returns value of ID field.
func (d *MessagesDeleteScheduledMessagesRequest) GetID() (value []int) {
	return d.ID
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteScheduledMessagesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteScheduledMessages#59ae2b16 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteScheduledMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteScheduledMessages#59ae2b16: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteScheduledMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteScheduledMessages#59ae2b16 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteScheduledMessages#59ae2b16: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteScheduledMessages#59ae2b16: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.deleteScheduledMessages#59ae2b16: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDeleteScheduledMessagesRequest.
var (
	_ bin.Encoder     = &MessagesDeleteScheduledMessagesRequest{}
	_ bin.Decoder     = &MessagesDeleteScheduledMessagesRequest{}
	_ bin.BareEncoder = &MessagesDeleteScheduledMessagesRequest{}
	_ bin.BareDecoder = &MessagesDeleteScheduledMessagesRequest{}
)

// MessagesDeleteScheduledMessages invokes method messages.deleteScheduledMessages#59ae2b16 returning error if any.
// Delete scheduled messages
//
// See https://core.telegram.org/method/messages.deleteScheduledMessages for reference.
func MessagesDeleteScheduledMessages(ctx context.Context, rpc Invoker, request *MessagesDeleteScheduledMessagesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
