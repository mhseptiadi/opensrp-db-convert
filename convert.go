package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/slug"
	_ "github.com/lib/pq"

	lib "./lib"
)

const (
	DB_USER     = "opensrp_admin"
	DB_PASSWORD = "C4nT(T0ucH)Th1S"
	DB_NAME     = "opensrp"
	PORT        = "5431"
	RELPATH     = "./"
	SCHEMA      = "sid"
	LIMIT       = 1000
)

func main() {

	lib.RELPATH = RELPATH
	lib.SCHEMA = SCHEMA

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=127.0.0.1 port=%s",
		DB_USER, DB_PASSWORD, DB_NAME, PORT)

	// dbinfo := fmt.Sprintf("postgres://%s:%s@127.0.0.1/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	lib.CheckErr(err)
	defer db.Close()

	eventID := lib.LastEventID()
	limit := LIMIT

	fmt.Println("# Querying event id > " + strconv.Itoa(eventID))

	mainQuery := fmt.Sprintf(` select 
	core.event_metadata.document_id, core.event_metadata.base_entity_id, core.event_metadata.location_id,
	core.event.id, core.event.json, core.event_metadata.event_type,
	core.client_metadata.first_name, core.client_metadata.last_name, core.client_metadata.birth_date,
	core.event_metadata.date_created, core.event_metadata.event_date, core.event_metadata.provider_id
	from core.event_metadata 
	join core.event on core.event.id = core.event_metadata.id 
	join core.client_metadata on core.event_metadata.base_entity_id = core.client_metadata.base_entity_id 
	where core.event.id > %d order by core.event.id asc limit %d`, eventID, limit)
	rows, err := db.Query(mainQuery)

	batchQuery := ""

	lib.CheckErr(err)
	for rows.Next() {
		var event_id int
		var document_id string
		var base_entity_id string
		var location_id string
		var jsonString string
		var event_type string
		var first_name *string
		var last_name *string
		var birth_date *string
		var date_created string
		var event_date string
		var provider_id string
		err = rows.Scan(&document_id, &base_entity_id, &location_id, &event_id, &jsonString, &event_type, &first_name, &last_name, &birth_date, &date_created, &event_date, &provider_id)
		eventID = event_id
		lib.CheckErr(err)

		// fmt.Println(jsonString)
		var eventData lib.Event
		json.Unmarshal([]byte(jsonString), &eventData)
		// fmt.Println(eventData)

		insertQuery := `
		INSERT INTO "` + SCHEMA + `"."tableName" 
		`

		tableName := slug.Make(event_type)
		query := strings.Replace(insertQuery, "tableName", tableName, -1)
		query = strings.Replace(query, "-", "_", -1)

		clientVersionTime := time.Unix(0, eventData.Version*int64(time.Millisecond))
		clientVersionSubmissionDate := strings.Replace(clientVersionTime.String(), " WIB", "", -1)
		serverVersionTime := time.Unix(0, eventData.Server_version*int64(time.Millisecond))
		serverVersionSubmissionDate := strings.Replace(serverVersionTime.String(), " WIB", "", -1)

		queryFields := `(
			"first_name", "last_name", "birth_date", 
			"event_id", "document_id", "base_entity_id", "location_id",
			"date_created", "event_date", "clientversionsubmissiondate", "serverversionsubmissiondate", 
			"provider_id",  `
		queryValues := fmt.Sprintf("VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%v', '%v', '%v', '%v', '%s', ",
			lib.IsNull(first_name), lib.IsNull(last_name), lib.IsNull(birth_date),
			strconv.Itoa(event_id), document_id, base_entity_id, location_id,
			date_created, event_date, clientVersionSubmissionDate, serverVersionSubmissionDate,
			provider_id)

		fields := lib.EventParser(eventData.Obs)

		for key, val := range fields {
			queryFields += fmt.Sprintf("\"%s\", ", key)
			queryValues += fmt.Sprintf("'%s', ", val)
		}

		queryFields = queryFields[:len(queryFields)-2] + ") "
		queryValues = queryValues[:len(queryValues)-2] + "); "
		query += queryFields + queryValues

		fmt.Println(eventID)
		fmt.Println(tableName)

		batchQuery += query

		//check client query
		if tableName == "identitas-ibu" {
			fmt.Println("insert client ibu")
			query = fmt.Sprintf(lib.InsertClientIbu(), document_id, date_created, base_entity_id, "", lib.IsNull(first_name), lib.IsNull(last_name),
				fields["province"], fields["district"], fields["sub_district"], fields["village"], fields["sub_village"],
				lib.IsNull(birth_date), "", "", provider_id)
			// fmt.Println(query)
			batchQuery += query
		}
		if tableName == "child-registration" {
			fmt.Println("insert client anak")
			q1 := `select core.event.json from core.event join core.event_metadata on core.event.id = core.event_metadata.id
			where event_metadata.base_entity_id = '` + base_entity_id + `' and core.event_metadata.event_type = 'Child Registration'`
			var json1 string
			db.QueryRow(q1).Scan(&json1)
			var event1 lib.Event
			json.Unmarshal([]byte(json1), &event1)
			fields1 := lib.EventParser(event1.Obs)

			q2 := `select core.client.json from core.client join core.client_metadata on core.client.id = core.client_metadata.id
			where client_metadata.base_entity_id = '` + base_entity_id + `'`
			var json2 string
			db.QueryRow(q2).Scan(&json2)
			var client2 lib.Client
			json.Unmarshal([]byte(json2), &client2)

			gender := client2.Gender
			ibucaseid := client2.Relationships.IbuCaseId[0]
			namabayi := fields1["namabayi"]

			query = fmt.Sprintf(lib.InsertClientAnak(), document_id, date_created, base_entity_id, "",
				lib.IsNull(birth_date), gender, ibucaseid, provider_id, namabayi)

			// fmt.Println(query)
			batchQuery += query
		}
		if tableName == "edit-ibu" {
			fmt.Println("update client ibu")
			query = fmt.Sprintf(lib.EditClientIbu(), fields["province"], fields["district"], fields["sub_district"], fields["village"], fields["sub_village"], base_entity_id)
			// fmt.Println(query)
			batchQuery += query
		}
		if tableName == "edit-bayi" {
			fmt.Println("update client anak")
			query = fmt.Sprintf(lib.EditClientAnak(), fields["namabayi"], base_entity_id)
			// fmt.Println(query)
			batchQuery += query
		}

	}

	fmt.Println("RUN batchQuery")
	// fmt.Println(batchQuery)

	_, err = db.Query(batchQuery)
	lib.CheckErr(err)

	lib.WriteEventID(eventID)
}
