package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Host function serves API endpoints and pages
func (app *App) Host() {
	r := gin.Default()
	r.LoadHTMLGlob("./pages/html/*")

	// CSS/JS
	r.Static("/css", "./pages/css")
	r.Static("/js", "./pages/js")
	r.Static("/fvc", "./pages/fvc")

	// Pages
	r.GET("/", home)
	r.GET("/info/:id", pokemonInfo)

	// Endpoints
	r.GET("/all", getAllPokemons)
	r.GET("/pokemon/:id", getPokemon)

	// Pokemon sprites
	r.Static("/sprt/", "./pkmn/sprites")

	// Pokemon images
	r.Static("/imgs/", "./pkmn/images")

	gin.SetMode(app.Env)
	err := r.Run(":" + app.Port)
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
	returnJSON, err := ReadJSON("./pkmn/pokedex.json")

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	ctx.JSON(http.StatusOK, returnJSON)
}

func getPokemon(ctx *gin.Context) {
	idPokemon, err := strconv.Atoi(ctx.Param("id"))
	foundPokemon := false

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	returnJSON, err := ReadJSON("./pkmn/pokedex.json")
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	for _, ele := range returnJSON {
		if ele.ID == idPokemon {
			ctx.JSON(http.StatusOK, ele)
			foundPokemon = true
		}
	}

	if !foundPokemon {
		ctx.JSON(http.StatusNotFound, "Pokémon not found.")
	}

}
