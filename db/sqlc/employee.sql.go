// Code generated by sqlc. DO NOT EDIT.
// source: employee.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addAdmin = `-- name: AddAdmin :execresult
INSERT INTO
  employees (
	type,
	mail,
	social_security_number,
	standard_tax_deductions,
	other_deductions,
	phone_number,
	salary_rate,
  root
  )
VALUES
  ('salaried', ? ,'', 0.02, 1, ? , 1 , 1)
`

type AddAdminParams struct {
	Mail        string `json:"mail"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) AddAdmin(ctx context.Context, arg AddAdminParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addAdmin, arg.Mail, arg.PhoneNumber)
}

const addEmployee = `-- name: AddEmployee :execresult
INSERT INTO
  employees (
    type,
    mail,
    social_security_number,
    standard_tax_deductions,
    other_deductions,
    phone_number,
    salary_rate
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?)
`

type AddEmployeeParams struct {
	Type                  EmployeesType `json:"type"`
	Mail                  string        `json:"mail"`
	SocialSecurityNumber  string        `json:"social_security_number"`
	StandardTaxDeductions string        `json:"standard_tax_deductions"`
	OtherDeductions       string        `json:"other_deductions"`
	PhoneNumber           string        `json:"phone_number"`
	SalaryRate            string        `json:"salary_rate"`
}

func (q *Queries) AddEmployee(ctx context.Context, arg AddEmployeeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addEmployee,
		arg.Type,
		arg.Mail,
		arg.SocialSecurityNumber,
		arg.StandardTaxDeductions,
		arg.OtherDeductions,
		arg.PhoneNumber,
		arg.SalaryRate,
	)
}

const addOrderInfo = `-- name: AddOrderInfo :exec
INSERT INTO order_info(order_id,product_id,amount) 
  VALUES(?,?,?)
`

type AddOrderInfoParams struct {
	OrderID   int64  `json:"order_id"`
	ProductID int64  `json:"product_id"`
	Amount    string `json:"amount"`
}

func (q *Queries) AddOrderInfo(ctx context.Context, arg AddOrderInfoParams) error {
	_, err := q.db.ExecContext(ctx, addOrderInfo, arg.OrderID, arg.ProductID, arg.Amount)
	return err
}

const addPurchaseOrder = `-- name: AddPurchaseOrder :execresult
INSERT INTO purchase_order(emp_id,customer_contact,customer_address,date)
VALUES(?,?,?,?)
`

type AddPurchaseOrderParams struct {
	EmpID           int64     `json:"emp_id"`
	CustomerContact string    `json:"customer_contact"`
	CustomerAddress string    `json:"customer_address"`
	Date            time.Time `json:"date"`
}

func (q *Queries) AddPurchaseOrder(ctx context.Context, arg AddPurchaseOrderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addPurchaseOrder,
		arg.EmpID,
		arg.CustomerContact,
		arg.CustomerAddress,
		arg.Date,
	)
}

const addTimecard = `-- name: AddTimecard :execresult
INSERT INTO timecard(emp_id) VALUES (?)
`

func (q *Queries) AddTimecard(ctx context.Context, empID int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTimecard, empID)
}

const addTimecardRecord = `-- name: AddTimecardRecord :exec
INSERT INTO timecard_record(charge_number,card_id,hours,date) VALUES (?,?,?,?)
`

type AddTimecardRecordParams struct {
	ChargeNumber int64     `json:"charge_number"`
	CardID       int64     `json:"card_id"`
	Hours        int32     `json:"hours"`
	Date         time.Time `json:"date"`
}

func (q *Queries) AddTimecardRecord(ctx context.Context, arg AddTimecardRecordParams) error {
	_, err := q.db.ExecContext(ctx, addTimecardRecord,
		arg.ChargeNumber,
		arg.CardID,
		arg.Hours,
		arg.Date,
	)
	return err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
UPDATE employees SET deleted = 1 where id = ?
`

func (q *Queries) DeleteEmployee(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, id)
	return err
}

const deleteOrderInfoById = `-- name: DeleteOrderInfoById :exec
DELETE FROM order_info WHERE order_id = ?
`

func (q *Queries) DeleteOrderInfoById(ctx context.Context, orderID int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrderInfoById, orderID)
	return err
}

const deletePurchaseOrderById = `-- name: DeletePurchaseOrderById :exec
DELETE FROM purchase_order WHERE id = ?
`

func (q *Queries) DeletePurchaseOrderById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePurchaseOrderById, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, password, type, mail, social_security_number, standard_tax_deductions, other_deductions, phone_number, salary_rate, hour_limit, payment_method, deleted, root FROM employees WHERE id = ?
`

func (q *Queries) GetUser(ctx context.Context, id int64) (Employee, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Type,
		&i.Mail,
		&i.SocialSecurityNumber,
		&i.StandardTaxDeductions,
		&i.OtherDeductions,
		&i.PhoneNumber,
		&i.SalaryRate,
		&i.HourLimit,
		&i.PaymentMethod,
		&i.Deleted,
		&i.Root,
	)
	return i, err
}

