package helper

import (
	"os"
)

// ReadFileData read file for test
func ReadFileData(path string, buf []byte) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.Read(buf)
	if err != nil {
		return err
	}

	return nil
}
