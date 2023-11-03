package workflow

import (
	"reflect"
	"testing"
)

func TestParseParams(t *testing.T) {
	cases := []struct {
		spec     map[string]any
		params   map[string]string
		expected map[string]any
	}{
		{
			spec: map[string]any{
				"description": "${bobName}",
			},
			params: map[string]string{
				"bobName": "bob",
			},
			expected: map[string]any{
				"description": "bob",
			},
		},
		{
			spec: map[string]any{
				"deepAttr": map[string]any{
					"deepKey": "${replace}",
				},
			},
			params: map[string]string{
				"replace": "newValue",
			},
			expected: map[string]any{
				"deepAttr": map[string]any{
					"deepKey": "newValue",
				},
			},
		},
	}

	for _, c := range cases {
		actual, err := parseParams(c.spec, c.params)
		if err != nil {
			t.Errorf("parse params error: %v", err)
		}

		if reflect.DeepEqual(actual, c.expected) == false {
			t.Errorf("parse params error: expected %v, got %v", c.expected, actual)
		}
	}
}
