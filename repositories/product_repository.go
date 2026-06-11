package repositories

import (
	"database/sql"
	"go-rest-api/config"
	"go-rest-api/models"
	"strings"
	
)

func GetProducts(limit int, offset int, search string, sort string, order string) ([]models.Product, error){
	query := "SELECT id, name, price, stock FROM products WHERE name LIKE ? ORDER BY " + sort + " " + order + " LIMIT ? OFFSET ?"

	rows, err := config.DB.Query(
		query,
		"%"+search+"%",
		limit,
		offset,
	)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}


	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	var product models.Product

	err := config.DB.QueryRow("SELECT id, name, price, stock FROM products WHERE id = ?", id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
	)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func CreateProduct(product *models.Product) error {
	result, err := config.DB.Exec(
		"INSERT INTO products (name, price, stock) VALUES (?, ?, ?)",
		product.Name,
		product.Price,
		product.Stock,
	)
	if err != nil {
		return err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = int(insertedID)
	return nil
}

func UpdateProduct(product *models.Product) error {
	result, err := config.DB.Exec(
		"UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?",
		product.Name,
		product.Price,
		product.Stock,
		product.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeleteProduct(id int) error {
	result, err := config.DB.Exec(
		"DELETE FROM products WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func containsIgnoreCase (text, search string) bool {
	return strings.Contains (
		strings.ToLower (text),
		strings.ToLower (search),
	)
}