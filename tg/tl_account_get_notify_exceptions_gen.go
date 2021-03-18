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

// AccountGetNotifyExceptionsRequest represents TL type `account.getNotifyExceptions#53577479`.
// Returns list of chats with non-default notification settings
//
// See https://core.telegram.org/method/account.getNotifyExceptions for reference.
type AccountGetNotifyExceptionsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If true, chats with non-default sound will also be returned
	CompareSound bool
	// If specified, only chats of the specified category will be returned
	//
	// Use SetPeer and GetPeer helpers.
	Peer InputNotifyPeerClass
}

// AccountGetNotifyExceptionsRequestTypeID is TL type id of AccountGetNotifyExceptionsRequest.
const AccountGetNotifyExceptionsRequestTypeID = 0x53577479

func (g *AccountGetNotifyExceptionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.CompareSound == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetNotifyExceptionsRequest) String() string {
	if g == nil {
		return "AccountGetNotifyExceptionsRequest(nil)"
	}
	type Alias AccountGetNotifyExceptionsRequest
	return fmt.Sprintf("AccountGetNotifyExceptionsRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetNotifyExceptionsRequest from given interface.
func (g *AccountGetNotifyExceptionsRequest) FillFrom(from interface {
	GetCompareSound() (value bool)
	GetPeer() (value InputNotifyPeerClass, ok bool)
}) {
	g.CompareSound = from.GetCompareSound()
	if val, ok := from.GetPeer(); ok {
		g.Peer = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetNotifyExceptionsRequest) TypeID() uint32 {
	return AccountGetNotifyExceptionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetNotifyExceptionsRequest) TypeName() string {
	return "account.getNotifyExceptions"
}

// TypeInfo returns info about TL type.
func (g *AccountGetNotifyExceptionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getNotifyExceptions",
		ID:   AccountGetNotifyExceptionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CompareSound",
			SchemaName: "compare_sound",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
			Null:       !g.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetNotifyExceptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifyExceptions#53577479 as nil")
	}
	b.PutID(AccountGetNotifyExceptionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetNotifyExceptionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifyExceptions#53577479 as nil")
	}
	if !(g.CompareSound == false) {
		g.Flags.Set(1)
	}
	if !(g.Peer == nil) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		if g.Peer == nil {
			return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field peer is nil")
		}
		if err := g.Peer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.getNotifyExceptions#53577479: field peer: %w", err)
		}
	}
	return nil
}

// SetCompareSound sets value of CompareSound conditional field.
func (g *AccountGetNotifyExceptionsRequest) SetCompareSound(value bool) {
	if value {
		g.Flags.Set(1)
		g.CompareSound = true
	} else {
		g.Flags.Unset(1)
		g.CompareSound = false
	}
}

// GetCompareSound returns value of CompareSound conditional field.
func (g *AccountGetNotifyExceptionsRequest) GetCompareSound() (value bool) {
	return g.Flags.Has(1)
}

// SetPeer sets value of Peer conditional field.
func (g *AccountGetNotifyExceptionsRequest) SetPeer(value InputNotifyPeerClass) {
	g.Flags.Set(0)
	g.Peer = value
}

// GetPeer returns value of Peer conditional field and
// boolean which is true if field was set.
func (g *AccountGetNotifyExceptionsRequest) GetPeer() (value InputNotifyPeerClass, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Peer, true
}

// Decode implements bin.Decoder.
func (g *AccountGetNotifyExceptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifyExceptions#53577479 to nil")
	}
	if err := b.ConsumeID(AccountGetNotifyExceptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetNotifyExceptionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifyExceptions#53577479 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: field flags: %w", err)
		}
	}
	g.CompareSound = g.Flags.Has(1)
	if g.Flags.Has(0) {
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getNotifyExceptions#53577479: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetNotifyExceptionsRequest.
var (
	_ bin.Encoder     = &AccountGetNotifyExceptionsRequest{}
	_ bin.Decoder     = &AccountGetNotifyExceptionsRequest{}
	_ bin.BareEncoder = &AccountGetNotifyExceptionsRequest{}
	_ bin.BareDecoder = &AccountGetNotifyExceptionsRequest{}
)

// AccountGetNotifyExceptions invokes method account.getNotifyExceptions#53577479 returning error if any.
// Returns list of chats with non-default notification settings
//
// See https://core.telegram.org/method/account.getNotifyExceptions for reference.
func AccountGetNotifyExceptions(ctx context.Context, rpc Invoker, request *AccountGetNotifyExceptionsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
