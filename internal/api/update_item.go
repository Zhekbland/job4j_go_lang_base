package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type UpdateItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var req UpdateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	validationErrors := make([]string, 0)
	if id == "" {
		validationErrors = append(validationErrors, "id is required")
	}
	if req.Name == "" {
		validationErrors = append(validationErrors, "Name is required")
	}
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "validation failed",
			"details": validationErrors,
		})
	}

	err := s.Repository.Update(c.Context(), tracker.Item{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		if errors.Is(err, tracker.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "item not found")
		}
		log.Errorw("s.Repository.Update", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
