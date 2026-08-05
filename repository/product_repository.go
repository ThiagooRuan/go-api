package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
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

	defer rows.Close()

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING ID")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer query.Close()

	err = query.QueryRow(product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return product.ID, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	defer query.Close()

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &produto, nil
}

func (pr *ProductRepository) UpdateProductByID(id_product int, product model.Product) (*model.Product, error) {
	query, err := pr.connection.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3 RETURNING id, product_name, price")
	if err != nil {
		return nil, err
	}

	defer query.Close()

	var produto model.Product

	err = query.QueryRow(
		product.Name,
		product.Price,
		id_product,
	).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &produto, nil
}

func (pr *ProductRepository) DeleteProductById(id_product int) (int, error) {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		return 0, err
	}

	defer query.Close()

	result, err := query.Exec(id_product)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	fmt.Println("Rows affected:", rowsAffected)

	if rowsAffected == 0 {
		return 0, fmt.Errorf("produto não encontrado")
	}

	return id_product, nil
}
