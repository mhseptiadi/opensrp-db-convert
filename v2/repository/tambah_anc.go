package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	lib "../lib"
	model "../model"
)

type (
	TambahAnc interface {
		Run()
	}

	tambahAnc struct {
		db *sql.DB
	}
)

func InitTambahAnc(d *sql.DB) TambahAnc {
	return &tambahAnc{
		db: d,
	}
}

func (ta tambahAnc) Run() {
	eventID := lib.LastEventID("tambah_anc")
	limit := LIMIT

	fmt.Println("# Querying event id > " + strconv.Itoa(eventID))

	mainQuery := fmt.Sprintf(` 
	select *
	from sid3.tambah_anc 
	where CAST(nullif(tambah_anc.event_id, '') AS integer) > %d
	order by CAST(nullif(tambah_anc.event_id, '') AS integer) asc 
	limit %d`, eventID, limit)
	rows, err := ta.db.Query(mainQuery)

	lib.CheckErr(err)

	insertQuery := `
	INSERT INTO "sid_v2"."tambah_anc"(FIELDS_REPLACE) VALUES(VALUES_REPLACE) ;
	`

	columns, _ := rows.Columns()

	batchQuery := ""

	for rows.Next() {

		queryFields := ""
		queryValues := ""

		row := make([]interface{}, len(columns))
		for idx := range columns {
			row[idx] = new(lib.MetalScanner)
		}

		err := rows.Scan(row...)
		if err != nil {
			fmt.Println(err)
		}

		for idx, column := range columns {
			var scanner = row[idx].(*lib.MetalScanner)
			if lib.Contains(model.Tambah_anc_v2, column) {
				queryFields += `"` + column + `", `
				if scanner.Type == "nil" {
					queryValues += `NULL, `
				} else if scanner.Type == "time.Time" {
					timeStr := scanner.Value.(string)
					timeArr := strings.Split(timeStr, " +")
					queryValues += `'` + timeArr[0] + `', `
				} else {
					queryValues += `'` + scanner.Value.(string) + `', `
				}

				if column == "penyakitkronis" && scanner.Type != "nil" {
					namaPenyakit := scanner.Value.(string)
					for _, val := range model.Tambah_anc_v2 {
						if strings.Contains(namaPenyakit, val) {
							queryFields += `"` + val + `", `
							queryValues += `TRUE, `
						}
					}
				}
			}

			if column == "event_id" {
				eventID, err = strconv.Atoi(scanner.Value.(string))
			}
			// fmt.Println(column, ":", scanner.Value)
		}

		queryFields = queryFields[:len(queryFields)-2]
		queryValues = queryValues[:len(queryValues)-2]

		query := strings.Replace(insertQuery, "FIELDS_REPLACE", queryFields, -1)
		query = strings.Replace(query, "VALUES_REPLACE", queryValues, -1)

		batchQuery += query
	}

	_, err = ta.db.Query(batchQuery)
	lib.CheckErr(err)

	lib.WriteEventID(eventID, "tambah_anc")
}
