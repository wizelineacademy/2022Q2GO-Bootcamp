package boardgame

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByID(id int) (*BoardGame, error) {
	return s.repo.FindBoardGame(id)
}
