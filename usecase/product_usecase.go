package usecase

import (
	"go-products/model"
	"go-products/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (p *ProductUsecase) GetAllProducts() ([]model.Product, error) {
	products, err := p.repository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductUsecase) AddProduct(product model.Product) (model.Product, error) {
	productId, err := p.repository.AddProduct(product)

	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (p *ProductUsecase) GetProductById(productId int) (model.Product, error) {
	product, err := p.repository.GetProductById(productId)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
