package usecase

import "github.com/devfullcycle/imersao12-go-esquenta/internal/entity"

type ListProductsOutputDto struct {
    ID    string
    Name  string
    Price float64
}

type ListProductsUseCase struct {
    ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
    return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
    products, err := u.ProductRepository.FindAll()
    if err != nil {
        return nil, err
    }

    var productsOutput []*ListProductsOutputDto
    for _, product := range products {
        productsOutput = append(productsOutput, &ListProductsOutputDto{
            ID:    product.ID,
            Name:  product.Name,
            Price: product.Price,
        })
    }
    return productsOutput, nil
}