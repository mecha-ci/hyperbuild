package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mecha-ci/hyperbuild/internal/app"
	"github.com/mecha-ci/hyperbuild/internal/config"
	"github.com/mecha-ci/hyperbuild/internal/model"
)

var (
	ErrCannotParseYAMLFile = fmt.Errorf("cannot parse YAML file")
	ErrCannotRunModel      = fmt.Errorf("cannot run the model")
)

func NewRunCommand(ctr *app.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run the build of the given file",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			m, err := config.ParseYAMLFile(args[0])
			if err != nil {
				return fmt.Errorf("%w: %v", ErrCannotParseYAMLFile, err)
			}

			res, err := model.Run(m)
			if err != nil {
				return fmt.Errorf("%w: %v", ErrCannotRunModel, err)
			}

			fmt.Printf("done: %v", res)

			return nil
		},
	}
}
