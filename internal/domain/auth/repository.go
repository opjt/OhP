package auth

type AuthRepository interface {
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return authRepository{}
}
