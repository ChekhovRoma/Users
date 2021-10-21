package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
	"users/models"
	mock_service "users/pkg/handler/mocks"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorizationService, user models.User)

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
			mockBehavior: func(s *mock_service.MockAuthorizationService, user models.User) {
				s.EXPECT().SignUp(user.Name, user.Email, user.Password, user.Role).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Empty Fields",
			inputBody:           `{"name": "Test","password": "qwerty","role": "Boss"}`,
			mockBehavior:        func(s *mock_service.MockAuthorizationService, user models.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name": "Test","email": "test","password": "qwerty","role": "Boss"}`,
			inputUser: models.User{
				Name:     "Test",
				Email:    "test",
				Password: "qwerty",
				Role:     "Boss",
			},
			mockBehavior: func(s *mock_service.MockAuthorizationService, user models.User) {
				s.EXPECT().SignUp(user.Name, user.Email, user.Password, user.Role).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorizationService(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			h := NewHandler(auth)

			r := gin.New()
			r.POST("/sign-up", h.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
