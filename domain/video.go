package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var videoValidate *validator.Validate

func init() {
	videoValidate = validator.New()
}

type Video struct {
	ID         string    `json:"encoded_video_folder" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	ResourceID string    `json:"resource_id" validate:"required" gorm:"type:varchar(255)"`
	FilePath   string    `json:"file_path" validate:"required" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"-" validate:"required"`
	Jobs       []*Job    `json:"-" validate:"-" gorm:"ForeignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:cascade"`
}

func NewVideo(resourceID string, filePath string) *Video {
	return &Video{
		ID:         uuid.New().String(),
		ResourceID: resourceID,
		FilePath:   filePath,
		CreatedAt:  time.Now(),
	}
}

func (v *Video) Validate() error {
	err := videoValidate.Struct(v)
	return err
}