const insertBank = `-- name: InsertBank :exec
INSERT INTO employee_account(id,bank_name,account_number)
	VALUES (?,?,?)
`

type InsertBankParams struct {
	ID            int64  `json:"id"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
}

func (q *Queries) InsertBank(ctx context.Context, arg InsertBankParams) error {
	_, err := q.db.ExecContext(ctx, insertBank, arg.ID, arg.BankName, arg.AccountNumber)
	return err
}

const listEmployees = `-- name: ListEmployees :many
SELECT
  id, name, password, type, mail, social_security_number, standard_tax_deductions, other_deductions, phone_number, salary_rate, hour_limit, payment_method, deleted, root
FROM
  employees
ORDER BY
  id
`

func (q *Queries) ListEmployees(ctx context.Context) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Password,
			&i.Type,
			&i.Mail,
			&i.SocialSecurityNumber,
			&i.StandardTaxDeductions,
			&i.OtherDeductions,
			&i.PhoneNumber,
			&i.SalaryRate,
			&i.HourLimit,
			&i.PaymentMethod,
			&i.Deleted,
			&i.Root,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectActiveTimecard = `-- name: SelectActiveTimecard :one
SELECT id, emp_id, start_date, committed FROM timecard WHERE emp_id = ?
`

func (q *Queries) SelectActiveTimecard(ctx context.Context, empID int64) (Timecard, error) {
	row := q.db.QueryRowContext(ctx, selectActiveTimecard, empID)
	var i Timecard
	err := row.Scan(
		&i.ID,
		&i.EmpID,
		&i.StartDate,
		&i.Committed,
	)
	return i, err
}

const selectEmployeeById = `-- name: SelectEmployeeById :one
SELECT
  id, name, password, type, mail, social_security_number, standard_tax_deductions, other_deductions, phone_number, salary_rate, hour_limit, payment_method, deleted, root
from
  employees
where
  id = ?
LIMIT
  1
`

func (q *Queries) SelectEmployeeById(ctx context.Context, id int64) (Employee, error) {
	row := q.db.QueryRowContext(ctx, selectEmployeeById, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Type,
		&i.Mail,
		&i.SocialSecurityNumber,
		&i.StandardTaxDeductions,
		&i.OtherDeductions,
		&i.PhoneNumber,
		&i.SalaryRate,
		&i.HourLimit,
		&i.PaymentMethod,
		&i.Deleted,
		&i.Root,
	)
	return i, err
}

const selectEmployeeByMail = `-- name: SelectEmployeeByMail :one
SELECT id, name, password, type, mail, social_security_number, standard_tax_deductions, other_deductions, phone_number, salary_rate, hour_limit, payment_method, deleted, root FROM employees WHERE mail = ?
`

func (q *Queries) SelectEmployeeByMail(ctx context.Context, mail string) (Employee, error) {
	row := q.db.QueryRowContext(ctx, selectEmployeeByMail, mail)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Type,
		&i.Mail,
		&i.SocialSecurityNumber,
		&i.StandardTaxDeductions,
		&i.OtherDeductions,
		&i.PhoneNumber,
		&i.SalaryRate,
		&i.HourLimit,
		&i.PaymentMethod,
		&i.Deleted,
		&i.Root,
	)
	return i, err
}

const selectOrderById = `-- name: SelectOrderById :one
select id, emp_id, customer_contact, customer_address, date, closed from purchase_order where id = ?
`

func (q *Queries) SelectOrderById(ctx context.Context, id int64) (PurchaseOrder, error) {
	row := q.db.QueryRowContext(ctx, selectOrderById, id)
	var i PurchaseOrder
	err := row.Scan(
		&i.ID,
		&i.EmpID,
		&i.CustomerContact,
		&i.CustomerAddress,
		&i.Date,
		&i.Closed,
	)
	return i, err
}

const selectOrderInfoById = `-- name: SelectOrderInfoById :one
select order_id, product_id, amount from order_info where order_id = ?
`

func (q *Queries) SelectOrderInfoById(ctx context.Context, orderID int64) (OrderInfo, error) {
	row := q.db.QueryRowContext(ctx, selectOrderInfoById, orderID)
	var i OrderInfo
	err := row.Scan(&i.OrderID, &i.ProductID, &i.Amount)
	return i, err
}

const updateEmployee = `-- name: UpdateEmployee :exec
UPDATE employees SET type = ?,mail = ?,social_security_number=?,
	standard_tax_deductions=?,other_deductions=?,phone_number = ?,
	salary_rate=?,hour_limit=? where id = ?
`

type UpdateEmployeeParams struct {
	Type                  EmployeesType `json:"type"`
	Mail                  string        `json:"mail"`
	SocialSecurityNumber  string        `json:"social_security_number"`
	StandardTaxDeductions string        `json:"standard_tax_deductions"`
	OtherDeductions       string        `json:"other_deductions"`
	PhoneNumber           string        `json:"phone_number"`
	SalaryRate            string        `json:"salary_rate"`
	HourLimit             sql.NullInt32 `json:"hour_limit"`
	ID                    int64         `json:"id"`
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployee,
		arg.Type,
		arg.Mail,
		arg.SocialSecurityNumber,
		arg.StandardTaxDeductions,
		arg.OtherDeductions,
		arg.PhoneNumber,
		arg.SalaryRate,
		arg.HourLimit,
		arg.ID,
	)
	return err
}

const updateOrderInfo = `-- name: UpdateOrderInfo :exec
UPDATE order_info SET product_id = ?,amount = ?  where order_id = ?
`

type UpdateOrderInfoParams struct {
	ProductID int64  `json:"product_id"`
	Amount    string `json:"amount"`
	OrderID   int64  `json:"order_id"`
}

func (q *Queries) UpdateOrderInfo(ctx context.Context, arg UpdateOrderInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateOrderInfo, arg.ProductID, arg.Amount, arg.OrderID)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE employees SET password = ?
WHERE id = ?
`

type UpdatePasswordParams struct {
	Password sql.NullString `json:"password"`
	ID       int64          `json:"id"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.Password, arg.ID)
	return err
}

const updatePaymentMethod = `-- name: UpdatePaymentMethod :exec
UPDATE employees SET payment_method = ? where id = ?
`

type UpdatePaymentMethodParams struct {
	PaymentMethod EmployeesPaymentMethod `json:"payment_method"`
	ID            int64                  `json:"id"`
}

func (q *Queries) UpdatePaymentMethod(ctx context.Context, arg UpdatePaymentMethodParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentMethod, arg.PaymentMethod, arg.ID)
	return err
}

const updatePaymentMethodWithMail = `-- name: UpdatePaymentMethodWithMail :exec
UPDATE employees SET payment_method = ?,mail = ?  where id = ?
`

type UpdatePaymentMethodWithMailParams struct {
	PaymentMethod EmployeesPaymentMethod `json:"payment_method"`
	Mail          string                 `json:"mail"`
	ID            int64                  `json:"id"`
}

func (q *Queries) UpdatePaymentMethodWithMail(ctx context.Context, arg UpdatePaymentMethodWithMailParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentMethodWithMail, arg.PaymentMethod, arg.Mail, arg.ID)
	return err
}

const updatePurchaseOrder = `-- name: UpdatePurchaseOrder :exec
UPDATE purchase_order SET customer_contact = ?,customer_address =? , date = ? WHERE id = ?
`

type UpdatePurchaseOrderParams struct {
	CustomerContact string    `json:"customer_contact"`
	CustomerAddress string    `json:"customer_address"`
	Date            time.Time `json:"date"`
	ID              int64     `json:"id"`
}

func (q *Queries) UpdatePurchaseOrder(ctx context.Context, arg UpdatePurchaseOrderParams) error {
	_, err := q.db.ExecContext(ctx, updatePurchaseOrder,
		arg.CustomerContact,
		arg.CustomerAddress,
		arg.Date,
		arg.ID,
	)
	return err
}
