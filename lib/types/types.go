package types

import "github.com/gin-gonic/gin"

type Controller func(e *gin.Engine)

type Service func()

type Model func()
