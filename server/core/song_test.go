package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSongId(t *testing.T) {
	testCases := []struct {
		name        string
		songName    string
		artistNames []string
		albumName   string
		expected    string
	}{
		{
			name:        "happy path",
			songName:    "Billie Jean",
			artistNames: []string{"Michael Jackson"},
			albumName:   "Thriller",
			expected:    "42546cc0-2143-59b6-a236-cc1eeb577c9c",
		},
	}
	for _, tt := range testCases {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.expected, GetSongId(tt.songName, tt.artistNames, tt.albumName))
			},
		)
	}
}
