package main

import (
	"os"

	"github.com/spf13/cobra"

	"cookie-shop-ci/lib/files"
)

func NewValidate() *cobra.Command {
	return &cobra.Command{
		Use:   "validate [targetDir]",
		Short: "Validate files",
		RunE: func(cmd *cobra.Command, args []string) error {
			targetDir, err := os.Getwd()
			if err != nil {
				return err
			}

			if len(args) >= 1 {
				targetDir = args[0]
			}

			fs, err := files.Open(targetDir)
			if err != nil {
				return err
			}

			if err := fs.Validate(); err != nil {
				return err
			}

			return nil
		},
	}
}
