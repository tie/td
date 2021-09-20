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

// UserEmpty represents TL type `userEmpty#d3bc4b7a`.
// Empty constructor, non-existent user.
//
// See https://core.telegram.org/constructor/userEmpty for reference.
type UserEmpty struct {
	// User identifier or 0
	ID int64
}

// UserEmptyTypeID is TL type id of UserEmpty.
const UserEmptyTypeID = 0xd3bc4b7a

func (u *UserEmpty) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserEmpty) String() string {
	if u == nil {
		return "UserEmpty(nil)"
	}
	type Alias UserEmpty
	return fmt.Sprintf("UserEmpty%+v", Alias(*u))
}

// FillFrom fills UserEmpty from given interface.
func (u *UserEmpty) FillFrom(from interface {
	GetID() (value int64)
}) {
	u.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserEmpty) TypeID() uint32 {
	return UserEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*UserEmpty) TypeName() string {
	return "userEmpty"
}

// TypeInfo returns info about TL type.
func (u *UserEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userEmpty",
		ID:   UserEmptyTypeID,
	}
	if u == nil {
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
func (u *UserEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userEmpty#d3bc4b7a as nil")
	}
	b.PutID(UserEmptyTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserEmpty) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userEmpty#d3bc4b7a as nil")
	}
	b.PutLong(u.ID)
	return nil
}

// GetID returns value of ID field.
func (u *UserEmpty) GetID() (value int64) {
	return u.ID
}

// Decode implements bin.Decoder.
func (u *UserEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userEmpty#d3bc4b7a to nil")
	}
	if err := b.ConsumeID(UserEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userEmpty#d3bc4b7a: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserEmpty) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userEmpty#d3bc4b7a to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode userEmpty#d3bc4b7a: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// construct implements constructor of UserClass.
func (u UserEmpty) construct() UserClass { return &u }

// Ensuring interfaces in compile-time for UserEmpty.
var (
	_ bin.Encoder     = &UserEmpty{}
	_ bin.Decoder     = &UserEmpty{}
	_ bin.BareEncoder = &UserEmpty{}
	_ bin.BareDecoder = &UserEmpty{}

	_ UserClass = &UserEmpty{}
)

// User represents TL type `user#3ff6ecb0`.
// Indicates info about a certain user
//
// See https://core.telegram.org/constructor/user for reference.
type User struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this user indicates the currently logged in user
	Self bool
	// Whether this user is a contact
	Contact bool
	// Whether this user is a mutual contact
	MutualContact bool
	// Whether the account of this user was deleted
	Deleted bool
	// Is this user a bot?
	Bot bool
	// Can the bot see all messages in groups?
	BotChatHistory bool
	// Can the bot be added to groups?
	BotNochats bool
	// Whether this user is verified
	Verified bool
	// Access to this user must be restricted for the reason specified in restriction_reason
	Restricted bool
	// See min¹
	//
	// Links:
	//  1) https://core.telegram.org/api/min
	Min bool
	// Whether the bot can request our geolocation in inline mode
	BotInlineGeo bool
	// Whether this is an official support user
	Support bool
	// This may be a scam user
	Scam bool
	// If set, the profile picture for this user should be refetched
	ApplyMinPhoto bool
	// Fake field of User.
	Fake bool
	// ID of the user
	ID int64
	// Access hash of the user
	//
	// Use SetAccessHash and GetAccessHash helpers.
	AccessHash int64
	// First name
	//
	// Use SetFirstName and GetFirstName helpers.
	FirstName string
	// Last name
	//
	// Use SetLastName and GetLastName helpers.
	LastName string
	// Username
	//
	// Use SetUsername and GetUsername helpers.
	Username string
	// Phone number
	//
	// Use SetPhone and GetPhone helpers.
	Phone string
	// Profile picture of user
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo UserProfilePhotoClass
	// Online status of user
	//
	// Use SetStatus and GetStatus helpers.
	Status UserStatusClass
	// Version of the bot_info field in userFull¹, incremented every time it changes
	//
	// Links:
	//  1) https://core.telegram.org/constructor/userFull
	//
	// Use SetBotInfoVersion and GetBotInfoVersion helpers.
	BotInfoVersion int
	// Contains the reason why access to this user must be restricted.
	//
	// Use SetRestrictionReason and GetRestrictionReason helpers.
	RestrictionReason []RestrictionReason
	// Inline placeholder for this inline bot
	//
	// Use SetBotInlinePlaceholder and GetBotInlinePlaceholder helpers.
	BotInlinePlaceholder string
	// Language code of the user
	//
	// Use SetLangCode and GetLangCode helpers.
	LangCode string
}

