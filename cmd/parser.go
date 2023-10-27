package cmd

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"path"

	"github.com/IlyaYP/cdrp/cmd/common"
	"github.com/IlyaYP/cdrp/model"
	"github.com/IlyaYP/cdrp/storage"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newParseCmd parses CDR files
func newParseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "Parse CDR files",
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

			// if err := st.CreateTableRecords(ctx); err != nil {
			// 	return err
			// }

			log.Print("parser started in ", config.CdrPath)

			files, err := os.ReadDir(config.CdrPath)
			if err != nil {
				return err
			}

			for _, file := range files {
				if file.IsDir() {
					continue
				}

				filename := path.Join(config.CdrPath, file.Name())
				err := parseFile(ctx, filename, st)
				if err != nil {
					log.Err(err).Msgf("parse file %s", filename)
					continue
				}

				// if err := os.Remove(filename); err != nil {
				// 	log.Err(err).Msgf("remove file %s", filename)
				// }
			}

			return nil
		},
	}

	return cmd
}

func parseFile(ctx context.Context, filename string, st storage.RecordsStorage) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		record, err := model.NewRecordFromStrSlice(rec)
		if err != nil {
			log.Err(err).Msg("NewRecordFromStrSlice")
			continue
		}

		if _, err := st.CreateRecord(ctx, *record); err != nil {
			log.Err(err).Msg("CreateRecord")
		}

	}

	return nil
}
