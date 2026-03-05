
package auction

import "ssp/internal/openrtb"

func SelectWinner(bids []openrtb.Bid, floor float64) *openrtb.Bid {

	var winner *openrtb.Bid
	highest := floor

	for _, bid := range bids {

		if bid.Price > highest {

			highest = bid.Price
			winner = &bid
		}
	}

	return winner
}
