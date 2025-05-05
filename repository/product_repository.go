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
	query := "SELECT id, product_name, price, stock from product"
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
			&productObj.Price,
			&productObj.Stock,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		"VALUES ($1, $2) RETURNING ID")
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

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

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

	query.Close()
	return &produto, nil
}

func (pr *ProductRepository) UpdateProduct(id_product int, request model.Product) (*model.Product, error) {

	query, err := pr.connection.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(request.Name, request.Price, id_product)
	if err != nil {
		fmt.Println(err)
		return &model.Product{}, err
	}

	updateProduct := &model.Product{
		ID:    id_product,
		Name:  request.Name,
		Price: request.Price,
	}

	query.Close()
	return updateProduct, nil
}

func (pr *ProductRepository) UpdateStock(id_product int, request model.Product) (*model.Product, error) {
	query := "UPDATE product SET stock = $1 WHERE id = $2  RETURNING id, product_name, price, stock"
	row := pr.connection.QueryRow(query, request.Stock, id_product)

	var requestStock model.Product

	err := row.Scan(
		&requestStock.ID,
		&requestStock.Name,
		&requestStock.Price,
		&requestStock.Stock,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &requestStock, nil

}

func (pr *ProductRepository) DeleteProduct(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("DELETE  FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	var produto model.Product

	_, err = query.Exec(id_product)
	if err != nil {
		return &model.Product{}, err
	}

	query.Close()
	return &produto, nil
}
