package persistance

import (
	"api/domain/model"
	"api/usecase/repository"

	"fmt"

	"github.com/jinzhu/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) repository.IProjectRepository {
	fmt.Println("REPOSITORY : NewProjectRepository OK")
	return &projectRepository{db}
}

func (projectRepo *projectRepository) FindByID(id string) (*model.Project, error) {
	project := model.Project{ProjectName: id}
	err := projectRepo.db.First(&project).Error

	// First ERROR
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (projectRepo *projectRepository) Store(project *model.Project) error {
	return projectRepo.db.Save(project).Error
}

func (projectRepo *projectRepository) Update(project *model.Project) error {
	// 値を更新していなくても、SQLの更新を実行すると、すべてのフィールドが上書きされます。
	return projectRepo.db.Model(&model.Project{ProjectName: project.ProjectName}).Updates(project).Error
}

func (projectRepo *projectRepository) Delete(project *model.Project) error {
	return projectRepo.db.Delete(project).Error
}

func (projectRepo *projectRepository) FindAll() ([]*model.Project, error) {
	projects := []*model.Project{}

	err := projectRepo.db.Find(&projects).Error

	// Find ERROR
	if err != nil {
		return nil, err
	}

	return projects, nil
}
