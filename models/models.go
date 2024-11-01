package models

type LoginData struct {
	ID       string // 사용자 ID
	Password string // 사용자 비밀번호 (보안상 실제 비밀번호는 저장하지 않는 것이 좋습니다)
	Error    string // 에러 메시지
}

type DBConfig struct {
	Host     string // 호스트
	Port     int    // 포트
	User     string // 유저명
	Password string // 패스워드
	DBName   string // 디비명
}
