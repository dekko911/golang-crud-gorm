package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/what-crud/initializers"
	"github.com/what-crud/models"
	"github.com/what-crud/utils"
	"gorm.io/gorm"
)

func GetUsers(ctx *gin.Context) {
	res, err := gorm.G[*models.User](initializers.DB).Find(ctx)
	if err != nil {
		ctx.JSON(utils.ISE, gin.H{
			"code":    utils.ISE,
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(utils.OK, gin.H{
		"code":    utils.OK,
		"message": "OK",
		"data":    res,
	})
}

func GetUserByID(ctx *gin.Context) {
	paramId := ctx.Param("id")
	res, err := gorm.G[*models.User](initializers.DB).Where("id = ?", paramId).First(ctx)
	if err != nil {
		ctx.JSON(utils.NF, gin.H{
			"code":    utils.NF,
			"message": "User Not Found!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(utils.OK, gin.H{
		"code":    utils.OK,
		"message": "OK",
		"data":    res,
	})
}

func StoreUser(ctx *gin.Context) {
	var req models.UserPayload

	// get data from request body
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	// validate the request
	if err := utils.Validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   errors,
		})
		return
	}

	// hash the password
	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		ctx.JSON(utils.ISE, gin.H{
			"code":    utils.ISE,
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	// then create user
	if err := gorm.G[models.User](initializers.DB).Create(ctx, &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}); err != nil {
		ctx.JSON(utils.ISE, gin.H{
			"code":    utils.ISE,
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(utils.CRD, gin.H{
		"code":    utils.CRD,
		"message": "User Created!",
	})
}

func UpdateUser(ctx *gin.Context) {
	paramId := ctx.Param("id")

	var req models.UserUpdatePayload

	// get data from request body
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	// validate the request
	if err := utils.Validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   errors,
		})
		return
	}

	// hash the password
	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		ctx.JSON(utils.ISE, gin.H{
			"code":    utils.ISE,
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	// then update user
	if _, err := gorm.G[*models.User](initializers.DB).Where("id = ?", paramId).Updates(ctx, &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}); err != nil {
		ctx.JSON(utils.ISE, gin.H{
			"code":    utils.ISE,
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(utils.OK, gin.H{
		"code":    utils.OK,
		"message": "User Updated!",
	})
}

func DestroyUser(ctx *gin.Context) {
	paramId := ctx.Param("id")
	_, err := gorm.G[*models.User](initializers.DB).Where("id = ?", paramId).Delete(ctx)
	if err != nil {
		ctx.JSON(utils.NF, gin.H{
			"code":    utils.NF,
			"message": "User Not Found!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(utils.OK, gin.H{
		"code":    utils.OK,
		"message": "User Deleted!",
	})
}
