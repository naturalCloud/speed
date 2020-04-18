package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"speed/app/http/model"
	"speed/app/lib/log"
)

type HelloController struct {
	Controller
}

var HelloC = &HelloController{}

func (h *HelloController) Index(ctx *gin.Context) {

	fmt.Println(model.Users{}.GetMore())
	log.WithCtx(ctx).Info("hello word")
	h.success(ctx, map[string]interface{}{"hello": "hello word"})
}
