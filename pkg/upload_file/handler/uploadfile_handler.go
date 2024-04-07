package handler

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Pugpaprika21/pkg/upload_file/dto"
	"github.com/Pugpaprika21/pkg/upload_file/repository"
	"github.com/labstack/echo/v4"
)

type handler struct {
	repository repository.IUploadFileRepository
}

func NewHandler(repository repository.IUploadFileRepository) *handler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) SaveFiles(c echo.Context) error {
	clientFiles, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.UploadfileRespones{
			Message:    err.Error(),
			StatusBool: false,
			Data:       nil,
		})
	}

	files := clientFiles.File
	for _, fileHeaders := range files {
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
					Message:    err.Error(),
					StatusBool: false,
					Data:       nil,
				})
			}
			defer file.Close()

			extension := filepath.Ext(fileHeader.Filename)
			encodedFilename := h.md5Filename(fileHeader.Filename)

			uploadDir := "./upload"
			if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
				os.Mkdir(uploadDir, os.ModePerm)
			}

			path := filepath.Join(uploadDir, encodedFilename+extension)
			dst, err := os.Create(path)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
					Message:    err.Error(),
					StatusBool: false,
					Data:       nil,
				})
			}

			defer dst.Close()

			if _, err = io.Copy(dst, file); err != nil {
				return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
					Message:    err.Error(),
					StatusBool: false,
					Data:       nil,
				})
			}

			err = h.repository.Save(&dto.UploadfileDTO{
				Filename:  encodedFilename,
				Path:      path,
				Size:      fileHeader.Size,
				MimeType:  fileHeader.Header.Get("Content-Type"),
				Extension: extension,
			})

			if err != nil {
				return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
					Message:    err.Error(),
					StatusBool: false,
					Data:       nil,
				})
			}
		}
	}

	return c.JSON(http.StatusOK, dto.UploadfileRespones{
		Message:    "save file success",
		StatusBool: true,
		Data:       nil,
	})
}

func (h *handler) GetFileByID(c echo.Context) error {
	id := c.Param("id")

	result := h.repository.GetByID(id)
	if result.ID == 0 {
		return c.JSON(http.StatusOK, dto.UploadfileRespones{
			Message:    "",
			StatusBool: false,
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, dto.UploadfileRespones{
		Message:    "",
		StatusBool: true,
		Data:       result,
	})
}

func (h *handler) GetFileAll(c echo.Context) error {
	results := h.repository.GetAll()
	return c.JSON(http.StatusOK, dto.UploadfileRespones{
		Message:    "",
		StatusBool: true,
		Data:       results,
	})
}

func (h *handler) UpdateByID(c echo.Context) error {
	id := c.Param("id")

	clientFiles, err := c.FormFile("client_files")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.UploadfileRespones{
			Message:    err.Error(),
			StatusBool: false,
			Data:       nil,
		})
	}

	file, err := clientFiles.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
			Message:    err.Error(),
			StatusBool: false,
			Data:       nil,
		})
	}
	defer file.Close()

	extension := filepath.Ext(clientFiles.Filename)
	encodedFilename := h.md5Filename(clientFiles.Filename)

	uploadDir := "./upload"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	path := filepath.Join(uploadDir, encodedFilename+extension)
	dst, err := os.Create(path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
			Message:    err.Error(),
			StatusBool: false,
			Data:       nil,
		})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
			Message:    err.Error(),
			StatusBool: false,
			Data:       nil,
		})
	}

	body := &dto.UploadfileDTO{
		Filename:  encodedFilename,
		Path:      path,
		Size:      clientFiles.Size,
		MimeType:  clientFiles.Header.Get("Content-Type"),
		Extension: extension,
	}

	err = h.repository.UpdateByID(id, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
			Message:    err.Error(),
			StatusBool: false,
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, dto.UploadfileRespones{
		Message:    "update file success",
		StatusBool: true,
		Data:       nil,
	})
}

func (h *handler) DeleteFileByID(c echo.Context) error {
	id := c.Param("id")

	err := h.repository.DeleteByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.UploadfileRespones{
			Message:    err.Error(),
			StatusBool: false,
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, dto.UploadfileRespones{
		Message:    "delete file success",
		StatusBool: true,
		Data:       nil,
	})
}

func (h *handler) md5Filename(filename string) string {
	hash := md5.New()
	hash.Write([]byte(filename + time.Now().String()))
	return hex.EncodeToString(hash.Sum(nil))
}
