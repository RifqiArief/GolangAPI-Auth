package data

type Register struct {
	Nama  string `json="nama"`
	Email string `json="email"`
	Token string `json="token"`
}

type Login struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Kota   string `json:"kota"`
	NoTelp string `json:"no_telp"`
	Email  string `json:"email"`
	Image  string `json:"image"`
	// Status   bool   `json:"status"`
	Token string `json:"token`
}
