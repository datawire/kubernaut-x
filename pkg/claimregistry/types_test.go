//+build fast

package claimregistry

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClaimID(t *testing.T) {
	inputs := []struct {
		expected      *ClaimID
		expectedError error
		claimName     string
		holder        ClaimHolder
	}{
		{&ClaimID{name: "claim000", holderID: "user0"}, nil, "claim000", ClaimHolder{ID: "user0", Name: "philip"}},
		{&ClaimID{name: "claim000", holderID: "user0"}, nil, "CLAIM000", ClaimHolder{ID: "USER0", Name: "philip"}},
		{nil, errors.New("claim name is empty"), "", ClaimHolder{ID: "USER0", Name: "philip"}},
		{nil, errors.New("claim name is empty"), "    ", ClaimHolder{ID: "USER0", Name: "philip"}},
		{nil, errors.New("claim holder ID is empty"), "claim000", ClaimHolder{ID: "", Name: "philip"}},
		{nil, errors.New("claim holder ID is empty"), "claim000", ClaimHolder{ID: "    ", Name: "philip"}},
	}

	for _, in := range inputs {
		claimID, err := NewClaimID(in.claimName, in.holder)
		if err != nil {
			assert.Equal(t, in.expectedError, err)
		}

		assert.Equal(t, in.expected, claimID)
	}
}
