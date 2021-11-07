package handler

import (
	"bytes"
	"context"
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
			name:                "Missing Fields",
			inputBody:           `{"email": "Test"}`,
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"required field is missing"}`,
		},
		// todo после реализации валидации добавить/переработать тест
		//{
		//	name:                "Not Valid Fields",
		//	inputBody:           `{"name": "Test","email": "","password": "qwerty","role": "Boss"}`,
		//	expectedStatusCode:  400,
		//	expectedRequestBody: `{"message":"invalid input body"}`,
		//},
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
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorizationService(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(auth, testCase.inputUser)
			}

			h := NewHandler(auth)

			// Test server
			r := gin.New()
			r.POST("/sign-up", h.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	testTable := []struct {
		name                string
		inputBody           string
		userEmail           string
		userPassword        string
		mockBehavior        func(s *mock_service.MockAuthorizationService, userEmail string, userPassword string)
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:         "OK",
			inputBody:    `{"email": "test","password": "qwerty"}`,
			userEmail:    "test",
			userPassword: "qwerty",
			mockBehavior: func(s *mock_service.MockAuthorizationService, userEmail string, userPassword string) {
				s.EXPECT().SignIn(context.Background(), userEmail, userPassword).
					Return(models.Tokens{AccessToken: "string", RefreshToken: "string"}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"accessToken":"string","refreshToken":"string"}`,
		},
		{
			name:                "Missing Fields",
			inputBody:           `{"email": "Test"}`,
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"required field is missing"}`,
		},
		{
			name:         "SignIn Failure",
			inputBody:    `{"email": "test","password": "qwerty"}`,
			userEmail:    "test",
			userPassword: "qwerty",
			mockBehavior: func(s *mock_service.MockAuthorizationService, userEmail string, userPassword string) {
				s.EXPECT().SignIn(context.Background(), userEmail, userPassword).
					Return(models.Tokens{}, errors.New("get user: record not found"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"get user: record not found"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorizationService(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(auth, testCase.userEmail, testCase.userPassword)
			}
			h := NewHandler(auth)

			// Test server
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.POST("/sign-in", h.SignIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
