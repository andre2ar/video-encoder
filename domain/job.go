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
	ID               string    `validate:"required,uuid4"`
	OutputBucketPath string    `validate:"required"`
	Status           string    `validate:"required"`
	Video            *Video    `validate:"required"`
	Error            string    `validate:"-"`
	CreatedAt        time.Time `validate:"required"`
	UpdatedAt        time.Time `validate:"required"`
}

func NewJob(outputBucketPath string, status string, video *Video) (*Job, error) {
	now := time.Now()
	job := Job{
		ID:               uuid.New().String(),
		OutputBucketPath: outputBucketPath,
		Status:           status,
		Video:            video,
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
