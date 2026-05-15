package migrate_cmd

import (
	"github.com/spf13/cobra"
	down_cmd "github.com/vydon-io/vydon/backend/internal/cmds/mgmt/migrate/down"
	up_cmd "github.com/vydon-io/vydon/backend/internal/cmds/mgmt/migrate/up"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Parent command for migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(up_cmd.NewCmd())
	cmd.AddCommand(down_cmd.NewCmd())
	return cmd
}
