package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func SeeAllCategories(ctx *gin.Context) {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	if limit < 1 {
		limit = 7
	}

	if page < 1 {
		page = 1
	}

	result := models.FindAllCategories(search, limit, page)

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "List All Categories",
		Results: result,
	})
}

func SeeOneCategoryById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataCategory := models.FindOneCategory(id)
	fmt.Println(dataCategory)

	if dataCategory.Id != 0 {

		ctx.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: "Category Found",
			Results: dataCategory,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "Category Not Found",
		})
	}
}

func CreateCategory(ctx *gin.Context) {
	search := ctx.Query("search")
	newCategory := models.Category{}
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	result := models.FindAllCategories(search, limit, offset)

	if err := ctx.ShouldBind(&newCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	ids := 0
	for _, v := range result {
		ids = v.Id
	}
	newCategory.Id = ids + 1

	err := models.CreateCategory(newCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Server{
			Success: false,
			Message: "Failed to create Category",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Category created successfully",
		Results: newCategory,
	})
}

func EditCategory(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	search := c.Query("search")
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	data := models.FindAllCategories(search, limit, offset)

	category := models.Category{}
	err := c.Bind(&category)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := models.Category{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id == 0 {
		c.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "Category with id " + param + " not found",
		})
		return
	}

	ids := 0
	for _, v := range data {
		ids = v.Id
	}
	category.Id = ids

	models.EditCategory(category.Name, param)

	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Category with id " + param + " Edit Success",
		Results: category,
	})
}

func DeleteCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataCategory := models.FindOneCategory(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid Category ID",
		})
		return
	}

	err = models.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Id Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Category deleted successfully",
		Results: dataCategory,
	})
}

func ListAllFilterCategory(c *gin.Context) {
	category := c.Query("category")

	products, err := models.GetAllCategoryWithFilter(category)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Id Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Category Has Been Filtered",
		Results: products,
	})
}
