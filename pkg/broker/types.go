package broker

import (
	"github.com/datawire/kubernaut/pkg/claimregistry"
	"time"
)

type NodeType int

const (
	NODETYPE_UNKNOWN = iota
	NODETYPE_MASTER
	NODETYPE_WORKER
)

func (n NodeType) String() string {
	return [...]string{"UNKNOWN", "MASTER", "WORKER"}[n]
}

type ClaimJSON struct {
	ID        string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Holder    claimregistry.ClaimHolder
	StartTime time.Time
	Duration  time.Duration
	Cluster   Cluster
}

type Cluster struct {
	ID     string `json:",omitempty"`
	Config string `json:",omitempty"`
}

type ClusterNode struct {
	ID                string   `json:",omitempty"`
	Type              NodeType `json:",omitempty"`
	PublicIPAddressV4 string   `json:",omitempty"`
	PublicIPAddressV6 string   `json:",omitempty"`
}

type ClaimList struct {
	Claims []ClaimJSON `json:",omitempty"`
}
