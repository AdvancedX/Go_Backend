package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface{}
type CommentRepo interface{}
type TagRepo interface{}

type SocialUsecase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo

	log *log.Helper
}

func NewSocialUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}
func (uc *UserUsecase) CreateArticles(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {

	}
	return nil
}
