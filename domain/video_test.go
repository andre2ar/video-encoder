package domain_test

import (
	"testing"

	"github.com/andre2ar/video-encoder/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo("", "")
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo("a", "path")

	video.ID = "not-uuid"

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo("a", "path")

	video.ID = uuid.New().String()

	err := video.Validate()
	require.Nil(t, err)
}
