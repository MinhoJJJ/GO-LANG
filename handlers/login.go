// handlers/login.go
package handlers

import (
"database/sql"
	"fmt"
	"html/template"
"log"
"net/http"
)

type LoginHandler struct {
	db *sql.DB
}

type LoginData struct {
	ID       string
	Password string
	Error    string
}

// NewLoginHandler 생성자 함수
func NewLoginHandler(db *sql.DB) *LoginHandler {
	return &LoginHandler{db: db}
}

// ServeHTTP implements http.Handler
func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.handlePost(w, r)
	} else {
		h.handleGet(w, r)
	}
}

func (h *LoginHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	password := r.FormValue("password")

	var dbPassword string
	err := h.db.QueryRow("SELECT password FROM USER_INFO WHERE id = $1", id).Scan(&dbPassword)

	switch {
	case err == sql.ErrNoRows:
		h.renderTemplate(w, "login.html", LoginData{
			ID:    id,
			Error: "아이디가 존재하지 않습니다.",
		})
		return
	case err != nil:
		log.Printf("Database error: %v", err)
		h.renderTemplate(w, "login.html", LoginData{
			ID:    id,
			Error: "시스템 오류가 발생했습니다. 잠시 후 다시 시도해주세요.",
		})
		return
	}

	if password == dbPassword {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	h.renderTemplate(w, "login.html", LoginData{
		ID:    id,
		Error: "비밀번호가 올바르지 않습니다.",
	})
}

func (h *LoginHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "login.html", LoginData{})
}

func (h *LoginHandler) renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handlers/index.go
package handlers

import (
"fmt"
"net/http"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}