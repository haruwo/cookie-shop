package main

import (
	"cookie-shop-ci/lib/files"

	"github.com/spf13/cobra"
)

func NewValidate() *cobra.Command {
	return &cobra.Command{
		Use:   "validate [targetDir]",
		Short: "Validate files",
		RunE: func(cmd *cobra.Command, args []string) error {
			targetDir := "."
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
