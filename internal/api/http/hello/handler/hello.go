package handler

import (
	"go-gin-api/internal/api/context"
)

// Hello GO.
func Hello(ctx *context.Context) {
	ctx.Success("Hello GO.")
}



