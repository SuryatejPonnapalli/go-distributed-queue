package users

import (
	"fmt"
	"net/http"

	"github.com/SuryatejPonnapalli/go_project/internal/common/utils"
	"github.com/gin-gonic/gin"
)


type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (ctr *Controller) Register(c *gin.Context){
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		utils.RequestError(c, err)
		return
	}

	resp, err := ctr.service.Register(c.Request.Context(), req)

	if err != nil{
		fmt.Println("REGISTER ERROR:", err)
		utils.InternalError(c)
		return
	}

	c.JSON(http.StatusCreated, resp)
}