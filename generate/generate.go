package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gosimple/slug"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "opensrp_admin"
	DB_PASSWORD = "C4nT(T0ucH)Th1S"
	DB_NAME     = "opensrp"
	PORT        = "5431"
	SCHEMA      = "sid3"
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

type form_json struct {
	Title string `json:"title"`
}

type form_definition_json struct {
	Form struct {
		Fields []struct {
			Name string `json:"name"`
			Bind string `json:"bind"`
		} `json:"fields"`
	} `json:"form"`
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

	var files []string
	root := "./form"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {

		if file != "./form" {
			// fmt.Println(file)
			tableName := ""
			dat, _ := ioutil.ReadFile(file + "/form.json")
			if dat != nil {
				var jsonArr form_json
				err = json.Unmarshal(dat, &jsonArr)
				checkErr(err)
				tableName = jsonArr.Title
			} else {
				dat, _ := ioutil.ReadFile(file + "/form.xml")
				if dat != nil {
					r := regexp.MustCompile("<h3 id=\"form-title\">(.+)</h3>")
					matches := r.FindStringSubmatch(string(dat))
					if len(matches) > 1 {
						tableName = matches[1]
					}
				}
			}

			tableName = slug.Make(tableName)
			tableName = strings.Replace(tableName, "-", "_", -1)
			tableName = strings.ToLower(tableName)
			query := strings.Replace(Generate, "tableName", tableName, -1)
			query = strings.Replace(query, "-", "_", -1)
			// fmt.Println("tableName " + tableName)

			dat, err = ioutil.ReadFile(file + "/form_definition.json")
			checkErr(err)
			var fieldsArr form_definition_json
			err = json.Unmarshal(dat, &fieldsArr)
			checkErr(err)

			fields := make(map[string]string)

			for _, fieldPath := range fieldsArr.Form.Fields {
				if fieldPath.Bind != "" {
					// fieldPathArr := strings.Split(fieldPath.Bind, "/")
					// count := len(fieldPathArr)
					// fieldName := fieldPathArr[count-1]
					fieldName := fieldPath.Name
					if fieldName != "start" && fieldName != "end" {
						// example IMPORTANT
						// refleks_patela_ibu >> reflekspatelaibu
						// sub-village >> sub_village
						fieldName = strings.Replace(fieldName, "_", "", -1)
						fieldName = strings.Replace(fieldName, "-", "_", -1)
						fieldName = strings.ToLower(fieldName)
						// fmt.Println("fieldName " + fieldName)
						fields[fieldName] = fieldName // to make sure it unique
					}
				}
			}

			for _, fieldName := range fields {
				query += fieldName + " text,\n"
			}

			query = query[:len(query)-2] + ");"
			fmt.Println(query)
			fmt.Println("---------------------------------------")

			// _, err := db.Query(query)
			// checkErr(err)

		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
