package app

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fe-umb/gokedex/models"
)

func ReadJSON(filename string) (models.Pokemon, error) {

	var pkList models.Pokemon

	jsonFile, err := os.Open(filename)
	if err != nil {
		return pkList, err
	}
	defer jsonFile.Close()

	byteV, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return pkList, err
	}

	json.Unmarshal(byteV, &pkList)

	return pkList, nil
}
