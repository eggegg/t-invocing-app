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
	"golang.org/x/crypto/bcrypt"
)

// Login - Login Handler will take a username and password from the request
// hash the password, verify it matches in the database and respond with a token
func Login(c echo.Context) error {
	resp := renderings.LoginResponse{}
	lr := new(bindings.LoginRequest)

	if err := c.Bind(lr); err != nil {
		resp.Success = false
		resp.Message = "Unable to bind request for login"
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := lr.Validate(c); err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	} // get DB from context
	db := c.Get(models.DBContextKey).(*sql.DB)
	// get user by username from models
	user, err := models.GetUserByUsername(db, lr.Username)
	if err != nil {
		resp.Success = false
		resp.Message = "Username or Password incorrect"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	if err := bcrypt.CompareHashAndPassword(
		user.PasswordHash, []byte(lr.Password)); err != nil {
		resp.Success = false
		resp.Message = "Username or Password incorrect"
		return c.JSON(http.StatusUnauthorized, resp)
	} // need to make a token, successful login

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
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Success = true
	resp.Token = ss
	resp.User = user

	return c.JSON(http.StatusOK, resp)
}
