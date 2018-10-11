package bindings

import "github.com/labstack/echo"

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func (lr RegisterRequest) Validate(c echo.Context) error {
	errs := new(RequestErrors)
	if lr.Username == "" {
		errs.Append(ErrUsernameEmpty)
	}
	if lr.Password == "" {
		errs.Append(ErrPasswordEmpty)
	}
	if lr.Email == "" {
		errs.Append(ErrEmailEmpty)
	}
	if errs.Len() == 0 {
		return nil
	}
	return errs
}
