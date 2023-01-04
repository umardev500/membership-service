package domain

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
	BankCharge() (interface{}, error)
}

type PGRepository interface {
	BankPermataCharge() (PermataResponse, error)
}
