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
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.Items = append(t.Items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) DeleteItem(id string) bool {
	for idx, item := range t.Items {
		if item.ID == id {
			t.Items = append(t.Items[:idx], t.Items[idx+1:]...)
			return true
		}
	}

	return false
}

func (t *Tracker) UpdateItem(id string, newName string) bool {
	for idx, item := range t.Items {
		if item.ID == id {
			t.Items[idx].Name = newName
			return true
		}
	}

	return false
}

func (t *Tracker) FindItemByPartName(partName string) []Item {
	res := make([]Item, 0)
	for _, item := range t.Items {
		if strings.Contains(item.Name, partName) {
			res = append(res, item)
		}
	}

	return res
}
