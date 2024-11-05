package middleware

import (
	"net/http"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config model.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx,config)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		token ,_ := utils.ExtractToken(ctx)
		ctx.Header("Authorization", token)
		ctx.Next()
	}
}
