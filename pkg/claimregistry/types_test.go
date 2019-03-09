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
		{&ClaimID{name: "claim000", holderID: "user0"}, nil, "claim000", ClaimHolder{ID: "user0", Alias: "philip"}},
		{&ClaimID{name: "claim000", holderID: "user0"}, nil, "CLAIM000", ClaimHolder{ID: "USER0", Alias: "philip"}},
		{nil, errors.New("claim name is empty"), "", ClaimHolder{ID: "USER0", Alias: "philip"}},
		{nil, errors.New("claim name is empty"), "    ", ClaimHolder{ID: "USER0", Alias: "philip"}},
		{nil, errors.New("claim holder ID is empty"), "claim000", ClaimHolder{ID: "", Alias: "philip"}},
		{nil, errors.New("claim holder ID is empty"), "claim000", ClaimHolder{ID: "    ", Alias: "philip"}},
	}

	for _, in := range inputs {
		claimID, err := NewClaimID(in.claimName, in.holder)
		if err != nil {
			assert.Equal(t, in.expectedError, err)
		}

		assert.Equal(t, in.expected, claimID)
	}
}

func TestClaimID_IsHolder(t *testing.T) {
	inputs := []struct {
		expected bool
		claimID  ClaimID
		holder   ClaimHolder
	}{
		{true, ClaimID{name: "claim000", holderID: "user0"}, ClaimHolder{ID: "user0", Alias: "philip"}},
		{false, ClaimID{name: "claim000", holderID: "user9"}, ClaimHolder{ID: "user0", Alias: "philip"}},
	}

	for _, in := range inputs {
		assert.Equal(t, in.expected, in.claimID.IsHolder(in.holder))
	}
}

func TestClaimHolder_String(t *testing.T) {
	inputs := []struct {
		expected string
		holder   ClaimHolder
	}{
		{`user0["philip"]`, ClaimHolder{ID: "user0", Alias: "philip"}},
		{`user0`, ClaimHolder{ID: "user0", Alias: ""}},
	}

	for _, in := range inputs {
		assert.Equal(t, in.expected, in.holder.String())
	}
}
