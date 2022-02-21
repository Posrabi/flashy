package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// nolint:gochecknoinits
func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func New() *cobra.Command {
	app := &cobra.Command{
		Use:   "users",
		Short: "users backend service",
	}

	app.AddCommand(newClientCmd())
	app.AddCommand(newServerCmd())

	return app
}
