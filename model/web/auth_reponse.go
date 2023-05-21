package web

type AuthResponse struct {
	NamaLengkap string `json:"nama_lengkap"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
}
