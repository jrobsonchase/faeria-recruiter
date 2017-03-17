package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	db *DB
}

func NewHandler(dbPath string, init bool) (*Handler, error) {
	db, err := NewDB(dbPath)
	if err != nil {
		return nil, err
	}

	if init {
		err = db.Init()
		if err != nil {
			return nil, err
		}
	}

	handler := &Handler{
		db: db,
	}

	return handler, nil
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	mainHeader.Execute(w, nil)
	homeTemplate.Execute(w, nil)
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		mainHeader.Execute(w, nil)
		redirectTemplate.Execute(w, map[string]string{
			"Msg": "Bad Request",
			"Url": "/adduser",
		})
		return
	}

	if user, ok := r.Form["user"]; ok {
		user := user[0]
		resp, err := http.Get(fmt.Sprintf("https://boards.faeria.com/users/%s/activity", user))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mainHeader.Execute(w, nil)
			msg := "Error checking user validity"
			redirectTemplate.Execute(w, map[string]string{
				"Msg": msg,
				"Url": "/adduser",
			})
			return
		}

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusBadRequest)
			mainHeader.Execute(w, nil)
			redirectTemplate.Execute(w, map[string]string{
				"Msg": "Invalid User",
				"Url": "/adduser",
			})
			return
		}

		err = h.db.AddUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mainHeader.Execute(w, nil)
			msg := "Internal Server Error"
			if err == ErrDuplicateUser {
				msg = "Duplicate User"
			}
			redirectTemplate.Execute(w, map[string]string{
				"Msg": msg,
				"Url": "/adduser",
			})
			return
		}

		mainHeader.Execute(w, nil)
		userSuccessTemplate.Execute(w, nil)
		return
	}

	mainHeader.Execute(w, nil)
	addUserTemplate.Execute(w, nil)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.db.RandomUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mainHeader.Execute(w, nil)
		redirectTemplate.Execute(w, map[string]string{
			"Msg": "Internal Server Error",
			"Url": "/",
		})
		return
	}

	mainHeader.Execute(w, nil)
	userTemplate.Execute(w, user)
	return
}
