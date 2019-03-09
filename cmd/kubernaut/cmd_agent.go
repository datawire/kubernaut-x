package kubernaut

import (
	"encoding/json"
	knautagent "github.com/datawire/kubernaut/pkg/agent"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

var brokerAddress string
var id string
var kubeconfig string
var reset bool
var token string

func createAgentCommand(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent",
		Short: "Start a Kubernaut agent",
		Run:   runAgent,
	}

	cmd.Flags().StringVar(&brokerAddress, "broker-addr", "", "set the address of the broker")
	cmd.Flags().StringVar(&id, "id", "", "set the kubernaut agent ID")
	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "set the path to the kubernetes cluster kubeconfig")
	cmd.Flags().BoolVar(&reset, "reset", false, "remove previous state before starting")
	cmd.Flags().StringVar(&token, "token", "", "set the broker token")

	return cmd
}

func runAgent(cmd *cobra.Command, args []string) {
	brokerAddress, err := url.Parse(brokerAddress)
	if err != nil {
		logger.Fatalln("parse broker url failed", err)
	}

	if kubeconfig == "" {

	}

	stateFile := path.Join(DataDirectory, "agent.json")
	if reset {
		if err := ensurePreviousStateIsRemoved(stateFile); err != nil {
			logger.WithField("file", stateFile).Fatalln("previous agent state could not be removed", err)
		}
	}

	agent, err := loadState(stateFile)
	if err != nil {
		logger.WithField("file", stateFile).Fatalln("load previous agent state failed", err)
	}

	if agent == nil || reset {
		newID, err := uuid.NewRandom()
		if err != nil {
			logger.Errorln("create agent id failed")
		}

		agent = &knautagent.Agent{
			ID:            newID.String(),
			Broker:        &knautagent.Broker{Address: *brokerAddress, Token: token},
			ClusterStatus: knautagent.UNKNOWN,
			Cluster:       knautagent.Cluster{},
		}
	}
}

func ensurePreviousStateIsRemoved(path string) error {
	err := os.Remove(path)
	if err != nil && os.IsNotExist(err) {
		logger.Debugln("previous agent state not found")
		return nil
	}

	return err
}

func loadState(stateFile string) (*knautagent.Agent, error) {
	var res *knautagent.Agent
	var err error

	jsonFile, err := os.Open(stateFile)
	if err != nil && os.IsNotExist(err) {
		return nil, nil // non-existence is OK
	} else if err != nil {
		return res, err
	}

	defer func() {
		if err := jsonFile.Close(); err != nil {
			logger.Errorln("closing state file closed")
		}
	}()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(jsonBytes, res)
	return res, err
}
