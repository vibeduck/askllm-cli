package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Ask(t *testing.T) {
	tests := []struct {
		name           string
		serverResponse string
		serverStatus   int
		question       string
		wantErr        bool
	}{
		{
			name: "successful response",
			serverResponse: `{
				"choices": [
					{
						"message": {
							"content": "Hello, World!"
						}
					}
				]
			}`,
			serverStatus: http.StatusOK,
			question:     "Say hello",
			wantErr:      false,
		},
		{
			name:           "server error",
			serverResponse: `{"error": "Internal Server Error"}`,
			serverStatus:   http.StatusInternalServerError,
			question:       "Say hello",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverStatus)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			client := NewClient("test-api-key")
			client.endpoint = server.URL

			response, err := client.Ask(tt.question)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Ask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && response == "" {
				t.Error("Client.Ask() returned empty response for successful case")
			}
		})
	}
}
