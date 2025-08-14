package repositories

import (
	"fmt"

	"github.com/andre2ar/video-encoder/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{
		Db: db,
	}
}

func (jr *JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {
	if job.ID == "" {
		job.ID = uuid.New().String()
	}

	err := jr.Db.Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (jr *JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job

	jr.Db.Preload("Video").First(&job, "id = ?", id)
	if job.ID == "" {
		return nil, fmt.Errorf("job not found with id: %s", id)
	}

	return &job, nil
}

func (jr *JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := jr.Db.Save(&job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}
