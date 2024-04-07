package handler

import "github.com/labstack/echo/v4"

type IUploadfileHandler interface {
	SaveFiles(c echo.Context) error
	GetFileByID(c echo.Context) error
	GetFileAll(c echo.Context) error
	UpdateByID(c echo.Context) error
	DeleteFileByID(c echo.Context) error
	md5Filename(filename string) string
}
