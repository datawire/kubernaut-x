package claimregistry

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Registry interface {
	GetClaim(ID ClaimID) (Claim, bool)
	CreateClaim(claim Claim) error
	DeleteClaim(ID ClaimID) bool
	Process(wg *sync.WaitGroup)
}

type ClaimID struct {
	name     string
	holderID string
}

func (i *ClaimID) String() string {
	return fmt.Sprintf("%s::%s", i.holderID, i.name)
}

// NewClaimID produces a valid ClaimID from a claim name and the ID of a ClaimHolder. If the claim name or claim holder
// ID are empty strings then an error is returned.
func NewClaimID(claimName string, holder ClaimHolder) (*ClaimID, error) {
	canonicalName := strings.ToLower(strings.TrimSpace(claimName))
	if canonicalName == "" {
		return nil, errors.New("claim name is empty")
	}

	canonicalHolder := strings.ToLower(strings.TrimSpace(holder.ID))
	if canonicalHolder == "" {
		return nil, errors.New("claim holder ID is empty")
	}

	return &ClaimID{name: canonicalName, holderID: canonicalHolder}, nil
}

// IsHolder returns true if the ID is associated to the passed in ClaimHolder and false otherwise.
func (i *ClaimID) IsHolder(holder ClaimHolder) bool {
	return i.holderID == holder.ID
}

type ClaimHolder struct {
	ID    string
	Alias string
}

func (h *ClaimHolder) String() string {
	res := strings.ToLower(h.ID)

	if h.Alias != "" {
		res += fmt.Sprintf("[%q]", h.Alias)
	}

	return res
}

type Claim struct {
	ID        ClaimID
	Name      string
	Holder    ClaimHolder
	StartTime time.Time
	Duration  time.Duration
	ClusterID string
}
