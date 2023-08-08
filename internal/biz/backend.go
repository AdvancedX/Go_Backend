package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mime/multipart"
	"time"
)

type Backend struct {
	Hello string
}

func NewBackendUsecase(repo BackendRepo, logger log.Logger) *BackendUsecase {
	return &BackendUsecase{repo: repo, log: log.NewHelper(logger)}
}

type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}

type FileLocalRepo interface {
	SaveLocalFile(relativePath string, file *multipart.FileHeader) error
}
type FileRepo interface {
	Save(ctx context.Context, video *File) error
	Exists(ctx context.Context, videoID string) (*File, bool, error)
	Update(ctx context.Context, video *File) error
	ListByType(ctx context.Context, videoType string) ([]*File, error)
	DeleteOne(ctx context.Context, videoID string) error
	ListTagsByType(ctx context.Context, videoType string) ([]string, error)
}
type BackendUsecase struct {
	repo BackendRepo
	log  *log.Helper
}
type BackendRepo interface {
	Save(context.Context, *Backend) (*Backend, error)
	Update(context.Context, *Backend) (*Backend, error)
	FindByID(context.Context, int64) (*Backend, error)
	ListByHello(context.Context, string) ([]*Backend, error)
	ListAll(context.Context) ([]*Backend, error)
}
type File struct {
	ID           string
	Type         string
	Title        string
	Description  string
	Tags         []string
	FilePart     *multipart.FileHeader
	UpdateTime   *time.Time
	RelativePath *string
}
