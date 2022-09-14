package response

import (
	"github.com/FaisalMashuri/submission-golang/app/domain"
	"github.com/google/uuid"
)

type ProductResponse struct {
	Id          uuid.UUID `json:"id" form:"id"`
	CodeProduct string    `json:"codeProduct" form:"codeProduct"`
	NameProduct string    `json:"nameProduct" form:"nameProduct"`
	Description string    `json:"description" form:"description"`
	Price       int       `json:"price" form:"price"`
	UOM         string    `json:"uom" form:"uom"`
}

func ProductFromDomain(domain domain.ProductDomain) ProductResponse {
	return ProductResponse{
		Id:          domain.Id,
		CodeProduct: domain.CodeProduct,
		NameProduct: domain.NameProduct,
		Description: domain.Description,
		Price:       domain.Price,
		UOM:         domain.UOM,
	}
}
