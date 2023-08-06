package core

import (
	"net/http"
	"strings"
	"time"

	"github.com/diphantxm/ozon-api-client/ozon/products"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/logger"
	"github.com/diphantxm/ozon-api-client/ozon/category"
	"github.com/diphantxm/ozon-api-client/ozon/product"
	"github.com/pkg/errors"
)

const (
	defaultURI = "https://api-seller.ozon.ru"
)

type ClientOpts struct {
	h *http.Client

	uri string

	clientID string
	apiKey   string

	// Request timeout
	timeout time.Duration
}

type Opts func(c *ClientOpts)

func WithURI(uri string) Opts {
	return func(c *ClientOpts) {
		c.uri = uri
	}
}

func WithClientID(clientID string) Opts {
	return func(c *ClientOpts) {
		c.clientID = clientID
	}
}

func WithApiKey(apiKey string) Opts {
	return func(c *ClientOpts) {
		c.apiKey = apiKey
	}
}

func WithClient(h *http.Client) Opts {
	return func(c *ClientOpts) {
		c.h = h
	}
}

func WithTimeout(dur time.Duration) Opts {
	return func(c *ClientOpts) {
		c.timeout = dur
	}
}

var (
	ErrClientIDRequired = errors.New("clientID is required")
	ErrAPIKeyRequired   = errors.New("apiKey is required")
)

func NewClient(opts ...Opts) (*Client, error) {
	c := new(ClientOpts)

	for _, opt := range opts {
		opt(c)
	}

	if c.h == nil {
		c.h = http.DefaultClient
	}

	if c.h.Transport == nil {
		c.h.Transport = http.DefaultTransport
	}

	if strings.EqualFold(c.uri, "") {
		c.uri = defaultURI
	}

	if strings.EqualFold(c.clientID, "") {
		return nil, ErrClientIDRequired
	}

	if strings.EqualFold(c.apiKey, "") {
		return nil, ErrAPIKeyRequired
	}

	if c.h.Timeout == 0 && c.timeout > 0 {
		c.h.Timeout = c.timeout
	}

	c.h.Transport = auth.NewRoundTripper(
		logger.NewRoundTripper(c.h.Transport),
		c.clientID,
		c.apiKey,
	)

	return &Client{
		category: category.New(c.h, c.uri),
		product:  product.New(c.h, c.uri),
		products: products.New(c.h, c.uri),
	}, nil
}

type Client struct {
	category *category.Category
	product  *product.Product
	products *products.Products
}

func (c Client) Category() *category.Category {
	return c.category
}

func (c Client) Product() *product.Product {
	return c.product
}

func (c Client) Products() *products.Products {
	return c.products
}
