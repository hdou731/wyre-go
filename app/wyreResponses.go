package main

type CreateWalletResponse struct {
	id string
	srn string
}

type GetTransferResponse struct {
	exchangeRate string
	status string
}