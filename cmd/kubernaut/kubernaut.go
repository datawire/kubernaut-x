package kubernaut

import (
	"fmt"
	"github.com/datawire/kubernaut/pkg/broker"
	"github.com/datawire/kubernaut/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var verbose bool
var dataDirectory string

var rootCmd = &cobra.Command{
	Use:   "kubernaut",
	Short: "Kubernaut",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.Logger.SetLevel(logrus.DebugLevel)
		}
	},
	Run: nil,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "log more information (debug level)")

	versionCmd := VersionCmd()
	versionCmd.Flags().BoolVarP(&UseShortFormat, "short", "", false, "format version in 'short format'")

	rootCmd.AddCommand(versionCmd)

	brokerCmd := BrokerCmd()
	rootCmd.AddCommand(brokerCmd)

	agentCmd := AgentCmd()
	agentCmd.Flags().StringVarP(&Token, "broker-token", "", "", "broker authentication token")
	_ = agentCmd.MarkFlagRequired("broker-token")

	agentCmd.Flags().StringVarP(&BrokerBaseURL, "broker", "", "", "address of the broker")
	_ = agentCmd.MarkFlagRequired("broker")

	rootCmd.AddCommand(agentCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func AgentCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "agent",
		Short: "Start a Kubernaut agent",
		Run:   runAgent,
	}
}

func BrokerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "broker",
		Short: "Start a Kubernaut broker",
		Run: func(cmd *cobra.Command, args []string) {
			b := broker.NewBroker()
			if err := b.Run(7000, 7001); err != nil {
				fmt.Println(err)
			}
		},
	}
}
