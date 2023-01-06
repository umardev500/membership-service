package repository

import (
	"bytes"
	"encoding/json"
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

func (p *PGRepository) BankCharge(orderId string, payment map[string]interface{}) (bankResponse domain.BankResponse, err error) {
	var bankTfRequest *domain.BankTransferRequest

	jsonStr, err := json.Marshal(payment)
	if err != nil {
		return
	}
	if err = json.Unmarshal(jsonStr, &bankTfRequest); err != nil {
		return
	}

	transactionDetails := bankTfRequest.TransactionDetails
	transactionDetails.OrderId = orderId                  // overide order id
	bankTfRequest.TransactionDetails = transactionDetails // overide transaction details

	jsonData, err := json.Marshal(&bankTfRequest)
	if err != nil {
		return
	}

	postData := bytes.NewBuffer(jsonData)

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
