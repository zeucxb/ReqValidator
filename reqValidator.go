package reqValidator

import (
	"fmt"
	"reflect"
)

type structure interface {
	Map() map[string]interface{}
}

// Validate the types and return a bool
func Validate(structure structure, itemsMap map[string]interface{}) bool {
	stMap := structure.Map()

	if len(stMap) != len(itemsMap) {
		fmt.Println(len(stMap), len(itemsMap))
		return false
	}

	for name, value := range stMap {
		if _, ok := itemsMap[name]; !ok {
			fmt.Println(name)
			return false
		}

		if reflect.TypeOf(value).Kind() != reflect.TypeOf(itemsMap[name]).Kind() {
			fmt.Println(name)
			return false
		}
	}

	return true
}

// ValidateMap Verify if items exists in map
func ValidateMap(items map[interface{}]interface{}, names ...interface{}) bool {

	for _, name := range names {
		_, ok := items[name]
		if !ok {
			return false
		}
	}

	return true
}

// ValidateForm Verify if items exists in the Form
func ValidateForm(items ...[]string) bool {

	for _, item := range items {
		if len(item) <= 0 {
			return false
		}
	}

	return true
}
