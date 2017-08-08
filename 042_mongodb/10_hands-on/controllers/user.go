package controllers

import (
	//"net/http"
	//"golang.org/x/crypto/bcrypt"
	//"github.com/satori/go.uuid"
	//"time"
	"net/http"
	"golang-web-dev/042_mongodb/10_hands-on/session"
	"html/template"
	"github.com/satori/go.uuid"
	"time"
	"golang.org/x/crypto/bcrypt"
	"golang-web-dev/042_mongodb/10_hands-on/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type UserController struct {
	sess session.SessionManager
}

func NewUserController() UserController {
	uc := UserController{}
	uc.sess = session.NewSessionManager()
	return uc
}

func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	u := uc.sess.GetUser(w, req)
	uc.sess.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (uc UserController) Bar(w http.ResponseWriter, req *http.Request) {
	u := uc.sess.GetUser(w, req)
	if !uc.sess.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	uc.sess.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func (uc UserController) Signup(w http.ResponseWriter, req *http.Request) {
	if uc.sess.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := uc.sess.DbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.SessionLength
		http.SetCookie(w, c)
		uc.sess.DbSessions[c.Value] = session.Session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		uc.sess.DbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	uc.sess.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (uc UserController) Login(w http.ResponseWriter, req *http.Request) {
	if uc.sess.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := uc.sess.DbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.SessionLength
		http.SetCookie(w, c)
		uc.sess.DbSessions[c.Value] = session.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	uc.sess.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (uc UserController) Logout(w http.ResponseWriter, req *http.Request) {
	if !uc.sess.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(uc.sess.DbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(uc.sess.DbSessionsCleaned) > (time.Second * 30) {
		go uc.sess.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
