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

// ContactsGetSavedRequest represents TL type `contacts.getSaved#82f1e39f`.
// Get all contacts
//
// See https://core.telegram.org/method/contacts.getSaved for reference.
type ContactsGetSavedRequest struct {
}

// ContactsGetSavedRequestTypeID is TL type id of ContactsGetSavedRequest.
const ContactsGetSavedRequestTypeID = 0x82f1e39f

func (g *ContactsGetSavedRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetSavedRequest) String() string {
	if g == nil {
		return "ContactsGetSavedRequest(nil)"
	}
	type Alias ContactsGetSavedRequest
	return fmt.Sprintf("ContactsGetSavedRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsGetSavedRequest) TypeID() uint32 {
	return ContactsGetSavedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsGetSavedRequest) TypeName() string {
	return "contacts.getSaved"
}

// TypeInfo returns info about TL type.
func (g *ContactsGetSavedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.getSaved",
		ID:   ContactsGetSavedRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *ContactsGetSavedRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getSaved#82f1e39f as nil")
	}
	b.PutID(ContactsGetSavedRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ContactsGetSavedRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getSaved#82f1e39f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetSavedRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getSaved#82f1e39f to nil")
	}
	if err := b.ConsumeID(ContactsGetSavedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getSaved#82f1e39f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ContactsGetSavedRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getSaved#82f1e39f to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetSavedRequest.
var (
	_ bin.Encoder     = &ContactsGetSavedRequest{}
	_ bin.Decoder     = &ContactsGetSavedRequest{}
	_ bin.BareEncoder = &ContactsGetSavedRequest{}
	_ bin.BareDecoder = &ContactsGetSavedRequest{}
)

// ContactsGetSaved invokes method contacts.getSaved#82f1e39f returning error if any.
// Get all contacts
//
// Possible errors:
//  403 TAKEOUT_REQUIRED: A takeout session has to be initialized, first
//
// See https://core.telegram.org/method/contacts.getSaved for reference.
func ContactsGetSaved(ctx context.Context, rpc Invoker) ([]SavedPhoneContact, error) {
	var result SavedPhoneContactVector

	request := &ContactsGetSavedRequest{}
	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return []SavedPhoneContact(result.Elems), nil
}
