
package bidder

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ssp/internal/openrtb"
)

type ORTBBidder struct {
	endpoint string
}

func NewORTBBidder(url string) *ORTBBidder {
	return &ORTBBidder{endpoint: url}
}

func (b *ORTBBidder) Name() string {
	return "ortb"
}

func (b *ORTBBidder) Request(req openrtb.BidRequest) ([]openrtb.Bid, error) {

	body, _ := json.Marshal(req)

	resp, err := http.Post(b.endpoint,"application/json",bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var bidResp openrtb.BidResponse

	json.NewDecoder(resp.Body).Decode(&bidResp)

	var bids []openrtb.Bid

	for _, sb := range bidResp.SeatBid {
		bids = append(bids, sb.Bid...)
	}

	return bids, nil
}
