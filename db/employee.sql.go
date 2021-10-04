// Code generated by sqlc. DO NOT EDIT.
// source: employee.sql

package db

import (
	"context"
	"database/sql"
)

const createEmployee = `-- name: CreateEmployee :execresult
INSERT INTO
  employees (
    type,
    mail,
    social_security_number,
    standard_tax_deductions,
    other_duductions,
    phone_number,
    rate
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?)
`

type CreateEmployeeParams struct {
	Type                  EmployeesType `json:"type"`
	Mail                  string        `json:"mail"`
	SocialSecurityNumber  string        `json:"social_security_number"`
	StandardTaxDeductions string        `json:"standard_tax_deductions"`
	OtherDuductions       string        `json:"other_duductions"`
	PhoneNumber           string        `json:"phone_number"`
	Rate                  string        `json:"rate"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createEmployee,
		arg.Type,
		arg.Mail,
		arg.SocialSecurityNumber,
		arg.StandardTaxDeductions,
		arg.OtherDuductions,
		arg.PhoneNumber,
		arg.Rate,
	)
}

const listEmployees = `-- name: ListEmployees :many
SELECT
  id, name, password, type, mail, social_security_number, standard_tax_deductions, other_duductions, phone_number, rate, hour_limit, payment_method
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
			&i.OtherDuductions,
			&i.PhoneNumber,
			&i.Rate,
			&i.HourLimit,
			&i.PaymentMethod,
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
