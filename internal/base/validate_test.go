package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ValidateRequest is nil", func(t *testing.T) {
		t.Parallel()

		var in *base.ValidateRequest
		rsl := base.Validate(in)
		expected := []string{"req is nil"}

		assert.Equal(t, expected, rsl)
	})

	t.Run("all fields are empty", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{}
		rsl := base.Validate(&in)
		expected := []string{
			"UserID is empty",
			"Title is empty",
			"Description is empty",
		}

		assert.Equal(t, expected, rsl)
	})

	t.Run("UserID is empty", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{
			Description: "Secret information",
			Title:       "Title",
		}
		rsl := base.Validate(&in)
		expected := []string{
			"UserID is empty",
		}

		assert.Equal(t, expected, rsl)
	})

	t.Run("Title is empty", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{
			UserID:      "1",
			Description: "Secret information",
		}
		rsl := base.Validate(&in)
		expected := []string{
			"Title is empty",
		}

		assert.Equal(t, expected, rsl)
	})

	t.Run("Description is empty", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{
			UserID: "1",
			Title:  "Title",
		}
		rsl := base.Validate(&in)
		expected := []string{
			"Description is empty",
		}

		assert.Equal(t, expected, rsl)
	})
}
