package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/IlyaYP/cdrp/cmd/common"
	"github.com/IlyaYP/cdrp/pkg/axl"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newUserCmd creates a new user cmd.
func newDbCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "DB sub-commands",
	}

	cmd.AddCommand(newMigrateCmd())
	cmd.AddCommand(newDbCreateTableCmd())
	cmd.AddCommand(newDbDropTableCmd())
	cmd.AddCommand(newDbResetTableCmd())
	cmd.AddCommand(newDbQueryCUCM())

	return cmd
}

// newDbQueryCUCM creates a new user.create cmd.
func newDbQueryCUCM() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "querycucm",
		Short: "Query CUCM",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init
			config := common.GetConfigFromCmdCtx(cmd)
			// timeout := common.GetTimeoutToCmdCtx(cmd)

			// log.Printf("deleteFlag is %v", deleteFlag)
			// log.Printf("axl is %s", config.AXL.CUCM)
			// log.Printf("DSN is %s", config.PSQLStorage.DSN)

			// st, err := config.BuildPsqlStorage()
			// if err != nil {
			// 	return err
			// }
			// defer func() {
			// 	if err := st.Close(); err != nil {
			// 		log.Error().Err(err).Msg("Shutting down the app")
			// 	}
			// }()

			// Exec
			// ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
			// defer ctxCancel()

			query := "<executeSQLQuery><sql>select pkid,name,description from device</sql></executeSQLQuery>"
			r := strings.NewReader(query)

			client := config.BuildAXLClient()

			resp, err := client.AXLRequest(r)
			result := "success"
			if err != nil {
				//if err, ok := err.(*axl.AXLError); ok{
				var e *axl.AXLError
				if errors.As(err, &e) {
					result = fmt.Sprintf("%s (%d)", e.AXLErrorMessage, e.AXLErrorCode)
				} else {
					log.Error().Err(err).Msg("error from AXLRequest")
					return err
				}
			}

			log.Printf("Result: %s", result)

			return nil
		},
	}

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

// newDbDropTableCmd creates a new user.create cmd.
func newDbDropTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "droptable",
		Short: "Drop Table",
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

			if err := st.DropTableRecords(ctx); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

// newDbResetTableCmd creates a new user.create cmd.
func newDbResetTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resettable",
		Short: "Reset Table",
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

			if err := st.ResetTableRecords(ctx); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
