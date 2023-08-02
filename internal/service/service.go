package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewRealWorldService(uc *biz.UserUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
