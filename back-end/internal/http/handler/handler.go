package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func (h *Handler) GetPersonalInformation(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) GetListOfProjects(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) GetASingleProject(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusAccepted)
}

func (h *Handler) Register(api fiber.Router) {
	api.Get("/personalInfo", h.GetPersonalInformation)
	api.Get("/projects", h.GetListOfProjects)
	api.Get("/projects/:id", h.GetASingleProject)
}