// UserTypeID is TL type id of User.
const UserTypeID = 0x3ff6ecb0

func (u *User) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Self == false) {
		return false
	}
	if !(u.Contact == false) {
		return false
	}
	if !(u.MutualContact == false) {
		return false
	}
	if !(u.Deleted == false) {
		return false
	}
	if !(u.Bot == false) {
		return false
	}
	if !(u.BotChatHistory == false) {
		return false
	}
	if !(u.BotNochats == false) {
		return false
	}
	if !(u.Verified == false) {
		return false
	}
	if !(u.Restricted == false) {
		return false
	}
	if !(u.Min == false) {
		return false
	}
	if !(u.BotInlineGeo == false) {
		return false
	}
	if !(u.Support == false) {
		return false
	}
	if !(u.Scam == false) {
		return false
	}
	if !(u.ApplyMinPhoto == false) {
		return false
	}
	if !(u.Fake == false) {
		return false
	}
	if !(u.ID == 0) {
		return false
	}
	if !(u.AccessHash == 0) {
		return false
	}
	if !(u.FirstName == "") {
		return false
	}
	if !(u.LastName == "") {
		return false
	}
	if !(u.Username == "") {
		return false
	}
	if !(u.Phone == "") {
		return false
	}
	if !(u.Photo == nil) {
		return false
	}
	if !(u.Status == nil) {
		return false
	}
	if !(u.BotInfoVersion == 0) {
		return false
	}
	if !(u.RestrictionReason == nil) {
		return false
	}
	if !(u.BotInlinePlaceholder == "") {
		return false
	}
	if !(u.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *User) String() string {
	if u == nil {
		return "User(nil)"
	}
	type Alias User
	return fmt.Sprintf("User%+v", Alias(*u))
}

