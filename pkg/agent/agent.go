package agent

import (
	brokerclient "github.com/datawire/kubernaut/pkg/broker/client"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

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

	// KubernetesCluster contains important information about the cluster that the agent is handling such as the
	// cluster ID and kubeconfig.
	Cluster KubernetesCluster

	// The state of the cluster
	ClusterStatus string

	broker *brokerclient.Client
	logger *logrus.Logger
}

func NewAgent(logger *logrus.Logger, brokerBaseURL url.URL, brokerToken string) *Agent {
	brokerClient := brokerclient.NewBrokerClient(brokerBaseURL, brokerToken)
	return &Agent{
		broker: brokerClient,
		logger: logger,
	}
}

type KubernetesCluster struct {
	ID      string
	Config  string
	Flavor  string
	Version string
}

func (a *Agent) Run() {
	heartbeatFrequency := 5 * time.Second

	for {
		a.logger.Infoln("Sending heartbeat to broker")

		a.logger.WithField("frequency", heartbeatFrequency).Infoln("Waiting %s before next heartbeat")
		time.Sleep(heartbeatFrequency)
	}
}

func (a *Agent) setup() {

}
