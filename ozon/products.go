package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type GetStocksInfoParams struct {
	// Identifier of the last value on the page. Leave this field blank in the first request.
	//
	// To get the next values, specify last_id from the response of the previous request.
	LastId string `json:"last_id,omitempty"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int64 `json:"limit,omitempty"`

	// Filter by product
	Filter GetStocksInfoFilter `json:"filter,omitempty"`
}

type GetStocksInfoFilter struct {
	// Filter by the offer_id parameter. It is possible to pass a list of values
	OfferId string `json:"offer_id,omitempty"`

	// Filter by the product_id parameter. It is possible to pass a list of values
	ProductId int64 `json:"product_id,omitempty"`

	// Filter by product visibility
	Visibility string `json:"visibility,omitempty"`
}

type GetStocksInfoResponse struct {
	core.CommonResponse

	// Method Result
	Result struct {
		// Identifier of the last value on the page
		//
		// To get the next values, specify the recieved value in the next request in the last_id parameter
		LastId string `json:"last_id,omitempty"`

		// The number of unique products for which information about stocks is displayed
		Total int32 `json:"total,omitempty"`

		// Product details
		Items []struct {
			// Product identifier in the seller's system
			OfferId string `json:"offer_id,omitempty"`

			// Product identifier
			ProductId int64 `json:"product_id,omitempty"`

			// Stock details
			Stocks []struct {
				// In a warehouse
				Present int32 `json:"present,omitempty"`

				// Reserved
				Reserved int32 `json:"reserved,omitempty"`

				// Warehouse type
				Type string `json:"type,omitempty" default:"ALL"`
			} `json:"stocks,omitempty"`
		} `json:"items,omitempty"`
	} `json:"result,omitempty"`
}

// Returns information about the quantity of products in stock:
//
// * how many items are available,
//
// * how many are reserved by customers.
func (c Client) GetStocksInfo(params *GetStocksInfoParams) (*GetStocksInfoResponse, error) {
	url := "/v3/product/info/stocks"

	resp := &GetStocksInfoResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductDetailsParams struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type GetProductDetailsResponse struct {
	core.CommonResponse

	// Request results
	Result struct {
		// Barcode
		Barcode string `json:"barcode"`

		// All product barcodes
		Barcodes []string `json:"barcodes"`

		// Main offer price on Ozon.
		//
		// The field is deprecated. Returns an empty string ""
		BuyboxPrice string `json:"buybox_price"`

		// Category identifier
		CategoryId int64 `json:"category_id"`

		// Marketing color
		ColorImage string `json:"color_image"`

		// Commission fees details
		Commissions []struct {
			// Delivery cost
			DeliveryAmount float64 `json:"deliveryAmount"`

			// Minimum commission fee
			MinValue float64 `json:"minValue"`

			// Commission percentage
			Percent float64 `json:"percent"`

			// Return cost
			ReturnAmount float64 `json:"returnAmount"`

			// Sale scheme
			SaleSchema string `json:"saleSchema"`

			// Commission fee amount
			Value float64 `json:"value"`
		} `json:"commissions"`

		// Date and time when the product was created
		CreatedAt time.Time `json:"created_at"`

		// SKU of the product that is sold from the Ozon warehouse (FBO)
		FBOSKU int64 `json:"fbo_sku"`

		// SKU of the product that is sold from the seller's warehouse (FBS and rFBS)
		FBSSKU int64 `json:"fbs_sku"`

		// Document generation task number
		Id int64 `json:"id"`

		// An array of links to images. The images in the array are arranged in the order of their arrangement on the site. If the `primary_image` parameter is not specified, the first image in the list is the main one for the product
		Images []string `json:"images"`

		// Main product image
		PrimaryImage string `json:"primary_image"`

		// Array of 360 images
		Images360 []string `json:"images360"`

		// true if the product has markdown equivalents at the Ozon warehouse
		HasDiscountedItem bool `json:"has_discounted_item"`

		// Indication of a markdown product:
		//
		// * true if the product was created by the seller as a markdown
		//
		// * false if the product is not markdown or was marked down by Ozon
		IsDiscounted bool `json:"is_discounted"`

		// Markdown products stocks
		DiscountedStocks struct {
			// Quantity of products to be supplied
			Coming int32 `json:"coming"`

			// Quantity of products in warehouse
			Present int32 `json:"present"`

			// Quantity of products reserved
			Reserved int32 `json:"reserved"`
		} `json:"discounted_stocks"`

		// Indication of a bulky product
		IsKGT bool `json:"is_kgt"`

		// Indication of mandatory prepayment for the product:
		//
		// * true — to buy a product, you need to make a prepayment.
		//
		// * false—prepayment is not required
		IsPrepayment bool `json:"is_prepayment"`

		// If prepayment is possible, the value is true
		IsPrepaymentAllowed bool `json:"is_prepayment_allowed"`

		// Currency of your prices. It matches the currency set in the personal account settings
		CurrencyCode string `json:"currency_code"`

		// The price of the product including all promotion discounts. This value will be shown on the Ozon storefront
		MarketingPrice string `json:"marketing_price"`

		// Minimum price for similar products on Ozon.
		//
		// The field is deprecated. Returns an empty string ""
		MinOzonPrice string `json:"min_ozon_price"`

		// Minimum product price with all promotions applied
		MinPrice string `json:"min_price"`

		// Name
		Name string `json:"name"`

		// Product identifier in the seller's system
		OfferId string `json:"offer_id"`

		// Price before discounts. Displayed strikethrough on the product description page
		OldPrice string `json:"old_price"`

		// Price for customers with an Ozon Premium subscription
		PremiumPrice string `json:"premium_price"`

		// Product price including discounts. This value is shown on the product description page
		Price string `json:"price"`

		// Price index. Learn more in Help Center
		PriceIndex string `json:"price_idnex"`

		// Product price suggested by the system based on similar offers
		RecommendedPrice string `json:"recommended_price"`

		// Product state description
		Status struct {
			// Product state
			State string `json:"state"`

			// Product state on the transition to which an error occurred
			StateFailed string `json:"state_failed"`

			// Moderation status
			ModerateStatus string `json:"moderate_status"`

			// Product decline reasons
			DeclineReasons []string `json:"decline_reasons"`

			// Validation status
			ValidationsState string `json:"validation_state"`

			// Product status name
			StateName string `json:"state_name"`

			// Product state description
			StateDescription string `json:"state_description"`

			// Indiction that there were errors while creating products
			IsFailed bool `json:"is_failed"`

			// Indiction that the product was created
			IsCreated bool `json:"is_created"`

			// Tooltips for the current product state
			StateTooltip string `json:"state_tooltip"`

			// Product loading errors
			ItemErrors []GetProductDetailsResponseItemError `json:"item_errors"`

			// The last time product state changed
			StateUpdatedAt time.Time `json:"state_updated_at"`
		} `json:"status"`

		// Details about the sources of similar offers. Learn more in Help Сenter
		Sources []struct {
			// Indication that the source is taken into account when calculating the market value
			IsEnabled bool `json:"is_enabled"`

			// Product identifier in the Ozon system, SKU
			SKU int64 `json:"sku"`

			// Link to the source
			Source string `json:"source"`
		} `json:"sources"`

		// Details about product stocks
		Stocks struct {
			// Supply expected
			Coming int32 `json:"coming"`

			// Currently at the warehouse
			Present int32 `json:"present"`

			// Reserved
			Reserved int32 `json:"reserved"`
		} `json:"stocks"`

		// Date of the last product update
		UpdatedAt time.Time `json:"updated_at"`

		// Product VAT rate
		VAT string `json:"vat"`

		// Product visibility settings
		VisibilityDetails struct {
			// If the product is active, the value is true
			ActiveProduct bool `json:"active_product"`

			// If the price is set, the value is true
			HasPrice bool `json:"has_price"`

			// If there is stock at the warehouses, the value is true
			HasStock bool `json:"has_stock"`
		} `json:"visibility_details"`

		// If the product is on sale, the value is true
		Visible bool `json:"visible"`

		// Product volume weight
		VolumeWeight float64 `json:"volume_weights"`
	} `json:"Result"`
}

type GetProductDetailsResponseItemError struct {
	// Error code
	Code string `json:"code"`

	// Product state in which an error occurred
	State string `json:"state"`

	// Error level
	Level string `json:"level"`

	// Error description
	Description string `json:"description"`

	// Error field
	Field string `json:"field"`

	// Error attribute identifier
	AttributeId int64 `json:"attribute_id"`

	// Attribute name
	AttributeName string `json:"attribute_name"`

	// Additional fields for error description
	OptionalDescriptionElements struct {
		// Additional field for error description
		PropertyName string `json:"property_name"`
	} `json:"optional_description_elements"`
}

// Get product details
func (c Client) GetProductDetails(params *GetProductDetailsParams) (*GetProductDetailsResponse, error) {
	url := "/v2/product/info"

	resp := &GetProductDetailsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateStocksParams struct {
	// Stock details
	Stocks []UpdateStocksStock `json:"stocks"`
}

// Stock detail
type UpdateStocksStock struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Quantity of products in stock
	Stock int64 `json:"stocks"`
}

type UpdateStocksResponse struct {
	core.CommonResponse

	// Request results
	Result []struct {
		// An array of errors that occurred while processing the request
		Errors []struct {
			// Error code
			Code string `json:"code"`

			// Error reason
			Message string `json:"message"`
		} `json:"errors"`

		// Product identifier in the seller's system
		OfferId string `json:"offer_id"`

		// Product identifier
		ProductId int64 `json:"product_id"`

		// If the product details have been successfully updated — true
		Updated bool `json:"updated"`
	}
}

// Allows you to change the products in stock quantity. The method is only used for FBS and rFBS warehouses.
//
// With one request you can change the availability for 100 products. You can send up to 80 requests in a minute.
//
// Availability can only be set after the product status has been changed to processed.
func (c Client) UpdateStocks(params *UpdateStocksParams) (*UpdateStocksResponse, error) {
	url := "/v1/product/import/stocks"

	resp := &UpdateStocksResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateQuantityStockProductsParams struct {
	// Information about the products at the warehouses
	Stocks []UpdateQuantityStockProductsStock `json:"stocks"`
}

type UpdateQuantityStockProductsStock struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Quantity
	Stock int64 `json:"stock"`

	// Warehouse identifier derived from the /v1/warehouse/list method
	WarehouseId int64 `json:"warehouse_id"`
}

type UpdateQuantityStockProductsResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// An array of errors that occurred while processing the request
		Errors []struct {
			// Error code
			Code string `json:"code"`

			// Error reason
			Message string `json:"message"`
		} `json:"errors"`

		// Product identifier in the seller's system
		Offerid string `json:"offer_id"`

		// Product identifier
		ProductId int64 `json:"product_id"`

		// If the request was completed successfully and the stocks are updated — true
		Updated bool `json:"updated"`

		// Warehouse identifier derived from the /v1/warehouse/list method
		WarehouseId int64 `json:"warehouse_id"`
	} `json:"result"`
}

// Allows you to change the products in stock quantity.
//
// With one request you can change the availability for 100 products. You can send up to 80 requests in a minute.
//
// You can update the stock of one product in one warehouse only once in 2 minutes, otherwise there will be the TOO_MANY_REQUESTS error in the response.
//
// Availability can only be set after the product status has been changed to processed.
//
// Bulky products stock can only be updated in the warehouses for bulky products.
func (c Client) UpdateQuantityStockProducts(params *UpdateQuantityStockProductsParams) (*UpdateQuantityStockProductsResponse, error) {
	url := "/v2/products/stocks"

	resp := &UpdateQuantityStockProductsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type StocksInSellersWarehouseParams struct {
	// SKU of the product that is sold from the seller's warehouse (FBS and RFBS schemes).
	//
	// Get fbs_sku in the /v2/product/info and /v2/product/info/list methods response.
	//
	// The maximum number of SKUs per request is 500.
	FBSSKU []string `json:"fbs_sku"`
}

type StocksInSellersWarehouseResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// SKU of the product that is sold from the seller's warehouse (FBS and RFBS schemes)
		FBSSKU int64 `json:"fbs_sku"`

		// Total number of items in the warehouse
		Present int64 `json:"present"`

		// The product identifier in the seller's system
		ProductId int64 `json:"product_id"`

		// The number of reserved products in the warehouse
		Reserved int64 `json:"reserved"`

		// Warehouse identifier
		WarehouseId int64 `json:"warehouse_id"`

		// Warehouse name
		WarehouseName string `json:"warehouse_name"`
	}
}

// Get stocks in seller's warehouse
func (c Client) StocksInSellersWarehouse(params *StocksInSellersWarehouseParams) (*StocksInSellersWarehouseResponse, error) {
	url := "/v1/product/info/stocks-by-warehouse/fbs"

	resp := &StocksInSellersWarehouseResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdatePricesParams struct {
	// Product prices details
	Prices []UpdatePricesPrice `json:"prices"`
}

// Product price details
type UpdatePricesPrice struct {
	// Attribute for enabling and disabling promos auto-application
	AutoActionEnabled string `json:"auto_action_enabled"`

	// Currency of your prices. The passed value must be the same as the one set in the personal account settings.
	// By default, the passed value is RUB, Russian ruble
	CurrencyCode string `json:"currency_code"`

	// Minimum product price with all promotions applied
	MinPrice string `json:"min_price"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Price before discounts. Displayed strikethrough on the product description page.
	// Specified in rubles. The fractional part is separated by decimal point, up to two digits after the decimal point.
	OldPrice string `json:"old_price"`

	// Product price including discounts. This value is displayed on the product description page.
	//
	// If the current price of the product is from 400 to 10 000 rubles inclusive, the difference between the values of price and old_price fields should be more than 5%, but not less than 20 rubles.
	Price string `json:"price"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type UpdatePricesResponse struct {
	core.CommonResponse

	Result []struct {
		// An array of errors that occurred while processing the request
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`

		// Product identifier in the seller's system
		OfferId string `json:"offer_id"`

		// Product ID
		ProductId int64 `json:"product_id"`

		// If the product details have been successfully updated — true
		Updated bool `json:"updated"`
	} `json:"result"`
}

// Allows you to change a price of one or more products.
// You can change prices for 1000 products in one request.
// To reset old_price or premium_price set these parameters to 0.
//
// A new price must differ from the old one by at least 5%.
func (c Client) UpdatePrices(params *UpdatePricesParams) (*UpdatePricesResponse, error) {
	url := "/v1/product/import/prices"

	resp := &UpdatePricesResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}