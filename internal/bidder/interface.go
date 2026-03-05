
package bidder

import "ssp/internal/openrtb"

type Bidder interface {
	Name() string
	Request(req openrtb.BidRequest) ([]openrtb.Bid, error)
}