// FillFrom fills User from given interface.
func (u *User) FillFrom(from interface {
	GetSelf() (value bool)
	GetContact() (value bool)
	GetMutualContact() (value bool)
	GetDeleted() (value bool)
	GetBot() (value bool)
	GetBotChatHistory() (value bool)
	GetBotNochats() (value bool)
	GetVerified() (value bool)
	GetRestricted() (value bool)
	GetMin() (value bool)
	GetBotInlineGeo() (value bool)
	GetSupport() (value bool)
	GetScam() (value bool)
	GetApplyMinPhoto() (value bool)
	GetFake() (value bool)
	GetID() (value int64)
	GetAccessHash() (value int64, ok bool)
	GetFirstName() (value string, ok bool)
	GetLastName() (value string, ok bool)
	GetUsername() (value string, ok bool)
	GetPhone() (value string, ok bool)
	GetPhoto() (value UserProfilePhotoClass, ok bool)
	GetStatus() (value UserStatusClass, ok bool)
	GetBotInfoVersion() (value int, ok bool)
	GetRestrictionReason() (value []RestrictionReason, ok bool)
	GetBotInlinePlaceholder() (value string, ok bool)
	GetLangCode() (value string, ok bool)
}) {
	u.Self = from.GetSelf()
	u.Contact = from.GetContact()
	u.MutualContact = from.GetMutualContact()
	u.Deleted = from.GetDeleted()
	u.Bot = from.GetBot()
	u.BotChatHistory = from.GetBotChatHistory()
	u.BotNochats = from.GetBotNochats()
	u.Verified = from.GetVerified()
	u.Restricted = from.GetRestricted()
	u.Min = from.GetMin()
	u.BotInlineGeo = from.GetBotInlineGeo()
	u.Support = from.GetSupport()
	u.Scam = from.GetScam()
	u.ApplyMinPhoto = from.GetApplyMinPhoto()
	u.Fake = from.GetFake()
	u.ID = from.GetID()
	if val, ok := from.GetAccessHash(); ok {
		u.AccessHash = val
	}

	if val, ok := from.GetFirstName(); ok {
		u.FirstName = val
	}

	if val, ok := from.GetLastName(); ok {
		u.LastName = val
	}

	if val, ok := from.GetUsername(); ok {
		u.Username = val
	}

	if val, ok := from.GetPhone(); ok {
		u.Phone = val
	}

	if val, ok := from.GetPhoto(); ok {
		u.Photo = val
	}

	if val, ok := from.GetStatus(); ok {
		u.Status = val
	}

	if val, ok := from.GetBotInfoVersion(); ok {
		u.BotInfoVersion = val
	}

	if val, ok := from.GetRestrictionReason(); ok {
		u.RestrictionReason = val
	}

	if val, ok := from.GetBotInlinePlaceholder(); ok {
		u.BotInlinePlaceholder = val
	}

	if val, ok := from.GetLangCode(); ok {
		u.LangCode = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*User) TypeID() uint32 {
	return UserTypeID
}

// TypeName returns name of type in TL schema.
func (*User) TypeName() string {
	return "user"
}

// TypeInfo returns info about TL type.
func (u *User) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "user",
		ID:   UserTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Self",
			SchemaName: "self",
			Null:       !u.Flags.Has(10),
		},
		{
			Name:       "Contact",
			SchemaName: "contact",
			Null:       !u.Flags.Has(11),
		},
		{
			Name:       "MutualContact",
			SchemaName: "mutual_contact",
			Null:       !u.Flags.Has(12),
		},
		{
			Name:       "Deleted",
			SchemaName: "deleted",
			Null:       !u.Flags.Has(13),
		},
		{
			Name:       "Bot",
			SchemaName: "bot",
			Null:       !u.Flags.Has(14),
		},
		{
			Name:       "BotChatHistory",
			SchemaName: "bot_chat_history",
			Null:       !u.Flags.Has(15),
		},
		{
			Name:       "BotNochats",
			SchemaName: "bot_nochats",
			Null:       !u.Flags.Has(16),
		},
		{
			Name:       "Verified",
			SchemaName: "verified",
			Null:       !u.Flags.Has(17),
		},
		{
			Name:       "Restricted",
			SchemaName: "restricted",
			Null:       !u.Flags.Has(18),
		},
		{
			Name:       "Min",
			SchemaName: "min",
			Null:       !u.Flags.Has(20),
		},
		{
			Name:       "BotInlineGeo",
			SchemaName: "bot_inline_geo",
			Null:       !u.Flags.Has(21),
		},
		{
			Name:       "Support",
			SchemaName: "support",
			Null:       !u.Flags.Has(23),
		},
		{
			Name:       "Scam",
			SchemaName: "scam",
			Null:       !u.Flags.Has(24),
		},
		{
			Name:       "ApplyMinPhoto",
			SchemaName: "apply_min_photo",
			Null:       !u.Flags.Has(25),
		},
		{
			Name:       "Fake",
			SchemaName: "fake",
			Null:       !u.Flags.Has(26),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "FirstName",
			SchemaName: "first_name",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "Username",
			SchemaName: "username",
			Null:       !u.Flags.Has(3),
		},
		{
			Name:       "Phone",
			SchemaName: "phone",
			Null:       !u.Flags.Has(4),
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
			Null:       !u.Flags.Has(5),
		},
		{
			Name:       "Status",
			SchemaName: "status",
			Null:       !u.Flags.Has(6),
		},
		{
			Name:       "BotInfoVersion",
			SchemaName: "bot_info_version",
			Null:       !u.Flags.Has(14),
		},
		{
			Name:       "RestrictionReason",
			SchemaName: "restriction_reason",
			Null:       !u.Flags.Has(18),
		},
		{
			Name:       "BotInlinePlaceholder",
			SchemaName: "bot_inline_placeholder",
			Null:       !u.Flags.Has(19),
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
			Null:       !u.Flags.Has(22),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *User) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode user#3ff6ecb0 as nil")
	}
	b.PutID(UserTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *User) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode user#3ff6ecb0 as nil")
	}
	if !(u.Self == false) {
		u.Flags.Set(10)
	}
	if !(u.Contact == false) {
		u.Flags.Set(11)
	}
	if !(u.MutualContact == false) {
		u.Flags.Set(12)
	}
	if !(u.Deleted == false) {
		u.Flags.Set(13)
	}
	if !(u.Bot == false) {
		u.Flags.Set(14)
	}
	if !(u.BotChatHistory == false) {
		u.Flags.Set(15)
	}
	if !(u.BotNochats == false) {
		u.Flags.Set(16)
	}
	if !(u.Verified == false) {
		u.Flags.Set(17)
	}
	if !(u.Restricted == false) {
		u.Flags.Set(18)
	}
	if !(u.Min == false) {
		u.Flags.Set(20)
	}
	if !(u.BotInlineGeo == false) {
		u.Flags.Set(21)
	}
	if !(u.Support == false) {
		u.Flags.Set(23)
	}
	if !(u.Scam == false) {
		u.Flags.Set(24)
	}
	if !(u.ApplyMinPhoto == false) {
		u.Flags.Set(25)
	}
	if !(u.Fake == false) {
		u.Flags.Set(26)
	}
	if !(u.AccessHash == 0) {
		u.Flags.Set(0)
	}
	if !(u.FirstName == "") {
		u.Flags.Set(1)
	}
	if !(u.LastName == "") {
		u.Flags.Set(2)
	}
	if !(u.Username == "") {
		u.Flags.Set(3)
	}
	if !(u.Phone == "") {
		u.Flags.Set(4)
	}
	if !(u.Photo == nil) {
		u.Flags.Set(5)
	}
	if !(u.Status == nil) {
		u.Flags.Set(6)
	}
	if !(u.BotInfoVersion == 0) {
		u.Flags.Set(14)
	}
	if !(u.RestrictionReason == nil) {
		u.Flags.Set(18)
	}
	if !(u.BotInlinePlaceholder == "") {
		u.Flags.Set(19)
	}
	if !(u.LangCode == "") {
		u.Flags.Set(22)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#3ff6ecb0: field flags: %w", err)
	}
	b.PutLong(u.ID)
	if u.Flags.Has(0) {
		b.PutLong(u.AccessHash)
	}
	if u.Flags.Has(1) {
		b.PutString(u.FirstName)
	}
	if u.Flags.Has(2) {
		b.PutString(u.LastName)
	}
	if u.Flags.Has(3) {
		b.PutString(u.Username)
	}
	if u.Flags.Has(4) {
		b.PutString(u.Phone)
	}
	if u.Flags.Has(5) {
		if u.Photo == nil {
			return fmt.Errorf("unable to encode user#3ff6ecb0: field photo is nil")
		}
		if err := u.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode user#3ff6ecb0: field photo: %w", err)
		}
	}
	if u.Flags.Has(6) {
		if u.Status == nil {
			return fmt.Errorf("unable to encode user#3ff6ecb0: field status is nil")
		}
		if err := u.Status.Encode(b); err != nil {
			return fmt.Errorf("unable to encode user#3ff6ecb0: field status: %w", err)
		}
	}
	if u.Flags.Has(14) {
		b.PutInt(u.BotInfoVersion)
	}
	if u.Flags.Has(18) {
		b.PutVectorHeader(len(u.RestrictionReason))
		for idx, v := range u.RestrictionReason {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode user#3ff6ecb0: field restriction_reason element with index %d: %w", idx, err)
			}
		}
	}
	if u.Flags.Has(19) {
		b.PutString(u.BotInlinePlaceholder)
	}
	if u.Flags.Has(22) {
		b.PutString(u.LangCode)
	}
	return nil
}

