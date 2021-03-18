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

// MessagesSetBotPrecheckoutResultsRequest represents TL type `messages.setBotPrecheckoutResults#9c2dd95`.
// Once the user has confirmed their payment and shipping details, the bot receives an updateBotPrecheckoutQuery¹ update.
// Use this method to respond to such pre-checkout queries.
// Note: Telegram must receive an answer within 10 seconds after the pre-checkout query was sent.
//
// Links:
//  1) https://core.telegram.org/constructor/updateBotPrecheckoutQuery
//
// See https://core.telegram.org/method/messages.setBotPrecheckoutResults for reference.
type MessagesSetBotPrecheckoutResultsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order, otherwise do not set it, and set the error field, instead
	Success bool
	// Unique identifier for the query to be answered
	QueryID int64
	// Required if the success isn't set. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
	//
	// Use SetError and GetError helpers.
	Error string
}

// MessagesSetBotPrecheckoutResultsRequestTypeID is TL type id of MessagesSetBotPrecheckoutResultsRequest.
const MessagesSetBotPrecheckoutResultsRequestTypeID = 0x9c2dd95

func (s *MessagesSetBotPrecheckoutResultsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Success == false) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.Error == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetBotPrecheckoutResultsRequest) String() string {
	if s == nil {
		return "MessagesSetBotPrecheckoutResultsRequest(nil)"
	}
	type Alias MessagesSetBotPrecheckoutResultsRequest
	return fmt.Sprintf("MessagesSetBotPrecheckoutResultsRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetBotPrecheckoutResultsRequest from given interface.
func (s *MessagesSetBotPrecheckoutResultsRequest) FillFrom(from interface {
	GetSuccess() (value bool)
	GetQueryID() (value int64)
	GetError() (value string, ok bool)
}) {
	s.Success = from.GetSuccess()
	s.QueryID = from.GetQueryID()
	if val, ok := from.GetError(); ok {
		s.Error = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetBotPrecheckoutResultsRequest) TypeID() uint32 {
	return MessagesSetBotPrecheckoutResultsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetBotPrecheckoutResultsRequest) TypeName() string {
	return "messages.setBotPrecheckoutResults"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetBotPrecheckoutResultsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setBotPrecheckoutResults",
		ID:   MessagesSetBotPrecheckoutResultsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Success",
			SchemaName: "success",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "Error",
			SchemaName: "error",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetBotPrecheckoutResultsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setBotPrecheckoutResults#9c2dd95 as nil")
	}
	b.PutID(MessagesSetBotPrecheckoutResultsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetBotPrecheckoutResultsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setBotPrecheckoutResults#9c2dd95 as nil")
	}
	if !(s.Success == false) {
		s.Flags.Set(1)
	}
	if !(s.Error == "") {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setBotPrecheckoutResults#9c2dd95: field flags: %w", err)
	}
	b.PutLong(s.QueryID)
	if s.Flags.Has(0) {
		b.PutString(s.Error)
	}
	return nil
}

// SetSuccess sets value of Success conditional field.
func (s *MessagesSetBotPrecheckoutResultsRequest) SetSuccess(value bool) {
	if value {
		s.Flags.Set(1)
		s.Success = true
	} else {
		s.Flags.Unset(1)
		s.Success = false
	}
}

// GetSuccess returns value of Success conditional field.
func (s *MessagesSetBotPrecheckoutResultsRequest) GetSuccess() (value bool) {
	return s.Flags.Has(1)
}

// GetQueryID returns value of QueryID field.
func (s *MessagesSetBotPrecheckoutResultsRequest) GetQueryID() (value int64) {
	return s.QueryID
}

// SetError sets value of Error conditional field.
func (s *MessagesSetBotPrecheckoutResultsRequest) SetError(value string) {
	s.Flags.Set(0)
	s.Error = value
}

// GetError returns value of Error conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotPrecheckoutResultsRequest) GetError() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Error, true
}

// Decode implements bin.Decoder.
func (s *MessagesSetBotPrecheckoutResultsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setBotPrecheckoutResults#9c2dd95 to nil")
	}
	if err := b.ConsumeID(MessagesSetBotPrecheckoutResultsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setBotPrecheckoutResults#9c2dd95: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetBotPrecheckoutResultsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setBotPrecheckoutResults#9c2dd95 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setBotPrecheckoutResults#9c2dd95: field flags: %w", err)
		}
	}
	s.Success = s.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotPrecheckoutResults#9c2dd95: field query_id: %w", err)
		}
		s.QueryID = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotPrecheckoutResults#9c2dd95: field error: %w", err)
		}
		s.Error = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetBotPrecheckoutResultsRequest.
var (
	_ bin.Encoder     = &MessagesSetBotPrecheckoutResultsRequest{}
	_ bin.Decoder     = &MessagesSetBotPrecheckoutResultsRequest{}
	_ bin.BareEncoder = &MessagesSetBotPrecheckoutResultsRequest{}
	_ bin.BareDecoder = &MessagesSetBotPrecheckoutResultsRequest{}
)

// MessagesSetBotPrecheckoutResults invokes method messages.setBotPrecheckoutResults#9c2dd95 returning error if any.
// Once the user has confirmed their payment and shipping details, the bot receives an updateBotPrecheckoutQuery¹ update.
// Use this method to respond to such pre-checkout queries.
// Note: Telegram must receive an answer within 10 seconds after the pre-checkout query was sent.
//
// Links:
//  1) https://core.telegram.org/constructor/updateBotPrecheckoutQuery
//
// Possible errors:
//  400 ERROR_TEXT_EMPTY: The provided error message is empty
//
// See https://core.telegram.org/method/messages.setBotPrecheckoutResults for reference.
// Can be used by bots.
func MessagesSetBotPrecheckoutResults(ctx context.Context, rpc Invoker, request *MessagesSetBotPrecheckoutResultsRequest) (bool, error) {
	var result BoolBox

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
