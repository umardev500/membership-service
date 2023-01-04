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

func (p *PGRepository) BankPermataCharge() (permataResponse domain.PermataResponse, err error) {

	jsonData := []byte(`{
		"payment_type": "bank_transfer",
		"bank_transfer": {
		  "bank": "permata",
		  "permata": {
			"recipient_name": "SUDARSONO"
		  }
		},
		"transaction_details": {
		  "order_id": "982738497234581",
		  "gross_amount": 145000
		}
	}`)
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

	err = json.NewDecoder(resp.Body).Decode(&permataResponse)
	if err != nil {
		return
	}

	return
}
