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
	r.LoadHTMLGlob("./pages/html/*")

	// CSS/JS
	r.Static("/css", "./pages/css")
	r.Static("/js", "./pages/js")

	// Pages
	r.GET("/", home)

	// Endpoints
	r.GET("/all", getAllPokemons)

	// Pokemon sprites
	r.Static("/sprt/", "./assets/sprites")

	gin.SetMode(gin.DebugMode)
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
