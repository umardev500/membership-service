package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"membership/domain"
	"net/http"
)

type PGRepository struct {
	chargeUrl  string
	authString string
}

func NewPGRepository(chargeURL, authString string) domain.PGRepository {
	return &PGRepository{
		chargeUrl:  chargeURL,
		authString: authString,
	}
}

func (p *PGRepository) BankPermataCharge(orderId string, payment map[string]interface{}) (bankResponse domain.BankResponse, err error) {
	var permataRequest domain.PermataRequest

	jsonStr, err := json.Marshal(payment)
	if err != nil {
		return
	}
	if err = json.Unmarshal(jsonStr, &permataRequest); err != nil {
		return
	}

	transactionDetails := permataRequest.TransactionDetails
	bankTf := permataRequest.BankTransfer

	jsonData := fmt.Sprintf(`{
		"payment_type": "bank_transfer",
		"bank_transfer": {
		  "bank": "permata",
		  "permata": {
			"recipient_name": "%s"
		  }
		},
		"transaction_details": {
		  "order_id": "%s",
		  "gross_amount": %d
		}
	}`, bankTf.Permata.RecipientName, orderId, transactionDetails.GrossAmount)

	jsonDataByte := []byte(jsonData)
	postData := bytes.NewBuffer(jsonDataByte)

	client := &http.Client{}
	request, err := http.NewRequest("POST", p.chargeUrl, postData)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", p.authString)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&bankResponse)
	if err != nil {
		return
	}

	return
}
