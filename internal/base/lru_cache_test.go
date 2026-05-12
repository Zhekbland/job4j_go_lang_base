package base_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func TestLruCache(t *testing.T) {
	t.Parallel()

	t.Run("LruCache is successfully used", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		lo.ForEach([]string{"1", "2", "3"}, func(idx string, _ int) {
			cache.Put("key"+idx, "value"+idx)
		})

		cache.Get("key1")
		cache.Put("key4", "value4")

		assert.Equal(t, "value1", *cache.Get("key1"))
		assert.Equal(t, "value3", *cache.Get("key3"))
		assert.Equal(t, "value4", *cache.Get("key4"))
		assert.Nil(t, cache.Get("key2"))
	})

	t.Run("LruCache shift keys and remove rarely used", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(5)

		lo.ForEach([]string{"1", "2", "3", "4", "5"}, func(idx string, _ int) {
			cache.Put("key"+idx, "value"+idx)
		})
		lo.ForEach([]string{"5", "4", "3", "2", "1"}, func(idx string, _ int) {
			cache.Get("key" + idx)
		})
		lo.ForEach([]string{"6", "7"}, func(idx string, _ int) {
			cache.Put("key"+idx, "value"+idx)
		})

		assert.Nil(t, cache.Get("key4"))
		assert.Nil(t, cache.Get("key5"))

	})

	t.Run("LruCache is successfully update value", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		cache.Put("key1", "value1")
		cache.Put("key1", "value11")

		assert.Equal(t, "value11", *cache.Get("key1"))
	})

	t.Run("LruCache is empty", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		assert.Nil(t, cache.Get("key1"))
	})
}
