package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUseCase) UpdateProductById(id_product int, product model.Product) (*model.Product, error) {

	productUpdated, err := pu.repository.UpdateProductByID(id_product, product)
	if err != nil {
		return nil, err
	}

	return productUpdated, nil

}

func (pu *ProductUseCase) DeleteProductById(id_product int) (int, error) {

	id_product, err := pu.repository.DeleteProductById(id_product)
	if err != nil {
		return 0, err
	}

	return id_product, nil

}