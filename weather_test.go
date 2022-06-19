package stormglass

import (
	"reflect"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestParamOptionsToList(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		var options ParamsOptions
		list := options.toList()

		assert.Len(t, list, 0)
	})

	t.Run("with all options", func(t *testing.T) {
		options := ParamsOptions{
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
		}

		list := options.toList()

		v := reflect.ValueOf(options)

		assert := assert.New(t)
		assert.Len(list, v.NumField())

		var expected []string
		// this assumes naming conventions for fields names
		// as camel case matches of expected values
		for i := 0; i < v.NumField(); i++ {
			expected = append(expected, lcFirstLetter(v.Type().Field(i).Name))
		}

		assert.Equal(expected, list)
	})
}

func TestSourceOptionsToList(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		var options SourcesOptions
		list := options.toList()

		assert.Len(t, list, 0)
	})

	t.Run("with all options", func(t *testing.T) {
		options := SourcesOptions{
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
		}

		list := options.toList()

		v := reflect.ValueOf(options)

		assert := assert.New(t)
		assert.Len(list, v.NumField())

		expected := []string{
			"icon",
			"noaa",
			"meteo",
			"meto",
			"fcoo",
			"fmi",
			"yr",
			"smhi",
			"sg",
		}

		assert.Equal(expected, list)
	})
}

func TestClientGetPoint(t *testing.T) {

}

// lower case first letter helper func
func lcFirstLetter(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
