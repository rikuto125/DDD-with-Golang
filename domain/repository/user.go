package repository

import "example.com/m/domain/model"

type UserRepository interface {
	Search(name string) ([]*model.User, error)
}
