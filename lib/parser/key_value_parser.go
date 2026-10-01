package parser

import (
	"regexp"
)

func ParseKeyValueByEqual(kv string) (string, string) {

	regex := regexp.MustCompile(`\=`)
	results := regex.Split(kv, 2)
	if !(ValidateKey(results[0])) {
		panic("Key should not conatain any special character")
	}
	return results[0], results[1]
}

func ParseKeyValueToMap(keyValueSlice []string) map[string]string {
	m := map[string]string{}
	for _, kv := range keyValueSlice {
		k, v := ParseKeyValueByEqual(kv)
		m[k] = v
	}
	return m
}

func ValidateKey(str string) bool {
	ok, err := regexp.MatchString(`^[a-zA-Z0-9_]+$`, str)
	if err != nil {
		panic(err)
	}
	return ok
}
