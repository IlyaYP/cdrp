package cmd

import (
	"context"

	"github.com/IlyaYP/cdrp/cmd/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newUserCmd creates a new user cmd.
func newDbCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "DB sub-commands",
	}

	cmd.AddCommand(newDbCreateTableCmd())

	return cmd
}

// newDbCreateTableCmd creates a new user.create cmd.
func newDbCreateTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createtable",
		Short: "Create Table",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init
			config := common.GetConfigFromCmdCtx(cmd)
			timeout := common.GetTimeoutToCmdCtx(cmd)

			st, err := config.BuildPsqlStorage()
			if err != nil {
				return err
			}
			defer func() {
				if err := st.Close(); err != nil {
					log.Error().Err(err).Msg("Shutting down the app")
				}
			}()

			// Exec
			ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
			defer ctxCancel()

			if err := st.CreateTableRecords(ctx); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
