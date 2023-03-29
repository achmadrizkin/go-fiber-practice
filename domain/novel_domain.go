package domain

import "go-fiber-practice/model"

type NovelRepo interface {
	GetAllNovel() ([]model.Novel, error)
	GetNovelById(id int) (model.Novel, error)
}

type NovelUseCase interface {
	GetAllNovel() ([]model.Novel, error)
	GetNovelById(id int) (model.Novel, error)
}
