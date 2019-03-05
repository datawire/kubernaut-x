package kubernaut

import (
	"fmt"
	"github.com/datawire/kubernaut/pkg/broker"
	"github.com/datawire/kubernaut/pkg/version"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kubernaut",
	Short: "Kubernaut",
	Run:   nil,
}

func init() {
	rootCmd.AddCommand(VersionCmd())

	brokerCmd := BrokerCmd()
	rootCmd.AddCommand(brokerCmd)

	AgentCmd := AgentCmd()
	rootCmd.AddCommand(AgentCmd)
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
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start a kubernaut agent")
		},
	}
}

func BrokerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "broker",
		Short: "Start a Kubernaut broker",
		Run: func(cmd *cobra.Command, args []string) {
			b := broker.NewBroker()
			if err := b.Run(7000); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Verify Kubernaut version",
		Long: `Use this command to check the version of Kubernaut.
	
	This command will return the version of the kubernaut binary.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", version.GetVersionJSON())
		},
	}
}
