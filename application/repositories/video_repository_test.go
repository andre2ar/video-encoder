package repositories_test

import (
	"testing"

	"github.com/andre2ar/video-encoder/application/repositories"
	"github.com/andre2ar/video-encoder/domain"
	"github.com/andre2ar/video-encoder/framework/database"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsertAndFind(t *testing.T) {
	db := database.NewDatabaseTest()
	defer database.CloseTestDB(db)

	video := domain.NewVideo("test", "test")

	videoRepository := repositories.VideoRepositoryDb{Db: db}
	insertedVideo, err := videoRepository.Insert(video)
	require.Nil(t, err)
	require.NotEmpty(t, insertedVideo)

	retrievedVideo, err := videoRepository.Find(video.ID)
	require.Nil(t, err)
	require.NotEmpty(t, retrievedVideo)

	require.Equal(t, insertedVideo.ID, retrievedVideo.ID)
}
