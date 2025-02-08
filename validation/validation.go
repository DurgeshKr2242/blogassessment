package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
)

var (
	errIsRequired = errors.New("is required")
	errMin5       = errors.New("should at least have 5 characters")
	errMax60      = errors.New("should not exceed 60 characters")
	errMin10      = errors.New("should at least have 10 characters")
	errMax300     = errors.New("should not exceed 300 characters")
	errUUID       = errors.New("must be a valid UUID")

	customErrors = map[string]error{
		"ID.required":          errIsRequired,
		"ID.uuid":              errUUID,
		"Title.required":       errIsRequired,
		"Title.min":            errMin5,
		"Title.max":            errMax60,
		"Description.required": errIsRequired,
		"Description.min":      errMin10,
		"Description.max":      errMax300,
		"Body.required":        errIsRequired,
		"Body.min":             errMin10,
	}
)

// CustomValidationError converts validator errors to a slice of custom error messages.
func CustomValidationError(err error) []map[string]string {
	errs := make([]map[string]string, 0)
	switch errTypes := err.(type) {

	case validator.ValidationErrors:
		for _, e := range errTypes {
			errorMap := make(map[string]string)
			// Create the key as Field.Tag (e.g., "Title.required", "ID.uuid")
			key := e.Field() + "." + e.Tag()
			if v, ok := customErrors[key]; ok {
				errorMap[e.Field()] = v.Error()
			} else {
				errorMap[e.Field()] = fmt.Sprintf("custom message is not available: %v", err)
			}
			errs = append(errs, errorMap)
		}
		return errs

	case *json.UnmarshalTypeError:
		errs = append(errs, map[string]string{
			errTypes.Field: fmt.Sprintf("%v cannot be a %v", errTypes.Field, errTypes.Value),
		})
		return errs
	}

	if errors.Is(err, io.EOF) {
		errs = append(errs, map[string]string{
			"body": "request body cannot be empty",
		})
	} else {
		errs = append(errs, map[string]string{
			"unknown": fmt.Sprintf("unsupported custom error for: %v", err),
		})
	}
	return errs
}
