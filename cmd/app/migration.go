package app

import (
	"github.com/imjuanleonard/palu-covid/pkg/db"
	"github.com/imjuanleonard/palu-covid/pkg/logger"
	"github.com/spf13/cobra"
)

func Migrate(_ *cobra.Command, _ []string) error {
	if err := db.Migrate(); err != nil {
		logger.Errorf("error while running migration: %+v", err)
	}
	logger.Info("migration success")
	return nil
}

func Rollback(_ *cobra.Command, _ []string) error {
	if err := db.Rollback(); err != nil {
		logger.Errorf("error while running rollback: %+v", err)
	}
	logger.Info("rollback success")
	return nil
}
