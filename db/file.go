package db

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type File struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name,omitempty"`
	Data      []byte    `gorm:"column:data" json:"data,omitempty"`
	CreatedBy uuid.UUID `gorm:"column:createdBy" json:"createdBy"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt" ignore:"true"`
	MimeType  string    `gorm:"column:mimeType" json:"mimeType"`
}

type FileStore interface {
	ReadFile(ctx context.Context, id string) (*[]File, error)
}

func (s Store) ReadFile(ctx context.Context, id string) (*[]File, error) {

	f := File{}
	uid := uuid.MustParse(id)
	//err := s.db.QueryRow(context.Background(), "select name, data, createdBy, createdAt, mimeType from files where id=$1", id).Scan(&f)
	result := s.db.DB.Table("files").Where("id", uid).Find(&f)
	if result.Error != nil {
		return nil, result.Error
	}
	files := make([]File, 0)
	files = append(files, f)
	return &files, nil
}
