package views

import (
	"context"

	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/internal/models"
	"github.com/jmoiron/sqlx"
)

// views interface has all the crud methods for the user and saved elements
type View_interface interface {
	get_userid_Byemail(ctx context.Context, eMail string) (*models.UserId, error)
	Create_user(ctx context.Context, names string, lastNames string, alias string, password string, eMail string, phoneNumber string, country string) error
	Read_userByid(ctx context.Context, id int) (*models.UserProfile, error)
	Update_userByid(ctx context.Context, names string, lastNames string, alias string, password string, eMail string, phoneNumber string, country string, id int) error
	Delete_userByid(ctx context.Context, id int) error
	get_requeststatus_Byid(ctx context.Context, id int) (*models.Request_Status, error)
	get_requeststatus_ByUser(ctx context.Context, idUser int) (*models.Request_Status, error)
	update_requeststatus_Byid(ctx context.Context, status int, id int) error
	delete_requests_ByUserid(ctx context.Context, idUser int) error
	create_requesttype(ctx context.Context, Status string) error
	create_cotizacion(ctx context.Context, idUser int, idRequest int, IAM_URL string, PDF_URL string, QUOTE_PDF_URL string) error
	delete_cotizacion_ByUserid(ctx context.Context, idUser int) error
	get_cotizacion_ByRequest(ctx context.Context, idRequest int) (*models.UsersElementsForQuotation, error)
}

// view is the implementation of the views interface
type View_struct struct {
	db *sqlx.DB
}

// New returns the implementation of the views interface
func New(db *sqlx.DB) View_interface {
	return &View_struct{db: db}
}
