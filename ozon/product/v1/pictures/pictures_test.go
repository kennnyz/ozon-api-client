package pictures_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/diphantxm/ozon-api-client/ozon/product/v1/pictures"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestImport_Success(t *testing.T) {
	c := pictures.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/attributes/update", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"color_image":"string","images":["string"],"images360":["string"],"product_id":0}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"pictures": [
									{
										"is_360": true,
										"is_color": true,
										"is_primary": true,
										"product_id": 0,
										"state": "imported",
										"url": "string"
									}
								]
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/attributes",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Import(context.Background(), &pictures.ImportRequest{
		ColorImage: "string",
		Images: []string{
			"string",
		},
		Images360: []string{
			"string",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &pictures.ImportResponse{
		Result: pictures.ImportResponseResult{
			Pictures: []pictures.ImportResponseResultPicture{
				{
					Is360:     true,
					IsColor:   true,
					IsPrimary: true,
					ProductID: 0,
					State:     pictures.ImportResponseResultPictureStateImported,
					URL:       "string",
				},
			},
		},
	}, resp)
}
