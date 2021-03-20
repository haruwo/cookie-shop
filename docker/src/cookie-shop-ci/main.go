package main

import (
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
	}

	rootCmd.AddCommand(NewValidate())

	rootCmd.Execute()
}

