package app

import "os"

func ReadJSON(filename string) (*os.File, error) {
	jsonFile, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return jsonFile, nil
}
