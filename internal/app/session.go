package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DanielTitkov/orrery/internal/configs"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/jfyne/live"
)

func (a *App) ResetSession(res http.ResponseWriter, req *http.Request) error {
	// all this crap doesn't work

	http.SetCookie(res, &http.Cookie{
		Name:     configs.LiveSessionName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	sess := live.NewSession()
	session, err := a.Store.Get(req, configs.LiveSessionName)
	if err != nil {
		return err
	}

	session.Values["_ls"] = sess
	err = session.Save(req, res)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) CreateOrUpdateUserSession(req *http.Request, user *domain.User, setActiveStatus bool) (*domain.UserSession, error) {
	fmt.Println("APP CREATE UPDATE SESSION") // FIXME
	// get session sid for request
	sid, err := a.LiveSessionID(req)
	if err != nil {
		return nil, nil
	}

	session := &domain.UserSession{
		SID:       sid,
		UserAgent: req.UserAgent(),
		IP:        req.RemoteAddr,
		UserID:    user.ID,
		Active:    setActiveStatus,
	}

	// create session record for user
	session, err = a.repo.CreateOrUpdateUserSession(req.Context(), session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (a *App) CreateOrUpdateUserSessionWithSID(req *http.Request, sid string, user *domain.User, setActiveStatus bool) (*domain.UserSession, error) {
	session := &domain.UserSession{
		SID:       sid,
		UserAgent: req.UserAgent(),
		IP:        req.RemoteAddr,
		UserID:    user.ID,
		Active:    setActiveStatus,
	}

	// create session record for user
	session, err := a.repo.CreateOrUpdateUserSession(req.Context(), session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (a *App) GetUserBySession(req *http.Request) (*domain.User, error) {
	// get session sid for request
	sid, err := a.LiveSessionID(req)
	if err != nil || sid == "" {
		return nil, nil
	}

	session := &domain.UserSession{
		SID:       sid,
		UserAgent: req.UserAgent(),
		IP:        req.RemoteAddr,
	}

	// check if session saved for some user
	registered, err := a.repo.IfSessionRegistered(req.Context(), session)
	if err != nil {
		return nil, err
	}

	var user *domain.User
	if !registered {
		// session is not registested
		// if user is not registered we need to create anonymous user
		user, err = a.CreateAnonymousUser(req.Context())
		if err != nil {
			return nil, err
		}

		// add or update session for user
		session, err = a.CreateOrUpdateUserSession(req, user, true)
		if err != nil {
			a.log.Error("failed to create or update user session (GetUserBySession)", err)
			return nil, err
		}
		if session == nil {
			return nil, nil
		}

		a.log.Debug("user session refreshed", fmt.Sprintf("email: %s, sid: %s", user.Email, session.SID))
	}
	// retrieve user and add to context
	user, err = a.repo.GetUserBySession(req.Context(), session)
	if err != nil {
		return nil, err
	}

	// update session activity
	err = a.repo.UpdateUserSessionLastActivityBySID(req.Context(), sid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *App) GetUserByLiveSession(req *http.Request, ls live.Session) (*domain.User, error) {
	fmt.Println("LIVE SESSION", ls) // FIXME
	// get session sid for request
	sid, err := a.LiveSessionIDFromSession(ls)
	if err != nil || sid == "" {
		fmt.Println("NIL SESSION (GET)")
		return nil, nil
	}

	fmt.Println("GetUserByLiveSession", sid) // FIXME

	session := &domain.UserSession{
		SID:       sid,
		UserAgent: req.UserAgent(),
		IP:        req.RemoteAddr,
	}

	// check if session saved for some user
	registered, err := a.repo.IfSessionRegistered(req.Context(), session)
	if err != nil {
		return nil, err
	}

	var user *domain.User
	if !registered {
		fmt.Println("SESSION IS NOT REGISTERED") // FIXME
		// session is not registested
		// if user is not registered we need to create anonymous user
		user, err = a.CreateAnonymousUser(req.Context())
		if err != nil {
			fmt.Println("FAILE ANON") // FIXME
			return nil, err
		}

		// add or update session for user
		session, err = a.CreateOrUpdateUserSessionWithSID(req, sid, user, true)
		if err != nil {
			a.log.Error("failed to create or update user session (GetUserBySession)", err)
			fmt.Println("FAILED CREATE SESSION", err) // FIXME
			return nil, err
		}
		if session == nil {
			fmt.Println("FAILED CREATE (NIL SESSION)", session, err) // FIXME
			return nil, nil
		}

		a.log.Debug("user session refreshed", fmt.Sprintf("email: %s, sid: %s", user.Email, session.SID))
	}
	// retrieve user and add to context
	fmt.Println("SESSION REGISTERED") // FIXME
	user, err = a.repo.GetUserBySession(req.Context(), session)
	if err != nil {
		return nil, err
	}

	// update session activity
	err = a.repo.UpdateUserSessionLastActivityBySID(req.Context(), sid)
	if err != nil {
		return nil, err
	}

	fmt.Println("RETURNING USER", user) // FIXME

	return user, nil
}

func (a *App) LiveSessionID(req *http.Request) (string, error) {
	ses, err := a.Store.Get(req, "go-live-session")
	if err != nil {
		return "", err
	}

	lsI := ses.Values["_ls"]
	ls, ok := lsI.(live.Session)
	if !ok {
		return "", fmt.Errorf("expected to get live.Session but got %T", lsI)
	}

	idI := ls["_lsid"]
	id, ok := idI.(string)
	if !ok {
		return "", fmt.Errorf("expected to get string but got %T", idI)
	}
	return id, nil
}

func (a *App) LiveSessionIDFromSession(ls live.Session) (string, error) {
	idI := ls["_lsid"]
	id, ok := idI.(string)
	if !ok {
		return "", fmt.Errorf("expected to get string but got %T", idI)
	}
	return id, nil
}
