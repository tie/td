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

// HelpAcceptTermsOfServiceRequest represents TL type `help.acceptTermsOfService#ee72f79a`.
// Accept the new terms of service
//
// See https://core.telegram.org/method/help.acceptTermsOfService for reference.
type HelpAcceptTermsOfServiceRequest struct {
	// ID of terms of service
	ID DataJSON
}

// HelpAcceptTermsOfServiceRequestTypeID is TL type id of HelpAcceptTermsOfServiceRequest.
const HelpAcceptTermsOfServiceRequestTypeID = 0xee72f79a

func (a *HelpAcceptTermsOfServiceRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ID.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *HelpAcceptTermsOfServiceRequest) String() string {
	if a == nil {
		return "HelpAcceptTermsOfServiceRequest(nil)"
	}
	type Alias HelpAcceptTermsOfServiceRequest
	return fmt.Sprintf("HelpAcceptTermsOfServiceRequest%+v", Alias(*a))
}

// FillFrom fills HelpAcceptTermsOfServiceRequest from given interface.
func (a *HelpAcceptTermsOfServiceRequest) FillFrom(from interface {
	GetID() (value DataJSON)
}) {
	a.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpAcceptTermsOfServiceRequest) TypeID() uint32 {
	return HelpAcceptTermsOfServiceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpAcceptTermsOfServiceRequest) TypeName() string {
	return "help.acceptTermsOfService"
}

// TypeInfo returns info about TL type.
func (a *HelpAcceptTermsOfServiceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.acceptTermsOfService",
		ID:   HelpAcceptTermsOfServiceRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *HelpAcceptTermsOfServiceRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode help.acceptTermsOfService#ee72f79a as nil")
	}
	b.PutID(HelpAcceptTermsOfServiceRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *HelpAcceptTermsOfServiceRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode help.acceptTermsOfService#ee72f79a as nil")
	}
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.acceptTermsOfService#ee72f79a: field id: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (a *HelpAcceptTermsOfServiceRequest) GetID() (value DataJSON) {
	return a.ID
}

// Decode implements bin.Decoder.
func (a *HelpAcceptTermsOfServiceRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode help.acceptTermsOfService#ee72f79a to nil")
	}
	if err := b.ConsumeID(HelpAcceptTermsOfServiceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.acceptTermsOfService#ee72f79a: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *HelpAcceptTermsOfServiceRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode help.acceptTermsOfService#ee72f79a to nil")
	}
	{
		if err := a.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.acceptTermsOfService#ee72f79a: field id: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpAcceptTermsOfServiceRequest.
var (
	_ bin.Encoder     = &HelpAcceptTermsOfServiceRequest{}
	_ bin.Decoder     = &HelpAcceptTermsOfServiceRequest{}
	_ bin.BareEncoder = &HelpAcceptTermsOfServiceRequest{}
	_ bin.BareDecoder = &HelpAcceptTermsOfServiceRequest{}
)

// HelpAcceptTermsOfService invokes method help.acceptTermsOfService#ee72f79a returning error if any.
// Accept the new terms of service
//
// See https://core.telegram.org/method/help.acceptTermsOfService for reference.
func HelpAcceptTermsOfService(ctx context.Context, rpc Invoker, id DataJSON) (bool, error) {
	var result BoolBox

	request := &HelpAcceptTermsOfServiceRequest{
		ID: id,
	}
	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
