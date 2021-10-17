package app

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fe-umb/gokedex/models"
)

type App struct {
	Port string
	Env  string
}

func NewAppCtx() App {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("GIN_ENV")
	if env != "release" {
		env = "debug"
	}

	var ctx = App{
		Port: port,
		Env:  env,
	}

	return ctx
}

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
