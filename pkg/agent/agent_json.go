package agent

import (
	"encoding/json"
	"github.com/datawire/kubernaut/pkg/broker/client"
	"net/url"
)

func (b *Broker) MarshalJSON() ([]byte, error) {
	type Alias Broker
	s := struct {
		Address string `json:""`
		*Alias
	}{
		Address: b.Address.String(),
		Alias:   (*Alias)(b),
	}

	return json.MarshalIndent(s, "", "    ")
}

func (b *Broker) UnmarshalJSON(data []byte) error {
	type Alias Broker
	s := struct {
		Address string `json:""`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	brokerURL, err := url.Parse(s.Address)
	if err != nil {
		return err
	}

	b.Address = *brokerURL
	b.client = client.NewBrokerClient(b.Address, b.Token)

	return nil
}

func (a *Agent) MarshalJSON() ([]byte, error) {
	type Alias Agent
	s := struct{ *Alias }{Alias: (*Alias)(a)}
	return json.MarshalIndent(s, "", "    ")
}

func (a *Agent) UnmarshalJSON(data []byte) error {
	type Alias Agent
	s := struct{ *Alias }{Alias: (*Alias)(a)}

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return nil
}
