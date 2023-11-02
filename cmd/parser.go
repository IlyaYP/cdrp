package cmd

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/IlyaYP/cdrp/cmd/common"
	"github.com/IlyaYP/cdrp/model"
	"github.com/IlyaYP/cdrp/storage"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	flagDeleteFiles = "delete"
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

			deleteFlag, err := cmd.Flags().GetBool(flagDeleteFiles)
			if err != nil {
				return fmt.Errorf("app initialization: reading flag %s: %w", flagOperationTimeout, err)
			}
			log.Printf("deleteFlag is %v", deleteFlag)

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

			log.Print("parser started in ", config.CdrPath)

			files, err := os.ReadDir(config.CdrPath)
			if err != nil {
				return err
			}

			numFiles := len(files)
			log.Printf("found %v files", numFiles)

			for _, file := range files {
				select {
				case <-ctx.Done():
					log.Error().Err(ctx.Err()).Msg("Exiting")
					return nil
				default:
					if file.IsDir() {
						continue
					}

					filename := path.Join(config.CdrPath, file.Name())
					err := parseFile(ctx, filename, st)
					if err != nil {
						log.Error().Err(err).Msgf("parse file %s", filename)
						continue
					}

					if deleteFlag {
						if err := os.Remove(filename); err != nil {
							log.Error().Err(err).Msgf("remove file %s", filename)
						}
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolP(flagDeleteFiles, "d", false, "Delete parsed files")

	return cmd
}

func parseFile(ctx context.Context, filename string, st storage.RecordsStorage) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	// Read first two rows to check if it CDR format file
	rec, err := csvReader.Read()
	if err != nil {
		return err
	}
	if !model.CompareHeader(rec) {
		err := errors.New("wrong file format")
		return err
	}
	_, err = csvReader.Read()
	if err != nil {
		return err
	}

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
			log.Error().Err(err).Msgf("NewRecordFromStrSlice:%s", filename)
			continue
		}

		if _, err := st.CreateRecord(ctx, *record); err != nil {
			log.Debug().Err(err).Msgf("CreateRecord:%s", record.Pkid)
		}
	}

	return nil
}
