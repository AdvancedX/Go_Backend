package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (rep *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (rep *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, req *v1.UnFollowUserRequest) (rep *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (rep *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (rep *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) GetArticles(ctx context.Context, req *v1.GetArticlesRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) CreateArticles(ctx context.Context, req *v1.CreateArticlesRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) UpdateArticles(ctx context.Context, req *v1.UpdateArticlesRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) DeleteArticles(ctx context.Context, req *v1.DeleteArticlesRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (rep *v1.SingleCommentRelpy, err error) {
	return &v1.SingleCommentRelpy{}, nil
}
func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (rep *v1.MultipleCommentsReply, err error) {
	return &v1.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (rep *v1.SingleCommentRelpy, err error) {
	return &v1.SingleCommentRelpy{}, nil
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *v1.UnFavoriteArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagRequest) (rep *v1.TagListReply, err error) {
	return &v1.TagListReply{}, nil
}
