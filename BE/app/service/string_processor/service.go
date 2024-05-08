package string_processor

import "context"

type IService interface {
	ProcessString(ctx context.Context, request StringProcessRequest) (StringProcessResponse, error)
}
type service struct {
}

var _ IService = &service{}

func NewService() IService {
	return &service{}
}

func (s *service) ProcessString(ctx context.Context, request StringProcessRequest) (StringProcessResponse, error) {
	return StringProcessResponse{request.Value + request.Value}, nil
}
