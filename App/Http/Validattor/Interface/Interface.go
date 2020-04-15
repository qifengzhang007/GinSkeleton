package Interface

import "github.com/gin-gonic/gin"

type ValidatorInterface interface {
	CheckParams(context *gin.Context)
}
