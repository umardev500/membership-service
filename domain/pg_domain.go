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

type BankResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	Currency          string `json:"currency"`
	VaNumbers         []struct {
		Bank     string `json:"bank"`
		VaNumber string `json:"va_number"`
	} `json:"va_numbers"`
	PermataVaNumber string `json:"permata_va_number"`
	MerchantID      string `json:"merchant_id"`
}

type PGUsecase interface {
	BankCharge(orderId string, payment map[string]interface{}) (*BankPaymentResponse, error)
}

type PGRepository interface {
	BankPermataCharge(orderId string, payment map[string]interface{}) (BankResponse, error)
}
