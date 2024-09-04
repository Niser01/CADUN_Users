package dtos

//dtos stands for data transfer objects and are used to transfer data between the client and the server

type Create_User struct {
	Names        string `json:"names"  validate:"required" `
	LastNames    string `json:"lastNames" validate:"required"`
	Alias        string `json:"alias" validate:"required"`
	Password     string `json:"password" validate:"required ,min=8"`
	EMail        string `json:"eMail" validate:"required ,min=8"`
	PhoneNumber  string `json:"phoneNumber" validate:"required,min=10,max=10"`
	Country      string `json:"country" validate:"required"`
	Home_address string `json:"home_address" validate:"required"`
}

type Get_userid_Byemail struct {
	EMail string `json:"eMail" validate:"required ,min=8"`
}

type Get_password_Byemail struct {
	EMail string `json:"eMail" validate:"required ,min=8"`
}

type RevisarPassword struct {
	EMail    string `json:"eMail" validate:"required ,min=8"`
	Password string `json:"password" validate:"required ,min=8"`
}

type Read_userByid struct {
	Id int `json:"id" validate:"required"`
}

type Update_userByid struct {
	Id           int    `json:"id" validate:"required"`
	Names        string `json:"names" validate:"required"`
	LastNames    string `json:"lastNames" validate:"required"`
	Alias        string `json:"alias" validate:"required"`
	Password     string `json:"password" validate:"required ,min=8"`
	EMail        string `json:"eMail" validate:"required ,min=8"`
	PhoneNumber  string `json:"phoneNumber" validate:"required,min=10,max=10"`
	Country      string `json:"country" validate:"required"`
	Home_address string `json:"home_address" validate:"required"`
}

type Delete_userByid struct {
	Id int `json:"id" validate:"required"`
}

type Get_requeststatus_Byid struct {
	Id int `json:"id" validate:"required"`
}

type Get_requeststatus_ByUser struct {
	Id int `json:"id" validate:"required"`
}

type Update_requeststatus_Byid struct {
	Id            int `json:"id" validate:"required"`
	RequestStatus int `json:"request_status" validate:"required"`
}

type Delete_requests_ByUserid struct {
	Id int `json:"id" validate:"required"`
}

type Create_request struct {
	Id            int    `json:"idUser" validate:"required"`
	RequestStatus int    `json:"request_status" validate:"required"`
	IAM_URL       string `json:"IAM_URL"`
	PDF_URL       string `json:"PDF_URL"`
	QUOTE_PDF_URL string `json:"QUOTE_PDF_URL"`
}

type Create_requesttype struct {
	Status string `json:"Status" validate:"required"`
}

type Create_cotizacion struct {
	ID          int    `json:"id" validate:"required"`
	IDUser      int    `json:"idUser" validate:"required"`
	IDRequest   int    `json:"idRequest" validate:"required"`
	IAMURL      string `json:"IAM_URL" validate:"required"`
	PDFURL      string `json:"PDF_URL" validate:"required"`
	QuotePDFURL string `json:"QUOTE_PDF_URL" validate:"required"`
}

type Delete_cotizacion_ByUserid struct {
	Id int `json:"id" validate:"required"`
}

type Get_cotizacion_ByRequest struct {
	IDRequest int `json:"idRequest" validate:"required"`
}

type Get_cotizacion_data struct {
	Names         string `json:"names"`
	LastNames     string `json:"lastNames"`
	EMail         string `json:"eMail"`
	PhoneNumber   string `json:"phoneNumber"`
	Home_address  string `json:"home_address"`
	IAM_URL       string `json:"IAM_URL"`
	PDF_URL       string `json:"PDF_URL"`
	QUOTE_PDF_URL string `json:"QUOTE_PDF_URL"`
}

type Get_cotizacion struct {
	Id int `json:"id" validate:"required"`
}
