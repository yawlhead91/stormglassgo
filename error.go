package stormglass

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Error response representation.
type Error struct {
	StatusCode int                    `json:"-"`
	Errors     map[string]interface{} `json:"errors"`
}

// Error returns formatted errors as a string.
func (e Error) Error() string {
	messages := make([]string, 0, len(e.Errors))
	for k, v := range e.Errors {
		messages = append(messages, fmt.Sprintf("%s:%s", k, v))
	}

	return fmt.Sprintf("%d: %s", e.StatusCode, strings.Join(messages, ","))
}

// NewError creates a new Error from an API response.
func NewError(resp *http.Response) error {
	apiErr := Error{
		StatusCode: resp.StatusCode,
		Errors:     map[string]interface{}{},
	}

	data, err := io.ReadAll(resp.Body)
	if err == nil && data != nil {
		if err = json.Unmarshal(data, &apiErr); err != nil {
			apiErr.Errors["unknown"] = []string{"unknown_error_format"}
		}
	}

	return &apiErr
}
