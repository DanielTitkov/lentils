package entgo

import (
	"context"
	"errors"
	"time"

	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/usersession"

	"github.com/google/uuid"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
)

// CreateOrUpdateUserSession must be used only with authentication like oauth
// because it updated ip and user agent and in other uses may be vulnerable
// to session steal
func (r *EntgoRepository) CreateOrUpdateUserSession(ctx context.Context, s *domain.UserSession) (*domain.UserSession, error) {
	if s == nil {
		return nil, nil
	}

	if s.UserID == uuid.Nil {
		return nil, errors.New("user id required to create or update user session")
	}

	ses, err := r.client.UserSession.
		Query().
		Where(
			usersession.And(
				usersession.SidEQ(s.SID),
				// FIXME
				// usersession.HasUserWith(
				// 	user.IDEQ(s.UserID),
				// ),
			),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create session
		ses, err := r.client.UserSession.
			Create().
			SetSid(s.SID).
			SetUserID(s.UserID).
			SetIP(s.IP).
			SetUserAgent(s.UserAgent).
			SetMeta(s.Meta).
			SetActive(s.Active).
			SetLastActivity(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return entToDomainUserSession(ses), nil
	}

	// update session
	ses, err = ses.Update().
		SetIP(s.IP).
		SetUserAgent(s.UserAgent).
		SetLastActivity(time.Now()).
		SetMeta(s.Meta).
		SetActive(s.Active).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUserSession(ses), nil
}

func (r *EntgoRepository) IfSessionRegistered(ctx context.Context, s *domain.UserSession) (bool, error) {
	exists, err := r.client.UserSession.
		Query().
		Where(
			usersession.And(
				usersession.Active(true),
				usersession.SidEQ(s.SID),
				// usersession.IPEQ(s.IP), // FIXME for commented out for debug
				// usersession.UserAgentEQ(s.UserAgent), // FIXME for commented out for debug
			),
		).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *EntgoRepository) UpdateUserSessionLastActivityBySID(ctx context.Context, sid string) error {
	ses, err := r.client.UserSession.Query().Where(usersession.SidEQ(sid)).Only(ctx)
	if err != nil {
		return err
	}

	_, err = ses.Update().SetLastActivity(time.Now()).Save(ctx)
	return err
}

func (r *EntgoRepository) GetUserBySession(ctx context.Context, s *domain.UserSession) (*domain.User, error) {
	user, err := r.client.UserSession.
		Query().
		Where(
			usersession.And(
				usersession.Active(true),
				usersession.SidEQ(s.SID),
				// usersession.IPEQ(s.IP), // FIXME commented out for debug
				// usersession.UserAgentEQ(s.UserAgent), // FIXME commented out for debug
			),
		).
		QueryUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func entToDomainUserSession(s *ent.UserSession) *domain.UserSession {
	var userID uuid.UUID
	if s.Edges.User != nil {
		userID = s.Edges.User.ID
	}

	return &domain.UserSession{
		ID:           s.ID,
		UserID:       userID,
		SID:          s.Sid,
		IP:           s.IP,
		Active:       s.Active,
		UserAgent:    s.UserAgent,
		CreateTime:   s.CreateTime,
		UpdateTime:   s.UpdateTime,
		LastActivity: s.LastActivity,
		Meta:         s.Meta,
	}
}
