
package bidder

import (
	"io"
	"net/http"
	"ssp/internal/openrtb"
)

type VASTBidder struct {
	tag string
}

func NewVASTBidder(tag string) *VASTBidder {
	return &VASTBidder{tag: tag}
}

func (b *VASTBidder) Name() string {
	return "vast"
}

func (b *VASTBidder) Request(req openrtb.BidRequest) ([]openrtb.Bid, error) {

	resp, err := http.Get(b.tag)

	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)

	bid := openrtb.Bid{
		ID: "vast1",
		ImpID: "1",
		Price: 1.0,
		Adm: string(body),
	}

	return []openrtb.Bid{bid}, nil
}
