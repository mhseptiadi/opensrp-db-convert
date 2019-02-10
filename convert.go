package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gosimple/slug"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "opensrp_admin"
	DB_PASSWORD = "C4nT(T0ucH)Th1S"
	DB_NAME     = "opensrp"
	PORT        = "5431"
	RELPATH     = "./"
)

type Event struct {
	Id                 string       `json:"_id"`
	Rev                string       `json:"_rev"`
	Base_entity_id     string       `json:"baseEntityId"`
	Date_created       string       `json:"dateCreated"`
	Duration           int          `json:"duration"`
	Entity_type        string       `json:"entityType"`
	Event_date         string       `json:"eventDate"`
	Event_type         string       `json:"eventType"`
	Form_submission_id string       `json:"formSubmissionId"`
	Location_id        string       `json:"locationId"`
	Provider_id        string       `json:"providerId"`
	Server_version     int          `json:"serverVersion"`
	Type               string       `json:"type"`
	Version            int          `json:"version"`
	Obs                []Obs_Struct `json:"obs"`
}

type Obs_Struct struct {
	Obs_relational_id         string
	Obs_field_code            string   `json:"fieldCode"`
	Obs_field_data_type       string   `json:"fieldDataType"`
	Obs_field_type            string   `json:"fieldType"`
	Obs_form_submission_field string   `json:"formSubmissionField"`
	Obs_human_readable_values []string `json:"humanReadableValues"`
	Obs_parent_code           string   `json:"parentCode"`
	Obs_values                []string `json:"values"`
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=127.0.0.1 port=%s",
		DB_USER, DB_PASSWORD, DB_NAME, PORT)

	// dbinfo := fmt.Sprintf("postgres://%s:%s@127.0.0.1/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	eventID := lastEventID()
	limit := 1000

	fmt.Println("# Querying event id > " + strconv.Itoa(eventID))

	mainQuery := fmt.Sprintf(` select 
	core.event_metadata.document_id, core.event_metadata.base_entity_id,
	core.event.id, core.event.json, core.event_metadata.event_type,
	core.client_metadata.first_name, core.client_metadata.last_name, core.client_metadata.birth_date,
	core.event_metadata.date_created, core.event_metadata.event_date, core.event_metadata.provider_id
	from core.event_metadata 
	join core.event on core.event.id = core.event_metadata.id 
	join core.client_metadata on core.event_metadata.base_entity_id = core.client_metadata.base_entity_id 
	where core.event.id > %d order by core.event.id asc limit %d`, eventID, limit)
	rows, err := db.Query(mainQuery)

	batchQuery := ""

	checkErr(err)
	for rows.Next() {
		var event_id int
		var document_id string
		var base_entity_id string
		var jsonString string
		var event_type string
		var first_name *string
		var last_name *string
		var birth_date *string
		var date_created string
		var event_date string
		var provider_id string
		err = rows.Scan(&document_id, &base_entity_id, &event_id, &jsonString, &event_type, &first_name, &last_name, &birth_date, &date_created, &event_date, &provider_id)
		eventID = event_id
		checkErr(err)

		// fmt.Println(jsonString)
		var eventData Event
		json.Unmarshal([]byte(jsonString), &eventData)
		// fmt.Println(eventData)

		insertQuery := `
		INSERT INTO "sid2"."tableName" 
		`

		tableName := slug.Make(event_type)
		query := strings.Replace(insertQuery, "tableName", tableName, -1)
		query = strings.Replace(query, "-", "_", -1)

		queryFields := `("first_name", "last_name", "birth_date", "event_id", "date_created", "event_date", "provider_id", "document_id", "base_entity_id", `
		queryValues := fmt.Sprintf("VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', ", isNull(first_name), isNull(last_name), isNull(birth_date), strconv.Itoa(event_id), date_created, event_date, provider_id, document_id, base_entity_id)

		fields := make(map[string]string)

		for _, ObsData := range eventData.Obs {

			Obs_human_readable_values := ""
			for _, Val := range ObsData.Obs_human_readable_values {
				Obs_human_readable_values += Val
			}

			Obs_values := ""
			for _, Val := range ObsData.Obs_values {
				Obs_values += Val
			}

			values := Obs_values
			if Obs_human_readable_values != "" {
				values = Obs_human_readable_values
			}

			fields[strings.Replace(strings.ToLower(ObsData.Obs_form_submission_field), "-", "_", -1)] = values

		}

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
		// fmt.Println(query)
		// _, err = db.Query(query)
		// checkErr(err)

		// writeEventID(eventID)

	}

	fmt.Println("RUN batchQuery")

	_, err = db.Query(batchQuery)
	checkErr(err)

	writeEventID(eventID)
}

func isNull(str *string) string {
	if str != nil {
		return *str
	}
	return string("")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func lastEventID() (id int) {
	b, err := ioutil.ReadFile(RELPATH + "event_id.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	i, _ := strconv.Atoi(str)
	return i
}

func writeEventID(id int) error {
	d1 := []byte(strconv.Itoa(id))
	err := ioutil.WriteFile(RELPATH+"event_id.txt", d1, 0644)
	return err
}
