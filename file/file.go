package file

import (
	"fmt"
	"os"
	"time"
)

// CreateFileLog membuat file log
func CreateFileLog(logName string) (*os.File, error) {
	fileName := fmt.Sprintf("%s%d-%s-%d.log", logName, time.Now().Day(), time.Now().Month().String(), time.Now().Year())
	if !CheckFile(fileName) {
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// CheckFile return true is exist and false is not exist
func CheckFile(file string) bool {
	var ada bool
	_, err := os.Stat(file)
	if err == nil {
		ada = true
	}
	return ada
}
