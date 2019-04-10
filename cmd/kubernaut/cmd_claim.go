package kubernaut

import (
	"fmt"
	"github.com/spf13/cobra"
)

func createClaimCommand(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claims",
		Short: "Manage Kubernetes cluster claims",
		Run:   nil,
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a claim",
		Run:   createClaim,
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a claim",
		Run:   deleteClaim,
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get information about a claim",
		Run:   getClaim,
	}

	getCredentialsCmd := &cobra.Command{
		Use:   "get-credentials",
		Short: "Get the credentials (\"kubeconfig\") for a claimed cluster",
		Run:   getCredentials,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all your claims",
		Run:   listClaims,
	}

	cmd.AddCommand(createCmd)
	cmd.AddCommand(deleteCmd)
	cmd.AddCommand(getCmd)
	cmd.AddCommand(getCredentialsCmd)
	cmd.AddCommand(listCmd)

	return cmd
}

func createClaim(cmd *cobra.Command, args []string) {
	fmt.Println("create claim")
}

func deleteClaim(cmd *cobra.Command, args []string) {
	fmt.Println("delete claim")
}

func getClaim(cmd *cobra.Command, args []string) {
	fmt.Println("get claim")
}

func getCredentials(cmd *cobra.Command, args []string) {
	fmt.Println("get claim")
}

func listClaims(cmd *cobra.Command, args []string) {
	fmt.Println("list claims")
}
