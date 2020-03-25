package cmd

import (
	"github.com/imjuanleonard/palu-covid/cmd/app"
	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	cli := &cobra.Command{
		Use:   "palu-covid",
		Short: "Simple REST API for tracking covid status in Palu",
	}
	cli.AddCommand(newServerCommand())
	cli.AddCommand(newMigrateCommand())
	cli.AddCommand(newRollbackCommand())

	return cli
}

func newServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "server",
		Short:   "Start Server",
		Aliases: []string{"serve", "start"},
		RunE:    app.Server,
	}
}

func newMigrateCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "migrate",
		Short:   "Migrate Database",
		Aliases: []string{"migrate"},
		RunE:    app.Migrate,
	}
}

func newRollbackCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "rollback",
		Short:   "Rollback Database",
		Aliases: []string{"rollback"},
		RunE:    app.Rollback,
	}
}
