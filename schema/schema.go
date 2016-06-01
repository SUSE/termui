package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/cloudfoundry/cli/cf/terminal"
)

type SchemaParser struct {
	ui terminal.UI
}

func NewSchemaParser(userInterface terminal.UI) *SchemaParser {
	return &SchemaParser{
		ui: userInterface,
	}
}

func (p *SchemaParser) ParseSchema(schema string) (string, error) {
	var jsontype map[string]interface{}
	json.Unmarshal([]byte(schema), &jsontype)
	res, err := p.parseObject("", jsontype)
	if err != nil {
		return "", err
	}
	jsonResult, err := json.Marshal(res)
	if err != nil {
		return "", err
	}
	return string(jsonResult), nil
}

func (p *SchemaParser) parseObject(key string, input map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{} = make(map[string]interface{})

	if prop, ok := input["properties"]; ok {
		for k, v := range prop.(map[string]interface{}) {
			t := v.(map[string]interface{})["type"]
			required := ""
			if isRequired(k, input) {
				required = " [required]"
			}
			switch t {
			case "string":
				if k == "password" {
					result[k] = p.ui.AskForPassword(fmt.Sprintf("Insert string value for %s/%s%s", key, k, required))
				} else {
					result[k] = p.ui.Ask(fmt.Sprintf("Insert string value for %s/%s%s", key, k, required))
				}
			case "integer":
				{
					i, err := strconv.Atoi(p.ui.Ask(fmt.Sprintf("Insert integer value for %s/%s%s", key, k, required)))
					if err != nil {
						return nil, err
					}
					result[k] = i
				}
			case "boolean":
				{
					i, err := strconv.ParseBool(p.ui.Ask(fmt.Sprintf("Insert boolean value for %s/%s%s", key, k, required)))
					if err != nil {
						return nil, err
					}
					result[k] = i
				}
			case "number":
				i, err := strconv.ParseFloat(p.ui.Ask(fmt.Sprintf("Insert numeric value for %s/%s%s", key, k, required)), 64)
				if err != nil {
					return nil, err
				}
				result[k] = i
			case "array":
				return nil, errors.New("Array parsing not implemented")
			case "object":
				{
					res, err := p.parseObject(fmt.Sprintf("%s/%s", key, k), v.(map[string]interface{}))
					if err != nil {
						return nil, err
					}
					result[k] = res
				}
			default:
				result[k] = nil
			}
		}
	}

	return result, nil
}

func isRequired(key string, input map[string]interface{}) bool {
	required := input["required"].([]interface{})
	for _, v := range required {
		if v == key {
			return true
		}
	}
	return false
}
