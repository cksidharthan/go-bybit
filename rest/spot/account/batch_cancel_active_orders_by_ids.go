package account

import (
	"context"
	"net/http"

	"github.com/cksidharthan/go-bybit/bybit"
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
)

func (c *Client) BatchCancelActiveOrdersByIDs(ctx context.Context, params spot.BatchCancelActiveOrdersByIDsParams) (res *spot.BatchCancelActiveOrdersByIDsResponse, err error) {
	err = c.transporter.SignedRequest(ctx, http.MethodDelete, bybit.PrivateSpotBatchCancelByIdsPath, params, &res)
	if err != nil {
		return
	}
	return
}
