package main

import (
	"github.com/gorilla/mux"
	"github.com/hdou731/wyre-go/wyre"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", testHandler)
	r.HandleFunc("/getTransfer/{transferToken}", wyre.GetTransfer)
	r.HandleFunc("/createOrder/{account}", wyre.CreateWalletOrderReservation)
	r.HandleFunc("/executeOrder/{reservationId}", wyre.ExecuteWalletOrderReservation)
	r.HandleFunc("/payout/{amount}", wyre.Payout)
	r.HandleFunc("/getWallet/{walletToken}", wyre.GetWallet)
	r.HandleFunc("/createWallet/{walletName}", wyre.CreateWallet)

	http.ListenAndServe(":3000", r)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, from Service: Wyre-GO"))
}
