
package openrtb

type BidResponse struct {
	ID string `json:"id"`
	SeatBid []SeatBid `json:"seatbid"`
}

type SeatBid struct {
	Bid []Bid `json:"bid"`
}

type Bid struct {
	ID string `json:"id"`
	ImpID string `json:"impid"`
	Price float64 `json:"price"`
	Adm string `json:"adm"`
	CrID string `json:"crid"`
}
