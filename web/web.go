package web

import (
	"fmt"
	"net/http"
	"strconv"

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
	r.GET("/info/:id", pokemonInfo)

	// Endpoints
	r.GET("/all", getAllPokemons)
	r.GET("/pokemon/:id", getPokemon)

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
		"home.html",
		gin.H{
			"title": "Gokédex • Home",
		},
	)
}

func pokemonInfo(ctx *gin.Context) {
	title := fmt.Sprintf(`Gokédex • Pokémon #%s`, ctx.Param("id"))
	ctx.HTML(
		http.StatusOK,
		"info.html",
		gin.H{
			"title": title,
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

func getPokemon(ctx *gin.Context) {
	idPokemon, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	returnJSON, err := app.ReadJSON("./assets/pokedex.json")
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	for _, ele := range returnJSON {
		if ele.ID == idPokemon {
			ctx.JSON(http.StatusOK, ele)
		}
	}

	ctx.JSON(http.StatusNotFound, "Pokémon not found.")

}
