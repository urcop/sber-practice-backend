package cli

import (
	"github.com/spf13/cobra"
	"github.com/urcop/sber-practice-backend/pkg/logger"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd() {
	c := cobra.Command{}

	c.AddCommand(NewServeCmd())

	if err := c.Execute(); err != nil {
		logger.Fatal(err.Error())
	}
}
