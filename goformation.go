package goformation

import (
	"encoding/json"
	"fmt"
	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/bridgecrewio/goformation/v4/intrinsics"
	"github.com/sanathkr/yaml"
	"io/ioutil"
	"strings"
)

//go:generate generate/generate.sh

// Open and parse a AWS CloudFormation template from file.
// Works with either JSON or YAML formatted templates.
func Open(filename string) (*cloudformation.Template, error) {
	return OpenWithOptions(filename, nil)
}

// OpenWithOptions opens and parse a AWS CloudFormation template from file.
// Works with either JSON or YAML formatted templates.
// Parsing can be tweaked via the specified options.
func OpenWithOptions(filename string, options *intrinsics.ProcessorOptions) (*cloudformation.Template, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(filename, ".json") {
		// This is definitely JSON
		return ParseJSONWithOptions(data, options)
	}

	return ParseYAMLWithOptions(data, options)
}

// ParseYAML an AWS CloudFormation template (expects a []byte of valid YAML)
func ParseYAML(data []byte) (*cloudformation.Template, error) {
	return ParseYAMLWithOptions(data, nil)
}

func StringifyInnerValues(iMapData interface{}, keyPath []string) interface{} {
	if len(keyPath) == 0 {
		return fmt.Sprintf("%v", iMapData)
	}

	mapData, ok := iMapData.(map[string]interface{})
	if !ok {
		return iMapData
	}

	currProperty := keyPath[0]
	if currProperty != "*" {
		innerMap, ok := mapData[currProperty]
		if !ok {
			return mapData
		}
		mapData[currProperty] = StringifyInnerValues(innerMap, keyPath[1:])
	} else {
		for key, val := range mapData {
			mapData[key] = StringifyInnerValues(val, keyPath[1:])
		}
	}

	return mapData
}

func StringifyValues(data []byte, keyPaths []string) ([]byte, error) {
	var structData interface{}
	err := yaml.Unmarshal(data, &structData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %v", err)
	}

	mapData, ok := structData.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to convert sturct to map for stringifying YAML elements")
	}

	for _, keyPath := range keyPaths {
		iMapData := StringifyInnerValues(mapData, strings.Split(keyPath, "/"))
		mapData, ok = iMapData.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to stringify paths in YAML")
		}
	}

	updatedData, err := yaml.Marshal(&mapData)

	return updatedData, err
}

// ParseYAMLWithOptions an AWS CloudFormation template (expects a []byte of valid YAML)
// Parsing can be tweaked via the specified options.
func ParseYAMLWithOptions(data []byte, options *intrinsics.ProcessorOptions) (*cloudformation.Template, error) {
	var err error
	if options != nil && len(options.StringifyPaths) > 0 {
		data, err = StringifyValues(data, options.StringifyPaths)
		if err != nil {
			return nil, err
		}
	}

	// Process all AWS CloudFormation intrinsic functions (e.g. Fn::Join)
	intrinsified, err := intrinsics.ProcessYAML(data, options)
	if err != nil {
		return nil, err
	}

	return unmarshal(intrinsified)

}

// ParseJSON an AWS CloudFormation template (expects a []byte of valid JSON)
func ParseJSON(data []byte) (*cloudformation.Template, error) {
	return ParseJSONWithOptions(data, nil)
}

// ParseJSONWithOptions an AWS CloudFormation template (expects a []byte of valid JSON)
// Parsing can be tweaked via the specified options.
func ParseJSONWithOptions(data []byte, options *intrinsics.ProcessorOptions) (*cloudformation.Template, error) {

	// Process all AWS CloudFormation intrinsic functions (e.g. Fn::Join)
	intrinsified, err := intrinsics.ProcessJSON(data, options)
	if err != nil {
		return nil, err
	}

	return unmarshal(intrinsified)

}

func unmarshal(data []byte) (*cloudformation.Template, error) {

	template := &cloudformation.Template{}
	if err := json.Unmarshal(data, template); err != nil {
		return nil, err
	}

	return template, nil

}
