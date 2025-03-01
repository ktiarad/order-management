package repository

import (
	"context"
	"order-management/internal/model"
)

func (repo *repo) GetCustomerDetail(
	/*req*/ ctx context.Context, id int) (
	/*res*/ customer model.Customer, err error,
) {
	query := "SELECT id, name FROM customers WHERE id = $1"

	statemnt, err := repo.db.Prepare(query)
	if err != nil {
		return customer, err
	}

	row, err := statemnt.Query(id)
	if err != nil {
		return customer, err
	}

	for row.Next() {
		var id int
		var name string

		err = row.Scan(&id, &name)
		if err != nil {
			return customer, err
		}

		customer = model.Customer{
			ID:   id,
			Name: name,
		}
	}

	return customer, nil
}

func (repo *repo) GetCustomerByName(
	/*req*/ ctx context.Context, name string) (
	/*res*/ customers model.Customers, err error,
) {
	query := "SELECT id, name FROM customers WHERE name LIKE '%' || $1 || '%'"

	statemnt, err := repo.db.Prepare(query)
	if err != nil {
		return customers, err
	}

	defer statemnt.Close()

	rows, err := statemnt.Query(name)
	if err != nil {
		return customers, err
	}

	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			return customers, err
		}

		customer := model.Customer{
			ID:   id,
			Name: name,
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (repo *repo) GetAllCustomers(
	/*req*/ ctx context.Context, request model.GetAllCustomersRequest) (
	/*res*/ customers model.Customers, err error,
) {
	query := "SELECT id, name FROM customers ORDER BY id LIMIT $1 OFFSET $2"

	statemnt, err := repo.db.Prepare(query)
	if err != nil {
		return customers, err
	}

	defer statemnt.Close()

	rows, err := statemnt.Query(request.Limit, request.Page)
	if err != nil {
		return customers, err
	}

	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			return customers, err
		}

		customer := model.Customer{
			ID:   id,
			Name: name,
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (repo *repo) CreateCustomer(
	/*req*/ ctx context.Context, request model.CreateCustomerRequest) (
	/*res*/ err error,
) {
	query := "INSERT INTO customers (name) VALUES ($1)"

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(request.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) UpdateCustomer(
	/*req*/ ctx context.Context, request model.Customer) (
	/*res*/ err error,
) {
	query := "UPDATE customers SET name = $1 WHERE id = $2"

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(request.Name, request.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) DeleteCustomer(
	/*req*/ ctx context.Context, id int) (
	/*res*/ err error,
) {
	query := "DELETE FROM customers WHERE id = $1"

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
