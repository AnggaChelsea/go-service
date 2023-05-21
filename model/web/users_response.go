package web

type UsersResponse struct {
	Id           int    `json:"id"`
	NamaLengkap  string `json:"nama_lengkap"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	IsPerusahaan bool   `json:"isPerusahaan"`
	Password     string `json:"password"`
}
