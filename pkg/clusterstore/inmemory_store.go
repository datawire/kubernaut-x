package clusterstore

import (
	"github.com/pkg/errors"
	"sync"
	"time"
)

type InMemoryClusterStore struct {
	records     map[string]Cluster
	recordsLock *sync.Mutex
	heartbeats  map[string]heartbeat
}

type heartbeat struct {
	// Time of the last heartbeat
	time time.Time

	// Expired is a flag that indicates the last heartbeat was too far in the past. An expired cluster cannot be updated
	expired bool
}

func (s *InMemoryClusterStore) GetCluster(ID string) (res Cluster, found bool) {
	s.recordsLock.Lock()
	defer s.recordsLock.Unlock()

	res, found = s.records[ID]
	return
}

func (s *InMemoryClusterStore) RemoveCluster(ID string) {
	s.recordsLock.Lock()
	defer s.recordsLock.Unlock()

	delete(s.records, ID)
	delete(s.heartbeats, ID)
}

func (s *InMemoryClusterStore) PutCluster(cluster Cluster) error {
	if cluster.ID == "" {
		return errors.New("cluster ID cannot be empty or blank string")
	}

	s.recordsLock.Lock()
	defer s.recordsLock.Unlock()

	lastHeartbeat := s.heartbeats[cluster.ID]
	if lastHeartbeat.expired {
		return errors.New("cluster is expired and cannot be added or modified")
	}

	s.records[cluster.ID] = cluster
	s.heartbeats[cluster.ID] = heartbeat{time: time.Now().UTC(), expired: false}

	return nil
}

func (s *InMemoryClusterStore) GetAndMarkExpired(cutoff time.Time) []Cluster {
	s.recordsLock.Lock()
	defer s.recordsLock.Unlock()

	result := make([]Cluster, 0)
	for ID, lastHeartbeat := range s.heartbeats {
		if lastHeartbeat.expired || lastHeartbeat.time.Before(cutoff.UTC()) {
			lastHeartbeat.expired = true
			s.heartbeats[ID] = lastHeartbeat
			result = append(result, s.records[ID])
		}
	}

	return result
}
