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

// MessagesUnpinAllMessagesRequest represents TL type `messages.unpinAllMessages#f025bc8b`.
// Unpin¹ all pinned messages
//
// Links:
//  1) https://core.telegram.org/api/pin
//
// See https://core.telegram.org/method/messages.unpinAllMessages for reference.
type MessagesUnpinAllMessagesRequest struct {
	// Chat where to unpin
	Peer InputPeerClass
}

// MessagesUnpinAllMessagesRequestTypeID is TL type id of MessagesUnpinAllMessagesRequest.
const MessagesUnpinAllMessagesRequestTypeID = 0xf025bc8b

func (u *MessagesUnpinAllMessagesRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUnpinAllMessagesRequest) String() string {
	if u == nil {
		return "MessagesUnpinAllMessagesRequest(nil)"
	}
	type Alias MessagesUnpinAllMessagesRequest
	return fmt.Sprintf("MessagesUnpinAllMessagesRequest%+v", Alias(*u))
}

// FillFrom fills MessagesUnpinAllMessagesRequest from given interface.
func (u *MessagesUnpinAllMessagesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	u.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUnpinAllMessagesRequest) TypeID() uint32 {
	return MessagesUnpinAllMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUnpinAllMessagesRequest) TypeName() string {
	return "messages.unpinAllMessages"
}

// TypeInfo returns info about TL type.
func (u *MessagesUnpinAllMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.unpinAllMessages",
		ID:   MessagesUnpinAllMessagesRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *MessagesUnpinAllMessagesRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.unpinAllMessages#f025bc8b as nil")
	}
	b.PutID(MessagesUnpinAllMessagesRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *MessagesUnpinAllMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.unpinAllMessages#f025bc8b as nil")
	}
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#f025bc8b: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#f025bc8b: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (u *MessagesUnpinAllMessagesRequest) GetPeer() (value InputPeerClass) {
	return u.Peer
}

// Decode implements bin.Decoder.
func (u *MessagesUnpinAllMessagesRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.unpinAllMessages#f025bc8b to nil")
	}
	if err := b.ConsumeID(MessagesUnpinAllMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.unpinAllMessages#f025bc8b: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *MessagesUnpinAllMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.unpinAllMessages#f025bc8b to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.unpinAllMessages#f025bc8b: field peer: %w", err)
		}
		u.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUnpinAllMessagesRequest.
var (
	_ bin.Encoder     = &MessagesUnpinAllMessagesRequest{}
	_ bin.Decoder     = &MessagesUnpinAllMessagesRequest{}
	_ bin.BareEncoder = &MessagesUnpinAllMessagesRequest{}
	_ bin.BareDecoder = &MessagesUnpinAllMessagesRequest{}
)

// MessagesUnpinAllMessages invokes method messages.unpinAllMessages#f025bc8b returning error if any.
// Unpin¹ all pinned messages
//
// Links:
//  1) https://core.telegram.org/api/pin
//
// See https://core.telegram.org/method/messages.unpinAllMessages for reference.
// Can be used by bots.
func MessagesUnpinAllMessages(ctx context.Context, rpc Invoker, peer InputPeerClass) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	request := &MessagesUnpinAllMessagesRequest{
		Peer: peer,
	}
	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
