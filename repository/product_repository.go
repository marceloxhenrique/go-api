package repository

import (
	"database/sql"
	"fmt"
	"go-products/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err

	}
	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}

func (pr *ProductRepository) AddProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(productId int) (model.Product, error) {

	query := "SELECT id, product_name, price FROM product WHERE id = $1"

	var product model.Product
	err := pr.connection.QueryRow(query, productId).Scan(
		&product.ID, &product.Name, &product.Price,
	)
	fmt.Println("ERROR", err == sql.ErrNoRows)
	fmt.Println("ERROR2")
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	return product, nil
}

func (pr *ProductRepository) DeleteProductById(productId int) (int64, error) {
	query := "DELETE FROM product WHERE id = $1"
	res, err := pr.connection.Exec(query, productId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil

}
