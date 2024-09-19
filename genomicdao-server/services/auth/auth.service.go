package auth

import "github.com/quytm/blockchain-engineer-interview/genomicdao-server/pkg/utils"

type (
	IService interface {
		Register([]byte) string
	}

	serviceImpl struct {
		users map[string]User // In-memory DB, instead of `users` table in DB
	}
)

func NewAuthService() IService {
	return &serviceImpl{
		users: make(map[string]User),
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *serviceImpl) Register(publicKey []byte) string {
	// 1. Generate User ID
	userID := utils.RandString(10)

	// 2. Save to DB
	s.users[userID] = User{
		ID:        userID,
		PublicKey: publicKey,
	}

	// 3. Return
	return userID
}
