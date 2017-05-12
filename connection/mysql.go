package connection

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

var db *sql.DB


const default_port = 3306

type MysqlConnection struct {
	host string
	user string
	password string
	database_name string
	port int
}

func (con *MysqlConnection) Open() bool {

	con.loadDbSettings()

	settings := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", con.user, con.password, con.host, con.port, con.database_name)

	connection, err := sql.Open("mysql", settings)
	if err != nil {
		panic(err)
	}
	if err := connection.Ping(); err != nil {
		panic(err)
	}
	db = connection
	return true
}

func (con *MysqlConnection) Close() bool {
	db.Close()
	return true
}

func (con *MysqlConnection) GetDB() interface{} {
	if db == nil {
		con.Open()
	}
	return db
}

// loadDbSettings loads the db.json file from the systems home directory/.api-skeleton/ folder
func (m *MysqlConnection) loadDbSettings() {

	homeDir := os.Getenv("HOME") // *nix
	if homeDir == "" {               // Windows
		homeDir = os.Getenv("USERPROFILE")
	}
	if homeDir == "" {
		panic("Can not find HOME env directory")
	}

	f := filepath.Join(homeDir, ".api-skeleton", "db.json")

	file, e := ioutil.ReadFile(f)
	if e != nil {
		panic(e)
	}

	config := make(map[string]interface{})

	json.Unmarshal(file, &config)

	port, ok := config["port"]
	if !ok {
		port = default_port
	}
	m.host = config["host"].(string)
	m.user = config["user"].(string)
	m.password = config["password"].(string)
	m.database_name = config["database_name"].(string)

	m.port = port.(int)
}
