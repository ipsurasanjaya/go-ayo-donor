package pql

import (
	"database/sql"
	"go-ayo-donor/model/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"github.com/newrelic/go-agent/v3/integrations/nrpgx5"
)

func CreateSQLDB(cfg domain.Config) (db *sql.DB, err error) {
	connstr := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName + "?sslmode=" + cfg.SSLMode + "&timezone=" + cfg.TimeZone
	cf, err := pgx.ParseConfig(connstr)
	if err != nil {
		return nil, err
	}
	cf.Tracer = nrpgx5.NewTracer()
	db = stdlib.OpenDB(*cf)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return
}
