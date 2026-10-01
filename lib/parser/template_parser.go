package parser

import (
	"errors"
	"os"
	"regexp"
	"text/template"
)

func ParseFieldsFromFile(file string) (map[string]struct{}, error){
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	text := string(f)
	regex := regexp.MustCompile(`{{\.([a-zA-Z0-9_]+)}}`)
	result := regex.FindAllStringSubmatch(text, -1)
	if len(result) < 1 {
		return nil, errors.New("No Field found")
	}
	m := map[string]struct{}{}
	for _, v := range result {
		m[v[1]] = struct{}{}
	}
	return m, nil
}

func ParseTemplate(srcFile string, destFile string, dataMap map[string]string) (map[string]struct{}, error) {
	if destFile == "" {
		return nil, errors.New("No output file specified.")
	}
	if srcFile == "" {
		return nil, errors.New("No input file specified!")
	}

	intersectedMap := map[string]string{}
	ignoredMap := map[string]struct{}{}
	fieldsMap, err := ParseFieldsFromFile(srcFile)
	if err != nil {
		return nil, err
	}
	// to set a intersected map of fields and provided substitute
	for k, v := range dataMap {
		_, ok := fieldsMap[k]
		if ok {
			intersectedMap[k] = v
		} else {
			return nil, errors.New("Your provided field `" + k + "` is not defined.")
		}
	}
	// to generate a map of ignored field which will cause <no value> tag in boilerplate
	for k, _ := range fieldsMap {
		_, ok := dataMap[k]
		if !ok {
			ignoredMap[k] = struct{}{}
		}
	}
	
	t, err := template.ParseFiles(srcFile)
	if err != nil {
		return nil, err
	}
	createFile, err := os.Create(destFile)
	if err != nil {
		return nil, err
	}
	err = t.Execute(createFile, intersectedMap)
	if err != nil {
		return nil, err
	}
	return ignoredMap, nil
}
