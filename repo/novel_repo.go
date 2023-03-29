package repo

import (
	"context"
	"encoding/json"
	"go-fiber-practice/domain"
	"go-fiber-practice/model"

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

// GetAllNovel implements domain.NovelRepo
func (n *novelRepo) GetAllNovel() ([]model.Novel, error) {
	var novel []model.Novel
	var ctx = context.Background()

	// check if data is available in Redis or not
	// if available, retrieve data from Redis
	// if not available, retrieve data from database and save to Redis
	resultNovels, _ := n.rdb.Get(ctx, "novel").Result()
	if len(resultNovels) > 0 {
		// data is available in Redis, decode it from JSON and return
		err := json.Unmarshal([]byte(resultNovels), &novel)
		return novel, err
	} else {
		// data is not available in Redis, retrieve it from database and save to Redis
		errMySQL := n.db.Model(model.Novel{}).Select("id", "name", "description", "description", "author").Find(&novel)
		if errMySQL != nil {
			// encode the books slice into JSON before saving to Redis
			jsonBytes, err := json.Marshal(novel)
			jsonString := string(jsonBytes)

			// set the JSON-encoded value in Redis
			n.rdb.Set(ctx, "novel", jsonString, 0)

			return novel, err
		} else {
			return novel, nil
		}
	}
}
