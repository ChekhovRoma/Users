package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
	mock_service "users/pkg/service/mocks"
)

func TestManager_Parse(t *testing.T) {
	testTable := []struct {
		name            string
		accessToken     string
		mockBehavior    func(tm *mock_service.MockTokenManager, accessToken string)
		expectedRequest string
		expectedError   error
	}{
		{
			name:        "OK",
			accessToken: "token",
			mockBehavior: func(tm *mock_service.MockTokenManager, accessToken string) {
				tm.EXPECT().Parse(accessToken).Return("1", nil)
			},
			expectedRequest: "1",
			expectedError:   nil,
		},
		{
			name:        "Wrong Signing Method",
			accessToken: "notHS256token",
			mockBehavior: func(tm *mock_service.MockTokenManager, accessToken string) {
				tm.EXPECT().Parse(accessToken).Return("", fmt.Errorf("unexpected signing method: RSAPSS"))
			},
			expectedRequest: "",
			expectedError:   fmt.Errorf("unexpected signing method: RSAPSS"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			tm := mock_service.NewMockTokenManager(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(tm, testCase.accessToken)
			}

			result, err := tm.Parse(testCase.accessToken)

			assert.Equal(t, result, testCase.expectedRequest)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}
