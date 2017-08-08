package session

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
	"golang-web-dev/042_mongodb/10_hands-on/models"
)

type Session struct {
	Un           string
	LastActivity time.Time
}

const SessionLength int = 30

type SessionManager struct {
	DbSessions map[string]Session // session ID, session
	DbSessionsCleaned time.Time
	DbUsers map[string]models.User       // user ID, user
}

func NewSessionManager() SessionManager {
	sess := SessionManager{}
	sess.DbSessionsCleaned = time.Now()
	sess.DbSessions = map[string]Session{}
	sess.DbUsers = map[string]models.User{}
	return sess
}

func (sess SessionManager) GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = SessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := sess.DbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		sess.DbSessions[c.Value] = s
		u = sess.DbUsers[s.Un]
	}
	return u
}

func (sess SessionManager) AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := sess.DbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		sess.DbSessions[c.Value] = s
	}
	_, ok = sess.DbUsers[s.Un]
	// refresh session
	c.MaxAge = SessionLength
	http.SetCookie(w, c)
	return ok
}

func (sess SessionManager) CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	sess.ShowSessions()              // for demonstration purposes
	for k, v := range sess.DbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(sess.DbSessions, k)
		}
	}
	sess.DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	sess.ShowSessions()             // for demonstration purposes
}

// for demonstration purposes
func (sess SessionManager) ShowSessions() {
	fmt.Println("********")
	for k, v := range sess.DbSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
