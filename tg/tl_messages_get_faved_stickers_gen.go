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

// MessagesGetFavedStickersRequest represents TL type `messages.getFavedStickers#21ce0b0e`.
// Get faved stickers
//
// See https://core.telegram.org/method/messages.getFavedStickers for reference.
type MessagesGetFavedStickersRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetFavedStickersRequestTypeID is TL type id of MessagesGetFavedStickersRequest.
const MessagesGetFavedStickersRequestTypeID = 0x21ce0b0e

func (g *MessagesGetFavedStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetFavedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetFavedStickersRequest(nil)"
	}
	type Alias MessagesGetFavedStickersRequest
	return fmt.Sprintf("MessagesGetFavedStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetFavedStickersRequest from given interface.
func (g *MessagesGetFavedStickersRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetFavedStickersRequest) TypeID() uint32 {
	return MessagesGetFavedStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetFavedStickersRequest) TypeName() string {
	return "messages.getFavedStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetFavedStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getFavedStickers",
		ID:   MessagesGetFavedStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetFavedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFavedStickers#21ce0b0e as nil")
	}
	b.PutID(MessagesGetFavedStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetFavedStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFavedStickers#21ce0b0e as nil")
	}
	b.PutInt(g.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetFavedStickersRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetFavedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFavedStickers#21ce0b0e to nil")
	}
	if err := b.ConsumeID(MessagesGetFavedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getFavedStickers#21ce0b0e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetFavedStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFavedStickers#21ce0b0e to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getFavedStickers#21ce0b0e: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetFavedStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetFavedStickersRequest{}
	_ bin.Decoder     = &MessagesGetFavedStickersRequest{}
	_ bin.BareEncoder = &MessagesGetFavedStickersRequest{}
	_ bin.BareDecoder = &MessagesGetFavedStickersRequest{}
)

// MessagesGetFavedStickers invokes method messages.getFavedStickers#21ce0b0e returning error if any.
// Get faved stickers
//
// See https://core.telegram.org/method/messages.getFavedStickers for reference.
func MessagesGetFavedStickers(ctx context.Context, rpc Invoker, hash int) (MessagesFavedStickersClass, error) {
	var result MessagesFavedStickersBox

	request := &MessagesGetFavedStickersRequest{
		Hash: hash,
	}
	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.FavedStickers, nil
}
