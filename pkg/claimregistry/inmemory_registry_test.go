//+build fast

package claimregistry

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCreateClaim(t *testing.T) {
	wg := &sync.WaitGroup{}

	r := NewInMemoryClaimRegistry()
	go r.Process(wg)

	cHolder := ClaimHolder{Alias: "philip", ID: "user001"}
	cID, _ := NewClaimID("test", cHolder)
	c := Claim{
		ID: *cID,
	}

	if err := r.CreateClaim(c); err != nil {
		t.Fatal(err)
	}

	res, found := r.GetClaim(*cID)
	assert.True(t, found)
	assert.Equal(t, c, res)

	wg.Wait()
}
