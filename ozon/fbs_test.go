package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestListUnprocessedShipments(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListUnprocessedShipmentsParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListUnprocessedShipmentsParams{
				Direction: "ASC",
				Filter: ListUnprocessedShipmentsFilter{
					CutoffFrom: core.TimeFromString(t, "2021-08-24T14:15:22Z"),
					CutoffTo:   core.TimeFromString(t, "2021-08-31T14:15:22Z"),
					Status:     "awaiting_packaging",
				},
				Limit: 100,
				With: ListUnprocessedShipmentsWith{
					AnalyticsData: true,
					Barcodes:      true,
					FinancialData: true,
					Translit:      true,
				},
			},
			`{
				"result": {
				  "postings": [
					{
					  "posting_number": "23713478-0018-3",
					  "order_id": 559293114,
					  "order_number": "33713378-0051",
					  "status": "awaiting_packaging",
					  "delivery_method": {
						"id": 15110442724000,
						"name": "Ozon Логистика курьеру, Москва",
						"warehouse_id": 15110442724000,
						"warehouse": "Склад на Ленина",
						"tpl_provider_id": 24,
						"tpl_provider": "Ozon Логистика"
					  },
					  "tracking_number": "",
					  "tpl_integration_type": "ozon",
					  "in_process_at": "2021-08-25T10:48:38Z",
					  "shipment_date": "2021-08-26T10:00:00Z",
					  "delivering_date": null,
					  "cancellation": {
						"cancel_reason_id": 0,
						"cancel_reason": "",
						"cancellation_type": "",
						"cancelled_after_ship": false,
						"affect_cancellation_rating": false,
						"cancellation_initiator": ""
					  },
					  "customer": null,
					  "products": [
						{
						  "currency_code": "RUB",
						  "price": "1259",
						  "offer_id": "УТ-0001365",
						  "name": "Мяч, цвет: черный, 5 кг",
						  "sku": 140048123,
						  "quantity": 1,
						  "mandatory_mark": []
						}
					  ],
					  "addressee": null,
					  "barcodes": {
						"upper_barcode": "%101%806044518",
						"lower_barcode": "23024930500000"
					  },
					  "analytics_data": {
						"region": "Санкт-Петербург",
						"city": "Санкт-Петербург",
						"delivery_type": "PVZ",
						"is_premium": false,
						"payment_type_group_name": "Карты оплаты",
						"warehouse_id": 15110442724000,
						"warehouse": "Склад на Ленина",
						"tpl_provider_id": 24,
						"tpl_provider": "Ozon Логистика",
						"delivery_date_begin": "2022-08-28T14:00:00Z",
						"delivery_date_end": "2022-08-28T18:00:00Z",
						"is_legal": false
					  },
					  "financial_data": {
						"products": [
						  {
							"commission_amount": 0,
							"commission_percent": 0,
							"payout": 0,
							"product_id": 140048123,
							"old_price": 1888,
							"price": 1259,
							"total_discount_value": 629,
							"total_discount_percent": 33.32,
							"actions": [
							  "Системная виртуальная скидка селлера"
							],
							"picking": null,
							"quantity": 1,
							"client_price": "",
							"item_services": {
							  "marketplace_service_item_fulfillment": 0,
							  "marketplace_service_item_pickup": 0,
							  "marketplace_service_item_dropoff_pvz": 0,
							  "marketplace_service_item_dropoff_sc": 0,
							  "marketplace_service_item_dropoff_ff": 0,
							  "marketplace_service_item_direct_flow_trans": 0,
							  "marketplace_service_item_return_flow_trans": 0,
							  "marketplace_service_item_deliv_to_customer": 0,
							  "marketplace_service_item_return_not_deliv_to_customer": 0,
							  "marketplace_service_item_return_part_goods_customer": 0,
							  "marketplace_service_item_return_after_deliv_to_customer": 0
							}
						  }
						],
						"posting_services": {
						  "marketplace_service_item_fulfillment": 0,
						  "marketplace_service_item_pickup": 0,
						  "marketplace_service_item_dropoff_pvz": 0,
						  "marketplace_service_item_dropoff_sc": 0,
						  "marketplace_service_item_dropoff_ff": 0,
						  "marketplace_service_item_direct_flow_trans": 0,
						  "marketplace_service_item_return_flow_trans": 0,
						  "marketplace_service_item_deliv_to_customer": 0,
						  "marketplace_service_item_return_not_deliv_to_customer": 0,
						  "marketplace_service_item_return_part_goods_customer": 0,
						  "marketplace_service_item_return_after_deliv_to_customer": 0
						}
					  },
					  "is_express": false,
					  "requirements": {
						"products_requiring_gtd": [],
						"products_requiring_country": []
					  }
					}
				  ],
				  "count": 55
				}
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.ListUnprocessedShipments(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetFBSShipmentsList(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetFBSShipmentsListParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetFBSShipmentsListParams{
				Direction: "ASC",
				Filter: GetFBSShipmentsListFilter{
					Since:  core.TimeFromString(t, "2021-11-01T00:00:00.000Z"),
					To:     core.TimeFromString(t, "2021-12-01T23:59:59.000Z"),
					Status: "awaiting_packaging",
				},
				Limit:  100,
				Offset: 0,
				With: GetFBSShipmentsListWith{
					AnalyticsData: true,
					FinancialData: true,
					Translit:      true,
				},
			},
			`{
				"result": {
				  "postings": [
					{
					  "posting_number": "05708065-0029-1",
					  "order_id": 680420041,
					  "order_number": "05708065-0029",
					  "status": "awaiting_deliver",
					  "delivery_method": {
						"id": 21321684811000,
						"name": "Ozon Логистика самостоятельно, Красногорск",
						"warehouse_id": 21321684811000,
						"warehouse": "Стим Тойс Нахабино",
						"tpl_provider_id": 24,
						"tpl_provider": "Ozon Логистика"
					  },
					  "tracking_number": "",
					  "tpl_integration_type": "ozon",
					  "in_process_at": "2022-05-13T07:07:32Z",
					  "shipment_date": "2022-05-13T10:00:00Z",
					  "delivering_date": null,
					  "cancellation": {
						"cancel_reason_id": 0,
						"cancel_reason": "",
						"cancellation_type": "",
						"cancelled_after_ship": false,
						"affect_cancellation_rating": false,
						"cancellation_initiator": ""
					  },
					  "customer": null,
					  "products": [
						{
						  "currency_code": "RUB",
						  "price": "1390.000000",
						  "offer_id": "205953",
						  "name": " Электронный конструктор PinLab Позитроник",
						  "sku": 358924380,
						  "quantity": 1,
						  "mandatory_mark": []
						}
					  ],
					  "addressee": null,
					  "barcodes": null,
					  "analytics_data": null,
					  "financial_data": null,
					  "is_express": false,
					  "requirements": {
						"products_requiring_gtd": [],
						"products_requiring_country": [],
						"products_requiring_mandatory_mark": []
					  }
					}
				  ],
				  "has_next": true
				}
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetFBSShipmentsList(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}