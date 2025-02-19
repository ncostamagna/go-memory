package result

import (
	"github.com/ncostamagna/go-memory/layers/domain"
)

type (
	Service interface {
		Run() error
	}

	service struct {
		repo Repository
	}

	Repository interface {
		Run() (*domain.Shit, error)
	}

	repository struct {}
)

func NewRepository() Repository {
	return &repository{}
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// go:noinline
func (s *service) Run() error {
	_, err := s.repo.Run()
	return err
}

// go:noinline
func (r *repository) Run() (*domain.Shit, error) {
	return &domain.Shit{
		ID: "1",
		Title: "test",
		Content: "test",
	}, nil
}
