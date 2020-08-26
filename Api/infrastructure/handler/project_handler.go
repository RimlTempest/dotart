package handler

import (
	"api/domain/model"
	"api/usecase/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectHandler struct {
	projectRepo repository.IProjectRepository
}

// ProjectHandler 動きのインターフェース
type ProjectHandler interface {
	// 新規作成
	CreateProject(ctx *gin.Context)
	// 更新
	UpdateProject(ctx *gin.Context)
	// 削除
	DeleteProject(ctx *gin.Context)
	// 取得系
	GetProject(ctx *gin.Context)
	GetProjects(ctx *gin.Context)
}

// NewProjectHandler 生成
func NewProjectHandler(projectRepo repository.IProjectRepository) ProjectHandler {
	fmt.Println("HANDLER : NewProjectHandler OK")
	return &projectHandler{projectRepo: projectRepo}
}

// 作成
func (projectHandle *projectHandler) CreateProject(ctx *gin.Context) {
	project := &model.Project{}

	// Bind ERROR
	if err := ctx.Bind(project); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// 4xx
	if err := projectHandle.projectRepo.Store(project); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// Success
	ctx.JSON(http.StatusOK, project)
}

// 更新
func (projectHandle *projectHandler) UpdateProject(ctx *gin.Context) {
	paramID := ctx.Param("id")
	//id, err := paramId //strconv.Atoi(paramId)

	// 4xx
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	//	return
	//}

	project := &model.Project{ProjectName: paramID}

	// Bind ERROR
	if err := ctx.Bind(project); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	// Update ERROR
	if err := projectHandle.projectRepo.Update(project); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, project)
}

// 削除
func (projectHandle *projectHandler) DeleteProject(ctx *gin.Context) {
	paramID := ctx.Param("id")
	//id, err := paramId //strconv.Atoi(paramId)

	// id ERROR
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	//	return
	//}

	// Delete ERROR
	if err := projectHandle.projectRepo.Delete(&model.Project{ProjectName: paramID}); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "ok"})
}

// 取得
func (projectHandle *projectHandler) GetProject(ctx *gin.Context) {
	paramID := ctx.Param("id")
	//id, err := paramId //strconv.Atoi(paramId)

	// id ERROR
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	//	return
	//}

	project, err := projectHandle.projectRepo.FindByID(paramID)

	// 400
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, project)
}

// 全取得
func (projectHandle *projectHandler) GetProjects(ctx *gin.Context) {

	project, err := projectHandle.projectRepo.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, project)
}
