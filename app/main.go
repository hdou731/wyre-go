package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", testHandler)
	r.HandleFunc("/getTransfer/{transferToken}", getTransfer)

	http.ListenAndServe(":3000", r)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, from Service: Wyre-GO"))
}

func getTransfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["transferToken"]

	url := fmt.Sprintf("https://api.testwyre.com/v3/transfers/%s", key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}

func getWalletBalance(w http.ResponseWriter, r *http.Request) {
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
