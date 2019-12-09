package jsonfunc

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// ReadFromFile reads json file with given fileName, then save into object
func ReadFromFile(fileName string, v interface{}) error {
	var err error
	ss := strings.Split(fileName, ".")
	if ss[len(ss)-1] != "json" {
		err = errors.New("File Extension Mismatched")
		return err
	}

	file, osErr := os.Open(fileName)
	if osErr != nil {
		return osErr
	}
	defer file.Close()
	osErr = json.NewDecoder(file).Decode(&v)
	if osErr != nil {
		return osErr
	}
	return nil
}