// SetSelf sets value of Self conditional field.
func (u *User) SetSelf(value bool) {
	if value {
		u.Flags.Set(10)
		u.Self = true
	} else {
		u.Flags.Unset(10)
		u.Self = false
	}
}

// GetSelf returns value of Self conditional field.
func (u *User) GetSelf() (value bool) {
	return u.Flags.Has(10)
}

// SetContact sets value of Contact conditional field.
func (u *User) SetContact(value bool) {
	if value {
		u.Flags.Set(11)
		u.Contact = true
	} else {
		u.Flags.Unset(11)
		u.Contact = false
	}
}

// GetContact returns value of Contact conditional field.
func (u *User) GetContact() (value bool) {
	return u.Flags.Has(11)
}

// SetMutualContact sets value of MutualContact conditional field.
func (u *User) SetMutualContact(value bool) {
	if value {
		u.Flags.Set(12)
		u.MutualContact = true
	} else {
		u.Flags.Unset(12)
		u.MutualContact = false
	}
}

// GetMutualContact returns value of MutualContact conditional field.
func (u *User) GetMutualContact() (value bool) {
	return u.Flags.Has(12)
}

// SetDeleted sets value of Deleted conditional field.
func (u *User) SetDeleted(value bool) {
	if value {
		u.Flags.Set(13)
		u.Deleted = true
	} else {
		u.Flags.Unset(13)
		u.Deleted = false
	}
}

