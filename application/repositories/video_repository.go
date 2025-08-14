package repositories

import (
	"fmt"

	"github.com/andre2ar/video-encoder/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{
		Db: db,
	}
}

func (jb *VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.New().String()
	}

	err := jb.Db.Create(video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (jb *VideoRepositoryDb) Find(id string) (*domain.Video, error) {
	var video domain.Video

	jb.Db.Preload("Jobs").First(&video, "id = ?", id)
	if video.ID == "" {
		return nil, fmt.Errorf("video not found with id: %s", id)
	}

	return &video, nil
}
