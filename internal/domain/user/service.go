package user

type UserService struct {
	repo userRepository
}

func NeUserService(repo userRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
