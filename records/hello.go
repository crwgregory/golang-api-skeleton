package records

import (
	"fmt"
	"github.com/crwgregory/golang-api-skeleton/utils"
	"strings"
)

// Hello The record struct
// the json meta data is used for the name of the column in the db, as well as returning the data in the api response
type Hello struct {
	MySQLRecord
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	// queries
	find_one_hello = "SELECT %s FROM hello WHERE id = %d"
)

// Load the record
func (h *Hello) Load(id int) (err error) {
	// get the columns as a []string
	columns := utils.GetRecordColumns(h)
	// format the query
	query := fmt.Sprintf(find_one_hello, strings.Join(columns, ","), id)
	// get the sql connection
	db := h.GetDb()
	// query for the data
	row := db.QueryRow(query)
	// parse the data from the row
	data, err := utils.ScanRow(row, columns)
	// set the data of the record
	utils.SetData(h, data)
	return
}
