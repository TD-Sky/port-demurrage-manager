package state

import (
	"github.com/jmoiron/sqlx"
)

type State struct {
	Database *sqlx.DB
}
