package main

import (
	"encoding/json"
	"log"

	"io/ioutil"
	"net/http"
)

type wyre struct {
	log *log.Logger
}

func (_ wyre) getWalletBalance(w http.ResponseWriter, r *http.Request) error {
	url := "https://api.testwyre.com/v2/wallet/WA_8FPWBHUXMWR"
	//fmt.Sprintf("https://api.testwyre.com/v2/wallet/%s", ctx.Value("walletToken"))

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return json.NewEncoder(w).Encode(string(body))
}

func (_ wyre) getTransfer(w http.ResponseWriter, r *http.Request) error {
	url := "https://api.testwyre.com/v3/transfers/TF_YFP9E9QZFCB"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return json.NewEncoder(w).Encode(string(body))
}