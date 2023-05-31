package user

import (
	"strconv"

	. "sx-api/handler"
	"sx-api/model"
	"sx-api/pkg/errno"
	"sx-api/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	log.Infof("Update function called.X-Request-Id: %s", util.GetReqID(c))
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
