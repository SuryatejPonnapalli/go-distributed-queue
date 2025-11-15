package users

import (
	"fmt"
	"net/http"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common/utils"
	"github.com/gin-gonic/gin"
)


type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (ctr *Controller) Register(c *gin.Context){
	var req AuthRequest

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

func (ctr *Controller) Login(c *gin.Context){
	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		utils.RequestError(c, err)
		return
	}

	resp, err := ctr.service.Login(c.Request.Context(), req)

	if err != nil{
		if err.Error() == "email not found"{
			utils.Unauthorized(c, err.Error())
			return
		}
		if err.Error() == "wrong password"{
			utils.Unauthorized(c, err.Error())
			return
		}
		utils.InternalError(c)
		return
	}

	c.SetCookie("token", resp.Token, 60*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusAccepted, resp)
}