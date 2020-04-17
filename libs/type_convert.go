package libs

import (
	"encoding/json"
	"reflect"
)

func Map2Json(m map[string]interface{}) (string, error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func Json2Map(j string) (map[string]interface{}, error) {
	b := []byte(j)
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Struct2Json(s interface{}) (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Struct2Map(s interface{}) map[string]interface{} {
	o1 := reflect.TypeOf(s)
	o2 := reflect.ValueOf(s)

	var m = make(map[string]interface{})
	for i := 0; i < o1.NumField(); i++ {
		tags := o1.Field(i).Tag
		var field string
		if tags.Get("map") != "" {
			field = tags.Get("map")
		} else {
			field = o1.Field(i).Name
		}
		m[field] = o2.Field(i).Interface()
	}
	return m
}
