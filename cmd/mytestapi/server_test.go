package mytestapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"mytestapi/cmd/mytestapi/mocks"
	"mytestapi/cmd/mytestapi/models"
)

func TestServer_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProfileService := mocks.NewMockIProfileService(ctrl)

	server := &Server{
		ProfileService: mockProfileService,
	}

	tests := []struct {
		name           string
		method         string
		username       string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Method not allowed",
			method:         http.MethodPost,
			username:       "",
			mockSetup:      func() {},
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed",
		},
		{
			name:     "Get all profiles",
			method:   http.MethodGet,
			username: "",
			mockSetup: func() {
				mockProfileService.EXPECT().GetProfiles().Return([]models.UserProfile{
					{ID: 1, Username: "user1", FirstName: "FirstName", LastName: "LastName"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"username":"user1","first_name":"FirstName","last_name":"LastName","city":"","school":""}]`,
		},
		{
			name:     "Get profile by username",
			method:   http.MethodGet,
			username: "user1",
			mockSetup: func() {
				mockProfileService.EXPECT().GetProfileByUsername("user1").Return(&models.UserProfile{
					ID: 1, Username: "user1", FirstName: "FirstName", LastName: "LastName",
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"username":"user1","first_name":"FirstName","last_name":"LastName","city":"","school":""}`,
		},
		{
			name:     "Profile not found",
			method:   http.MethodGet,
			username: "user2",
			mockSetup: func() {
				mockProfileService.EXPECT().GetProfileByUsername("user2").Return(nil, ErrNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "no profile found with user ID user2",
		},
		{
			name:     "Internal server error",
			method:   http.MethodGet,
			username: "user3",
			mockSetup: func() {
				mockProfileService.EXPECT().GetProfileByUsername("user3").Return(nil, errors.New("internal error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Error retrieving user profile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest(tt.method, "/profile", nil)
			q := req.URL.Query()
			if tt.username != "" {
				q.Add("username", tt.username)
			}
			req.URL.RawQuery = q.Encode()
			w := httptest.NewRecorder()
			server.GetProfile().ServeHTTP(w, req)
			resp := w.Result()
			defer resp.Body.Close()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response body: %v", err)
			}

			if strings.TrimSpace(string(body)) != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, body)
			}
		})
	}

}
