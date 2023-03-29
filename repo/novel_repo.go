package repo

import (
	"fmt"
	"go-fiber-practice/domain"
	"go-fiber-practice/model"

	"gorm.io/gorm"
)

type novelRepo struct {
	db *gorm.DB
}

func NewNovelRepo(db *gorm.DB) domain.NovelRepo {
	return &novelRepo{
		db: db,
	}
}

// GetOrganizationAll implements domain.NovelRepo
func (n *novelRepo) GetAllNovel() ([]model.Novel, error) {
	var organizations []model.Novel
	err := n.db.Model(model.Novel{}).Select("id", "name", "description", "description", "author").Find(&organizations).Error
	fmt.Println(organizations)
	return organizations, err
}
