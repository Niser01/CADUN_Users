package dtos

//dtos stands for data transfer objects and are used to transfer data between the client and the server

type Create_User struct {
	Names       string `json:"names"  validate:"required" `
	LastNames   string `json:"lastNames" validate:"required"`
	Alias       string `json:"alias" validate:"required"`
	Password    string `json:"password" validate:"required ,min=8"`
	EMail       string `json:"eMail" validate:"required ,min=8"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=10"`
	Country     string `json:"country" validate:"required"`
}

type Get_userid_Byemail struct {
	EMail string `json:"eMail" validate:"required ,min=8"`
}

type Read_userByid struct {
	Id int `json:"id" validate:"required"`
}

type Update_userByid struct {
	Id          int    `json:"id" validate:"required"`
	Names       string `json:"names" validate:"required"`
	LastNames   string `json:"lastNames" validate:"required"`
	Alias       string `json:"alias" validate:"required"`
	Password    string `json:"password" validate:"required ,min=8"`
	EMail       string `json:"eMail" validate:"required ,min=8"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=10"`
	Country     string `json:"country" validate:"required"`
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
