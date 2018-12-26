package gateways

import (
	"context"

	"github.com/int128/amefuriso/domain"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const userKind = "User"

func newUserKey(ctx context.Context, id domain.UserID) *datastore.Key {
	return datastore.NewKey(ctx, userKind, string(id), 0, nil)
}

type userEntity struct {
	YahooClientID string
}

type UserRepository struct{}

func (r *UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	q := datastore.NewQuery(userKind)
	var entities []userEntity
	keys, err := q.GetAll(ctx, &entities)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting entities")
	}
	var ret []domain.User
	for i, e := range entities {
		k := keys[i]
		ret = append(ret, domain.User{
			ID:            domain.UserID(k.StringID()),
			YahooClientID: domain.YahooClientID(e.YahooClientID),
		})
	}
	return ret, nil
}

func (r *UserRepository) FindById(ctx context.Context, id domain.UserID) (*domain.User, error) {
	k := newUserKey(ctx, id)
	var e userEntity
	if err := datastore.Get(ctx, k, &e); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, domain.ErrNoSuchUser{ID: id}
		}
		return nil, errors.Wrapf(err, "error while getting entity")
	}
	return &domain.User{
		ID:            id,
		YahooClientID: domain.YahooClientID(e.YahooClientID),
	}, nil
}

func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	if user.ID == "" {
		return errors.Errorf("User.ID must not be empty: %+v", user)
	}
	k := newUserKey(ctx, user.ID)
	e := userEntity{
		YahooClientID: string(user.YahooClientID),
	}
	if _, err := datastore.Put(ctx, k, &e); err != nil {
		return errors.Wrapf(err, "error while saving entity")
	}
	return nil
}
