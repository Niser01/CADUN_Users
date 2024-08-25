package models

// UserProfile db model
type UserProfile struct {
	ID          int    `db:"id"`
	Names       string `db:"names"`
	LastNames   string `db:"lastNames"`
	Alias       string `db:"alias"`
	Password    string `db:"password"`
	EMail       string `db:"eMail"`
	PhoneNumber string `db:"phoneNumber"`
	Country     string `db:"country"`
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

// UsersElementsForQuotation db model
type UsersElementsForQuotation struct {
	ID          int    `db:"id"`
	IDUser      int    `db:"idUser"`
	IAMURL      string `db:"IAM_URL"`
	PDFURL      string `db:"PDF_URL"`
	QuotePDFURL string `db:"QUOTE_PDF_URL"`
}
