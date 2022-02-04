package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", testHandler)
	http.HandleFunc("/getTransfer", getTransfer)

	http.ListenAndServe(":3000", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, from Service: Wyre-GO"))
}

func getTransfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	url := "https://api.testwyre.com/v3/transfers/TF_YFP9E9QZFCB"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer SK-WUEY3J78-3FJBFVEM-MAZM7XHC-NGCW2G4F")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Write(body)
}
