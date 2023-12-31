package database

import (
	"database/sql"
	"fmt"

	"github.com/allgurgel/FC-GO-EXPERT-DESAFIO-CLEAN-ARCH/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) List(page, limit int, sort string) ([]entity.Order, error) {
	var orders []entity.Order
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	query := "SELECT * from orders"

	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		query = fmt.Sprintf("SELECT * FROM orders ORDER BY final_price %s LIMIT %d OFFSET %d", sort, limit, offset)
	}
	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil

}
