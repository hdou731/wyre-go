package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hdou731/wyre-go/circle"
	"github.com/hdou731/wyre-go/tables"
	"github.com/hdou731/wyre-go/wyre"
	"github.com/jinzhu/gorm"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// our initial migration function
func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&tables.Transfer{})
}

func main() {
	initialMigration()
	r := mux.NewRouter()

	r.HandleFunc("/", testHandler)
	r.HandleFunc("/wyre/getTransfer/{transferToken}", wyre.GetTransfer)
	r.HandleFunc("/wyre/createOrder/{account}", wyre.CreateWalletOrderReservation)
	r.HandleFunc("/wyre/executeOrder/{reservationId}", wyre.ExecuteWalletOrderReservation)
	r.HandleFunc("/wyre/payout/{amount}", wyre.Payout)
	r.HandleFunc("/wyre/getWallet/{walletToken}", wyre.GetWallet)
	r.HandleFunc("/wyre/createWallet/{walletName}", wyre.CreateWallet)

	r.HandleFunc("/circle/createBankAccount", circle.CreateBankAccount)
	r.HandleFunc("/circle/createWirePayment/{trackingRef}", circle.CreateWirePayment)
	r.HandleFunc("/circle/getBalance", circle.GetBalance)
	r.HandleFunc("/circle/getWallets", circle.GetWallets)

	http.ListenAndServe(":3000", r)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, from Service: Wyre-GO"))
}
