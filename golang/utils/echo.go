package utils

import (
	"fmt"
	"reflect"

	"github.com/logrusorgru/aurora"
)

// PrintPtrStructObject : print pointer structure object
func PrintPtrStructObject(obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	numField := v.NumField()
	for i := 0; i < numField; i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.FieldByName(fieldName)
		fmt.Println(aurora.BrightCyan(fieldName), ":", aurora.BrightGreen(fieldValue))
	}
}
