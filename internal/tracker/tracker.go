package tracker

import (
	"fmt"
	"strings"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.items = append(t.items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) DeleteItem(id string) bool {
	idx, ok := t.indexOf(id)
	if !ok {
		return false
	}

	t.items = append(t.items[:idx], t.items[idx+1:]...)
	return true
}

func (t *Tracker) UpdateItem(id string, newName string) bool {
	idx, ok := t.indexOf(id)
	if !ok {
		return false
	}

	t.items[idx].Name = newName
	return true
}

func (t *Tracker) FindItemByPartName(partName string) []Item {
	res := make([]Item, 0)
	for _, item := range t.items {
		if strings.Contains(item.Name, partName) {
			res = append(res, item)
		}
	}

	return res
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for idx, item := range t.items {
		if item.ID == id {
			return idx, true
		}
	}
	return -1, false
}
