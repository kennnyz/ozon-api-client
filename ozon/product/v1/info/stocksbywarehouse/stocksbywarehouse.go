package stocksbywarehouse

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/andmetoo/ozon-api-client/internal/request"
	"github.com/pkg/errors"
)

func New(
	h *http.Client,
	uri string,
) *StocksByWarehouse {
	return &StocksByWarehouse{
		h:   h,
		uri: uri,
	}
}

type StocksByWarehouse struct {
	h   *http.Client
	uri string
}

func (c StocksByWarehouse) FBS(ctx context.Context, req *FBSRequest) (*FBSResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "FBSRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/fbs", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "FBSRequest.NewRequest")
	}

	return request.Send[FBSResponse](c.h, r, request.ContentTypeApplicationJson)
}
