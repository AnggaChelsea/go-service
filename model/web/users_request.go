package web

type UsersCreateRequest struct {
	NamaLengkap  string `json:"nama_lengkap"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	IsPerusahaan bool   `json:"isPerusahaan"`
	Password     string `json:"password"`
}
