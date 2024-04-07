package dto

type UploadfileDTO struct {
	Filename  string
	Path      string
	Size      int64
	MimeType  string
	Extension string
}

type UploadfileRecord struct {
	ID        uint   `gorm:"id" json:"id"`
	Filename  string `gorm:"filename" json:"filename"`
	Path      string `gorm:"path" json:"path"`
	Size      int64  `gorm:"size" json:"size"`
	MimeType  string `gorm:"mime_type" json:"mime_type"`
	Extension string `gorm:"extension" json:"extension"`
}

type UploadfileRespones struct {
	Message    string      `json:"message"`
	StatusBool bool        `json:"status_bool"`
	Data       interface{} `json:"data,omitempty"`
}
