package utils

import (
	"fmt"
	"reflect"

	"github.com/logrusorgru/aurora"
)

// PrintPtrStructObject : print pointer structure object
func PrintPtrStructObject(obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.FieldByName(fieldName)
		fmt.Println(aurora.BrightCyan(t.Field(i).Tag.Get("json")), ":", aurora.BrightGreen(fieldValue))
	}
}

// PrintObject : if found, print object. else print no result
func PrintObject(obj interface{}, isFound bool, key, value string) {
	if !isFound {
		NoResult(obj, key, value)
	} else {
		PrintPtrStructObject(obj)
	}
}

// NoResult : print no results found
func NoResult(obj interface{}, key, value string) {
	dataType := reflect.TypeOf(obj).Elem().Name()
	fmt.Println(aurora.Red("Searching " + dataType + "s for " + key + " with a value of " + value))
	fmt.Println(aurora.Red("No results found"))
}

// PrintPtrStructObjectTags : print pointer structure object tags
func PrintPtrStructObjectTags(obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(aurora.BrightGreen(t.Field(i).Tag.Get("json")))
	}
}

// PrintSeparation : print separation
func PrintSeparation() {
	fmt.Println("------------------------------------------------")
}
