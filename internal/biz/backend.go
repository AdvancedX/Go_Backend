package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Backend struct {
	Hello string
}

func NewBackendUsecase(repo BackendRepo, logger log.Logger) *BackendUsecase {
	return &BackendUsecase{repo: repo, log: log.NewHelper(logger)}
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
