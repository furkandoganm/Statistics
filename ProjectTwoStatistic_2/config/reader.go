package config

import (
	"encoding/json"
	"fmt"
	"os"
	"projects/ProjectTwoStatistic_2/model"
)

func Reader(str string) (string, error) {
	var cString model.Connection

	file, err := os.ReadFile("./file/connectionString.json")
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(file, &cString)
	if err != nil {
		return "", err
	}

	for _, c := range cString.ConnectionString {
		if str == c.Name {
			return c.CString, nil
		}
	}

	er := fmt.Errorf("not found any string have this name")
	return "", er
}
