package handler

import (
	"api/domain/model"
	"api/usecase/repository"
	"api/utils"
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// userHandler DBのインターフェース
type userHandler struct {
	userRepo repository.IUserRepository
}

// UserHandler 動作のインターフェース
type UserHandler interface {
	// 新規作成
	CreateUser(ctx *gin.Context)
	// 更新
	UpdateUser(ctx *gin.Context)
	// 削除
	DeleteUser(ctx *gin.Context)
	// 取得系
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetMe(ctx *gin.Context)
}

// NewUserHandler 生成
func NewUserHandler(userRepo repository.IUserRepository) UserHandler {
	fmt.Println("HANDLER : NewUserHandler OK")
	return &userHandler{userRepo: userRepo}
}

// 新規作成
func (userHandle *userHandler) CreateUser(ctx *gin.Context) {
	// モデル
	user := &model.User{}

	// Bind
	if err := ctx.Bind(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// パスワードのハッシュ化
	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	// ハッシュ化後格納
	user.Password = hash
	//fmt.Println(user.Password)

	// Store ERROR
	if err := userHandle.userRepo.Store(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// Success
	ctx.JSON(http.StatusOK, user)
}

// 更新
// TODO: UpdateUser実装
func (userHandle *userHandler) UpdateUser(ctx *gin.Context) {
	paramID := ctx.Param("id")
	//id, err := strconv.Atoi(paramId)

	// ID ERROR
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	//	return
	//}
	user := &model.User{UserId: paramID}

	// Bind ERROR
	if err := ctx.Bind(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// Update ERROR
	if err := userHandle.userRepo.Update(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// Success
	ctx.JSON(http.StatusOK, user)
}

// 削除
// TODO: Delete実装
func (userHandle *userHandler) DeleteUser(ctx *gin.Context) {
	paramID := ctx.Param("id")
	//id, err := strconv.Atoi(paramId)

	// ID ERROR
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	//	return
	//}

	// DELETE ERROR
	if err := userHandle.userRepo.Delete(&model.User{UserId: paramID}); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// Success
	ctx.JSON(http.StatusNoContent, gin.H{"message": "ok"})
}

// 取得
// TODO: GetUser実装
func (userHandle *userHandler) GetUser(ctx *gin.Context) {
	paramID := ctx.Param("id")
	//id, err := strconv.Atoi(paramId)

	// ID ERROR
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	//	return
	//}

	user, err := userHandle.userRepo.FindByID(paramID)

	// Find ERROR
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// 全取得
func (userHandle *userHandler) GetUsers(ctx *gin.Context) {
	user, err := userHandle.userRepo.FindAll()

	// Find ERROR
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// Success
	ctx.JSON(http.StatusOK, user)
}

func (userHandle *userHandler) GetMe(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	ctx.JSON(200, gin.H{
		"userID": claims["id"],
	})
}
