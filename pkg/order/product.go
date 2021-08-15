package order

// Product represents a product inside an order
type Product struct {
	ID string `json:"id"`

	//request fields
	Quantity int `json:"quantity"`
	Unit string `json:"unit"`

	//response fields
	Price int `json:"price"`
	PriceUnit string `json:"price_unit"`
}
