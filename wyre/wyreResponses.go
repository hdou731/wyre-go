package wyre

type CreateWalletResponse struct {
	ID  string `json:"id"`
	SRN string `json:"srn"`
}

type GetTransferResponse struct {
	ExchangeRate int    `json:"exchangeRate"`
	Status       string `json:"status"`
}
