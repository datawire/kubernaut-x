package brokerapi

import (
	"github.com/datawire/kubernaut/pkg/broker"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

type ListQuery struct{}

func (c *Client) ListClaims(params ListQuery) broker.ClaimList {
	return broker.ClaimList{}
}

func (c *Client) CreateClaim(name string) (broker.Claim, error) {
	return broker.Claim{}, nil
}

func (c *Client) GetClaim(name string) (broker.Claim, error) {
	return broker.Claim{}, nil
}

func (c *Client) DeleteClaim(name string) bool {
	return true
}
