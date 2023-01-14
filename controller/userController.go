package controller

import (
	"Blog/ent"
	"Blog/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"Blog/service"

	"github.com/gin-gonic/gin"
)

func UserGetAllController(c *gin.Context) {

	users, err := service.NewUserOps(c.Request.Context()).UsersGetAll()
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, users)
}

func UserGetByIDController(c *gin.Context) {

	vars := c.Param("id")

	id, err := strconv.Atoi(vars)

	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}

	user, err := service.NewUserOps(c.Request.Context()).UserGetByID(id)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, user)
}

func UserCreateController(c *gin.Context) {

	var newUser ent.User
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}
	user, err := service.NewUserOps(c.Request.Context()).UserCreate(newUser)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, user)
}

func UserUpdateController(c *gin.Context) {

	var newUserData ent.User
	err := json.NewDecoder(c.Request.Body).Decode(&newUserData)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}

	vars := c.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}
	newUserData.ID = id

	updatedUser, err := service.NewUserOps(c.Request.Context()).UserUpdate(newUserData)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, updatedUser)
}

func UserDeleteController(c *gin.Context) {

	vars := c.Param("id")

	id, err := strconv.Atoi(vars)

	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewUserOps(c.Request.Context()).UserDelete(id)

	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, deletedID)
}
