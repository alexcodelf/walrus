package workflow

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// parseParams parses the params into the spec.
// The params are the key-value pairs that are used to replace the keywords in the spec.
// The keywords are the strings that are wrapped by dollar sign and curly braces, like "${key}".
// The keyword is "${key}" will be replaced by the value of the key in the params.
func parseParams(spec map[string]any, params map[string]string) (map[string]any, error) {
	if err := CheckParams(params); err != nil {
		return nil, err
	}

	bs, err := json.Marshal(spec)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bs))

	for k, v := range params {
		paramReg := regexp.MustCompile(fmt.Sprintf(`\$\{\s*%s\s*\}`, k))
		bs = paramReg.ReplaceAll(bs, []byte(v))
	}

	var replacedSpec map[string]any

	err = json.Unmarshal(bs, &replacedSpec)
	if err != nil {
		return nil, err
	}

	return replacedSpec, nil
}

// CheckParams checks if the params contain keywords.
func CheckParams(params map[string]string) error {
	keywordsReg := regexp.MustCompile(`{{.*}}`)

	for k, v := range params {
		if keywordsReg.MatchString(v) {
			return fmt.Errorf("params contain keywords: %s=%s", k, v)
		}
	}

	return nil
}
