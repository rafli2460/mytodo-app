package datasources

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rafli024/mytodo-app/internal/constant"
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rs/zerolog/log"
)

func Init(cfg map[string]string) *contract.Datasources {
	var err error
	var writerDB *sqlx.DB
	var readerDB *sqlx.DB

	dsWriter, dsReader := parseDB(cfg)

	if writerDB, err = sqlx.Connect(cfg[constant.DBDriver], dsWriter); err == nil {
		writerDB.SetConnMaxLifetime(time.Duration(1) * time.Hour)
		writerDB.SetMaxOpenConns(10)
		writerDB.SetMaxIdleConns(10)

		log.Log().Msg("Initializing DB Writer : pass")
	} else {
		log.Panic().
			Str("Context", "Connecting to Writer DB").
			Err(err).Msg("")
	}

	if readerDB, err = sqlx.Connect(cfg[constant.DBDriver], dsReader); err == nil {
		readerDB.SetConnMaxLifetime(time.Duration(1) * time.Hour)
		readerDB.SetMaxOpenConns(10)
		readerDB.SetMaxIdleConns(10)

		log.Log().Msg("Initializing DB Reader : pass")
	} else {
		log.Panic().
			Str("Context", "Connecting to Reader DB").
			Err(err).Msg("")
	}

	ds := &contract.Datasources{
		WriterDB: writerDB,
		ReaderDb: readerDB,
	}

	return ds

}

func parseDB(config map[string]string) (dsWriter, dsReader string) {
	hostWriter := config[constant.DBHostWriter]
	hostReader := config[constant.DBHostReader]
	user := config[constant.DBUsername]
	pass := config[constant.DBPassword]
	name := config[constant.DBName]
	port := config[constant.DBPort]

	dsWriter = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", user, pass, hostWriter, port, name)
	dsReader = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", user, pass, hostReader, port, name)

	return
}
