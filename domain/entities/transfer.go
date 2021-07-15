package entities

type Transfer struct {
	Id                     string  `json: "transfer_id"`
	Account_origin_id      string  `json: "account_origin_id"`
	Account_destination_id string  `json: "account_destination_id"`
	Amount                 float64 `json: "amount"`
	Created_at             string  `json: "account_destination_id"`
}
