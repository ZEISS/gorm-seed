package seed_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	t.Parallel()

	s := seed.NewSeeder(nil)
	require.NotNil(t, s)
	require.Implements(t, (*seed.Seeder)(nil), s)
}

func TestSeedsAll(t *testing.T) {
	t.Parallel()

	s := seed.Seeds{
		seed.Seed{
			Name: "seed",
			Run: func(db *gorm.DB) error {
				return nil
			},
		},
	}

	require.Len(t, s.All(), 1)
}
