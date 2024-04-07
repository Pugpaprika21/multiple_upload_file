package repository

import "github.com/Pugpaprika21/pkg/upload_file/dto"

type IUploadFileRepository interface {
	Save(body *dto.UploadfileDTO) error
	GetByID(id string) *dto.UploadfileRecord
	GetAll() []*dto.UploadfileRecord
	UpdateByID(id string, body *dto.UploadfileDTO) error
	DeleteByID(id string) error
}
