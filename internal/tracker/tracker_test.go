package tracker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/tracker"
)

func TestTracker(t *testing.T) {

	t.Run("successfully add item", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		added, err := track.AddItem(item)
		assert.NoError(t, err)
		assert.Equal(t, added, item)
	})

	t.Run("successfully get item", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		added, err := track.AddItem(item)

		res := track.GetItems()

		assert.NoError(t, err)
		assert.Equal(t, item, added)
		assert.Equal(t, res[0], item)
	})

	t.Run("successfully update without error", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		added, err := track.AddItem(item)
		assert.NotEmpty(t, added)
		assert.NoError(t, err)

		item.Name = "Second Item"
		res, err := track.UpdateItem(item)

		assert.NoError(t, err)
		assert.Equal(t, item, res)
	})

	t.Run("successfully delete item", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		added, err := track.AddItem(item)
		assert.NotEmpty(t, added)
		assert.NoError(t, err)

		track.DeleteItem("1")
		res := track.GetItems()
		assert.Empty(t, res)
	})

	t.Run("successfully find item", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item1 := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		item2 := tracker.Item{
			ID:   "2",
			Name: "Second Item",
		}

		added, err := track.AddItem(item1)
		assert.NotEmpty(t, added)
		assert.NoError(t, err)

		added, err = track.AddItem(item2)
		assert.NotEmpty(t, added)
		assert.NoError(t, err)

		res := track.FindItemByPartName("Item")
		assert.Len(t, res, 2)
	})

	t.Run("error add item - already exist", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		firstAdded, err := track.AddItem(item)
		assert.NoError(t, err)
		assert.NotEmpty(t, firstAdded)

		secondAdded, err := track.AddItem(item)
		assert.NotEmpty(t, secondAdded)
		assert.ErrorIs(t, err, tracker.ErrAlreadyExists)
	})

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		res, err := track.UpdateItem(item)
		assert.NotEmpty(t, res)
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})
}
