
package openrtb

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BidRequest struct {
	ID string `json:"id"`
	Imp []Imp `json:"imp"`
	App *App `json:"app,omitempty"`
	Device Device `json:"device"`
}

type Imp struct {
	ID string `json:"id"`
	BidFloor float64 `json:"bidfloor"`
	Video *Video `json:"video,omitempty"`
}

type Video struct {
	Mimes []string `json:"mimes"`
	W int `json:"w"`
	H int `json:"h"`
}

type App struct {
	Bundle string `json:"bundle"`
}

type Device struct {
	UA string `json:"ua"`
	IP string `json:"ip"`
}

func BuildFromHTTP(c *fiber.Ctx) BidRequest {

	return BidRequest{
		ID: uuid.New().String(),
		Imp: []Imp{
			{
				ID: "1",
				BidFloor: 0.5,
				Video: &Video{
					Mimes: []string{"video/mp4"},
					W: 1920,
					H: 1080,
				},
			},
		},
		App: &App{
			Bundle: c.Query("bundle"),
		},
		Device: Device{
			UA: c.Get("User-Agent"),
			IP: c.IP(),
		},
	}
}
