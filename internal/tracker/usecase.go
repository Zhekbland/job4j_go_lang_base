package tracker

import "github.com/google/uuid"

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(in Input, out Output, tracker *Tracker) {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter uuid:")
	id := in.Get()
	out.Out("enter new name:")
	newName := in.Get()

	res := tracker.UpdateItem(id, newName)
	if !res {
		out.Out("Update failed or item not found")
		return
	}

	out.Out("Successfully updated")
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter uuid:")
	id := in.Get()

	res := tracker.DeleteItem(id)
	if !res {
		out.Out("Delete failed or item not found")
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
