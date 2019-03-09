package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	brokerclient "github.com/datawire/kubernaut/pkg/broker/client"
	"github.com/sirupsen/logrus"
	"net/url"
	"os/exec"
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

type Agent struct {
	ID            string        `json:",omitempty"`
	Broker        *Broker       `json:",omitempty"`
	Cluster       Cluster       `json:",omitempty"`
	ClusterStatus ClusterStatus `json:",omitempty"`
	logger        *logrus.Logger
}

type Broker struct {
	Address url.URL `json:",omitempty"`
	Token   string  `json:",omitempty"`
	client  *brokerclient.Client
}

type Cluster struct {
	ID      string `json:",omitempty"`
	Config  string `json:",omitempty"`
	Flavor  string `json:",omitempty"`
	Version string `json:",omitempty"`
}

func NewAgent(logger *logrus.Logger) *Agent {
	return &Agent{logger: logger}
}

func (a *Agent) ResolveClusterID(namespace string) error {
	ns, err := kubectlGetNamespace(namespace)
	if err != nil {
		return err
	}

	if ns == "" {
		return errors.New("namespace not found")
	}

	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(ns), raw); err != nil {
		return err
	}

	metadata := raw["metadata"]
	if metadata == nil {
		return fmt.Errorf("namespace %q is missing metadata information", namespace)
	}

	metadata, ok := metadata.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed metadata type transformation")
	}

	return nil
}

func (a *Agent) PerformPreflightChecks() error {
	return nil
}

func (a *Agent) Initialize() {
	if a.Broker.client == nil {
		a.Broker.client = brokerclient.NewBrokerClient(a.Broker.Address, a.Broker.Token)
	}
}

func (a *Agent) Run() {
	heartbeatFrequency := 5 * time.Second

	for {
		a.logger.Infoln("Sending heartbeat to broker")

		a.logger.WithField("frequency", heartbeatFrequency).Infoln("Waiting before next heartbeat")
		time.Sleep(heartbeatFrequency)
	}
}

func (a *Agent) setup() {

}

func kubectlGetNamespace(name string) (string, error) {
	cmd := exec.Command("kubectl", "get", "namespace", name, "--output=json", "--ignore-not-found")
	out, err := cmd.Output()
	return string(out), err
}
