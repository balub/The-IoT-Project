package utils

import (
	"fmt"
	"reflect"
)

type DefiniedModel struct {
	Name     string
	Type     string
	Required bool
}

func ModelValidator(fieldDefs []DefiniedModel, userData map[string]interface{}) (map[string]interface{}, error) {
	resData := make(map[string]interface{})

	for i := 0; i < len(fieldDefs); i++ {

		// check if it's required
		if fieldDefs[i].Required {

			// check if the field is present
			val, ok := userData[fieldDefs[i].Name]

			if !ok || val == nil {
				return nil, fmt.Errorf("missing required field: %s", userData[fieldDefs[i].Name])
			}

			fieldType := fieldDefs[i].Type

			fmt.Println(reflect.TypeOf(val).Name(), reflect.TypeOf(val).String())
			if fieldType != reflect.TypeOf(val).Name() {
				return nil, fmt.Errorf("Type mismatch on field: %s", userData[fieldDefs[i].Name])
			}

			resData["fields"] = val
		}
	}

	return resData, nil
}

// func ModelValidator(fieldDefs map[string]map[string]interface{}, userData map[string]interface{}) error {
// 	// Iterate over the field definitions
// 	for fieldName, fieldDef := range fieldDefs {
// 		// Check if the field is required
// 		required, ok := fieldDef["required"].(bool)
// 		if ok && required {
// 			// Field is required, check if it exists in the user data
// 			val, ok := userData[fieldName]
// 			if !ok || val == nilPrinPrintln()
// 				return fmt.Errorf("missing required field: %s", fieldName)
// 			}

// 			// Check if the field type matches the expected type
// 			fieldType, ok := fieldDef["type"].(string)

// 			fmt.Println(fieldType, reflect.TypeOf(val).Name(), val)
// 			if !ok || reflect.TypeOf(val).Name() != fieldType {
// 				return fmt.Errorf("field %s has incorrect type", fieldName)
// 			}
// 		}
// 	}

// 	return nil
// }
