package repositories_test

import (
	"testing"

	"github.com/andre2ar/video-encoder/application/repositories"
	"github.com/andre2ar/video-encoder/domain"
	"github.com/andre2ar/video-encoder/framework/database"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsertAndFind(t *testing.T) {
	db := database.NewDatabaseTest()
	defer database.CloseTestDB(db)

	video := domain.NewVideo("test", "test")
	videoRepository := repositories.VideoRepositoryDb{Db: db}
	videoRepository.Insert(video)

	job, err := domain.NewJob("test", "test", video)

	jobRepository := repositories.JobRepositoryDb{Db: db}
	insertedJob, err := jobRepository.Insert(job)
	require.Nil(t, err)
	require.NotEmpty(t, insertedJob)

	retrievedJob, err := jobRepository.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, retrievedJob)

	require.Equal(t, insertedJob.ID, retrievedJob.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDatabaseTest()
	defer database.CloseTestDB(db)

	video := domain.NewVideo("test", "test")
	videoRepository := repositories.VideoRepositoryDb{Db: db}
	videoRepository.Insert(video)

	job, err := domain.NewJob("test", "test", video)

	jobRepository := repositories.JobRepositoryDb{Db: db}
	insertedJob, err := jobRepository.Insert(job)
	require.Nil(t, err)
	require.NotEmpty(t, insertedJob)

	job.Status = "Complete"
	jobRepository.Update(job)

	retrievedJob, err := jobRepository.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, retrievedJob)

	require.Equal(t, insertedJob.ID, retrievedJob.ID)
	require.Equal(t, "Complete", job.Status)
}
