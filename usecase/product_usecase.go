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
	// return []model.Product{}, nil
	products, err := p.repository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
