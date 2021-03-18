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

// MessagesGetWebPageRequest represents TL type `messages.getWebPage#32ca8f91`.
// Get instant view¹ page
//
// Links:
//  1) https://instantview.telegram.org
//
// See https://core.telegram.org/method/messages.getWebPage for reference.
type MessagesGetWebPageRequest struct {
	// URL of IV page to fetch
	URL string
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetWebPageRequestTypeID is TL type id of MessagesGetWebPageRequest.
const MessagesGetWebPageRequestTypeID = 0x32ca8f91

func (g *MessagesGetWebPageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.URL == "") {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetWebPageRequest) String() string {
	if g == nil {
		return "MessagesGetWebPageRequest(nil)"
	}
	type Alias MessagesGetWebPageRequest
	return fmt.Sprintf("MessagesGetWebPageRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetWebPageRequest from given interface.
func (g *MessagesGetWebPageRequest) FillFrom(from interface {
	GetURL() (value string)
	GetHash() (value int)
}) {
	g.URL = from.GetURL()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetWebPageRequest) TypeID() uint32 {
	return MessagesGetWebPageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetWebPageRequest) TypeName() string {
	return "messages.getWebPage"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetWebPageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getWebPage",
		ID:   MessagesGetWebPageRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetWebPageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPage#32ca8f91 as nil")
	}
	b.PutID(MessagesGetWebPageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetWebPageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPage#32ca8f91 as nil")
	}
	b.PutString(g.URL)
	b.PutInt(g.Hash)
	return nil
}

// GetURL returns value of URL field.
func (g *MessagesGetWebPageRequest) GetURL() (value string) {
	return g.URL
}

// GetHash returns value of Hash field.
func (g *MessagesGetWebPageRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetWebPageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPage#32ca8f91 to nil")
	}
	if err := b.ConsumeID(MessagesGetWebPageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetWebPageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPage#32ca8f91 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: field url: %w", err)
		}
		g.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetWebPageRequest.
var (
	_ bin.Encoder     = &MessagesGetWebPageRequest{}
	_ bin.Decoder     = &MessagesGetWebPageRequest{}
	_ bin.BareEncoder = &MessagesGetWebPageRequest{}
	_ bin.BareDecoder = &MessagesGetWebPageRequest{}
)

// MessagesGetWebPage invokes method messages.getWebPage#32ca8f91 returning error if any.
// Get instant view¹ page
//
// Links:
//  1) https://instantview.telegram.org
//
// Possible errors:
//  400 WC_CONVERT_URL_INVALID: WC convert URL invalid
//
// See https://core.telegram.org/method/messages.getWebPage for reference.
func MessagesGetWebPage(ctx context.Context, rpc Invoker, request *MessagesGetWebPageRequest) (WebPageClass, error) {
	var result WebPageBox

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WebPage, nil
}
