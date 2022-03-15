package user

import (
	"github.com/gin-gonic/gin"
	"test/aulaInterface/businessRules/user"
	"test/aulaInterface/gin/endpoint/restful"
	"test/aulaInterface/interfaces"
	"test/aulaInterface/view/viewUser"
)

type DataSource struct {
	restful.Restful
	DataSource interfaces.InterfaceUser `json:"-"`
}

func (e *DataSource) UserByEmail(c *gin.Context) {
	var err error
	var userData viewUser.User
	var length int

	var email = c.Param("mail")

	menuBusinessRules := user.BusinessRules{}
	length, userData, err = menuBusinessRules.GetByEmail(email)

	e.Meta.Error = []string{}
	if err != nil {
		e.Meta.Success = false
		e.Meta.Error = []string{err.Error()}
		e.Object = []int{}
		c.JSON(200, e)
		return
	}

	e.Meta.Total = length
	e.Meta.Actual = length
	e.Meta.Success = true
	e.Object = []viewUser.User{
		userData,
	}

	c.JSON(200, e)
}
