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

		res, _ := track.AddItem(item)
		assert.Equal(t, res, item)
	})

	t.Run("successfully get item", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		_, _ = track.AddItem(item)

		res := track.GetItems()
		assert.Equal(t, res[0], item)
	})

	t.Run("successfully update without error", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		_, _ = track.AddItem(item)

		item.Name = "Second Item"
		_, err := track.UpdateItem(item)
		assert.Nil(t, err)
	})

	t.Run("successfully delete item", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		_, _ = track.AddItem(item)
		track.DeleteItem("1")

		res := track.GetItems()
		assert.Equal(t, len(res), 0)
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
		_, _ = track.AddItem(item1)
		_, _ = track.AddItem(item2)

		res := track.FindItemByPartName("Item")
		assert.Equal(t, len(res), 2)
	})

	t.Run("error add item - already exist", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		_, _ = track.AddItem(item)
		_, err := track.AddItem(item)

		assert.ErrorIs(t, err, tracker.ErrAlreadyExists)
	})

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		track := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		_, err := track.UpdateItem(item)
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})
}
