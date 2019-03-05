package agent

type ClusterStatus int

const (
	UNREGISTERED = iota
	REGISTERED
	CLAIMED
	DISCARDED
	UNKNOWN
)

func (c ClusterStatus) String() string {
	return [...]string{"UNREGISTERED", "REGISTERED", "CLAIMED", "DISCARDED", "UNKNOWN"}[c]
}

// Agent periodically contacts the kubernaut broker to indicate the status of a Kubernaut cluster.
type Agent struct {
	// ID is a unique identifier for the agent.
	ID string

	// Token is a shared secret used to authenticate the Agent with the Broker.
	Token string

	// KubernetesCluster contains important information about the cluster that the agent is handling such as the
	// cluster ID and kubeconfig.
	Cluster KubernetesCluster

	// The state of the cluster
	ClusterStatus string
}

type KubernetesCluster struct {
	ID      string
	Config  string
	Flavor  string
	Version string
}

func (a *Agent) Run() {

}
