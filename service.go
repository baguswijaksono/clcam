// service.go
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func serviceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "Manage services and access control",
	}

	cmd.AddCommand(addServiceCmd())
	cmd.AddCommand(allowAccessCmd())
	return cmd
}

func addServiceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [service-name]",
		Short: "Add a new service",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			serviceName := args[0]
			// TODO: Save service in the database
			fmt.Printf("Service %s added.\n", serviceName)
		},
	}
}

func allowAccessCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "allow [role-name] [service-name]",
		Short: "Allow a role to access a service",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			roleName := args[0]
			serviceName := args[1]
			// TODO: Grant access to the role for the service
			fmt.Printf("Role %s allowed access to service %s.\n", roleName, serviceName)
		},
	}
}

