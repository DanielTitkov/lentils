package entgo

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) GetUserCount(ctx context.Context) (int, error) {
	return r.client.User.Query().Count(ctx)
}

func (r *EntgoRepository) UpdateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	if u.ID == uuid.Nil {
		return nil, errors.New("used id is required")
	}

	usr, err := r.client.User.UpdateOneID(u.ID).
		SetUseDarkTheme(u.UseDarkTheme).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(usr), nil
}

func (r *EntgoRepository) IfEmailRegistered(ctx context.Context, email string) (bool, error) {
	exists, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *EntgoRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *EntgoRepository) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	var email *string
	if u.Email != "" {
		email = &u.Email
	}

	user, err := r.client.User.
		Create().
		SetName(u.Name).
		SetNillableEmail(email).
		SetPasswordHash(u.PasswordHash).
		SetPicture(u.Picture).
		SetLocale(user.Locale(u.Locale)).
		SetAnonymous(u.Anonymous).
		// TODO: not setting admin here
		SetMeta(u.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func entToDomainUser(user *ent.User) *domain.User {
	var email string
	if user.Email != nil {
		email = *user.Email
	}

	return &domain.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        email,
		Picture:      user.Picture,
		PasswordHash: user.PasswordHash,
		Meta:         user.Meta,
		Locale:       user.Locale.String(),
		Admin:        user.Admin,
		UseDarkTheme: user.UseDarkTheme,
		Anonymous:    user.Anonymous,
	}
}