// GetDeleted returns value of Deleted conditional field.
func (u *User) GetDeleted() (value bool) {
	return u.Flags.Has(13)
}

// SetBot sets value of Bot conditional field.
func (u *User) SetBot(value bool) {
	if value {
		u.Flags.Set(14)
		u.Bot = true
	} else {
		u.Flags.Unset(14)
		u.Bot = false
	}
}

// GetBot returns value of Bot conditional field.
func (u *User) GetBot() (value bool) {
	return u.Flags.Has(14)
}

// SetBotChatHistory sets value of BotChatHistory conditional field.
func (u *User) SetBotChatHistory(value bool) {
	if value {
		u.Flags.Set(15)
		u.BotChatHistory = true
	} else {
		u.Flags.Unset(15)
		u.BotChatHistory = false
	}
}

// GetBotChatHistory returns value of BotChatHistory conditional field.
func (u *User) GetBotChatHistory() (value bool) {
	return u.Flags.Has(15)
}

// SetBotNochats sets value of BotNochats conditional field.
func (u *User) SetBotNochats(value bool) {
	if value {
		u.Flags.Set(16)
		u.BotNochats = true
	} else {
		u.Flags.Unset(16)
		u.BotNochats = false
	}
}

// GetBotNochats returns value of BotNochats conditional field.
func (u *User) GetBotNochats() (value bool) {
	return u.Flags.Has(16)
}

// SetVerified sets value of Verified conditional field.
func (u *User) SetVerified(value bool) {
	if value {
		u.Flags.Set(17)
		u.Verified = true
	} else {
		u.Flags.Unset(17)
		u.Verified = false
	}
}

// GetVerified returns value of Verified conditional field.
func (u *User) GetVerified() (value bool) {
	return u.Flags.Has(17)
}

// SetRestricted sets value of Restricted conditional field.
func (u *User) SetRestricted(value bool) {
	if value {
		u.Flags.Set(18)
		u.Restricted = true
	} else {
		u.Flags.Unset(18)
		u.Restricted = false
	}
}

// GetRestricted returns value of Restricted conditional field.
func (u *User) GetRestricted() (value bool) {
	return u.Flags.Has(18)
}

// SetMin sets value of Min conditional field.
func (u *User) SetMin(value bool) {
	if value {
		u.Flags.Set(20)
		u.Min = true
	} else {
		u.Flags.Unset(20)
		u.Min = false
	}
}

// GetMin returns value of Min conditional field.
func (u *User) GetMin() (value bool) {
	return u.Flags.Has(20)
}

// SetBotInlineGeo sets value of BotInlineGeo conditional field.
func (u *User) SetBotInlineGeo(value bool) {
	if value {
		u.Flags.Set(21)
		u.BotInlineGeo = true
	} else {
		u.Flags.Unset(21)
		u.BotInlineGeo = false
	}
}

// GetBotInlineGeo returns value of BotInlineGeo conditional field.
func (u *User) GetBotInlineGeo() (value bool) {
	return u.Flags.Has(21)
}

// SetSupport sets value of Support conditional field.
func (u *User) SetSupport(value bool) {
	if value {
		u.Flags.Set(23)
		u.Support = true
	} else {
		u.Flags.Unset(23)
		u.Support = false
	}
}

// GetSupport returns value of Support conditional field.
func (u *User) GetSupport() (value bool) {
	return u.Flags.Has(23)
}

// SetScam sets value of Scam conditional field.
func (u *User) SetScam(value bool) {
	if value {
		u.Flags.Set(24)
		u.Scam = true
	} else {
		u.Flags.Unset(24)
		u.Scam = false
	}
}

// GetScam returns value of Scam conditional field.
func (u *User) GetScam() (value bool) {
	return u.Flags.Has(24)
}

