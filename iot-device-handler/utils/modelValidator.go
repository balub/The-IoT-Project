package utils

import (
	"fmt"
	"reflect"
)

func ModelValidator(fieldDefs map[string]map[string]interface{}, userData map[string]interface{}) error {
	// Iterate over the field definitions
	for fieldName, fieldDef := range fieldDefs {
		// Check if the field is required
		required, ok := fieldDef["required"].(bool)
		if ok && required {
			// Field is required, check if it exists in the user data
			val, ok := userData[fieldName]
			if !ok || val == nil {
				return fmt.Errorf("missing required field: %s", fieldName)
			}

			// Check if the field type matches the expected type
			fieldType, ok := fieldDef["type"].(string)

			fmt.Println(fieldType, reflect.TypeOf(val).Name(), val)
			if !ok || reflect.TypeOf(val).Name() != fieldType {
				return fmt.Errorf("field %s has incorrect type", fieldName)
			}
		}
	}

	return nil
}
