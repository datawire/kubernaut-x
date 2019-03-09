package kubernaut

import "github.com/spf13/cobra"

func createToolboxCommand(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "toolbox",
		Short: "A collection of useful but highly specialized tools",
		Run:   nil,
	}

	return cmd
}
