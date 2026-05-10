package base_test

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func TestLruCache(t *testing.T) {
	t.Parallel()

	t.Run("LruCache is successfully used", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		for i := 1; i <= 3; i++ {
			cache.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		cache.Get("key1")
		cache.Put("key4", "value4")

		assert.Equal(t, "value1", *cache.Get("key1"))
		assert.Equal(t, "value3", *cache.Get("key3"))
		assert.Equal(t, "value4", *cache.Get("key4"))
		assert.Nil(t, cache.Get("key2"))
	})

	t.Run("LruCache is frequently used", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(10)

		for i := 1; i <= 10; i++ {
			cache.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		for i := 10; i >= 1; i-- {
			cache.Get("key" + strconv.Itoa(i))
		}

		cache.Put("key11", "value11")
		cache.Put("key12", "value12")

		for i := 1; i <= 12; i++ {
			if i == 9 || i == 10 {
				assert.Nil(t, cache.Get("key"+strconv.Itoa(i)))
				continue
			}
			assert.Equal(t, "value"+strconv.Itoa(i), *cache.Get("key" + strconv.Itoa(i)))
		}

	})

	t.Run("LruCache is successfully used with match key", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		for i := 1; i <= 3; i++ {
			cache.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		cache.Get("key1")
		cache.Put("key2", "value22")
		cache.Put("key3", "value33")

		assert.Equal(t, "value1", *cache.Get("key1"))
		assert.Equal(t, "value22", *cache.Get("key2"))
		assert.Equal(t, "value33", *cache.Get("key3"))
	})

	t.Run("LruCache is empty", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)

		assert.Nil(t, cache.Get("key1"))
	})
}
