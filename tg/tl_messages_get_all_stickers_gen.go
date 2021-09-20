// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

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
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// MessagesGetAllStickersRequest represents TL type `messages.getAllStickers#b8a0a1a8`.
// Get all installed stickers
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
type MessagesGetAllStickersRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetAllStickersRequestTypeID is TL type id of MessagesGetAllStickersRequest.
const MessagesGetAllStickersRequestTypeID = 0xb8a0a1a8

func (g *MessagesGetAllStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAllStickersRequest) String() string {
	if g == nil {
		return "MessagesGetAllStickersRequest(nil)"
	}
	type Alias MessagesGetAllStickersRequest
	return fmt.Sprintf("MessagesGetAllStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetAllStickersRequest from given interface.
func (g *MessagesGetAllStickersRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAllStickersRequest) TypeID() uint32 {
	return MessagesGetAllStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAllStickersRequest) TypeName() string {
	return "messages.getAllStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAllStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAllStickers",
		ID:   MessagesGetAllStickersRequestTypeID,
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
func (g *MessagesGetAllStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAllStickers#b8a0a1a8 as nil")
	}
	b.PutID(MessagesGetAllStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetAllStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAllStickers#b8a0a1a8 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetAllStickersRequest) GetHash() (value int64) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAllStickers#b8a0a1a8 to nil")
	}
	if err := b.ConsumeID(MessagesGetAllStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAllStickers#b8a0a1a8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetAllStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAllStickers#b8a0a1a8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAllStickers#b8a0a1a8: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetAllStickersRequest{}
	_ bin.Decoder     = &MessagesGetAllStickersRequest{}
	_ bin.BareEncoder = &MessagesGetAllStickersRequest{}
	_ bin.BareDecoder = &MessagesGetAllStickersRequest{}
)

// MessagesGetAllStickers invokes method messages.getAllStickers#b8a0a1a8 returning error if any.
// Get all installed stickers
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
func (c *Client) MessagesGetAllStickers(ctx context.Context, hash int64) (MessagesAllStickersClass, error) {
	var result MessagesAllStickersBox

	request := &MessagesGetAllStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AllStickers, nil
}
