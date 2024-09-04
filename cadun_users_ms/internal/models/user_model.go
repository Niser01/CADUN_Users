package models

// UserProfile db model
type UserProfile struct {
	ID           int    `db:"id"`
	Names        string `db:"names"`
	LastNames    string `db:"lastNames"`
	Alias        string `db:"alias"`
	Password     string `db:"password"`
	EMail        string `db:"eMail"`
	PhoneNumber  string `db:"phoneNumber"`
	Country      string `db:"country"`
	Home_address string `db:"home_address"`
}

type Password struct {
	Password string `db:"password"`
}

type UserId struct {
	ID int `db:"id"`
}

// RequestType db model
type RequestType struct {
	ID     int    `db:"id"`
	Status string `db:"status"`
}

// Request db model
type Request struct {
	ID            int `db:"id"`
	IDUser        int `db:"idUser"`
	RequestStatus int `db:"request_status"`
}

type Request_Status struct {
	RequestStatus int `db:"request_status"`
}

type Status struct {
	Status string `db:"Status"`
}

type Get_cotizacion_data struct {
	Names         string `db:"names"`
	LastNames     string `db:"lastNames"`
	EMail         string `db:"eMail"`
	PhoneNumber   string `db:"phoneNumber"`
	Home_address  string `db:"home_address"`
	IAM_URL       string `db:"IAM_URL"`
	PDF_URL       string `db:"PDF_URL"`
	QUOTE_PDF_URL string `db:"QUOTE_PDF_URL"`
}
