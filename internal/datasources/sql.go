package datasources

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rafli024/mytodo-app/internal/constant"
	"github.com/rs/zerolog/log"
)

func Prepare(db *sqlx.DB, query string) *sqlx.Stmt {
	s, err := db.Preparex(query)
	if err != nil {
		log.Error().Stack().
			Str("Context", "preparing sql statement").
			Str("Query", query).
			Err(err).Msg("")
		os.Exit(constant.ExitPrepareStmtFail)
	}
	return s
}

func PrepareNamed(db *sqlx.DB, query string) *sqlx.NamedStmt {
	s, err := db.PrepareNamed(query)
	if err != nil {
		log.Error().Stack().
			Str("Context", "preparing sql statement").
			Str("Query", query).
			Err(err).Msg("")
		os.Exit(constant.ExitPrepareStmtFail)
	}
	return s
}
