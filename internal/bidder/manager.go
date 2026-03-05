
package bidder

import (
	"sync"
	"ssp/internal/openrtb"
)

type Manager struct {
	bidders []Bidder
}

func NewManager() *Manager {

	return &Manager{
		bidders: []Bidder{
			NewORTBBidder("https://example-dsp.com/openrtb"),
			NewVASTBidder("https://ads.network.com/vast"),
		},
	}
}

func (m *Manager) CallAll(req openrtb.BidRequest) []openrtb.Bid {

	var wg sync.WaitGroup

	bidChan := make(chan openrtb.Bid, 20)

	for _, bidder := range m.bidders {

		wg.Add(1)

		go func(b Bidder) {

			defer wg.Done()

			bids, err := b.Request(req)

			if err != nil {
				return
			}

			for _, bid := range bids {
				bidChan <- bid
			}

		}(bidder)
	}

	wg.Wait()

	close(bidChan)

	var bids []openrtb.Bid

	for b := range bidChan {
		bids = append(bids, b)
	}

	return bids
}
