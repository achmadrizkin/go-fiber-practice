package usecase

import (
	"errors"
	"go-fiber-practice/domain"
	"go-fiber-practice/model"
)

type novelUsecase struct {
	novelRepo domain.NovelRepo
}

func NewNovelUsecase(novelRepo domain.NovelRepo) domain.NovelUseCase {
	return &novelUsecase{
		novelRepo: novelRepo,
	}
}

// GetAllNovel implements domain.NovelUseCase
func (n *novelUsecase) GetAllNovel() ([]model.Novel, error) {
	res, err := n.novelRepo.GetAllNovel()
	if err != nil {
		return []model.Novel{}, errors.New("internal server error, get table novel")
	}
	return res, nil
}
