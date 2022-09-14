package entity

import (
	"time"

	"github.com/FaisalMashuri/submission-golang/app/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductEntity struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey"`
	CodeProduct string
	NameProduct string
	Description string
	Price       int
	UOM         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *ProductEntity) BeforeCreate(tx *gorm.DB) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

func (p *ProductEntity) ToDomain(data ProductEntity) domain.ProductDomain {
	return domain.ProductDomain{
		Id:          data.Id,
		CodeProduct: data.CodeProduct,
		NameProduct: data.NameProduct,
		Description: data.Description,
		Price:       data.Price,
		UOM:         data.UOM,
	}
}

func (p *ProductEntity) ToListDomain(data []ProductEntity) []domain.ProductDomain {
	var listDomain []domain.ProductDomain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain(item))
	}
	return listDomain
}
