package views

import (
	"context"
	"errors"

	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/models"
)

// querys for the database
const (
	queryCreate_User = `
	INSERT INTO USERS_PROFILE (names, lastNames, alias, password, eMail, phoneNumber, country) 
	VALUES (?, ?, ?, SHA2(?, 256), ?, ?, ?)`

	queryread_user_Byid = `
	SELECT *
	FROM USERS_PROFILE 
	WHERE id = ?`

	queryget_userid_Byemail = `
	SELECT id
	from USERS_PROFILE 
	WHERE eMail = ?`

	queryget_password_Byemail = `
	SELECT password
	from USERS_PROFILE 
	WHERE eMail = ?`

	queryupdate_user_Byid = `
	UPDATE USERS_PROFILE 
	SET names = ?, lastNames = ?, alias = ?, password = ?, eMail = ?, phoneNumber = ?, country = ?
	WHERE id = ?`

	querydelete_user_Byid = `
	DELETE FROM USERS_PROFILE 
	WHERE id = ?`

	queryget_requeststatus_Byid = `
	SELECT request_status
	from REQUEST 
	WHERE id = ?`

	queryget_status_byid = `
	SELECT Status
	FROM REQUEST_TYPES
	WHERE id = ?`

	queryget_requeststatus_ByUser = `
	SELECT request_status
	from REQUEST 
	WHERE idUser = ?`

	queryupdate_requeststatus_Byid = `
	UPDATE REQUEST 
	SET request_status = ?
	WHERE id = ?`

	querydelete_requests_ByUserid = `
	DELETE FROM REQUEST 
	WHERE idUser = ?`

	querycreate_requesttype = `
	INSERT INTO REQUEST_TYPES (Status)
	VALUES (?)`

	querycreate_request = `
	INSERT INTO REQUEST (idUser, request_status)
	VALUES (?, ?)`

	querycreate_cotizacion = `
	INSERT INTO USERS_PROFILE (idUser, idRequest, IAM_URL, PDF_URL, QUOTE_PDF_URL) 
	VALUES (?, ?, ?, ?, ?)`

	querydelete_cotizacion_ByUserid = `
	DELETE FROM USERS_ELEMENTS_FOR_QUOTATION 
	WHERE idUser = ?`

	queryget_cotizacion_ByRequest = `
	SELECT IAM_URL, PDF_URL, QUOTE_PDF_URL
	FORM USERS_ELEMENTS_FOR_QUOTATION
	WHERE idRequest = ?`
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

// Reads user from DB, since there can only be a unique email, the function is used first when creating a user it gets the system conext and gets the email
func (r *View_struct) Get_userid_Byemail(ctx context.Context, eMail string) (*models.UserId, error) {
	u := &models.UserId{}
	err := r.db.GetContext(ctx, u, queryget_userid_Byemail, eMail)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// create_user creates a new user in the database, it uses the ExcecContext method
func (r *View_struct) Create_user(ctx context.Context, names string, lastNames string, alias string, password string, eMail string, phoneNumber string, country string) error {
	u, _ := r.Get_userid_Byemail(ctx, eMail)

	if u != nil {
		return ErrUserAlreadyExists
	}
	_, err := r.db.ExecContext(ctx, queryCreate_User, names, lastNames, alias, password, eMail, phoneNumber, country)
	if err != nil {
		return err
	}
	return nil
}

// Reads the user info by their id
func (r *View_struct) Read_userByid(ctx context.Context, id int) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	err := r.db.GetContext(ctx, u, queryread_user_Byid, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *View_struct) Get_password_Byemail(ctx context.Context, eMail string) (*models.Password, error) {
	u := &models.Password{}
	err := r.db.GetContext(ctx, u, queryget_password_Byemail, eMail)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *View_struct) Get_status_byid(ctx context.Context, id int) (*models.Status, error) {
	u := &models.Status{}
	err := r.db.GetContext(ctx, u, queryget_status_byid, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// This function updates the user information, the user is selected by it´s id,  it uses the ExcecContext method
func (r *View_struct) Update_userByid(ctx context.Context, names string, lastNames string, alias string, password string, eMail string, phoneNumber string, country string, id int) error {
	_, err := r.db.ExecContext(ctx, queryupdate_user_Byid, names, lastNames, alias, password, eMail, phoneNumber, country, id)
	if err != nil {
		return err
	}
	return nil
}

// This funcition delets a user it uses the ExcecContext method
func (r *View_struct) Delete_userByid(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, querydelete_user_Byid, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *View_struct) Get_requeststatus_Byid(ctx context.Context, id int) (*models.Request_Status, error) {
	u := &models.Request_Status{}
	err := r.db.GetContext(ctx, u, queryget_requeststatus_Byid, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *View_struct) Update_requeststatus_Byid(ctx context.Context, status int, id int) error {
	_, err := r.db.ExecContext(ctx, queryupdate_requeststatus_Byid, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *View_struct) Get_requeststatus_ByUser(ctx context.Context, idUser int) (*models.Request_Status, error) {
	u := &models.Request_Status{}
	err := r.db.GetContext(ctx, u, queryget_requeststatus_ByUser, idUser)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// This function edits the status of a user  it uses the ExcecContext method
func (r *View_struct) Create_request(ctx context.Context, idUser int, request_status int) error {
	_, err := r.db.ExecContext(ctx, querycreate_request, idUser, request_status)
	if err != nil {
		return err
	}
	return nil
}

func (r *View_struct) Delete_requests_ByUserid(ctx context.Context, idUser int) error {
	_, err := r.db.ExecContext(ctx, querydelete_requests_ByUserid, idUser)
	if err != nil {
		return err
	}
	return nil
}

func (r *View_struct) Create_requesttype(ctx context.Context, Status string) error {
	_, err := r.db.ExecContext(ctx, querycreate_requesttype, Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *View_struct) Create_cotizacion(ctx context.Context, idUser int, idRequest int, IAM_URL string, PDF_URL string, QUOTE_PDF_URL string) error {
	_, err := r.db.ExecContext(ctx, querycreate_cotizacion, idUser, idRequest, IAM_URL, PDF_URL, QUOTE_PDF_URL)
	if err != nil {
		return err
	}
	return nil
}

func (r *View_struct) Delete_cotizacion_ByUserid(ctx context.Context, idUser int) error {
	_, err := r.db.ExecContext(ctx, querydelete_cotizacion_ByUserid, idUser)
	if err != nil {
		return err
	}
	return nil
}

func (r *View_struct) Get_cotizacion_ByRequest(ctx context.Context, idRequest int) (*models.UsersElementsForQuotation, error) {
	u := &models.UsersElementsForQuotation{}
	err := r.db.GetContext(ctx, u, queryget_cotizacion_ByRequest, idRequest)
	if err != nil {
		return nil, err
	}
	return u, nil
}
