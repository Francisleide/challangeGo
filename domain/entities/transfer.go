package entities

type Transfer struct {
	ID               string
	OriginAccountID  string
	DestineAccountID string
	Amount           float64
	CreatedAt        string
}