// SetApplyMinPhoto sets value of ApplyMinPhoto conditional field.
func (u *User) SetApplyMinPhoto(value bool) {
	if value {
		u.Flags.Set(25)
		u.ApplyMinPhoto = true
	} else {
		u.Flags.Unset(25)
		u.ApplyMinPhoto = false
	}
}

// GetApplyMinPhoto returns value of ApplyMinPhoto conditional field.
func (u *User) GetApplyMinPhoto() (value bool) {
	return u.Flags.Has(25)
}

// SetFake sets value of Fake conditional field.
func (u *User) SetFake(value bool) {
	if value {
		u.Flags.Set(26)
		u.Fake = true
	} else {
		u.Flags.Unset(26)
		u.Fake = false
	}
}

// GetFake returns value of Fake conditional field.
func (u *User) GetFake() (value bool) {
	return u.Flags.Has(26)
}

// GetID returns value of ID field.
func (u *User) GetID() (value int64) {
	return u.ID
}

// SetAccessHash sets value of AccessHash conditional field.
func (u *User) SetAccessHash(value int64) {
	u.Flags.Set(0)
	u.AccessHash = value
}

// GetAccessHash returns value of AccessHash conditional field and
// boolean which is true if field was set.
func (u *User) GetAccessHash() (value int64, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.AccessHash, true
}

// SetFirstName sets value of FirstName conditional field.
func (u *User) SetFirstName(value string) {
	u.Flags.Set(1)
	u.FirstName = value
}

// GetFirstName returns value of FirstName conditional field and
// boolean which is true if field was set.
func (u *User) GetFirstName() (value string, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.FirstName, true
}

// SetLastName sets value of LastName conditional field.
func (u *User) SetLastName(value string) {
	u.Flags.Set(2)
	u.LastName = value
}

// GetLastName returns value of LastName conditional field and
// boolean which is true if field was set.
func (u *User) GetLastName() (value string, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.LastName, true
}

// SetUsername sets value of Username conditional field.
func (u *User) SetUsername(value string) {
	u.Flags.Set(3)
	u.Username = value
}

// GetUsername returns value of Username conditional field and
// boolean which is true if field was set.
func (u *User) GetUsername() (value string, ok bool) {
	if !u.Flags.Has(3) {
		return value, false
	}
	return u.Username, true
}

// SetPhone sets value of Phone conditional field.
func (u *User) SetPhone(value string) {
	u.Flags.Set(4)
	u.Phone = value
}

// GetPhone returns value of Phone conditional field and
// boolean which is true if field was set.
func (u *User) GetPhone() (value string, ok bool) {
	if !u.Flags.Has(4) {
		return value, false
	}
	return u.Phone, true
}

