package repository

import (
	"context"
	"order-management/internal/model"
)

func (repo *repo) GetOrderDetail(
	/*req*/ ctx context.Context, id int) (
	/*res*/ order model.Order, err error,
) {
	query := "SELECT id, title, description, user_id, status, price FROM orders WHERE id = $1"

	statemnt, err := repo.db.Prepare(query)
	if err != nil {
		return order, err
	}

	row, err := statemnt.Query(id)
	if err != nil {
		return order, err
	}

	for row.Next() {
		var id, user_id int
		var title, description, status string
		var price float64

		err = row.Scan(&id, &title, &description, &user_id, &status, &price)
		if err != nil {
			return order, err
		}

		order = model.Order{
			ID:          id,
			Title:       title,
			Description: description,
			UserID:      user_id,
			Status:      status,
			Price:       price,
		}
	}

	return order, nil
}

func (repo *repo) GetOrderByTitle(
	/*req*/ ctx context.Context, title string) (
	/*res*/ orders model.Orders, err error,
) {
	query := "SELECT id, title FROM orders WHERE title LIKE '%' || $1 || '%'"

	statemnt, err := repo.db.Prepare(query)
	if err != nil {
		return orders, err
	}

	defer statemnt.Close()

	rows, err := statemnt.Query(title)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			return orders, err
		}

		order := model.Order{
			ID:    id,
			Title: title,
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *repo) GetAllOrders(
	/*req*/ ctx context.Context, request model.GetAllOrdersRequest) (
	/*res*/ orders model.Orders, err error,
) {
	query := "SELECT id, title, description, user_id, status, price FROM orders ORDER BY id LIMIT $1 OFFSET $2"

	statemnt, err := repo.db.Prepare(query)
	if err != nil {
		return orders, err
	}

	defer statemnt.Close()

	rows, err := statemnt.Query(request.Limit, request.Page)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var id, user_id int
		var title, description, status string
		var price float64

		err = rows.Scan(&id, &title, &description, &user_id, &status, &price)
		if err != nil {
			return orders, err
		}

		order := model.Order{
			ID:          id,
			Title:       title,
			Description: description,
			UserID:      user_id,
			Status:      status,
			Price:       price,
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *repo) CreateOrder(
	/*req*/ ctx context.Context, request model.Order) (
	/*res*/ err error,
) {
	query := "INSERT INTO orders (title, description, user_id, status, price) VALUES ($1, $2, $3, $4, $5)"

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(request.Title, request.Description, request.UserID, request.Status, request.Price)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) UpdateStatus(
	/*req*/ ctx context.Context, request model.Order) (
	/*res*/ err error,
) {
	query := "UPDATE orders SET status = $1 WHERE id = $2"

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(request.Status, request.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) DeleteOrder(
	/*req*/ ctx context.Context, id int) (
	/*res*/ err error,
) {
	query := "DELETE FROM orders WHERE id = $1"

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
