package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(id_product int, product model.Product) (model.Product, error) {

	updateProduct, err := pu.repository.UpdateProduct(id_product, product)
	if err != nil {
		return model.Product{}, err
	}
	return *updateProduct, nil
}

func (pu *ProductUsecase) UpdateStock(id_product int, product model.Product) (*model.Product, error) {

	updateStock, err := pu.repository.UpdateStock(id_product, product)
	if err != nil {
		return nil, err
	}
	return updateStock, nil
}

func (pu *ProductUsecase) DeleteProduct(id_product int) (*model.Product, error) {

	product, err := pu.repository.DeleteProduct(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
