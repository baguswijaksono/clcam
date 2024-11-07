// user.go
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func userCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users for access control",
	}

	cmd.AddCommand(addUserCmd())
	cmd.AddCommand(removeUserCmd())
	cmd.AddCommand(loginCmd())
	return cmd
}

func addUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [username] [password]",
		Short: "Add a new user",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			password := args[1]
			// TODO: Save the user details securely
			fmt.Printf("User %s added.\n", username)
		},
	}
}

func loginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login [username] [password]",
		Short: "Login as a user",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			password := args[1]
			// TODO: Authenticate the user
			fmt.Printf("User %s logged in.\n", username)
		},
	}
}
