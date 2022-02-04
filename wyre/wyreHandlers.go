package wyre

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["walletName"]

	url := "https://api.testwyre.com/v2/wallets"

	requestJson := fmt.Sprintf("{\"type\":\"DEFAULT\",\"name\":\"%s\"}", key)
	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var walletResponse CreateWalletResponse
	json.Unmarshal(body, &walletResponse)

	log.Println(walletResponse.ID)

	w.Write(body)
}

func CreateWalletOrderReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["account"]

	url := "https://api.testwyre.com/v3/orders/reserve"

	requestJson := fmt.Sprintf("{\"referrerAccountId\":\"%s\"}", key)
	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}

func ExecuteWalletOrderReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["reservationId"]

	url := "https://api.testwyre.com/v3/debitcard/process/partner"

	requestJson := fmt.Sprintf("{\"debitCard\":"+
		"{\"number\":\"4111111111111111\","+
		"\"year\":\"2023\","+
		"\"month\":\"01\","+
		"\"cvv\":\"123\"},"+
		"\"address\":"+
		"{\"city\":\"Los Angeles\","+
		"\"state\":\"CA\","+
		"\"postalCode\":\"91423\","+
		"\"street1\":\"2000 E Madison St\","+
		"\"country\":\"US\"},"+
		"\"reservationId\":\"%s\","+
		"\"amount\":\"100\","+
		"\"sourceCurrency\":\"USD\","+
		"\"destCurrency\":\"USDC\","+
		"\"dest\":\"wallet:WA_8FPWBHUXMWR\","+
		"\"referrerAccountId\":\"AC_NE6PC8GTUYT\","+
		"\"givenName\":\"Hellen\","+
		"\"familyName\":\"Bandicoot\","+
		"\"email\":\"fakey-fake@squareup.com\","+
		"\"phone\":\"8473343106\","+
		"\"ipAddress\":\"1.1.1.1\","+
		"\"referenceId\":\"AC_NE6PC8GTUYT\"}", key)
	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}

func Payout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["amount"]

	url := "https://api.testwyre.com/v2/transfers"

	requestJson := fmt.Sprintf("{"+
		"\"autoConfirm\":true,"+
		"\"source\":\"wallet:WA_8FPWBHUXMWR\","+
		"\"sourceCurrency\":\"USDC\","+
		"\"destCurrency\":\"USD\","+
		"\"destAmount\":%s,"+
		"\"dest\":{"+
		"\"paymentMethodType\":\"INTERNATIONAL_TRANSFER\","+
		"\"country\":\"US\","+
		"\"currency\":\"USD\","+
		"\"paymentType\":\"LOCAL_BANK_WIRE\","+
		"\"firstNameOnAccount\":\"Billy-Bob\","+
		"\"lastNameOnAccount\":\"Jones\","+
		"\"accountNumber\":\"0000000000000\","+
		"\"routingNumber\":\"0000000000\","+
		"\"accountType\":\"CHECKING\","+
		"\"bankName\":\"JP Morgan\""+
		"}}", key)
	payload := strings.NewReader(requestJson)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}

func GetTransfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["transferToken"]

	url := fmt.Sprintf("https://api.testwyre.com/v3/transfers/%s", key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var transferResponse GetTransferResponse
	err := json.Unmarshal(body, &transferResponse)
	if err != nil {
		log.Println(err)
	}

	log.Println(transferResponse.Status, transferResponse.ExchangeRate)

	w.Write(body)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["walletToken"]

	url := fmt.Sprintf("https://api.testwyre.com/v2/wallet/%s", key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}
