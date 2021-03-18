package telegram

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

func (c *Client) exportAuth(ctx context.Context, dcID int) (*tg.AuthExportedAuthorization, error) {
	export, err := tg.AuthExportAuthorization(ctx, c, dcID)
	if err != nil {
		return nil, xerrors.Errorf("export auth to %d: %w", dcID, err)
	}

	return export, nil
}

// transfer exports current authorization and imports it to another DC.
// See https://core.telegram.org/api/datacenter#authorization-transfer.
func (c *Client) transfer(ctx context.Context, to tg.Invoker, dc int) (tg.AuthAuthorizationClass, error) {
	auth, err := c.exportAuth(ctx, dc)
	if err != nil {
		return nil, xerrors.Errorf("export to %d: %w", dc, err)
	}

	req := &tg.AuthImportAuthorizationRequest{}
	req.FillFrom(auth)
	r, err := tg.AuthImportAuthorization(ctx, to, req)
	if err != nil {
		return nil, xerrors.Errorf("import from %d: %w", dc, err)
	}

	return r, nil
}
