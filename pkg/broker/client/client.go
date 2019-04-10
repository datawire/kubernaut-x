package client

import (
	"github.com/datawire/kubernaut/pkg/broker"
	"github.com/datawire/kubernaut/pkg/util"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL    url.URL
	token      string
	httpClient *http.Client
}

type CreateClaimParams struct {
	Name  string
	Class string
}

func NewBrokerClient(baseURL url.URL, token string) *Client {
	result := &Client{
		baseURL:    baseURL,
		token:      token,
		httpClient: util.NewHTTPClient(0, false),
	}

	return result
}

type ListQuery struct{}

func (c *Client) ListClaims(params ListQuery) broker.ClaimList {
	return broker.ClaimList{}
}

func (c *Client) CreateClaim(name string) (broker.ClaimJSON, error) {
	return broker.ClaimJSON{}, nil
}

func (c *Client) GetClaim(name string) (broker.ClaimJSON, error) {
	return broker.ClaimJSON{}, nil
}

func (c *Client) DeleteClaim(name string) bool {
	return true
}