// SetPhoto sets value of Photo conditional field.
func (u *User) SetPhoto(value UserProfilePhotoClass) {
	u.Flags.Set(5)
	u.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (u *User) GetPhoto() (value UserProfilePhotoClass, ok bool) {
	if !u.Flags.Has(5) {
		return value, false
	}
	return u.Photo, true
}

// SetStatus sets value of Status conditional field.
func (u *User) SetStatus(value UserStatusClass) {
	u.Flags.Set(6)
	u.Status = value
}

// GetStatus returns value of Status conditional field and
// boolean which is true if field was set.
func (u *User) GetStatus() (value UserStatusClass, ok bool) {
	if !u.Flags.Has(6) {
		return value, false
	}
	return u.Status, true
}

// SetBotInfoVersion sets value of BotInfoVersion conditional field.
func (u *User) SetBotInfoVersion(value int) {
	u.Flags.Set(14)
	u.BotInfoVersion = value
}

// GetBotInfoVersion returns value of BotInfoVersion conditional field and
// boolean which is true if field was set.
func (u *User) GetBotInfoVersion() (value int, ok bool) {
	if !u.Flags.Has(14) {
		return value, false
	}
	return u.BotInfoVersion, true
}

// SetRestrictionReason sets value of RestrictionReason conditional field.
func (u *User) SetRestrictionReason(value []RestrictionReason) {
	u.Flags.Set(18)
	u.RestrictionReason = value
}

// GetRestrictionReason returns value of RestrictionReason conditional field and
// boolean which is true if field was set.
func (u *User) GetRestrictionReason() (value []RestrictionReason, ok bool) {
	if !u.Flags.Has(18) {
		return value, false
	}
	return u.RestrictionReason, true
}

// SetBotInlinePlaceholder sets value of BotInlinePlaceholder conditional field.
func (u *User) SetBotInlinePlaceholder(value string) {
	u.Flags.Set(19)
	u.BotInlinePlaceholder = value
}

// GetBotInlinePlaceholder returns value of BotInlinePlaceholder conditional field and
// boolean which is true if field was set.
func (u *User) GetBotInlinePlaceholder() (value string, ok bool) {
	if !u.Flags.Has(19) {
		return value, false
	}
	return u.BotInlinePlaceholder, true
}

// SetLangCode sets value of LangCode conditional field.
func (u *User) SetLangCode(value string) {
	u.Flags.Set(22)
	u.LangCode = value
}

// GetLangCode returns value of LangCode conditional field and
// boolean which is true if field was set.
func (u *User) GetLangCode() (value string, ok bool) {
	if !u.Flags.Has(22) {
		return value, false
	}
	return u.LangCode, true
}

// Decode implements bin.Decoder.
func (u *User) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode user#3ff6ecb0 to nil")
	}
	if err := b.ConsumeID(UserTypeID); err != nil {
		return fmt.Errorf("unable to decode user#3ff6ecb0: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *User) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode user#3ff6ecb0 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field flags: %w", err)
		}
	}
	u.Self = u.Flags.Has(10)
	u.Contact = u.Flags.Has(11)
	u.MutualContact = u.Flags.Has(12)
	u.Deleted = u.Flags.Has(13)
	u.Bot = u.Flags.Has(14)
	u.BotChatHistory = u.Flags.Has(15)
	u.BotNochats = u.Flags.Has(16)
	u.Verified = u.Flags.Has(17)
	u.Restricted = u.Flags.Has(18)
	u.Min = u.Flags.Has(20)
	u.BotInlineGeo = u.Flags.Has(21)
	u.Support = u.Flags.Has(23)
	u.Scam = u.Flags.Has(24)
	u.ApplyMinPhoto = u.Flags.Has(25)
	u.Fake = u.Flags.Has(26)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field id: %w", err)
		}
		u.ID = value
	}
	if u.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field access_hash: %w", err)
		}
		u.AccessHash = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field first_name: %w", err)
		}
		u.FirstName = value
	}
	if u.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field last_name: %w", err)
		}
		u.LastName = value
	}
	if u.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field username: %w", err)
		}
		u.Username = value
	}
	if u.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field phone: %w", err)
		}
		u.Phone = value
	}
	if u.Flags.Has(5) {
		value, err := DecodeUserProfilePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field photo: %w", err)
		}
		u.Photo = value
	}
	if u.Flags.Has(6) {
		value, err := DecodeUserStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field status: %w", err)
		}
		u.Status = value
	}
	if u.Flags.Has(14) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field bot_info_version: %w", err)
		}
		u.BotInfoVersion = value
	}
	if u.Flags.Has(18) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field restriction_reason: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value RestrictionReason
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode user#3ff6ecb0: field restriction_reason: %w", err)
			}
			u.RestrictionReason = append(u.RestrictionReason, value)
		}
	}
	if u.Flags.Has(19) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field bot_inline_placeholder: %w", err)
		}
		u.BotInlinePlaceholder = value
	}
	if u.Flags.Has(22) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#3ff6ecb0: field lang_code: %w", err)
		}
		u.LangCode = value
	}
	return nil
}

// construct implements constructor of UserClass.
func (u User) construct() UserClass { return &u }

// Ensuring interfaces in compile-time for User.
var (
	_ bin.Encoder     = &User{}
	_ bin.Decoder     = &User{}
	_ bin.BareEncoder = &User{}
	_ bin.BareDecoder = &User{}

	_ UserClass = &User{}
)

// UserClass represents User generic type.
//
// See https://core.telegram.org/type/User for reference.
//
// Example:
//  g, err := tg.DecodeUser(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UserEmpty: // userEmpty#d3bc4b7a
//  case *tg.User: // user#3ff6ecb0
//  default: panic(v)
//  }
type UserClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() UserClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// User identifier or 0
	GetID() (value int64)

	// AsNotEmpty tries to map UserClass to User.
	AsNotEmpty() (*User, bool)
}

