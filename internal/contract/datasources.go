package contract

import "github.com/jmoiron/sqlx"

type Datasources struct {
	WriterDB *sqlx.DB `json:"writer-db"`
	ReaderDb *sqlx.DB `json:"reader-db"`
}
