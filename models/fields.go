package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type GORMJsonMapper map[string]interface{}

func (g *GORMJsonMapper) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("given value: %v is not []byte", value)
	}

	return json.Unmarshal(bytes, &g)
}

func (g GORMJsonMapper) Value() (driver.Value, error) {
	val, err := json.Marshal(g)
	if err != nil {
		return `{}`, err
	}
	return val, err
}