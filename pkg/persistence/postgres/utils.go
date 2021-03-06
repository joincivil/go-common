package postgres // import "github.com/joincivil/go-common/pkg/persistence/postgres"

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// If this is the field name in a struct db tag, it should be ignored
	// ex. `db:"-"`
	ignoredFieldName = "-"
)

// JsonbPayload is the jsonb payload
type JsonbPayload map[string]interface{}

// Value is the value interface implemented for the sql driver
func (jp JsonbPayload) Value() (driver.Value, error) {
	// Ensure the payload gets converted properly,
	// particularly proper case for common.Address
	toMarshal := map[string]interface{}{}
	for key, val := range jp {
		switch v := val.(type) {
		case common.Address:
			toMarshal[key] = v.Hex()
		default:
			toMarshal[key] = v
		}
	}
	return json.Marshal(toMarshal)
}

// Scan is the scan interface implemented for the sql driver
func (jp *JsonbPayload) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}
	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*jp, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

// StructFieldsForQuery is a generic Insert statement for any table
// tablePrefix="" is used to specify table name, i.e. "l" to get "l.name", etc.
func StructFieldsForQuery(exampleStruct interface{}, colon bool, tablePrefix string) (string, string) {
	var fields bytes.Buffer
	var fieldsWithColon bytes.Buffer
	valStruct := reflect.ValueOf(exampleStruct)
	typeOf := valStruct.Type()
	for i := 0; i < valStruct.NumField(); i++ {
		dbFieldName := typeOf.Field(i).Tag.Get("db")
		// Skip ignored fields
		if strings.TrimSpace(dbFieldName) == ignoredFieldName {
			continue
		}
		if tablePrefix != "" {
			fields.WriteString(fmt.Sprintf("%v.", tablePrefix)) // nolint: gosec
		}
		fields.WriteString(dbFieldName) // nolint: gosec
		if colon {
			fieldsWithColon.WriteString(":")         // nolint: gosec
			fieldsWithColon.WriteString(dbFieldName) // nolint: gosec
		}
		if i+1 < valStruct.NumField() {
			fields.WriteString(", ") // nolint: gosec
			if colon {
				fieldsWithColon.WriteString(", ") // nolint: gosec
			}
		}
	}
	return strings.Trim(fields.String(), ", "),
		strings.Trim(fieldsWithColon.String(), ", ")
}

// InsertIntoDBQueryString creates the query to insert a given struct into a given table
func InsertIntoDBQueryString(tableName string, dbModelStruct interface{}) string {
	fieldNames, fieldNamesColon := StructFieldsForQuery(dbModelStruct, true, "")
	queryString := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s);", tableName, fieldNames, fieldNamesColon) // nolint: gosec
	return queryString
}

// DbFieldNameFromModelName gets the field name from db given postgres model struct
func DbFieldNameFromModelName(exampleStruct interface{}, fieldName string) (string, error) {
	sType := reflect.TypeOf(exampleStruct)
	field, ok := sType.FieldByName(fieldName)
	if !ok {
		return "", fmt.Errorf("%s may not exist in struct", fieldName)
	}
	return field.Tag.Get("db"), nil
}
