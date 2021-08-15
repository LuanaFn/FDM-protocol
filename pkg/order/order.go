package order

// Order contains the information for an order request
type Order struct {
	URL string `json:"url"`

	// request fields
	MimeTypeAccepted []string `json:"mime_type_accepted"`
	Products []Product `json:"products"`
	Customer Customer `json:"customer"`
	Details string `json:"details"`
	ProcessOrder bool `json:"process_order"`

	// response fields
	ID string `json:"order_id"`
	InternalOrder string `json:"internal_order"`
	OrderMimeType string `json:"order_mime_type"`
	Payment string `json:"payment"`
	PaymentMimeType string `json:"payment_mime_type"`
	Errors []string `json:"errors"`
}
