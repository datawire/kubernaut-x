package agent

// Agent periodically contacts the kubernaut broker to indicate the status of a Kubernaut cluster.
type Agent struct {
	// ID is a unique identifier for the agent.
	ID string

	// KubernetesCluster contains important information about the cluster that the agent is handling such as the
	// cluster ID and kubeconfig.
	Cluster KubernetesCluster
}

type KubernetesCluster struct {
	ID      string
	Config  string
	State   string
	Flavor  string
	Version string
}

func (a *Agent) Run() {

}
