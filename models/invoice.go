package models

import (
	"strconv"
	"fmt"
	"database/sql"

	"github.com/pkg/errors"

)


type Invoice struct {
	ID           int
	Name     	string
	UserId        int
	Paid		int
}


func CreateInvoice(db *sql.DB, name string, userId int, paid int) (Invoice, error) {
	var invoice = Invoice{}
	
	stmt, err := db.Prepare("INSERT invoices SET name=?, user_id=?, paid=?")

	if err != nil {
		return invoice, errors.Wrap(err, "prepare db error:")
	}

	res, err := stmt.Exec(name, userId, paid)
	if err != nil {
		return invoice, errors.Wrap(err, "insert db error:")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return invoice, errors.Wrap(err, "last insert id error:")
	}

	invoice.ID = int(id)
	invoice.Name = name
	invoice.UserId = userId
	invoice.Paid = paid

	return invoice, nil

}


func GetUserInvoice(db *sql.DB, userId string) ([]*Invoice, error) {
	
	invoices := make([]*Invoice, 0)

	id, _ := strconv.Atoi(userId)

	// query
	rows, err := db.Query( fmt.Sprintf("SELECT id, name, user_id, paid FROM invoices WHERE user_id=%d", id) )
	if err != nil {
		return invoices, errors.Wrap(err, "select db error:")
	}

	for rows.Next() {
		in := &Invoice{}
		rows.Scan(&in.ID, &in.Name, &in.UserId, &in.Paid)
		invoices = append(invoices, in)
	}

	return invoices, nil
}


func GetOneInvoice(db *sql.DB, userId string, invoiceId string) (Invoice, error) {

	var invoice = Invoice{}

	uId, _ := strconv.Atoi(userId)
	iId, _ := strconv.Atoi(invoiceId)

	// query
	row := db.QueryRow( fmt.Sprintf("SELECT id, name, user_id, paid FROM invoices WHERE user_id=%d AND id=%d", uId, iId) )
	err := row.Scan(&invoice.ID, &invoice.Name, &invoice.UserId, &invoice.Paid)
	if err != nil && err != sql.ErrNoRows {
		return invoice, errors.Wrap(err, "query row fail:")
	}

	return invoice, nil
}