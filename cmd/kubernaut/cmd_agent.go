package kubernaut

import (
	"github.com/datawire/kubernaut/pkg/agent"
	"github.com/datawire/kubernaut/pkg/log"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"net/url"
)

var Token string
var BrokerBaseURL string

func runAgent(cmd *cobra.Command, args []string) {
	// if "--fresh" is set then discard any previous state and perform the following:
	//		1. generate a new agent ID
	//		2. create a new agent state directory
	//		3. write the new agent ID
	//		4. initialize a new Agent struct

	var logger = log.Logger
	logger.Infoln(Token)
	logger.Infoln(BrokerBaseURL)

	a := agent.NewAgent(logger, url.URL{}, "")
	a.Run()
}

func getPreviousAgentID() string {
	return ""
}

func createAgentID() string {
	return uuid.New().String()
}
