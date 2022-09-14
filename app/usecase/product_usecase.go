package usecase

import (
	"context"
	"time"

	"github.com/FaisalMashuri/submission-golang/app/domain"
	"github.com/FaisalMashuri/submission-golang/app/repository"
)

type ProductUsecaseInterface interface {
	CreateProduct(ctx context.Context, domain domain.ProductDomain) (productDomain domain.ProductDomain, err error)
	GetAllProduct(ctx context.Context, domain domain.ProductDomain) (allOrder []domain.ProductDomain, err error)
	GetByID(ctx context.Context, id int) (domain.ProductDomain, error)
	UpdateProduct(ctx context.Context, domain domain.ProductDomain) (domain.ProductDomain, error)
	DeleteProduct(ctx context.Context, domain domain.ProductDomain) (err error)
}

type ProductUsecase struct {
	Repo           repository.ProductRepository
	ContextTimeout time.Duration
}

func NewProductUsecase(repo repository.ProductRepository, timeout time.Duration) *ProductUsecase {
	return &ProductUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (pu *ProductUsecase) CreateProduct(ctx context.Context, product domain.ProductDomain) (productDomain domain.ProductDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	productDomain, err = pu.Repo.CreateProduct(ctx, product)
	if err != nil {
		return
	}
	return
}

func (pu *ProductUsecase) GetAllProduct(ctx context.Context) (allOrder []domain.ProductDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	allOrder, err = pu.Repo.GetAllProduct(ctx)
	if err != nil {
		return
	}

	return
}

func (pu *ProductUsecase) GetProductByCodeProduct(ctx context.Context, codeProduct string) (productDomain domain.ProductDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()
	productDomain, err = pu.Repo.GetById(ctx, codeProduct)
	if err != nil {
		return
	}

	return
}

func (pu *ProductUsecase) UpdateProduct(ctx context.Context, id string, product domain.ProductDomain) (productDomain domain.ProductDomain, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()
	productDomain, err = pu.Repo.UpdateProduct(ctx, id, product)
	if err != nil {
		return
	}

	return
}

func (pu *ProductUsecase) DeleteProduct(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	if err = pu.Repo.DeleteProduct(ctx, id); err != nil {
		return
	}

	return
}
