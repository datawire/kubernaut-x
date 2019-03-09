// +build fast

package agent

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestMarshalBroker(t *testing.T) {
	address, err := url.Parse("https://kubernaut.io/")
	if err != nil {
		t.Fatal(err)
	}

	b := &Broker{Address: *address, Token: "IAmTheWalrus"}

	bBytes, err := json.Marshal(b)
	if err != nil {
		t.Fatal(bBytes)
	}

	expected := `{"Address": "https://kubernaut.io/", "Token": "IAmTheWalrus"}`
	assert.JSONEq(t, expected, string(bBytes), "")
}

func TestUnmarshalBroker(t *testing.T) {
	input := `{"Address": "https://kubernaut.io/", "Token": "IAmTheWalrus"}`
	b := &Broker{}
	if err := json.Unmarshal([]byte(input), b); err != nil {
		t.Fatal(err)
	}

	address, err := url.Parse("https://kubernaut.io/")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *address, b.Address)
	assert.Equal(t, "IAmTheWalrus", b.Token)
	assert.NotNil(t, b.client)
}
