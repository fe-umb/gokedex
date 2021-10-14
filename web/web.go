package web

import (
	"fmt"
	"net/http"

	"github.com/fe-umb/gokedex/app"
	"github.com/gin-gonic/gin"
)

// Host function serves API endpoints and pages
func Host() {
	r := gin.Default()
	r.LoadHTMLGlob("./pages/index.html")
	gin.SetMode(gin.DebugMode)

	r.GET("/", home)
	r.GET("/all", getAllPokemons)
	err := r.Run(":8080")

	if err != nil {
		fmt.Println(err.Error())
	}

}

func home(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Gokédex",
		},
	)
}

func getAllPokemons(ctx *gin.Context) {
	returnJSON, err := app.ReadJSON("./assets/pokedex.json")

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	ctx.JSON(http.StatusOK, returnJSON)
}
