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

// StickersAddStickerToSetRequest represents TL type `stickers.addStickerToSet#8653febe`.
// Add a sticker to a stickerset, bots only. The sticker set must have been created by the bot.
//
// See https://core.telegram.org/method/stickers.addStickerToSet for reference.
type StickersAddStickerToSetRequest struct {
	// The stickerset
	Stickerset InputStickerSetClass
	// The sticker
	Sticker InputStickerSetItem
}

// StickersAddStickerToSetRequestTypeID is TL type id of StickersAddStickerToSetRequest.
const StickersAddStickerToSetRequestTypeID = 0x8653febe

func (a *StickersAddStickerToSetRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Stickerset == nil) {
		return false
	}
	if !(a.Sticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *StickersAddStickerToSetRequest) String() string {
	if a == nil {
		return "StickersAddStickerToSetRequest(nil)"
	}
	type Alias StickersAddStickerToSetRequest
	return fmt.Sprintf("StickersAddStickerToSetRequest%+v", Alias(*a))
}

// FillFrom fills StickersAddStickerToSetRequest from given interface.
func (a *StickersAddStickerToSetRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
	GetSticker() (value InputStickerSetItem)
}) {
	a.Stickerset = from.GetStickerset()
	a.Sticker = from.GetSticker()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersAddStickerToSetRequest) TypeID() uint32 {
	return StickersAddStickerToSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersAddStickerToSetRequest) TypeName() string {
	return "stickers.addStickerToSet"
}

// TypeInfo returns info about TL type.
func (a *StickersAddStickerToSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.addStickerToSet",
		ID:   StickersAddStickerToSetRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *StickersAddStickerToSetRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stickers.addStickerToSet#8653febe as nil")
	}
	b.PutID(StickersAddStickerToSetRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *StickersAddStickerToSetRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stickers.addStickerToSet#8653febe as nil")
	}
	if a.Stickerset == nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field stickerset is nil")
	}
	if err := a.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field stickerset: %w", err)
	}
	if err := a.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field sticker: %w", err)
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (a *StickersAddStickerToSetRequest) GetStickerset() (value InputStickerSetClass) {
	return a.Stickerset
}

// GetSticker returns value of Sticker field.
func (a *StickersAddStickerToSetRequest) GetSticker() (value InputStickerSetItem) {
	return a.Sticker
}

// Decode implements bin.Decoder.
func (a *StickersAddStickerToSetRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stickers.addStickerToSet#8653febe to nil")
	}
	if err := b.ConsumeID(StickersAddStickerToSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *StickersAddStickerToSetRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stickers.addStickerToSet#8653febe to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: field stickerset: %w", err)
		}
		a.Stickerset = value
	}
	{
		if err := a.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: field sticker: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersAddStickerToSetRequest.
var (
	_ bin.Encoder     = &StickersAddStickerToSetRequest{}
	_ bin.Decoder     = &StickersAddStickerToSetRequest{}
	_ bin.BareEncoder = &StickersAddStickerToSetRequest{}
	_ bin.BareDecoder = &StickersAddStickerToSetRequest{}
)

// StickersAddStickerToSet invokes method stickers.addStickerToSet#8653febe returning error if any.
// Add a sticker to a stickerset, bots only. The sticker set must have been created by the bot.
//
// Possible errors:
//  400 BOT_MISSING: This method can only be run by a bot
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/stickers.addStickerToSet for reference.
// Can be used by bots.
func StickersAddStickerToSet(ctx context.Context, rpc Invoker, request *StickersAddStickerToSetRequest) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
