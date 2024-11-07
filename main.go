package main

import (
	"log"
	"os"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "access-manager",
		Short: "Access Manager CLI for authentication and authorization",
	}

	rootCmd.AddCommand(userCmd(), serviceCmd(), roleCmd())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
