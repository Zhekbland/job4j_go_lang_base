package tracker

import (
	"github.com/google/uuid"
)

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()

	item, err := tracker.AddItem(Item{Name: name, ID: id})
	if err != nil {
		out.Out("Failed to add item: " + err.Error())
		return
	}

	out.Out("Successfully add item: " + item.toString())
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.items {
		out.Out(item.toString())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter uuid:")
	id := in.Get()
	out.Out("enter new name:")
	newName := in.Get()
	item := Item{
		ID:   id,
		Name: newName,
	}

	res, err := tracker.UpdateItem(item)
	if err != nil {
		out.Out("Failed to update item: " + err.Error())
		return
	}

	out.Out("Successfully updated " + res.toString())
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter uuid:")
	id := in.Get()

	res := tracker.DeleteItem(id)
	if !res {
		out.Out("Item not found")
		return
	}

	out.Out("Successfully deleted")
}

type FindNameUsecase struct{}

func (u FindNameUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter part of name:")
	partName := in.Get()

	res := tracker.FindItemByPartName(partName)
	if len(res) == 0 {
		out.Out("Items with such partName not found")
		return
	}

	for _, item := range res {
		out.Out(item.toString())
	}
}
