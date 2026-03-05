
package http

import (
	"github.com/gofiber/fiber/v2"
	"ssp/internal/auction"
	"ssp/internal/bidder"
	"ssp/internal/openrtb"
	"ssp/internal/vast"
)

func NewRouter() *fiber.App {

	app := fiber.New()

	manager := bidder.NewManager()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/vast/:tag", func(c *fiber.Ctx) error {

		req := openrtb.BuildFromHTTP(c)

		bids := manager.CallAll(req)

		winner := auction.SelectWinner(bids, req.Imp[0].BidFloor)

		if winner == nil {
			return c.SendStatus(204)
		}

		xml := vast.Build(winner)

		return c.Type("xml").SendString(xml)
	})

	return app
}
