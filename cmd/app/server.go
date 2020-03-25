package app

import (
	"github.com/imjuanleonard/palu-covid/pkg/logger"

	"github.com/imjuanleonard/palu-covid/pkg/server"
	"github.com/spf13/cobra"
)

func Server(cmd *cobra.Command, args []string) error {
	h := server.NewHandler()
	r := server.NewRouter(h)
	s, err := server.New(r)
	if err != nil {
		logger.Errorf("could not initialize server: %v", err)
		return err
	}

	// routes.SetupRoutes(s)
	s.StartHTTPServer()
	return nil
}
