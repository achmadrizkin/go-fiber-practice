package domain

import "go-fiber-practice/model"

type NovelRepo interface {
	GetAllNovel() ([]model.Novel, error)
}

type NovelUseCase interface {
	GetAllNovel() ([]model.Novel, error)
}
