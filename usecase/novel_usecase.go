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

// GetAllNovelById implements domain.NovelUseCase
func (n *novelUsecase) GetNovelById(id int) (model.Novel, error) {
	res, err := n.novelRepo.GetNovelById(id)
	if err != nil {
		return model.Novel{}, errors.New("internal server error, get table novel")
	}
	return res, nil
}

// CreateNovel implements domain.NovelRepo
func (n *novelUsecase) CreateNovel(createNovel model.Novel) error {
	if err := n.novelRepo.CreateNovel(createNovel); err != nil {
		return errors.New("internal server error, creating organization")
	}

	// delete redis
	if errRedis := n.novelRepo.DeleteNovelRedis("novel"); errRedis != nil {
		return errors.New("internal server error, cannot delete redis key:novel")
	}

	return nil
}

// UpdateNovel implements domain.NovelUseCase
func (n *novelUsecase) UpdateNovelById(id string, novel model.Novel) error {
	err := n.novelRepo.UpdateNovelById(id, novel)
	if err != nil {
		return errors.New("internal server error, cannot update table novel")
	}

	// delete redis by id
	if errRedis := n.novelRepo.DeleteNovelRedis("novel" + id); errRedis != nil {
		return errors.New("internal server error, cannot delete redis key:novel+id")
	}

	// delete all redis data
	if errRedisAll := n.novelRepo.DeleteNovelRedis("novel"); errRedisAll != nil {
		return errors.New("internal server error, cannot delete redis key:novel")
	}

	return nil
}
