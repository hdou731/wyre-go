package circle

type CreateBankAccountResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CreateWirePaymentResponse struct {
	TrackingRef string `json:"trackingRef"`
	Status      string `json:"status"`
}
