package repository

import (
	"api/domain/model"
)

// 処理のインターフェイス
type IProjectRepository interface {
	FindByID(id string) (*model.Project, error)
	Store(project *model.Project) error
	Update(project *model.Project) error
	Delete(project *model.Project) error
	FindAll() ([]*model.Project, error)
}
