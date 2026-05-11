package base_test

import (
	"fmt"
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func TestTracker(t *testing.T) {
	t.Parallel()
	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()
		tracker := base.NewTracker()
		getItems := tracker.GetItems()
		fmt.Println(getItems)
		item := base.Item{
			ID:   "1",
			Name: "First item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second item"

		assert.Equal(t,
			[]base.Item{item},
			tracker.GetItems(),
		)
	})

	t.Run("check multiple link leak", func(t *testing.T) {
		t.Parallel()
		tracker := base.NewTracker()
		items := []base.Item{
			{ID: "1", Name: "First"},
			{ID: "2", Name: "Second"},
			{ID: "3", Name: "Third"},
		}

		for _, item := range items {
			tracker.AddItem(item)
		}

		res := tracker.GetItems()
		for idx, item := range res {
			item.Name = "dummy" + strconv.Itoa(idx)
		}

		assert.Len(t, res, 3)
		assert.Equal(t,
			items,
			tracker.GetItems(),
		)
	})
}
