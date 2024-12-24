package user

import (
	"context"

	"github.com/morheus9/go_rest/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) {

}
