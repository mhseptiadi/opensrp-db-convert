package lib

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	SCHEMA  = ""
	RELPATH = ""
)

type Client struct {
	Gender        string `json:"gender"`
	Relationships struct {
		IbuCaseId []string `json:"ibuCaseId"`
	}
}

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
	Server_version     int64        `json:"serverVersion"`
	Type               string       `json:"type"`
	Version            int64        `json:"version"`
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

func IsNull(str *string) string {
	if str != nil {
		return *str
	}
	return string("")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func LastEventID() (id int) {
	b, err := ioutil.ReadFile(RELPATH + "event_id.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	i, _ := strconv.Atoi(str)
	return i
}

func WriteEventID(id int) error {
	d1 := []byte(strconv.Itoa(id))
	err := ioutil.WriteFile(RELPATH+"event_id.txt", d1, 0644)
	return err
}

func EventParser(obsData []Obs_Struct) map[string]string {
	fields := make(map[string]string)

	for _, ObsData := range obsData {

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

	return fields
}

func InsertClientIbu() string {
	return `INSERT INTO "` + SCHEMA + `"."client_ibu"
("docid", "datecreated", "baseentityid", "uniqueid", "namalengkap", "namasuami", "provinsi", "kabupaten", "kecamatan", "desa", "dusun", "birthdate", "nik", "noibu", "providerid") VALUES('%s', '%v', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%v', '%s', '%s', '%s');`
}
func EditClientIbu() string {
	return `UPDATE "` + SCHEMA + `"."client_ibu" SET "provinsi"='%s', "kabupaten"='%s', "kecamatan"='%s', "desa"='%s', "dusun"='%s' WHERE "client_ibu"."baseentityid" = '%s';`
}
func InsertClientAnak() string {
	return `INSERT INTO "` + SCHEMA + `"."client_anak"("docid", "datecreated", "baseentityid", "uniqueid",
"birthdate", "gender", "ibucaseid", "providerid", "namabayi") VALUES('%s', '%v', '%s', '%s', '%v', '%s', '%s', '%s', '%s');`
}
func EditClientAnak() string {
	return `UPDATE "` + SCHEMA + `"."client_anak" SET "namabayi"='%s' WHERE "client_anak"."baseentityid" = '%s';`
}
