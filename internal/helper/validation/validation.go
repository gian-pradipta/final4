package validation

import (
	"encoding/json"
	"errors"
)

func ValidateJSONParams(jsonByte []byte, allowedParams []string) error {
	var err error

	var validationMap map[string]interface{} = make(map[string]interface{})
	err = json.Unmarshal(jsonByte, &validationMap)
	if err != nil {
		return err
	}

	var allowedParamsMap map[string]bool = make(map[string]bool)
	for _, param := range allowedParams {
		allowedParamsMap[param] = true
	}

	for key := range validationMap {
		if allowedParamsMap[key] == false {
			err = errors.New("Invalid JSON Parameters")
			return err
		}
	}

	if len(allowedParams) != len(validationMap) {
		err = errors.New("Invalid JSON Parameters")
		return err
	}

	return err
}
