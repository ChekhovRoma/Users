package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
	"users/models"
	"users/pkg/service"
	mock_service "users/pkg/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user models.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test","email": "test","password": "qwerty","role": "Boss"}`,
			inputUser: models.User{
				Name:     "Test",
				Email:    "test",
				Password: "qwerty",
				Role:     "Boss",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().SignUp(user.Name, user.Email, user.Password, user.Role).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id": 1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.AuthorizationService{}
			h := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", h.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
