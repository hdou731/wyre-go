package circle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var BearerKey = "QVBJX0tFWTo4YTkzN2I3ZDE3NjIwYjFlN2I2ZmI0YzI3NjFhOWU5Njo2OTdmMDgwYjI1Y2RlNTE3OTg5NGEwNmNlY2M2YWVjMg=="

func CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// key := vars["walletName"]

	url := "https://api-sandbox.circle.com/v1/banks/wires"

	requestJson := "{ \"billingDetails\": " +
		"{ \"name\": \"Satoshi Nakamoto\", " +
		"\"city\": \"Boston\", " +
		"\"country\": \"US\", " +
		"\"line1\": \"100 Money Street\", " +
		"\"postalCode\": \"01234\", " +
		"\"district\": \"MA\" }, " +
		"\"bankAddress\": { \"country\": \"US\" }, " +
		"\"idempotencyKey\": \"ba943ff1-ca16-49b2-ba55-1057e70ca5c7\", " +
		"\"accountNumber\": \"12340010\", " +
		"\"routingNumber\": \"121000248\"}"

	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BearerKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var bankAccountResponse CreateBankAccountResponse
	json.Unmarshal(body, &bankAccountResponse)

	log.Println(bankAccountResponse.ID)

	w.Write(body)
}

func CreateWirePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["trackingRef"]

	url := "https://api-sandbox.circle.com/v1/mocks/payments/wire"

	requestJson := fmt.Sprintf("{\"amount\": "+
		"{\"amount\": \"50.00\", "+
		"\"currency\": \"USDC\"}, "+
		"\"trackingRef\": \"%s\"}", key)

	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BearerKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var paymentResponse CreateWirePaymentResponse
	json.Unmarshal(body, &paymentResponse)

	log.Println(paymentResponse.Status)

	w.Write(body)
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	url := "https://api-sandbox.circle.com/v1/businessAccount/balances"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+BearerKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}

func GetWallets(w http.ResponseWriter, r *http.Request) {
	url := "https://api-sandbox.circle.com/v1/wallets"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+BearerKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}

func InitiateTransfer(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//key := vars["walletAddress"]

	url := "https://api-sandbox.circle.com/v1/transfers"

	requestJson := "{\"source\": { " +
		"\"type\": \"wallet\", " +
		"\"id\": \"1000594591\" }, " +
		"\"destination\": { " +
		"\"type\": \"blockchain\", " +
		"\"address\": \"0xb3fa65823DCf60f68E21b9c365e52902f4685200\", " +
		"\"chain\": \"ETH\" }, " +
		"\"amount\": { \"amount\": \"5\", \"currency\": \"USD\" }, " +
		"\"idempotencyKey\": \"ba943ff1-ca16-49b2-ba55-1057e70ca5c7\"}"

	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BearerKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var paymentResponse CreateWirePaymentResponse
	json.Unmarshal(body, &paymentResponse)

	log.Println(paymentResponse.Status)

	w.Write(body)
}

func Payout(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//key := vars["walletAddress"]

	url := "https://api-sandbox.circle.com/v1/payouts"

	requestJson := "{ \"source\": { " +
		"\"type\": \"wallet\", " +
		"\"id\": \"1000594591\" }, " +
		"\"destination\": { " +
		"\"type\": \"wire\", " +
		"\"id\": \"f64e699d-8762-4025-942b-f645fb2d3c8a\" }, " +
		"\"amount\": { \"amount\": \"2.00\", \"currency\": \"USD\" }, " +
		"\"metadata\": { \"beneficiaryEmail\": \"satoshi@circle.com\" }, " +
		"\"idempotencyKey\": \"ba943ff1-ca16-49b2-ba55-1057e70ca577\"\n}"

	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BearerKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var paymentResponse CreateWirePaymentResponse
	json.Unmarshal(body, &paymentResponse)

	log.Println(paymentResponse.Status)

	w.Write(body)
}
