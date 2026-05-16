package tracker

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Input interface {
	Get() string
}

type Output interface {
	Out(string)
}

type Store interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context) ([]Item, error)
	Get(ctx context.Context, id string) (Item, error)
}

type Usecase interface {
	Done(ctx context.Context, in Input, out Output, store Store) error
}

type AddUsecase struct{}

func (u AddUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()

	if err := store.Create(
		ctx,
		Item{ID: id, Name: name},
	); err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}
	return nil
}

type GetUsecase struct{}

func (u GetUsecase) Done(
	ctx context.Context,
	_ Input,
	out Output,
	store Store,
) error {
	items, err := store.List(ctx)
	if err != nil {
		return fmt.Errorf("failed to get items: %w", err)
	}
	for _, item := range items {
		out.Out(item.ID + " " + item.Name)
	}
	return nil
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
