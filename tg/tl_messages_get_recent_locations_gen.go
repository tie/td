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

// MessagesGetRecentLocationsRequest represents TL type `messages.getRecentLocations#702a40e0`.
// Get live location history of a certain user
//
// See https://core.telegram.org/method/messages.getRecentLocations for reference.
type MessagesGetRecentLocationsRequest struct {
	// User
	Peer InputPeerClass
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetRecentLocationsRequestTypeID is TL type id of MessagesGetRecentLocationsRequest.
const MessagesGetRecentLocationsRequestTypeID = 0x702a40e0

func (g *MessagesGetRecentLocationsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetRecentLocationsRequest) String() string {
	if g == nil {
		return "MessagesGetRecentLocationsRequest(nil)"
	}
	type Alias MessagesGetRecentLocationsRequest
	return fmt.Sprintf("MessagesGetRecentLocationsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetRecentLocationsRequest from given interface.
func (g *MessagesGetRecentLocationsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetLimit() (value int)
	GetHash() (value int64)
}) {
	g.Peer = from.GetPeer()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetRecentLocationsRequest) TypeID() uint32 {
	return MessagesGetRecentLocationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetRecentLocationsRequest) TypeName() string {
	return "messages.getRecentLocations"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetRecentLocationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getRecentLocations",
		ID:   MessagesGetRecentLocationsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetRecentLocationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getRecentLocations#702a40e0 as nil")
	}
	b.PutID(MessagesGetRecentLocationsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetRecentLocationsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getRecentLocations#702a40e0 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getRecentLocations#702a40e0: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getRecentLocations#702a40e0: field peer: %w", err)
	}
	b.PutInt(g.Limit)
	b.PutLong(g.Hash)
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetRecentLocationsRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetLimit returns value of Limit field.
func (g *MessagesGetRecentLocationsRequest) GetLimit() (value int) {
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *MessagesGetRecentLocationsRequest) GetHash() (value int64) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetRecentLocationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getRecentLocations#702a40e0 to nil")
	}
	if err := b.ConsumeID(MessagesGetRecentLocationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getRecentLocations#702a40e0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetRecentLocationsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getRecentLocations#702a40e0 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentLocations#702a40e0: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentLocations#702a40e0: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentLocations#702a40e0: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetRecentLocationsRequest.
var (
	_ bin.Encoder     = &MessagesGetRecentLocationsRequest{}
	_ bin.Decoder     = &MessagesGetRecentLocationsRequest{}
	_ bin.BareEncoder = &MessagesGetRecentLocationsRequest{}
	_ bin.BareDecoder = &MessagesGetRecentLocationsRequest{}
)

// MessagesGetRecentLocations invokes method messages.getRecentLocations#702a40e0 returning error if any.
// Get live location history of a certain user
//
// See https://core.telegram.org/method/messages.getRecentLocations for reference.
func (c *Client) MessagesGetRecentLocations(ctx context.Context, request *MessagesGetRecentLocationsRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
