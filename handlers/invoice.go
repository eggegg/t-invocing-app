package handlers

import (
	_"strconv"
	_"fmt"
	"database/sql"
	"reflect"
	"net/http"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"

	"github.com/eggegg/t-invocing-app/models"
	"github.com/eggegg/t-invocing-app/renderings"

)

// Get user info from jwt token
func HelloWorld(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	id := int(claims["id"].(float64))

	c.Logger().Info(claims["id"])
	c.Logger().Print(reflect.TypeOf(username), username)
	c.Logger().Print(reflect.TypeOf(id), id)

	return c.String(http.StatusOK, "Welcome " + username )
}

// create
func CreateInvoice(c echo.Context) error {
	resp := renderings.CommonResponse{}

	// get user id from jwt token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))


	name := c.FormValue("name")
	paid := 0

	//get DB from context
	db := c.Get(models.DBContextKey).(*sql.DB)

	invoice, err := models.CreateInvoice(db, name, userId, paid)
	if err != nil {
		resp.Success = false
		resp.Message = "Create invoice failure"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	resp.Success = true
	resp.Message = "Create invoice success : " + invoice.Name

	return c.JSON(http.StatusOK, resp)
}


func GetUserInvoice(c echo.Context) error  {
	resp := renderings.ResultResponse{}

	userId := c.Param("user_id")

	//get DB from context
	db := c.Get(models.DBContextKey).(*sql.DB)

	invoices, err := models.GetUserInvoice(db, userId)
	if err != nil {
		resp.Success = false
		resp.Message = "Get user invoice failure"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	resp.Success = true
	resp.Message = "Get invoice success " 
	resp.Data = invoices

	return c.JSON(http.StatusOK, resp)
}

func GetOneInvoice(c echo.Context) error {
	resp := renderings.ResultResponse{}

	userId := c.Param("user_id")
	invoiceId := c.Param("invoce_id")


	//get DB from context
	db := c.Get(models.DBContextKey).(*sql.DB)

	invoice, err := models.GetOneInvoice(db, userId, invoiceId)
	if err != nil {
		resp.Success = false
		resp.Message = "Get user invoice failure"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	resp.Success = true
	resp.Message = "Get invoice success " 
	resp.Data = invoice

	return c.JSON(http.StatusOK, resp)
}

func SendInvoice(c echo.Context) error {
	resp := renderings.ResultResponse{}
	resp.Success = true
	resp.Message = "Get invoice success " 

	return c.JSON(http.StatusOK, resp)
}