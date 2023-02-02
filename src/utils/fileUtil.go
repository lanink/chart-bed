package utils

import "os"

func CheckAndCreateDir(dir string) (bool, error) {
	_, err := os.Stat(dir)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err == nil {
			return true, nil
		}
	}

	return false, err
}