// AsInputPeer tries to map User to InputPeerUser.
func (u *User) AsInputPeer() *InputPeerUser {
	value := new(InputPeerUser)
	value.UserID = u.GetID()
	if fieldValue, ok := u.GetAccessHash(); ok {
		value.AccessHash = fieldValue
	}

	return value
}

// AsInput tries to map User to InputUser.
func (u *User) AsInput() *InputUser {
	value := new(InputUser)
	value.UserID = u.GetID()
	if fieldValue, ok := u.GetAccessHash(); ok {
		value.AccessHash = fieldValue
	}

	return value
}

// AsNotEmpty tries to map UserEmpty to User.
func (u *UserEmpty) AsNotEmpty() (*User, bool) {
	return nil, false
}

// AsNotEmpty tries to map User to User.
func (u *User) AsNotEmpty() (*User, bool) {
	return u, true
}

// DecodeUser implements binary de-serialization for UserClass.
func DecodeUser(buf *bin.Buffer) (UserClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserEmptyTypeID:
		// Decoding userEmpty#d3bc4b7a.
		v := UserEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserClass: %w", err)
		}
		return &v, nil
	case UserTypeID:
		// Decoding user#3ff6ecb0.
		v := User{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserClass: %w", bin.NewUnexpectedID(id))
	}
}

// User boxes the UserClass providing a helper.
type UserBox struct {
	User UserClass
}

// Decode implements bin.Decoder for UserBox.
func (b *UserBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserBox to nil")
	}
	v, err := DecodeUser(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.User = v
	return nil
}

// Encode implements bin.Encode for UserBox.
func (b *UserBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.User == nil {
		return fmt.Errorf("unable to encode UserClass as nil")
	}
	return b.User.Encode(buf)
}

// UserClassArray is adapter for slice of UserClass.
type UserClassArray []UserClass

// Sort sorts slice of UserClass.
func (s UserClassArray) Sort(less func(a, b UserClass) bool) UserClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UserClass.
func (s UserClassArray) SortStable(less func(a, b UserClass) bool) UserClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UserClass.
func (s UserClassArray) Retain(keep func(x UserClass) bool) UserClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UserClassArray) First() (v UserClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserClassArray) Last() (v UserClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserClassArray) PopFirst() (v UserClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UserClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserClassArray) Pop() (v UserClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUserEmpty returns copy with only UserEmpty constructors.
func (s UserClassArray) AsUserEmpty() (to UserEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*UserEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUser returns copy with only User constructors.
func (s UserClassArray) AsUser() (to UserArray) {
	for _, elem := range s {
		value, ok := elem.(*User)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s UserClassArray) AppendOnlyNotEmpty(to []*User) []*User {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s UserClassArray) AsNotEmpty() (to []*User) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s UserClassArray) FirstAsNotEmpty() (v *User, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s UserClassArray) LastAsNotEmpty() (v *User, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *UserClassArray) PopFirstAsNotEmpty() (v *User, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *UserClassArray) PopAsNotEmpty() (v *User, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// UserEmptyArray is adapter for slice of UserEmpty.
type UserEmptyArray []UserEmpty

// Sort sorts slice of UserEmpty.
func (s UserEmptyArray) Sort(less func(a, b UserEmpty) bool) UserEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UserEmpty.
func (s UserEmptyArray) SortStable(less func(a, b UserEmpty) bool) UserEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UserEmpty.
func (s UserEmptyArray) Retain(keep func(x UserEmpty) bool) UserEmptyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UserEmptyArray) First() (v UserEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserEmptyArray) Last() (v UserEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserEmptyArray) PopFirst() (v UserEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UserEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserEmptyArray) Pop() (v UserEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UserArray is adapter for slice of User.
type UserArray []User

// Sort sorts slice of User.
func (s UserArray) Sort(less func(a, b User) bool) UserArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of User.
func (s UserArray) SortStable(less func(a, b User) bool) UserArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of User.
func (s UserArray) Retain(keep func(x User) bool) UserArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UserArray) First() (v User, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UserArray) Last() (v User, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UserArray) PopFirst() (v User, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero User
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UserArray) Pop() (v User, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
