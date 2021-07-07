package lib

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

var (
	SCHEMA  = ""
	RELPATH = ""
)

func IsNull(str *string) string {
	if str != nil {
		return *str
	}
	return string("")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func LastEventID(key string) (id int) {
	b, err := ioutil.ReadFile(RELPATH + key + ".id")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	i, _ := strconv.Atoi(str)
	return i
}

func WriteEventID(id int, key string) error {
	d1 := []byte(strconv.Itoa(id))
	err := ioutil.WriteFile(RELPATH+key+".id", d1, 0644)
	return err
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//https://github.com/lib/pq/issues/415
type MetalScanner struct {
	Value interface{}
	Type  string
	Valid bool
}

func (scanner *MetalScanner) getBytes(src interface{}) []byte {
	if a, ok := src.([]uint8); ok {
		return a
	}
	return nil
}
func (scanner *MetalScanner) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		if value, ok := src.(int64); ok {
			scanner.Value = value
			scanner.Valid = true
			scanner.Type = "int64"
		}
	case float64:
		if value, ok := src.(float64); ok {
			scanner.Value = value
			scanner.Valid = true
			scanner.Type = "float64"
		}
	case bool:
		if value, ok := src.(bool); ok {
			scanner.Value = value
			scanner.Valid = true
			scanner.Type = "bool"
		}
	case string:
		// value := scanner.getBytes(src)
		// scanner.Value = string(value)
		scanner.Value = src.(string)
		scanner.Valid = true
		scanner.Type = "string"
	case []byte:
		value := scanner.getBytes(src)
		scanner.Value = value
		scanner.Valid = true
		scanner.Type = "[]byte"
	case time.Time:
		if value, ok := src.(time.Time); ok {
			scanner.Value = value.String() //convert to string
			scanner.Valid = true
			scanner.Type = "time.Time"
		}
	case nil:
		scanner.Value = nil
		scanner.Valid = true
		scanner.Type = "nil"
	}
	return nil
}
