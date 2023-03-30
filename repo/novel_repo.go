package repo

import (
	"context"
	"encoding/json"
	"errors"
	"go-fiber-practice/domain"
	"go-fiber-practice/model"
	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type novelRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewNovelRepo(db *gorm.DB, rdb *redis.Client) domain.NovelRepo {
	return &novelRepo{
		db:  db,
		rdb: rdb,
	}
}

// CreateNovel implements domain.NovelRepo
func (n *novelRepo) CreateNovel(createNovel model.Novel) error {
	err := n.db.Create(&createNovel).Error
	return err
}

// GetAllNovel implements domain.NovelRepo
func (n *novelRepo) GetAllNovel() ([]model.Novel, error) {
	var novels []model.Novel
	var ctx = context.Background()

	// Check if data is available in Redis.
	result, err := n.rdb.Get(ctx, "novel").Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	// If data is available in Redis, decode it from JSON and return.
	if len(result) > 0 {
		err = json.Unmarshal([]byte(result), &novels)
		return novels, err
	}

	// If data is not available in Redis, retrieve it from database.
	err = n.db.Model(model.Novel{}).Select("id", "name", "description", "author").Find(&novels).Error
	if err != nil {
		return nil, err
	}

	// Encode the novels slice into JSON before saving to Redis.
	jsonBytes, err := json.Marshal(novels)
	if err != nil {
		return nil, err
	}
	jsonString := string(jsonBytes)

	// Set the JSON-encoded value in Redis.
	err = n.rdb.Set(ctx, "novel", jsonString, 0).Err()
	if err != nil {
		return nil, err
	}

	return novels, nil
}

// GetAllNovelById implements domain.NovelRepo
func (n *novelRepo) GetNovelById(id int) (model.Novel, error) {
	var novels model.Novel
	var ctx = context.Background()

	// Check if data is available in Redis.
	result, err := n.rdb.Get(ctx, "novel"+strconv.Itoa(id)).Result()
	if err != nil && err != redis.Nil {
		return novels, err
	}

	// If data is available in Redis, decode it from JSON and return.
	if len(result) > 0 {
		err = json.Unmarshal([]byte(result), &novels)
		return novels, err
	}

	// If data is not available in Redis, retrieve it from database.
	err = n.db.Model(model.Novel{}).Select("id", "name", "description", "author").Where("id=?", id).Find(&novels).Error
	if err != nil {
		return novels, err
	}

	// Encode the novels slice into JSON before saving to Redis.
	jsonBytes, err := json.Marshal(novels)
	if err != nil {
		return novels, err
	}
	jsonString := string(jsonBytes)

	// Set the JSON-encoded value in Redis.
	err = n.rdb.Set(ctx, "novel"+strconv.Itoa(id), jsonString, 0).Err()
	if err != nil {
		return novels, err
	}

	return novels, nil
}

// DeleteNovelRedis implements domain.NovelRepo
func (n *novelRepo) DeleteNovelRedis(key string) error {
	ctx := context.Background()
	if err := n.rdb.Del(ctx, key).Err(); err != nil {
		return errors.New("cannot delete key, maybe null value")
	}

	return nil
}

// UpdateNovel implements domain.NovelRepo
func (n *novelRepo) UpdateNovelById(id string, novel model.Novel) error {
	err := n.db.Model(&model.Novel{}).Where("id = ?", id).Updates(model.Novel{Name: novel.Name, Description: novel.Description}).Error
	if err != nil {
		return errors.New("cannot update novel")
	}
	return err
}
