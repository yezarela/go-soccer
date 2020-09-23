package player

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/yezarela/go-soccer/model"
	playerMock "github.com/yezarela/go-soccer/module/player/mock"
	"github.com/yezarela/go-soccer/pkg/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetByID(t *testing.T) {

	t.Run("Response OK", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := playerMock.NewMockRepository(ctrl)

		mockPlayer := model.Player{}
		mockPlayer.ID, _ = primitive.ObjectIDFromHex("5f6a5d6129b2289c40b7444b")

		mockResp, _ := json.Marshal(api.Response{
			Meta: api.ResponseMeta{
				Code: http.StatusOK,
			},
			Data: mockPlayer,
		})

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/players/:id")
		c.SetParamNames("id")
		c.SetParamValues(mockPlayer.ID.Hex())

		h := &Handler{
			playerRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().GetPlayer(ctx, mockPlayer.ID.Hex()).Return(&mockPlayer, nil)

		// Assertions
		if assert.NoError(t, h.GetByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(mockResp), strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})

	t.Run("Response Not Found", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := playerMock.NewMockRepository(ctrl)
		mockPlayerID := "5f6a5d6129b2289c40b7444b"

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/players/:id")
		c.SetParamNames("id")
		c.SetParamValues(mockPlayerID)

		h := &Handler{
			playerRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().GetPlayer(ctx, mockPlayerID).Return(nil, nil)

		// Assertions
		if assert.NoError(t, h.GetByID(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestPost(t *testing.T) {

	t.Run("Response Created", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := playerMock.NewMockRepository(ctrl)

		mockPlayer := model.Player{}
		mockPlayer.Name = "Ronaldo"
		mockPlayer.Position = "Captain"

		mockPayload, _ := json.Marshal(mockPlayer)
		mockResp, _ := json.Marshal(api.Response{
			Meta: api.ResponseMeta{
				Code: http.StatusCreated,
			},
			Data: mockPlayer,
		})

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/players", strings.NewReader(string(mockPayload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{
			playerRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().CreatePlayer(ctx, mockPlayer).Return(&mockPlayer, nil)

		// Assertions
		if assert.NoError(t, h.Post(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, string(mockResp), strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})

	t.Run("Response Bad Request", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := playerMock.NewMockRepository(ctrl)

		mockPlayer := model.Player{}
		mockPlayer.Name = "Ronaldo"

		mockPayload, _ := json.Marshal(mockPlayer)

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/players", strings.NewReader(string(mockPayload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{
			playerRepo: mockRepo,
		}

		// Assertions
		if assert.NoError(t, h.Post(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestGetAll(t *testing.T) {

	t.Run("Response OK", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := playerMock.NewMockRepository(ctrl)

		mockPlayer := model.Player{}
		mockPlayer.ID, _ = primitive.ObjectIDFromHex("5f6a5d6129b2289c40b7444b")
		mockPlayers := []model.Player{mockPlayer}

		mockResp, _ := json.Marshal(api.Response{
			Meta: api.ResponseMeta{
				Code: http.StatusOK,
			},
			Data: mockPlayers,
		})

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/players", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{
			playerRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().ListPlayer(ctx).Return(mockPlayers, nil)

		// Assertions
		if assert.NoError(t, h.GetAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(mockResp), strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})
}
