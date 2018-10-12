package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/eggegg/t-invocing-app/bindings"
	"github.com/eggegg/t-invocing-app/models"
	"github.com/eggegg/t-invocing-app/renderings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	resp := renderings.RegisterResponse{}
	lr := new(bindings.RegisterRequest)

	if err := c.Bind(lr); err != nil {
		resp.Success = false
		resp.Message = "Unable to bind request for register"
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := lr.Validate(c); err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	} // get DB from context
	db := c.Get(models.DBContextKey).(*sql.DB)

	// get user by username from models
	user, err := models.CreateUser(db, lr.Username, lr.Password, lr.Email)
	if err != nil {
		resp.Success = false
		resp.Message = "Create User failure"
		return c.JSON(http.StatusOK, resp)
	}
	
	signingKey := c.Get(models.SigningContextKey).([]byte)


	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	ss, err := token.SignedString(signingKey)
	if err != nil {
		resp.Success = false
		resp.Message = "Server Error"
		return c.JSON(http.StatusOK, resp)
	}

	resp.Success = true
	resp.Token = ss
	resp.User = user

	return c.JSON(http.StatusOK, resp)

}