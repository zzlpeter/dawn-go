package libs

import "testing"

func TestJson2Map(t *testing.T) {
	j := "{\"age\":12,\"name\":\"alex\"}"
	_, err := Json2Map(j)
	if err != nil {
		t.Errorf("Json2Map fails")
	}
}

func TestMap2Json(t *testing.T) {
	m := map[string]interface{}{
		"name": "alex",
		"age": 24,
	}
	_, err := Map2Json(m)
	if err != nil {
		t.Errorf("Map2Json fails")
	}
}

type User struct {
	Name	string		`json:"name"`
}

func TestStruct2Json(t *testing.T) {
	user := User{"alex"}
	_, err := Struct2Json(user)
	if err != nil {
		t.Errorf("Struct2Json fails")
	}
}