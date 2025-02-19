package param

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
		Run(shit *domain.Shit) error
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
	shit := &domain.Shit{
		ID: "1",
	}
	v := s.repo.Run(shit)
	//fmt.Println(shit)
	return v
}

// go:noinline
func (r *repository) Run(shit *domain.Shit) error {
	shit.Title = "test"
	shit.Content = "test"
	return nil
}
