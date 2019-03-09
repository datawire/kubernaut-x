package kubernaut

import (
	"fmt"
	"github.com/datawire/kubernaut/pkg/version"
	"github.com/spf13/cobra"
)

var simpleFormat bool

func createVersionCommand(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Verify Kubernaut version",
		Long: `Use this command to check the version of Kubernaut.
	
	This command will return the version of the kubernaut binary.`,
		Run: func(cmd *cobra.Command, args []string) {
			if simpleFormat {
				fmt.Printf("%s\n", version.GetVersionShort())
			} else {
				fmt.Printf("%s\n", version.GetVersionJSON())
			}
		},
	}

	cmd.Flags().BoolVarP(&simpleFormat, "simple", "s", false, "format version as just 'kubernaut <version>'")

	return cmd
}
