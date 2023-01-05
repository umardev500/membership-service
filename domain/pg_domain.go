package domain

type BankPaymentResponse struct {
	PaymentType string
	OrderId     string
	Bank        string
	VaNumber    string
	GrossAmount int64
}

type PermataRequest struct {
	PaymentType  string `json:"payment_type"`
	BankTransfer struct {
		Bank    string `json:"bank"`
		Permata struct {
			RecipientName string `json:"recipient_name"`
		} `json:"permata"`
	} `json:"bank_transfer"`
	TransactionDetails struct {
		OrderId     string `json:"order_id"`
		GrossAmount int64  `json:"gross_amount"`
	} `json:"transaction_details"`
}

type PermataResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionId     string `json:"transaction_id"`
	OrderId           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	Currency          string `json:"currency"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	PermataVaNumber   string `json:"permata_va_number"`
	MerchantId        string `json:"merchant_id"`
}

type PGUsecase interface {
	BankCharge(orderId string, payment map[string]interface{}) (interface{}, error)
}

type PGRepository interface {
	BankPermataCharge(orderId string, payment map[string]interface{}) (PermataResponse, error)
}
