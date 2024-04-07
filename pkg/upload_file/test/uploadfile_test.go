package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pugpaprika21/db"
	uploadFileRouter "github.com/Pugpaprika21/pkg/upload_file/router"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUploadfile(t *testing.T) {
	e := echo.New()

	db, err := db.New().UseDB()
	if err != nil {
		t.Fatal(err)
	}

	uploadFileRouter.Register(e, db)

	testCases := []struct {
		name     string
		method   string
		path     string
		expected int
	}{
		{"SaveFiles", http.MethodPost, "/api/v1/saveFiles", http.StatusOK},
		{"GetFileByID", http.MethodGet, "/api/v1/getFileByID/1", http.StatusOK},
		{"GetFileAll", http.MethodGet, "/api/v1/getFileAll", http.StatusOK},
		{"UpdateByID", http.MethodPut, "/api/v1/updateByID/1", http.StatusOK},
		{"DeleteFileByID", http.MethodDelete, "/api/v1/deleteFileByID/1", http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expected, rec.Code)
		})
	}
}
