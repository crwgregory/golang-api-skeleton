package records

import (
	"github.com/crwgregory/golang-api-skeleton/connection"
	"database/sql"
)

type Record struct {
	connection connection.ConnectionInterface
}

func (p *Record) GetConnection() connection.ConnectionInterface {
	if p.connection == nil {
		p.connection = new(connection.MysqlConnection)
	}
	return p.connection
}

func (p *Record) GetSqlDb() *sql.DB {
	return p.GetConnection().GetDB().(*sql.DB)
}