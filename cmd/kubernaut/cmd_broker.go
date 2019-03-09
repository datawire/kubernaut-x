package kubernaut

import (
	"github.com/datawire/kubernaut/pkg/broker"
	"github.com/datawire/kubernaut/pkg/claimregistry"
	"github.com/spf13/cobra"
	"sync"
)

var (
	port      int
	adminPort int
)

func runBroker(cmd *cobra.Command, args []string) {
	wg := &sync.WaitGroup{}

	claims := claimregistry.NewInMemoryClaimRegistry()
	claims.Process(wg)

	b := broker.NewBroker(claims)
	if err := b.Run(port, adminPort); err != nil {
		logger.Fatalln(err)
	}
}

func createBrokerCommand(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "broker",
		Short: "Start a Kubernaut broker",
		Run:   runBroker,
	}

	cmd.Flags().IntVar(&port, "port", 5000, "configure the api server to listen on this port")
	cmd.Flags().IntVar(&adminPort, "admin-port", 5001, "configure the admin api server to listen on this port")

	return cmd
}
