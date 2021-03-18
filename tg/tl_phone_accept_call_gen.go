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

// PhoneAcceptCallRequest represents TL type `phone.acceptCall#3bd2b4a0`.
// Accept incoming call
//
// See https://core.telegram.org/method/phone.acceptCall for reference.
type PhoneAcceptCallRequest struct {
	// The call to accept
	Peer InputPhoneCall
	// Parameter for E2E encryption key exchange »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/end-to-end/voice-calls
	GB []byte
	// Phone call settings
	Protocol PhoneCallProtocol
}

// PhoneAcceptCallRequestTypeID is TL type id of PhoneAcceptCallRequest.
const PhoneAcceptCallRequestTypeID = 0x3bd2b4a0

func (a *PhoneAcceptCallRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Peer.Zero()) {
		return false
	}
	if !(a.GB == nil) {
		return false
	}
	if !(a.Protocol.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *PhoneAcceptCallRequest) String() string {
	if a == nil {
		return "PhoneAcceptCallRequest(nil)"
	}
	type Alias PhoneAcceptCallRequest
	return fmt.Sprintf("PhoneAcceptCallRequest%+v", Alias(*a))
}

// FillFrom fills PhoneAcceptCallRequest from given interface.
func (a *PhoneAcceptCallRequest) FillFrom(from interface {
	GetPeer() (value InputPhoneCall)
	GetGB() (value []byte)
	GetProtocol() (value PhoneCallProtocol)
}) {
	a.Peer = from.GetPeer()
	a.GB = from.GetGB()
	a.Protocol = from.GetProtocol()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneAcceptCallRequest) TypeID() uint32 {
	return PhoneAcceptCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneAcceptCallRequest) TypeName() string {
	return "phone.acceptCall"
}

// TypeInfo returns info about TL type.
func (a *PhoneAcceptCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.acceptCall",
		ID:   PhoneAcceptCallRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "GB",
			SchemaName: "g_b",
		},
		{
			Name:       "Protocol",
			SchemaName: "protocol",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *PhoneAcceptCallRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode phone.acceptCall#3bd2b4a0 as nil")
	}
	b.PutID(PhoneAcceptCallRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *PhoneAcceptCallRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode phone.acceptCall#3bd2b4a0 as nil")
	}
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field peer: %w", err)
	}
	b.PutBytes(a.GB)
	if err := a.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (a *PhoneAcceptCallRequest) GetPeer() (value InputPhoneCall) {
	return a.Peer
}

// GetGB returns value of GB field.
func (a *PhoneAcceptCallRequest) GetGB() (value []byte) {
	return a.GB
}

// GetProtocol returns value of Protocol field.
func (a *PhoneAcceptCallRequest) GetProtocol() (value PhoneCallProtocol) {
	return a.Protocol
}

// Decode implements bin.Decoder.
func (a *PhoneAcceptCallRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode phone.acceptCall#3bd2b4a0 to nil")
	}
	if err := b.ConsumeID(PhoneAcceptCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *PhoneAcceptCallRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode phone.acceptCall#3bd2b4a0 to nil")
	}
	{
		if err := a.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field g_b: %w", err)
		}
		a.GB = value
	}
	{
		if err := a.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneAcceptCallRequest.
var (
	_ bin.Encoder     = &PhoneAcceptCallRequest{}
	_ bin.Decoder     = &PhoneAcceptCallRequest{}
	_ bin.BareEncoder = &PhoneAcceptCallRequest{}
	_ bin.BareDecoder = &PhoneAcceptCallRequest{}
)

// PhoneAcceptCall invokes method phone.acceptCall#3bd2b4a0 returning error if any.
// Accept incoming call
//
// Possible errors:
//  400 CALL_ALREADY_ACCEPTED: The call was already accepted
//  400 CALL_ALREADY_DECLINED: The call was already declined
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//  400 CALL_PROTOCOL_FLAGS_INVALID: Call protocol flags invalid
//
// See https://core.telegram.org/method/phone.acceptCall for reference.
func PhoneAcceptCall(ctx context.Context, rpc Invoker, request *PhoneAcceptCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
