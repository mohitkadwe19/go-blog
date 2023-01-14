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

func BlogGetByIDController(c *gin.Context) {

	vars := c.Param("id")

	id, err := strconv.Atoi(vars)

	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}

	blog, err := service.NewBlogOps(c.Request.Context()).BlogGetByID(id)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, blog)
}

func BlogCreateController(c *gin.Context) {

	var newTodo ent.Blog

	err := json.NewDecoder(c.Request.Body).Decode(&newTodo)

	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}

	createdBlog, err := service.NewBlogOps(c.Request.Context()).BlogCreate(newTodo)

	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, createdBlog)
}

func BlogUpdateController(c *gin.Context) {

	var newBlogData ent.Blog
	err := json.NewDecoder(c.Request.Body).Decode(&newBlogData)
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
	newBlogData.ID = id

	updatedBlog, err := service.NewBlogOps(c.Request.Context()).BlogUpdate(newBlogData)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, updatedBlog)
}

func BlogDeleteController(c *gin.Context) {

	vars := c.Param("id")

	id, err := strconv.Atoi(vars)

	if err != nil {
		utils.Return(c.Writer, false, http.StatusBadRequest, err, nil)
		return
	}

	deletedID, err := service.NewBlogOps(c.Request.Context()).BlogDelete(id)
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, deletedID)
}

func BlogGetAllController(c *gin.Context) {

	blogs, err := service.NewBlogOps(c.Request.Context()).BlogGetAll()
	if err != nil {
		utils.Return(c.Writer, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(c.Writer, true, http.StatusOK, nil, blogs)
}
