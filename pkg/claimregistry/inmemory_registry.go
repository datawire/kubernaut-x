package claimregistry

import (
	"errors"
	"github.com/datawire/kubernaut/pkg/log"
	"sync"
)

var logger = log.Logger

type getClaim struct {
	id    ClaimID
	claim chan<- *getClaimResult
}

type getClaimResult struct {
	claim Claim
	found bool
}

type listClaims struct {
	holder ClaimHolder
	result chan Claim
}

type createClaim struct {
	claim  Claim
	result chan error
}

type deleteClaim struct {
	id      ClaimID
	success chan bool
}

type InMemoryClaimRegistry struct {
	records map[ClaimID]Claim
	get     chan getClaim
	list    chan listClaims
	create  chan createClaim
	delete  chan deleteClaim
}

func NewInMemoryClaimRegistry() Registry {
	r := &InMemoryClaimRegistry{
		records: make(map[ClaimID]Claim),
		get:     make(chan getClaim),
		create:  make(chan createClaim),
		delete:  make(chan deleteClaim),
	}

	return r
}

func (r *InMemoryClaimRegistry) CreateClaim(claim Claim) error {
	reply := make(chan error)
	r.create <- createClaim{claim: claim, result: reply}

	var res error

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(replyChan chan error) {
		defer wg.Done()
		for e := range replyChan {
			if e != nil {
				res = e
			}
			close(replyChan)
		}
	}(reply)

	wg.Wait()
	return res
}

func (r *InMemoryClaimRegistry) DeleteClaim(ID ClaimID) bool {
	reply := make(chan bool)
	r.delete <- deleteClaim{id: ID, success: reply}

	var deleted bool

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(replyChan chan bool) {
		defer wg.Done()
		for b := range replyChan {
			deleted = b
			close(replyChan)
		}
	}(reply)

	wg.Wait()
	return deleted
}

func (r *InMemoryClaimRegistry) GetClaim(ID ClaimID) (Claim, bool) {
	reply := make(chan *getClaimResult)
	r.get <- getClaim{id: ID, claim: reply}

	var res Claim
	var found bool

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(replyChan chan *getClaimResult) {
		defer wg.Done()
		for v := range replyChan {
			found = v.found
			res = v.claim
		}
	}(reply)

	wg.Wait()
	return res, found
}

func (r *InMemoryClaimRegistry) Process(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case listClaims := <-r.list:
			for k, v := range r.records {
				if listClaims.holder.ID != "" && k.IsHolder(listClaims.holder) {
					listClaims.result <- v
				} else {
					listClaims.result <- v
				}
			}
			close(listClaims.result)
		case getClaim := <-r.get:
			if c, ok := r.records[getClaim.id]; ok {
				logger.WithField("id", c.ID).Infoln("claim retrieved")
				getClaim.claim <- &getClaimResult{found: true, claim: c}
			} else {
				logger.WithField("id", c.ID).Infoln("claim not found")
				getClaim.claim <- &getClaimResult{found: false, claim: Claim{}}
			}
			close(getClaim.claim)
		case deleteClaim := <-r.delete:
			if c, ok := r.records[deleteClaim.id]; ok {
				delete(r.records, c.ID)
				logger.WithField("id", c.ID).Infoln("deleted claim")
				deleteClaim.success <- true
			} else {
				logger.WithField("id", c.ID).Infoln("claim not found")
				deleteClaim.success <- true // delete always succeeds
			}
		case createClaim := <-r.create:
			if c, ok := r.records[createClaim.claim.ID]; !ok {
				r.records[createClaim.claim.ID] = createClaim.claim
				logger.WithField("id", createClaim.claim.ID.String()).Infoln("claim created")
				createClaim.result <- nil
			} else {
				logger.WithField("id", c.ID.String()).Warnln("claim already exists")
				createClaim.result <- errors.New("claim already exists")
			}
		}
	}
}
