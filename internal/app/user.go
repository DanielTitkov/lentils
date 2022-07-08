package app

import (
	"context"
	"errors"

	"github.com/markbates/goth"
	"github.com/sethvargo/go-password/password"

	"github.com/DanielTitkov/orrery/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	if u.Password == "" {
		// TODO: add password strength checks
		return nil, errors.New("user password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	u.PasswordHash = string(hash)

	if !domain.IsValidLocale(u.Locale) {
		// a.log.Warn(fmt.Sprintf("got unknown locale %s, setting to default %s", u.Locale, domain.LocaleEn), "unknown locale") // FIXME
		u.Locale = domain.LocaleEn
	}

	user, err := a.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *App) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return a.repo.UpdateUser(ctx, user)
}

// AuthenticateGothUser creates new user or returns existsing one.
// It relies on goth authetication to verify user has access
// to that profile and thus doesn't check password.
func (a *App) AuthenticateGothUser(ctx context.Context, gu *goth.User) (*domain.User, error) {
	exists, err := a.repo.IfEmailRegistered(ctx, gu.Email)
	if err != nil {
		return nil, err
	}

	if !exists {
		return a.CreateUserFromGoth(ctx, gu)
	}

	// TODO if user came from another provider add new data to meta
	return a.GetUserByEmail(ctx, gu.Email)
}

func (a *App) LogoutUser(ctx context.Context) error {
	return nil
}

func (a *App) CreateAnonymousUser(ctx context.Context) (*domain.User, error) {
	passw, err := password.Generate(16, 5, 0, false, true)
	if err != nil {
		return nil, err
	}

	meta := make(map[string]interface{})
	user := &domain.User{
		Name:      "anonymous",
		Email:     "",
		Password:  passw,
		Admin:     false,
		Anonymous: true,
		Meta:      meta,
	}

	user, err = a.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// user.Password = passw

	return user, nil
}

func (a *App) CreateUserFromGoth(ctx context.Context, gu *goth.User) (*domain.User, error) {
	passw, err := password.Generate(16, 5, 0, false, true)
	if err != nil {
		return nil, err
	}

	meta := make(map[string]interface{})
	meta[gu.Provider] = *gu

	user := &domain.User{
		Name:      gu.NickName,
		Email:     gu.Email,
		Picture:   gu.AvatarURL,
		Password:  passw,
		Admin:     false,
		Anonymous: false,
		Meta:      meta,
	}

	user, err = a.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Password = passw // TODO: this is to show or send password to the user

	return user, nil
}

func (a *App) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := a.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// we should not return password hash if this is not needed
	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Picture:   user.Picture,
		Admin:     user.Admin,
		Anonymous: user.Anonymous,
		Meta:      user.Meta,
	}, nil
}

func (a *App) ValidateUserPassword(ctx context.Context, u *domain.User) (bool, error) {
	user, err := a.repo.GetUserByEmail(ctx, u.Email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(u.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
