package team

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
	teamMock "github.com/yezarela/go-soccer/module/team/mock"
	"github.com/yezarela/go-soccer/pkg/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetByID(t *testing.T) {

	t.Run("Response OK", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := teamMock.NewMockRepository(ctrl)

		mockTeam := model.Team{}
		mockTeam.ID, _ = primitive.ObjectIDFromHex("5f6a5d6129b2289c40b7444b")

		mockResp, _ := json.Marshal(api.Response{
			Meta: api.ResponseMeta{
				Code: http.StatusOK,
			},
			Data: mockTeam,
		})

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/teams/:id")
		c.SetParamNames("id")
		c.SetParamValues(mockTeam.ID.Hex())

		h := &Handler{
			teamRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().GetTeam(ctx, mockTeam.ID.Hex()).Return(&mockTeam, nil)

		// Assertions
		if assert.NoError(t, h.GetByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(mockResp), strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})

	t.Run("Response Not Found", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := teamMock.NewMockRepository(ctrl)
		mockTeamID := "5f6a5d6129b2289c40b7444b"

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/teams/:id")
		c.SetParamNames("id")
		c.SetParamValues(mockTeamID)

		h := &Handler{
			teamRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().GetTeam(ctx, mockTeamID).Return(nil, nil)

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
		mockRepo := teamMock.NewMockRepository(ctrl)

		mockTeam := model.Team{}
		mockTeam.Name = "Arsenal"
		mockTeam.Location = "Jakarta"

		mockPayload, _ := json.Marshal(mockTeam)
		mockResp, _ := json.Marshal(api.Response{
			Meta: api.ResponseMeta{
				Code: http.StatusCreated,
			},
			Data: mockTeam,
		})

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/teams", strings.NewReader(string(mockPayload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{
			teamRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().CreateTeam(ctx, mockTeam).Return(&mockTeam, nil)

		// Assertions
		if assert.NoError(t, h.Post(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, string(mockResp), strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})

	t.Run("Response Bad Request", func(t *testing.T) {

		// Mock
		ctrl := gomock.NewController(t)
		mockRepo := teamMock.NewMockRepository(ctrl)

		mockTeam := model.Team{}
		mockTeam.Name = "Arsenal"

		mockPayload, _ := json.Marshal(mockTeam)

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/teams", strings.NewReader(string(mockPayload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{
			teamRepo: mockRepo,
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
		mockRepo := teamMock.NewMockRepository(ctrl)

		mockTeam := model.Team{}
		mockTeam.ID, _ = primitive.ObjectIDFromHex("5f6a5d6129b2289c40b7444b")
		mockTeams := []model.Team{mockTeam}

		mockResp, _ := json.Marshal(api.Response{
			Meta: api.ResponseMeta{
				Code: http.StatusOK,
			},
			Data: mockTeams,
		})

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/teams", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{
			teamRepo: mockRepo,
		}

		ctx := c.Request().Context()

		mockRepo.EXPECT().ListTeam(ctx).Return(mockTeams, nil)

		// Assertions
		if assert.NoError(t, h.GetAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(mockResp), strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})
}
