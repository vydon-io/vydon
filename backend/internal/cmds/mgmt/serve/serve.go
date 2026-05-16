package serve

import (
	"github.com/spf13/cobra"
	serve_connect "github.com/vydon-io/vydon/backend/internal/cmds/mgmt/serve/connect"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Parent command for serving",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(serve_connect.NewCmd())
	return cmd
}
