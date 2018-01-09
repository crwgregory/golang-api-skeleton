package records

import (
	"database/sql"
	"github.com/crwgregory/golang-api-skeleton/connection"
)

type MySQLRecord struct {
	Record
}

func (p *MySQLRecord) GetConnection() connection.ConnectionInterface {
	if p.connection == nil {
		p.connection = new(connection.MysqlConnection)
	}
	return p.connection
}

func (p *MySQLRecord) GetDb() *sql.DB {
	return p.GetConnection().GetDB().(*sql.DB)
}
