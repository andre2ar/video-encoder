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
	ID         string    `validate:"required,uuid"`
	ResourceID string    `validate:"required"`
	FilePath   string    `validate:"required"`
	CreatedAt  time.Time `validate:"required"`
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
