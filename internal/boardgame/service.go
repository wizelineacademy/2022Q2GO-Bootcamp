package boardgame

type reader interface {
	FindBoardGame(id int) (*BoardGame, error)
}

type repository interface {
	reader
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByID(id int) (*BoardGame, error) {
	return s.repo.FindBoardGame(id)
}
