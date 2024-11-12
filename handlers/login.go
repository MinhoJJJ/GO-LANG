package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler Gin 컨텍스트를 사용하는 로그인 핸들러
// 포인터(pointer)를 사용하여 gin.Context 객체를 함수의 매개변수로 전달
func LoginHandler(c *gin.Context) { // c는 gin.Context 객체를 가리키는 포인터입니다.
	// POST 요청인 경우 (login.do)
	if c.Request.Method == http.MethodPost {
		// 로그인 폼 데이터 받기
		username := c.PostForm("id")
		password := c.PostForm("password")

		log.Println("username: " + username)
		log.Println("password: " + password)

		// 입력값 검증
		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "사용자명과 비밀번호를 모두 입력해주세요.",
			})
			return
		}
		// DB 연결 가져오기
		db := c.MustGet("db").(*sql.DB)

		// 사용자 인증 로직
		var storedPassword string
		var userID string
		err := db.QueryRow("SELECT id,password FROM user_info WHERE id = $1", username).Scan(&userID, &storedPassword)

		// 디버깅을 위한 로그
		log.Printf("Query execution - error: %v", err)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "사용자를 찾을 수 없습니다.",
			})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "서버 오류가 발생했습니다.",
			})
			return
		}

		// 비밀번호 검증 (실제로는 암호화된 비밀번호를 비교해야 함)
		if password == storedPassword { // 실제 구현시에는 bcrypt 등을 사용하여 비교
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "로그인 성공",
				"user_id": userID,
			})
			c.HTML(http.StatusOK, "main.html", gin.H{
				"title": "로그인 페이지",
			})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "비밀번호가 일치하지 않습니다.",
		})
		return
	}

	// GET 요청인 경우 (/)
	// HTML 템플릿 렌더링
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "로그인 페이지",
	})
}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {

// 	db := config.GetDB()

// 	if db == nil {
// 		log.Println("Database connection is nil in LoginHandler")
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		// 폼에서 전송된 데이터 받기
// 		id := r.FormValue("id")
// 		password := r.FormValue("password")

// 		// 데이터베이스에서 사용자 확인
// 		var dbPassword string
// 		err := db.QueryRow("SELECT password FROM USER_INFO WHERE id = $1", id).Scan(&dbPassword)

// 		switch {
// 		case err == sql.ErrNoRows:
// 			// 사용자가 존재하지 않는 경우
// 			data := models.LoginData{
// 				ID:    id,
// 				Error: "아이디가 존재하지 않습니다.",
// 			}
// 			RenderTemplate(w, "login.html", data)
// 			return
// 		case err != nil:
// 			// 데이터베이스 오류
// 			log.Printf("Database error: %v", err)
// 			data := models.LoginData{
// 				ID:    id,
// 				Error: "시스템 오류가 발생했습니다. 잠시 후 다시 시도해주세요.",
// 			}
// 			RenderTemplate(w, "login.html", data)
// 			return
// 		}

// 		// 비밀번호 검증
// 		if password == dbPassword { // 실제로는 암호화된 비밀번호를 비교해야 합니다!
// 			// 로그인 성공
// 			http.Redirect(w, r, "/login", http.StatusSeeOther)
// 			return
// 		}

// 		// 비밀번호가 일치하지 않는 경우
// 		data := models.LoginData{ // 구조체 사용
// 			ID:    id,
// 			Error: "비밀번호가 올바르지 않습니다.",
// 		}
// 		RenderTemplate(w, "login.html", data)
// 	} else {
// 		// GET 요청 시 빈 로그인 폼 표시
// 		data := models.LoginData{} // 구조체 사용
// 		RenderTemplate(w, "login.html", data)
// 	}
// }
// config 패키지를 통해 DB 연결 가져오기
// db := config.GetDB()
// if db == nil {
// 	log.Println("Database connection is nil in LoginHandler")
// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		"status":  "error",
// 		"message": "데이터베이스 연결 오류",
// 	})
// 	return
// }
