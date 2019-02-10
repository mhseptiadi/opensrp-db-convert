package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "opensrp_admin"
	DB_PASSWORD = "C4nT(T0ucH)Th1S"
	DB_NAME     = "opensrp"
	PORT        = "5431"
	SCHEMA      = "sid"
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

	Generate := `CREATE TABLE IF NOT EXISTS ` + SCHEMA + `.tableName (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		`

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query(`SELECT event_type
	FROM core.event_metadata
	GROUP BY event_type;`)
	checkErr(err)

	for rows.Next() {
		var event_type string
		err = rows.Scan(&event_type)
		// fmt.Println(event_type)

		mainQuery := fmt.Sprintf(" select core.event.json from core.event join core.event_metadata on core.event.id = core.event_metadata.id where event_metadata.event_type = '%s' order by core.event.id desc", event_type)

		fields := make(map[string]string)

		subrows, err := db.Query(mainQuery)
		checkErr(err)
		for subrows.Next() {
			var jsonString string
			err = subrows.Scan(&jsonString)

			var eventData Event
			json.Unmarshal([]byte(jsonString), &eventData)

			//unique check
			for _, val := range eventData.Obs {
				// var re = regexp.MustCompile(`([a-z])([A-Z])`)
				// fieldName := re.ReplaceAllString(val.Obs_form_submission_field, `$1-$2`)
				// fieldName = strings.Replace(fieldName, "-", "_", -1)
				// fieldName = strings.ToLower(fieldName)

				fieldName := strings.ToLower(val.Obs_form_submission_field)
				fields[fieldName] = fieldName
			}
		}

		tableName := slug.Make(event_type)

		query := strings.Replace(Generate, "tableName", tableName, -1)
		query = strings.Replace(query, "-", "_", -1)

		for _, val := range fields {
			query += strings.Replace(val, "-", "_", -1) + " character varying(75),\n"
		}

		query = query[:len(query)-2] + ");"
		fmt.Println(query)

		_, err = db.Query(query)
		checkErr(err)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		// fmt.Println("========================================================================================================================")

	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
