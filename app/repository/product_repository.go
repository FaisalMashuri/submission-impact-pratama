package repository

import (
	"context"
	"time"

	"github.com/FaisalMashuri/submission-golang/app/domain"
	"github.com/FaisalMashuri/submission-golang/app/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, domain domain.ProductDomain) (productDomain domain.ProductDomain, err error)
	GetAllProduct(ctx context.Context) (allOrder []domain.ProductDomain, err error)
	GetById(ctx context.Context, id string) (domain.ProductDomain, error)
	UpdateProduct(ctx context.Context, id string, product domain.ProductDomain) (domain.ProductDomain, error)
	DeleteProduct(ctx context.Context, id string) (err error)
}

func NewProductRepositoryImpl(conn *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Conn: conn}
}

// ProductRepositoryImpl implementation for ProductRepository
type ProductRepositoryImpl struct {
	Conn *gorm.DB
}

func (p ProductRepositoryImpl) CreateProduct(ctx context.Context, product domain.ProductDomain) (productDomain domain.ProductDomain, err error) {
	//TODO implement me
	var productModel entity.ProductEntity

	createdProduct := entity.ProductEntity{
		Id:          product.Id,
		CodeProduct: product.CodeProduct,
		NameProduct: product.NameProduct,
		Description: product.Description,
		Price:       product.Price,
		UOM:         product.UOM,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = p.Conn.Create(&createdProduct).Error
	if err != nil {
		return product, err
	}
	return productModel.ToDomain(createdProduct), nil
}

func (p ProductRepositoryImpl) GetAllProduct(ctx context.Context) (allOrder []domain.ProductDomain, err error) {
	//TODO implement me
	var productModel []entity.ProductEntity
	if err := p.Conn.Find(&productModel).Error; err != nil {
		return nil, err
	}
	for _, item := range productModel {
		allOrder = append(allOrder, item.ToDomain(item))
	}
	return allOrder, nil
}

func (p ProductRepositoryImpl) GetById(ctx context.Context, id string) (domain.ProductDomain, error) {
	//TODO implement me
	var productModel entity.ProductEntity
	if err := p.Conn.Where("id = ?", id).First(&productModel).Error; err != nil {
		return domain.ProductDomain{}, err
	}
	return productModel.ToDomain(productModel), nil
}

func (p ProductRepositoryImpl) UpdateProduct(ctx context.Context, id string, product domain.ProductDomain) (domain.ProductDomain, error) {
	//TODO implement me
	var productModel entity.ProductEntity
	if err := p.Conn.Where("id = ?", id).First(&productModel).Error; err != nil {
		return domain.ProductDomain{}, err
	}
	productModel.CodeProduct = product.CodeProduct
	productModel.NameProduct = product.NameProduct
	productModel.Description = product.Description
	productModel.Price = product.Price
	productModel.UOM = product.UOM
	if err := p.Conn.Save(&productModel).Error; err != nil {
		return domain.ProductDomain{}, err
	}
	return productModel.ToDomain(productModel), nil

}

func (p ProductRepositoryImpl) DeleteProduct(ctx context.Context, id string) (err error) {
	//TODO implement me
	var productModel entity.ProductEntity
	if err = p.Conn.Where("id = ?", id).Delete(&productModel).Error; err != nil {
		return
	}
	return
}
