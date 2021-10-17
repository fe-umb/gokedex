package main

import (
	"github.com/fe-umb/gokedex/app"
)

func main() {
	ctx := app.NewAppCtx()
	ctx.Host()
}
