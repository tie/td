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

// UploadGetFileHashesRequest represents TL type `upload.getFileHashes#c7025931`.
// Get SHA256 hashes for verifying downloaded files
//
// See https://core.telegram.org/method/upload.getFileHashes for reference.
type UploadGetFileHashesRequest struct {
	// File
	Location InputFileLocationClass
	// Offset from which to get file hashes
	Offset int
}

// UploadGetFileHashesRequestTypeID is TL type id of UploadGetFileHashesRequest.
const UploadGetFileHashesRequestTypeID = 0xc7025931

func (g *UploadGetFileHashesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Location == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UploadGetFileHashesRequest) String() string {
	if g == nil {
		return "UploadGetFileHashesRequest(nil)"
	}
	type Alias UploadGetFileHashesRequest
	return fmt.Sprintf("UploadGetFileHashesRequest%+v", Alias(*g))
}

// FillFrom fills UploadGetFileHashesRequest from given interface.
func (g *UploadGetFileHashesRequest) FillFrom(from interface {
	GetLocation() (value InputFileLocationClass)
	GetOffset() (value int)
}) {
	g.Location = from.GetLocation()
	g.Offset = from.GetOffset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadGetFileHashesRequest) TypeID() uint32 {
	return UploadGetFileHashesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadGetFileHashesRequest) TypeName() string {
	return "upload.getFileHashes"
}

// TypeInfo returns info about TL type.
func (g *UploadGetFileHashesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.getFileHashes",
		ID:   UploadGetFileHashesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *UploadGetFileHashesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getFileHashes#c7025931 as nil")
	}
	b.PutID(UploadGetFileHashesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UploadGetFileHashesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getFileHashes#c7025931 as nil")
	}
	if g.Location == nil {
		return fmt.Errorf("unable to encode upload.getFileHashes#c7025931: field location is nil")
	}
	if err := g.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.getFileHashes#c7025931: field location: %w", err)
	}
	b.PutInt(g.Offset)
	return nil
}

// GetLocation returns value of Location field.
func (g *UploadGetFileHashesRequest) GetLocation() (value InputFileLocationClass) {
	return g.Location
}

// GetOffset returns value of Offset field.
func (g *UploadGetFileHashesRequest) GetOffset() (value int) {
	return g.Offset
}

// Decode implements bin.Decoder.
func (g *UploadGetFileHashesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getFileHashes#c7025931 to nil")
	}
	if err := b.ConsumeID(UploadGetFileHashesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getFileHashes#c7025931: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UploadGetFileHashesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getFileHashes#c7025931 to nil")
	}
	{
		value, err := DecodeInputFileLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFileHashes#c7025931: field location: %w", err)
		}
		g.Location = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getFileHashes#c7025931: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadGetFileHashesRequest.
var (
	_ bin.Encoder     = &UploadGetFileHashesRequest{}
	_ bin.Decoder     = &UploadGetFileHashesRequest{}
	_ bin.BareEncoder = &UploadGetFileHashesRequest{}
	_ bin.BareDecoder = &UploadGetFileHashesRequest{}
)

// UploadGetFileHashes invokes method upload.getFileHashes#c7025931 returning error if any.
// Get SHA256 hashes for verifying downloaded files
//
// Possible errors:
//  400 LOCATION_INVALID: The provided location is invalid
//
// See https://core.telegram.org/method/upload.getFileHashes for reference.
// Can be used by bots.
func UploadGetFileHashes(ctx context.Context, rpc Invoker, request *UploadGetFileHashesRequest) ([]FileHash, error) {
	var result FileHashVector

	if err := rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return []FileHash(result.Elems), nil
}
