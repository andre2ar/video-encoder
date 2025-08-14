package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var jobValidate *validator.Validate

func init() {
	jobValidate = validator.New()
}

type Job struct {
	ID               string    `json:"job_id" validate:"required,uuid4" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `json:"output_bucket_path" validate:"required"`
	Status           string    `json:"status" validate:"required"`
	Video            *Video    `json:"video" validate:"required"`
	VideoID          string    `json:"-" validate:"required" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `json:"error" validate:"-"`
	CreatedAt        time.Time `json:"created_at" validate:"required"`
	UpdatedAt        time.Time `json:"updated_at" validate:"required"`
}

func NewJob(outputBucketPath string, status string, video *Video) (*Job, error) {
	now := time.Now()
	job := Job{
		ID:               uuid.New().String(),
		OutputBucketPath: outputBucketPath,
		Status:           status,
		Video:            video,
		VideoID:          video.ID,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	err := job.Validate()
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (j *Job) Validate() error {
	err := jobValidate.Struct(j)
	return err
}
