package bamboleo

import (
	"encoding/json"
	"fmt"
	"strings"
)

type entity []kv

type kv map[string]interface{}

// ParseToJSON ...
func ParseToJSON(text string) []interface{} {

	r := strings.NewReplacer("{", "[",
		"}", "]",
		"acf6192e-81ca-46ef-93a6-5a6968b78663", "\"table\"")
	str := r.Replace(text)

	var e []interface{}
	b := []byte(str)
	json.Unmarshal(b, &e)
	return e

}

// Parse ...
func Parse(text string) entity {
	return parse(text)
}

func parse(text string) entity {

	r := strings.NewReplacer("{", "[",
		"}", "]",
		"acf6192e-81ca-46ef-93a6-5a6968b78663", "\"table\"")
	str := r.Replace(text)

	var e []interface{}
	if err := fetchJSON(str, &e); err != nil {
		fmt.Print(err)
		return nil
	}

	c := e[2].([]interface{})
	fieldNames := getFieldNames(c[1].([]interface{}))

	some := c[2].([]interface{})
	lsome := len(some)
	some2 := some[lsome-3].([]interface{})
	lsome2 := len(some2)

	in := int(some2[1].(float64))
	var result = make(entity, in)
	for i := 2; i < lsome2; i++ {
		result[i-2] = getMap(some2[i].([]interface{}), fieldNames)
	}
	return result
}
func fetchJSON(text string, e *[]interface{}) error {
	b := []byte(text)
	return json.Unmarshal(b, e)
}

func getFieldNames(data []interface{}) map[int]string {
	m := make(map[int]string)
	lData := len(data)
	for i := 1; i < lData; i++ {
		item := data[i].([]interface{})
		m[i] = getString(item, 1)
	}
	return m
}

func getString(d []interface{}, i int) string {
	return d[i].(string)
}

func getMap(row []interface{}, fieldNames map[int]string) kv {
	m := make(kv)
	lRow := len(row)
	for i := 3; i < lRow-1; i++ {
		item := row[i].([]interface{})
		itemType := getString(item, 0)
		fieldName := fieldNames[i-2]
		switch itemType {
		case "S":
			m[fieldName] = getString(item, 1)
		case "N":
			m[fieldName] = item[1].(float64)
		}
	}
	return m
}
