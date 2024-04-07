package uploadFileRouter

import (
	"github.com/Pugpaprika21/pkg/upload_file/handler"
	"github.com/Pugpaprika21/pkg/upload_file/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(e *echo.Echo, db *gorm.DB) {
	var uploadFileRepository repository.IUploadFileRepository = repository.NewRepository(db)
	var uploadFileHandler handler.IUploadfileHandler = handler.NewHandler(uploadFileRepository)

	group := e.Group("/api/v1")
	group.POST("/saveFiles", uploadFileHandler.SaveFiles)
	group.GET("/getFileByID/:id", uploadFileHandler.GetFileByID)
	group.GET("/getFileAll", uploadFileHandler.GetFileAll)
	group.PUT("/updateByID/:id", uploadFileHandler.UpdateByID)
	group.DELETE("/deleteFileByID/:id", uploadFileHandler.DeleteFileByID)
}
