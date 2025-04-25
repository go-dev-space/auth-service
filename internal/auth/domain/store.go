package domain

import "context"

type Store interface {
	Save(context.Context, *User) error
}
