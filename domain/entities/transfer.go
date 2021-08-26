package entities

type Transfer struct {
	ID               string
	OriginAccountID  string
	DestinationAccountID string
	Amount           float64
	CreatedAt        string
}
