package player

import (
	"github.com/labstack/echo/v4"
	"github.com/yezarela/go-soccer/model"
	"github.com/yezarela/go-soccer/pkg/api"
)

// Handler represents the httphandler for player
type Handler struct {
	playerRepo Repository
}

// NewHandler initializes endpoints for player
func NewHandler(e *echo.Echo, playerRepo Repository) {
	handler := &Handler{
		playerRepo: playerRepo,
	}

	e.GET("/players", handler.GetAll)
	e.POST("/players", handler.Post)
	e.GET("/players/:id", handler.GetByID)
}

// GetAll returns list of players
func (h *Handler) GetAll(c echo.Context) error {

	ctx := c.Request().Context()

	res, err := h.playerRepo.ListPlayer(ctx)
	if err != nil {
		return api.ResponseError(c, err)
	}

	return api.ResponseOK(c, res)
}

// GetByID returns a player by id
func (h *Handler) GetByID(c echo.Context) error {

	ctx := c.Request().Context()

	res, err := h.playerRepo.GetPlayer(ctx, c.Param("id"))
	if err != nil {
		return api.ResponseError(c, err)
	}

	if res == nil {
		return api.ResponseNotFound(c, "cannot find the requested player")
	}

	return api.ResponseOK(c, res)
}

// Post creates a new player
func (h *Handler) Post(c echo.Context) error {

	ctx := c.Request().Context()

	var body model.Player
	err := c.Bind(&body)
	if err != nil {
		return api.ResponseUnprocessableEntity(c, "invalid body")
	}

	if len(body.Name) <= 0 {
		return api.ResponseBadRequest(c, "name cannot be empty")
	}
	if len(body.Nickname) <= 0 {
		return api.ResponseBadRequest(c, "nickname cannot be empty")
	}
	if len(body.Position) <= 0 {
		return api.ResponseBadRequest(c, "position cannot be empty")
	}

	res, err := h.playerRepo.CreatePlayer(ctx, body)
	if err != nil {
		return api.ResponseError(c, err)
	}

	return api.ResponseOK(c, res)
}
