package team

import (
	"github.com/labstack/echo/v4"
	"github.com/yezarela/go-soccer/model"
	"github.com/yezarela/go-soccer/pkg/api"
)

// Handler represents the httphandler for team
type Handler struct {
	teamRepo Repository
}

// NewHandler initializes endpoints for team
func NewHandler(e *echo.Echo, teamRepo Repository) {
	handler := &Handler{
		teamRepo: teamRepo,
	}

	e.GET("/teams", handler.GetAll)
	e.POST("/teams", handler.Post)
	e.GET("/teams/:id", handler.GetByID)
}

// GetAll returns list of teams
func (h *Handler) GetAll(c echo.Context) error {

	ctx := c.Request().Context()

	res, err := h.teamRepo.ListTeam(ctx)
	if err != nil {
		return api.ResponseError(c, err)
	}

	return api.ResponseOK(c, res)
}

// GetByID returns a team by id
func (h *Handler) GetByID(c echo.Context) error {

	ctx := c.Request().Context()

	res, err := h.teamRepo.GetTeam(ctx, c.Param("id"))
	if err != nil {
		return api.ResponseError(c, err)
	}

	if res == nil {
		return api.ResponseNotFound(c, "cannot find the requested team")
	}

	return api.ResponseOK(c, res)
}

// Post creates a new team
func (h *Handler) Post(c echo.Context) error {

	ctx := c.Request().Context()

	var body model.Team
	err := c.Bind(&body)
	if err != nil {
		return api.ResponseUnprocessableEntity(c, "invalid body")
	}

	if len(body.Name) <= 0 {
		return api.ResponseBadRequest(c, "name cannot be empty")
	}
	if len(body.Location) <= 0 {
		return api.ResponseBadRequest(c, "location cannot be empty")
	}

	for _, p := range body.Players {
		if len(p.Name) <= 0 {
			return api.ResponseBadRequest(c, "player name cannot be empty")
		}
		if len(p.Position) <= 0 {
			return api.ResponseBadRequest(c, "player position cannot be empty")
		}
	}

	res, err := h.teamRepo.CreateTeam(ctx, body)
	if err != nil {
		return api.ResponseError(c, err)
	}

	return api.ResponseOK(c, res)
}
