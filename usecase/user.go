package usecase

import (
	"example.com/m/domain/model"
	"example.com/m/domain/repository"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	Search(name string) ([]*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// Userデータに関するUseCaseを生成
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

// 検索
func (uu userUseCase) Search(name string) (user []*model.User, err error) {
	user, err = uu.userRepository.Search(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
