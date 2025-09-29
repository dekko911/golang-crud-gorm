package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/what-crud/initializers"
	"github.com/what-crud/models"
	"github.com/what-crud/utils"
	"gorm.io/gorm"
)

func Login(ctx *gin.Context) {
	var req models.UserPayload

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   errors,
		})
		return
	}

	user, err := GetUserFromReqByEmail(req.Email, ctx)
	if err != nil {
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Invalid Email!",
			"error":   err.Error(),
		})
		return
	}

	if !utils.CompareHashPassword(user.Password, []byte(req.Password)) {
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Invalid Password!",
		})
		return
	}

	token, err := utils.CreateJWT(user.ID)
	if err != nil {
		ctx.JSON(utils.ISE, gin.H{
			"code":    utils.ISE,
			"message": "Invalid Token!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(utils.OK, gin.H{
		"code":    utils.OK,
		"message": "Login Successfully!",
		"token":   token,
	})
}

func Register(ctx *gin.Context) {
	var req models.UserPayload

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.JSON(utils.BR, gin.H{
			"code":    utils.BR,
			"message": "Bad Request",
			"error":   errors,
		})
		return
	}

	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		ctx.JSON(utils.ISE, gin.H{
			"code":    utils.ISE,
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

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
		"message": "User Registered!",
	})
}
