package lib

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

var (
	RELPATH = ""
)

func SetRELPATH(val string) {
	RELPATH = val
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

func InsertClientIbu() string {
	query := `INSERT INTO "sid"."client_ibu"
	("docid", "datecreated", "baseentityid", "uniqueid", "namalengkap", "namasuami", "provinsi", "kabupaten", "kecamatan", "desa", "dusun", "birthdate", "nik", "noibu", "providerid") VALUES('%s', '%v', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%v', '%s', '%s', '%s');`
	return query
}
