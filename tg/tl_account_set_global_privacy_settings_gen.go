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

// AccountSetGlobalPrivacySettingsRequest represents TL type `account.setGlobalPrivacySettings#1edaaac2`.
// Set global privacy settings
//
// See https://core.telegram.org/method/account.setGlobalPrivacySettings for reference.
type AccountSetGlobalPrivacySettingsRequest struct {
	// Global privacy settings
	Settings GlobalPrivacySettings
}

// AccountSetGlobalPrivacySettingsRequestTypeID is TL type id of AccountSetGlobalPrivacySettingsRequest.
const AccountSetGlobalPrivacySettingsRequestTypeID = 0x1edaaac2

func (s *AccountSetGlobalPrivacySettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetGlobalPrivacySettingsRequest) String() string {
	if s == nil {
		return "AccountSetGlobalPrivacySettingsRequest(nil)"
	}
	type Alias AccountSetGlobalPrivacySettingsRequest
	return fmt.Sprintf("AccountSetGlobalPrivacySettingsRequest%+v", Alias(*s))
}

// FillFrom fills AccountSetGlobalPrivacySettingsRequest from given interface.
func (s *AccountSetGlobalPrivacySettingsRequest) FillFrom(from interface {
	GetSettings() (value GlobalPrivacySettings)
}) {
	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSetGlobalPrivacySettingsRequest) TypeID() uint32 {
	return AccountSetGlobalPrivacySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSetGlobalPrivacySettingsRequest) TypeName() string {
	return "account.setGlobalPrivacySettings"
}

// TypeInfo returns info about TL type.
func (s *AccountSetGlobalPrivacySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.setGlobalPrivacySettings",
		ID:   AccountSetGlobalPrivacySettingsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSetGlobalPrivacySettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setGlobalPrivacySettings#1edaaac2 as nil")
	}
	b.PutID(AccountSetGlobalPrivacySettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSetGlobalPrivacySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setGlobalPrivacySettings#1edaaac2 as nil")
	}
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setGlobalPrivacySettings#1edaaac2: field settings: %w", err)
	}
	return nil
}

// GetSettings returns value of Settings field.
func (s *AccountSetGlobalPrivacySettingsRequest) GetSettings() (value GlobalPrivacySettings) {
	return s.Settings
}

// Decode implements bin.Decoder.
func (s *AccountSetGlobalPrivacySettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setGlobalPrivacySettings#1edaaac2 to nil")
	}
	if err := b.ConsumeID(AccountSetGlobalPrivacySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setGlobalPrivacySettings#1edaaac2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSetGlobalPrivacySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setGlobalPrivacySettings#1edaaac2 to nil")
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.setGlobalPrivacySettings#1edaaac2: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSetGlobalPrivacySettingsRequest.
var (
	_ bin.Encoder     = &AccountSetGlobalPrivacySettingsRequest{}
	_ bin.Decoder     = &AccountSetGlobalPrivacySettingsRequest{}
	_ bin.BareEncoder = &AccountSetGlobalPrivacySettingsRequest{}
	_ bin.BareDecoder = &AccountSetGlobalPrivacySettingsRequest{}
)

// AccountSetGlobalPrivacySettings invokes method account.setGlobalPrivacySettings#1edaaac2 returning error if any.
// Set global privacy settings
//
// See https://core.telegram.org/method/account.setGlobalPrivacySettings for reference.
// Can be used by bots.
func AccountSetGlobalPrivacySettings(ctx context.Context, rpc Invoker, settings GlobalPrivacySettings) (*GlobalPrivacySettings, error) {
	var result GlobalPrivacySettings

	request := &AccountSetGlobalPrivacySettingsRequest{
		Settings: settings,
	}
	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
