package kubernaut

import (
	"fmt"
	"github.com/datawire/kubernaut/pkg/version"
	"github.com/spf13/cobra"
)

var UseShortFormat bool

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Verify Kubernaut version",
		Long: `Use this command to check the version of Kubernaut.
	
	This command will return the version of the kubernaut binary.`,
		Run: func(cmd *cobra.Command, args []string) {
			if UseShortFormat {
				fmt.Printf("%s\n", version.GetVersionShort())
			} else {
				fmt.Printf("%s\n", version.GetVersionJSON())
			}
		},
	}
}
