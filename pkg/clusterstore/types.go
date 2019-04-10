package clusterstore

import (
	"net"
	"time"
)

type NodeType int

const (
	Unknown NodeType = iota
	Master
	Worker
	MasterWorker
)

func (c NodeType) String() string {
	choices := [...]string{"UNKNOWN", "MASTER", "WORKER", "MASTER_WORKER"}

	if c < Unknown || c > MasterWorker {
		return choices[0]
	}

	return choices[c]
}

type Cluster struct {
	ID          string
	Nodes       []Node
	Credentials string
}

type Node struct {
	ID        string
	MachineID string
	PublicIP  net.IP
	PrivateIP net.IP
	Type      NodeType
}

type ClusterStore interface {
	GetCluster(ID string) (Cluster, bool)
	PutCluster(cluster Cluster) error
	RemoveCluster(ID string)
	GetAndMarkExpired(cutoff time.Time) []Cluster
}
