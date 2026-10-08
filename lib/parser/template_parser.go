package parser

import (
	"errors"
	"io"
	"regexp"
	"text/template"
)

func ParseFieldsFromFile(text string) (map[string]struct{}, error){
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

func ParseTemplate(srcFile io.Reader, destFile io.Writer, dataMap map[string]string, forced bool) (map[string]struct{}, error) {
	intersectedMap := map[string]string{}
	ignoredMap := map[string]struct{}{}
	srcFileContens := []byte{}
	buf := make([]byte, 2048)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		srcFileContens = append(srcFileContens, buf[:n]...)
	}
	fieldsMap, err := ParseFieldsFromFile(string(srcFileContens))
	if err != nil {
		return nil, err
	}
	// to generate a map of ignored field which will cause <no value> tag in boilerplate
	for k, _ := range fieldsMap {
		_, ok := dataMap[k]
		if !ok {
			ignoredMap[k] = struct{}{}
		}
	}

	// to set a intersected map of fields and provided substitute
	for k, v := range dataMap {
		_, ok := fieldsMap[k]
		if ok {
			intersectedMap[k] = v
		} else {
			return ignoredMap, errors.New("Your provided field `" + k + "` is not defined.")
		}
	}
	if !forced && (len(ignoredMap) > 0) {
		return ignoredMap, errors.New("You have ignored some fields in template, use -f or --force-ignore to force.")
	}
	
	t := template.New("t1")
	t, err = t.Parse(string(srcFileContens))
	if err != nil {
		return nil, err
	}
	err = t.Execute(destFile, intersectedMap)
	if err != nil {
		return nil, err
	}
	return ignoredMap, nil
}
