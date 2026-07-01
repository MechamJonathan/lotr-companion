package theoneapi

import (
	"net/http"
	"time"

	"github.com/MechamJonathan/lotr-companion/lotrcache"
)

type Client struct {
	cache      lotrcache.Cache
	httpClient http.Client
	apiKey     string
}

func NewClient(timeout time.Duration, cacheInterval time.Duration, apiKey string) Client {
	return Client{
		cache: lotrcache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		apiKey: apiKey,
	}
}
