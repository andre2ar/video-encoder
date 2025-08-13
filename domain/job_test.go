package domain_test

import (
	"testing"

	"github.com/andre2ar/video-encoder/domain"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo("a", "path")
	job, err := domain.NewJob("output", "Converted", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}
