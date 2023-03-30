package domain

import "go-fiber-practice/model"

type NovelRepo interface {
	CreateNovel(createNovel model.Novel) error
	GetAllNovel() ([]model.Novel, error)
	GetNovelById(id int) (model.Novel, error)
	DeleteNovelRedis(key string) error
}

type NovelUseCase interface {
	CreateNovel(createNovel model.Novel) error
	GetAllNovel() ([]model.Novel, error)
	GetNovelById(id int) (model.Novel, error)
}
