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

// HelpGetSupportNameRequest represents TL type `help.getSupportName#d360e72c`.
// Get localized name of the telegram support user
//
// See https://core.telegram.org/method/help.getSupportName for reference.
type HelpGetSupportNameRequest struct {
}

// HelpGetSupportNameRequestTypeID is TL type id of HelpGetSupportNameRequest.
const HelpGetSupportNameRequestTypeID = 0xd360e72c

func (g *HelpGetSupportNameRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetSupportNameRequest) String() string {
	if g == nil {
		return "HelpGetSupportNameRequest(nil)"
	}
	type Alias HelpGetSupportNameRequest
	return fmt.Sprintf("HelpGetSupportNameRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetSupportNameRequest) TypeID() uint32 {
	return HelpGetSupportNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetSupportNameRequest) TypeName() string {
	return "help.getSupportName"
}

// TypeInfo returns info about TL type.
func (g *HelpGetSupportNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getSupportName",
		ID:   HelpGetSupportNameRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetSupportNameRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getSupportName#d360e72c as nil")
	}
	b.PutID(HelpGetSupportNameRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetSupportNameRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getSupportName#d360e72c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetSupportNameRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getSupportName#d360e72c to nil")
	}
	if err := b.ConsumeID(HelpGetSupportNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getSupportName#d360e72c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetSupportNameRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getSupportName#d360e72c to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetSupportNameRequest.
var (
	_ bin.Encoder     = &HelpGetSupportNameRequest{}
	_ bin.Decoder     = &HelpGetSupportNameRequest{}
	_ bin.BareEncoder = &HelpGetSupportNameRequest{}
	_ bin.BareDecoder = &HelpGetSupportNameRequest{}
)

// HelpGetSupportName invokes method help.getSupportName#d360e72c returning error if any.
// Get localized name of the telegram support user
//
// Possible errors:
//  403 USER_INVALID: Invalid user provided
//
// See https://core.telegram.org/method/help.getSupportName for reference.
func HelpGetSupportName(ctx context.Context, rpc Invoker) (*HelpSupportName, error) {
	var result HelpSupportName

	request := &HelpGetSupportNameRequest{}
	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
