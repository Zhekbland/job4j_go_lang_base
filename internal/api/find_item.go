package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (s *Server) FindByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	items, err := s.Repository.FindByName(c.Context(), name)
	if err != nil {
		log.Errorw("s.Repository.FindByName", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	res := make([]ItemRequest, 0, len(items))
	for _, item := range items {
		res = append(res, ItemRequest{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(GetItemsResponse{Items: res})
}
