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

// MessagesDeleteChatUserRequest represents TL type `messages.deleteChatUser#a2185cab`.
// Deletes a user from a chat and sends a service message on it.
//
// See https://core.telegram.org/method/messages.deleteChatUser for reference.
type MessagesDeleteChatUserRequest struct {
	// Flags field of MessagesDeleteChatUserRequest.
	Flags bin.Fields
	// RevokeHistory field of MessagesDeleteChatUserRequest.
	RevokeHistory bool
	// Chat ID
	ChatID int64
	// User ID to be deleted
	UserID InputUserClass
}

// MessagesDeleteChatUserRequestTypeID is TL type id of MessagesDeleteChatUserRequest.
const MessagesDeleteChatUserRequestTypeID = 0xa2185cab

func (d *MessagesDeleteChatUserRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.RevokeHistory == false) {
		return false
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteChatUserRequest) String() string {
	if d == nil {
		return "MessagesDeleteChatUserRequest(nil)"
	}
	type Alias MessagesDeleteChatUserRequest
	return fmt.Sprintf("MessagesDeleteChatUserRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteChatUserRequest from given interface.
func (d *MessagesDeleteChatUserRequest) FillFrom(from interface {
	GetRevokeHistory() (value bool)
	GetChatID() (value int64)
	GetUserID() (value InputUserClass)
}) {
	d.RevokeHistory = from.GetRevokeHistory()
	d.ChatID = from.GetChatID()
	d.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteChatUserRequest) TypeID() uint32 {
	return MessagesDeleteChatUserRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteChatUserRequest) TypeName() string {
	return "messages.deleteChatUser"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteChatUserRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteChatUser",
		ID:   MessagesDeleteChatUserRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RevokeHistory",
			SchemaName: "revoke_history",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteChatUserRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteChatUser#a2185cab as nil")
	}
	b.PutID(MessagesDeleteChatUserRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteChatUserRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteChatUser#a2185cab as nil")
	}
	if !(d.RevokeHistory == false) {
		d.Flags.Set(0)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteChatUser#a2185cab: field flags: %w", err)
	}
	b.PutLong(d.ChatID)
	if d.UserID == nil {
		return fmt.Errorf("unable to encode messages.deleteChatUser#a2185cab: field user_id is nil")
	}
	if err := d.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteChatUser#a2185cab: field user_id: %w", err)
	}
	return nil
}

// SetRevokeHistory sets value of RevokeHistory conditional field.
func (d *MessagesDeleteChatUserRequest) SetRevokeHistory(value bool) {
	if value {
		d.Flags.Set(0)
		d.RevokeHistory = true
	} else {
		d.Flags.Unset(0)
		d.RevokeHistory = false
	}
}

// GetRevokeHistory returns value of RevokeHistory conditional field.
func (d *MessagesDeleteChatUserRequest) GetRevokeHistory() (value bool) {
	return d.Flags.Has(0)
}

// GetChatID returns value of ChatID field.
func (d *MessagesDeleteChatUserRequest) GetChatID() (value int64) {
	return d.ChatID
}

// GetUserID returns value of UserID field.
func (d *MessagesDeleteChatUserRequest) GetUserID() (value InputUserClass) {
	return d.UserID
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteChatUserRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteChatUser#a2185cab to nil")
	}
	if err := b.ConsumeID(MessagesDeleteChatUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteChatUser#a2185cab: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteChatUserRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteChatUser#a2185cab to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.deleteChatUser#a2185cab: field flags: %w", err)
		}
	}
	d.RevokeHistory = d.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteChatUser#a2185cab: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteChatUser#a2185cab: field user_id: %w", err)
		}
		d.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDeleteChatUserRequest.
var (
	_ bin.Encoder     = &MessagesDeleteChatUserRequest{}
	_ bin.Decoder     = &MessagesDeleteChatUserRequest{}
	_ bin.BareEncoder = &MessagesDeleteChatUserRequest{}
	_ bin.BareDecoder = &MessagesDeleteChatUserRequest{}
)

// MessagesDeleteChatUser invokes method messages.deleteChatUser#a2185cab returning error if any.
// Deletes a user from a chat and sends a service message on it.
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 USER_ID_INVALID: The provided user ID is invalid
//  400 USER_NOT_PARTICIPANT: You're not a member of this supergroup/channel
//
// See https://core.telegram.org/method/messages.deleteChatUser for reference.
// Can be used by bots.
func (c *Client) MessagesDeleteChatUser(ctx context.Context, request *MessagesDeleteChatUserRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
