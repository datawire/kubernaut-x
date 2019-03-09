package kubernaut

import (
	"fmt"
	"github.com/datawire/kubernaut/pkg/log"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var logger = log.Logger

var verbose bool
var DataDirectory string

var rootCmd = &cobra.Command{
	Use:   "kubernaut",
	Short: "Kubernaut",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logger.SetLevel(logrus.DebugLevel)
			logger.Debugln("debug logging enabled")
		}

		if err := os.MkdirAll(DataDirectory, 0755); err != nil {
			logger.Errorln(err)
		}
	},
	Run: nil,
}

func init() {
	userHome, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set the default data directory
	DataDirectory = path.Join(userHome, ".kubernaut")

	rootCmd.PersistentFlags().BoolVarP(&verbose,
		"verbose", "v", false, "log more information (debug level)")

	rootCmd.PersistentFlags().StringVar(&DataDirectory,
		"data-dir", DataDirectory, "filesystem location where kubernaut stores data")

	commands := []*cobra.Command{
		createAgentCommand(rootCmd),
		createBrokerCommand(rootCmd),
		createToolboxCommand(rootCmd),
		createVersionCommand(rootCmd),
	}

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
