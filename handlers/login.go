package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

// LoginData stores the login form data and potential error messages
type LoginData struct {
	ID       string
	Password string
	Error    string
}

// LoginHandler handles all login-related requests
func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		// 폼에서 전송된 데이터 받기
		id := r.FormValue("id")
		password := r.FormValue("password")

		// 데이터베이스에서 사용자 확인
		var dbPassword string
		err := db.QueryRow("SELECT password FROM USER_INFO WHERE id = $1", id).Scan(&dbPassword)

		switch {
		case err == sql.ErrNoRows:
			// 사용자가 존재하지 않는 경우
			data := LoginData{
				ID:    id,
				Error: "아이디가 존재하지 않습니다.",
			}
			renderTemplate(w, "login.html", data)
			return
		case err != nil:
			// 데이터베이스 오류
			log.Printf("Database error: %v", err)
			data := LoginData{
				ID:    id,
				Error: "시스템 오류가 발생했습니다. 잠시 후 다시 시도해주세요.",
			}
			renderTemplate(w, "login.html", data)
			return
		}

		// 비밀번호 검증
		if password == dbPassword { // 실제로는 암호화된 비밀번호를 비교해야 합니다!
			// 로그인 성공
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// 비밀번호가 일치하지 않는 경우
		data := LoginData{
			ID:    id,
			Error: "비밀번호가 올바르지 않습니다.",
		}
		renderTemplate(w, "login.html", data)
	} else {
		// GET 요청 시 빈 로그인 폼 표시
		data := LoginData{}
		renderTemplate(w, "login.html", data)
	}
}

// renderTemplate handles template rendering
func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
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
