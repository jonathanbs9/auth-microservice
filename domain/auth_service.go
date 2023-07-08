package domain


type authService struct {
	storage Storage
	hasher  Hasher
}

func NewAuthService(s Storage, h Hasher) AuthService {
	return &authService(storage: s, hasher: h)
}


